// Copyright 2018 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aws

import (
	"fmt"
	"gonum.org/v1/gonum/graph"
	"math"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"

	path_graph "gonum.org/v1/gonum/graph/path"
	simple_graph "gonum.org/v1/gonum/graph/simple"
)

var SgAllowEmptyValues = []string{"tags."}

type void struct{}

var member void
var absent = math.Inf(1)

type SecurityGenerator struct {
	AWSService
}

type SecurityGroupRule struct {
	sourceNode int64
	id         string // Terraform format SECURITYGROUPID_TYPE_PROTOCOL_FROMPORT_TOPORT_SOURCE[_SOURCE]*
}

func (SecurityGenerator) createResources(securityGroups []*ec2.SecurityGroup) []terraform_utils.Resource {

	spanningTree := findSpanningTree(securityGroups)

	fmt.Printf("%+v\n", spanningTree)

	var resources []terraform_utils.Resource
	for _, sg := range securityGroups {
		if sg.VpcId == nil {
			continue
		}
		resources = append(resources, terraform_utils.NewSimpleResource(
			aws.StringValue(sg.GroupId),
			strings.Trim(aws.StringValue(sg.GroupName)+"_"+aws.StringValue(sg.GroupId), " "),
			"aws_security_group",
			"aws",
			SgAllowEmptyValues))
	}
	return resources
}

func findSpanningTree(securityGroups []*ec2.SecurityGroup) *simple_graph.WeightedDirectedGraph {
	// Edges are security groups, vertexes are rules. The task is to find correct set of rule definitions, so that we
	// won't have cycles. Direction in graph is placement of the edge definition.
	sourceGraph := simple_graph.NewWeightedDirectedGraph(-1, absent)
	idToSg := make(map[int64]*ec2.SecurityGroup)
	idToSecurityGroupRule := make(map[int]SecurityGroupRule)
	sgToIdx := make(map[string]int64)
	for idx, sg := range securityGroups {
		idToSg[int64(idx)] = sg
		sgToIdx[aws.StringValue(sg.GroupId)] = int64(idx)
		sourceGraph.AddNode(sourceGraph.NewNode())
	}
	for idx, sg := range securityGroups {
		for _, rule := range sg.IpPermissions {
			pairs := rule.UserIdGroupPairs
			for _, pair := range pairs {
				if pair.GroupId != nil {
					fromNode := sourceGraph.Node(int64(idx))
					toNode := sourceGraph.Node(sgToIdx[aws.StringValue(pair.GroupId)])
					if fromNode.ID() != toNode.ID() {
						i := len(idToSecurityGroupRule)
						idToSecurityGroupRule[i] = SecurityGroupRule{
							sourceNode: fromNode.ID(),
							id:         "IDDDD",
						}
						sourceGraph.SetWeightedEdge(sourceGraph.NewWeightedEdge(fromNode, toNode, float64(i)))
					}
				}
			}
		}
	}
	// we'll try to split edges that are connected to security group with lowest number of rules
	// ref https://stackoverflow.com/a/947519/3784897
	// TODO add weights on resulting graph's edges
	lineGraph := simple_graph.NewWeightedUndirectedGraph(-1, math.Inf(-1))
	allSourceEdges := sourceGraph.Edges()
	for allSourceEdges.Next() {
		lineGraph.AddNode(lineGraph.NewNode()) // create a new line node for each source edge
	}
	sourceNodes := sourceGraph.Nodes()
	for sourceNodes.Next() {
		sourceNode := sourceNodes.Node()
		set := make(map[graph.Node]void)
		allSourceEdges := sourceGraph.Edges()
		for allSourceEdges.Next() {
			sourceEdge := allSourceEdges.Edge()
			if sourceEdge.From().ID() == sourceNode.ID() {
				w, _ := sourceGraph.Weight(sourceEdge.From().ID(), sourceEdge.To().ID())
				if w != absent {
					node := lineGraph.NewNode()
					// TODO save reference between old edge and new node
					set[node] = member
				}
			}
		}
		group := idToSg[sourceNode.ID()]
		for k1 := range set { // create cliques
			for k2 := range set {
				if k1.ID() != k2.ID() {
					lineGraph.SetWeightedEdge(lineGraph.NewWeightedEdge(k1, k2, float64(len(group.IpPermissions))))
					// TODO save reference between old node and new edge (to keep track of security group references)
				}
			}
		}
	}
	result := simple_graph.NewWeightedDirectedGraph(-1, math.Inf(1))
	path_graph.Kruskal(result, lineGraph)

	return result
}

// Generate TerraformResources from AWS API,
// from each security group create 1 TerraformResource.
// Need GroupId as ID for terraform resource
func (g *SecurityGenerator) InitResources() error {
	sess := g.generateSession()
	svc := ec2.New(sess)
	var securityGroups []*ec2.SecurityGroup
	var err error

	err = svc.DescribeSecurityGroupsPages(&ec2.DescribeSecurityGroupsInput{}, func(securityGroupsOut *ec2.DescribeSecurityGroupsOutput, lastPage bool) bool {
		securityGroups = append(securityGroups, securityGroupsOut.SecurityGroups...)
		return !lastPage
	})
	if err != nil {
		return err
	}

	g.Resources = g.createResources(securityGroups)
	return nil
}

// PostGenerateHook - replace sg-xxxxx string to terraform ID in all security group
func (g *SecurityGenerator) PostConvertHook() error {
	for j, resource := range g.Resources {
		for _, typeOfRule := range []string{"ingress", "egress"} {
			if _, exist := resource.Item[typeOfRule]; !exist {
				continue
			}
			for i, k := range resource.Item[typeOfRule].([]interface{}) {
				ingresses := k.(map[string]interface{})
				for key, ingress := range ingresses {
					if key != "security_groups" {
						continue
					}
					securityGroups := ingress.([]interface{})
					renamedSecurityGroups := []string{}
					for _, securityGroup := range securityGroups {
						found := false
						for _, i := range g.Resources {
							if i.InstanceState.ID == securityGroup {
								renamedSecurityGroups = append(renamedSecurityGroups, "${"+i.InstanceInfo.Type+"."+i.ResourceName+".id}")
								found = true
								break
							}
						}
						if !found {
							renamedSecurityGroups = append(renamedSecurityGroups, securityGroup.(string))
						}
					}
					g.Resources[j].Item[typeOfRule].([]interface{})[i].(map[string]interface{})["security_groups"] = renamedSecurityGroups
				}
			}
		}
	}
	return nil
}
