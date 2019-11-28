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
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"

	simplegraph "gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

var SgAllowEmptyValues = []string{"tags."}

type void struct{}

var member void

type SecurityGenerator struct {
	AWSService
}

func (SecurityGenerator) createResources(securityGroups []ec2.SecurityGroup) []terraform_utils.Resource {
	sgsToMoveOut := findSgsToMoveOut(securityGroups)

	fmt.Printf("%+v\n", sgsToMoveOut)

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

// Let's try to find all cycles by applying Johnson's method on the directed graph (very naive implementation)
func findSgsToMoveOut(securityGroups []ec2.SecurityGroup) []*ec2.SecurityGroup {
	// Edges are security groups, vertexes are rules. The task is to find correct set of rule definitions, so that we
	// won't have cycles
	// TODO verify cross account rules (are they working fine?)
	sourceGraph := simplegraph.NewDirectedGraph()
	idToSg := make(map[int]ec2.SecurityGroup)
	sgToIdx := make(map[string]int64)
	for idx, sg := range securityGroups {
		idToSg[idx] = sg
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
						sourceGraph.SetEdge(sourceGraph.NewEdge(fromNode, toNode))
					}
				}
			}
		}
	}

	cyclesInLineGraph := topo.DirectedCyclesIn(sourceGraph)

	resultingSet := make(map[*ec2.SecurityGroup]void)

	for idx, sg := range securityGroups {
		for _, rule := range sg.IpPermissions {
			pairs := rule.UserIdGroupPairs
			for _, pair := range pairs {
				if pair.GroupId != nil {
					fromNode := sourceGraph.Node(int64(idx))
					toNode := sourceGraph.Node(sgToIdx[aws.StringValue(pair.GroupId)])
					if fromNode.ID() == toNode.ID() { // references to itself
						resultingSet[&sg] = member
					}
				}
			}
		}
	}

	for _, v := range cyclesInLineGraph {
		// Try to move out first node
		id := v[0].ID()
		group := idToSg[int(id)]
		// TODO select rule with SG with least number of rules
		// TODO check if any has been selected before as 1C
		resultingSet[&group] = member
	}

	result := make([]*ec2.SecurityGroup, len(resultingSet))
	i := 0
	for k := range resultingSet {
		// Try to move out first node
		result[i] = k
		i++
	}

	return result
}

// Generate TerraformResources from AWS API,
// from each security group create 1 TerraformResource.
// Need GroupId as ID for terraform resource
func (g *SecurityGenerator) InitResources() error {
	config, err := g.generateConfig()
	if err != nil {
		return err
	}
	svc := ec2.New(config)
	p := ec2.NewDescribeSecurityGroupsPaginator(svc.DescribeSecurityGroupsRequest(&ec2.DescribeSecurityGroupsInput{}))
	for p.Next(context.Background()) {
		g.Resources = append(g.Resources, g.createResources(p.CurrentPage().SecurityGroups)...)
	}

	if err := p.Err(); err != nil {
		return err
	}
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
