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
	"bytes"
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/hashicorp/terraform/flatmap"
	"gonum.org/v1/gonum/graph"
	simplegraph "gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

var SgAllowEmptyValues = []string{"tags."}

type void struct{}

var member void

type SecurityGenerator struct {
	AWSService
}

type ByGroupPair []types.UserIdGroupPair

func (b ByGroupPair) Len() int      { return len(b) }
func (b ByGroupPair) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b ByGroupPair) Less(i, j int) bool {
	if b[i].GroupId != nil && b[j].GroupId != nil {
		return *b[i].GroupId < *b[j].GroupId
	}
	if b[i].GroupName != nil && b[j].GroupName != nil {
		return *b[i].GroupName < *b[j].GroupName
	}

	panic("mismatched security group rules, may be a terraform bug")
}

func (SecurityGenerator) createResources(securityGroups []types.SecurityGroup) []terraformutils.Resource {
	var sgIDsToMoveOut []string
	_, shouldSplitRules := os.LookupEnv("SPLIT_SG_RULES")
	if shouldSplitRules {
		for _, sg := range securityGroups {
			sgIDsToMoveOut = append(sgIDsToMoveOut, *sg.GroupId)
		}
	} else {
		sgIDsToMoveOut = findSgsToMoveOut(securityGroups)
	}

	var resources []terraformutils.Resource
	for _, sg := range securityGroups {
		if sg.VpcId == nil {
			continue
		}
		ruleAttributes := map[string]interface{}{}
		// we must move out all of the rules - https://github.com/hashicorp/terraform/issues/11011#issuecomment-283076580
		for _, groupIDToMoveOut := range sgIDsToMoveOut {
			if groupIDToMoveOut == *sg.GroupId {
				ruleAttributes["clearRules"] = true
				for _, rule := range sg.IpPermissions {
					resources = processRule(rule, "ingress", sg, resources)
				}
				for _, rule := range sg.IpPermissionsEgress {
					resources = processRule(rule, "egress", sg, resources)
				}
			}
		}

		resources = append(resources, terraformutils.NewResource(
			StringValue(sg.GroupId),
			strings.Trim(StringValue(sg.GroupName)+"_"+StringValue(sg.GroupId), " "),
			"aws_security_group",
			"aws",
			map[string]string{},
			SgAllowEmptyValues,
			ruleAttributes))
	}
	return resources
}

func processRule(rule types.IpPermission, ruleType string, sg types.SecurityGroup, resources []terraformutils.Resource) []terraformutils.Resource {
	if rule.UserIdGroupPairs != nil && len(rule.UserIdGroupPairs) > 0 {
		if len(rule.IpRanges) > 0 { // we must unwind coupled CIDR IPv4 range + security group rules
			attributes := baseRuleAttributes(ruleType, rule, sg)
			resources = append(resources, terraformutils.NewResource(
				permissionID(*sg.GroupId, ruleType, "", rule),
				permissionID(*sg.GroupId, ruleType, "", rule),
				"aws_security_group_rule",
				"aws",
				flatmap.Flatten(attributes),
				SgAllowEmptyValues,
				map[string]interface{}{}))
		}
		if len(rule.Ipv6Ranges) > 0 { // we must unwind coupled CIDR IPv6 range + security group rules
			attributes := baseRuleAttributes(ruleType, rule, sg)
			resources = append(resources, terraformutils.NewResource(
				permissionID(*sg.GroupId, ruleType, "", rule),
				permissionID(*sg.GroupId, ruleType, "", rule),
				"aws_security_group_rule",
				"aws",
				flatmap.Flatten(attributes),
				SgAllowEmptyValues,
				map[string]interface{}{}))
		}
		for _, groupPair := range rule.UserIdGroupPairs {
			attributes := baseRuleAttributes(ruleType, rule, sg)
			delete(attributes, "cidr_blocks")
			delete(attributes, "ipv6_cidr_blocks")
			if *groupPair.GroupId == *sg.GroupId { // Solution to C1
				attributes["self"] = true
			} else {
				attributes["source_security_group_id"] = *groupPair.GroupId
			}

			resources = append(resources, terraformutils.NewResource(
				permissionID(*sg.GroupId, ruleType, *groupPair.GroupId, rule),
				permissionID(*sg.GroupId, ruleType, *groupPair.GroupId, rule),
				"aws_security_group_rule",
				"aws",
				flatmap.Flatten(attributes),
				SgAllowEmptyValues,
				map[string]interface{}{}))
		}
	} else {
		attributes := baseRuleAttributes(ruleType, rule, sg)
		resources = append(resources, terraformutils.NewResource(
			permissionID(*sg.GroupId, ruleType, "", rule),
			permissionID(*sg.GroupId, ruleType, "", rule),
			"aws_security_group_rule",
			"aws",
			flatmap.Flatten(attributes),
			SgAllowEmptyValues,
			map[string]interface{}{}))
	}
	return resources
}

func baseRuleAttributes(ruleType string, rule types.IpPermission, sg types.SecurityGroup) map[string]interface{} {
	attributes := map[string]interface{}{
		"type":              ruleType,
		"cidr_blocks":       ipRange(rule),
		"ipv6_cidr_blocks":  ip6Range(rule),
		"prefix_list_ids":   prefixes(rule),
		"from_port":         fromPort(rule),
		"protocol":          *rule.IpProtocol,
		"security_group_id": *sg.GroupId,
		"to_port":           toPort(rule),
	}
	return attributes
}

// Let's try to find all cycles by applying Johnson's method on the directed graph
// We cannot build a line graph and move out only rules because of hashicorp/terraform#11011
func findSgsToMoveOut(securityGroups []types.SecurityGroup) []string {
	// Vertexes are security groups, edges are rules. The task is to find correct set of rule definitions, so that we
	// won't have cycles
	sourceGraph := simplegraph.NewDirectedGraph()
	idToSg := make(map[int]types.SecurityGroup)
	sgToIdx := make(map[string]int64)
	for idx, sg := range securityGroups {
		idToSg[idx] = sg
		sgToIdx[StringValue(sg.GroupId)] = int64(idx)
		sourceGraph.AddNode(sourceGraph.NewNode())
	}
	for idx, sg := range securityGroups {
		for _, rule := range sg.IpPermissions {
			pairs := rule.UserIdGroupPairs
			for _, pair := range pairs {
				if pair.GroupId != nil {
					fromNode := sourceGraph.Node(int64(idx))
					toNode := sourceGraph.Node(sgToIdx[StringValue(pair.GroupId)])
					if fromNode.ID() != toNode.ID() {
						sourceGraph.SetEdge(sourceGraph.NewEdge(fromNode, toNode))
					}
				}
			}
		}
	}

	cyclesInLineGraph := topo.DirectedCyclesIn(sourceGraph) // C1 cycles won't be found but Terraform solves that issue
	resultingSet := make(map[string]void)

	for _, v := range cyclesInLineGraph {
		if elementAlreadyFound(resultingSet, v, idToSg) {
			continue
		}

		// Try to move out node with lowest number of rules
		group := idToSg[int(v[0].ID())]
		for _, vi := range v {
			viGroup := idToSg[int(vi.ID())]
			if len(viGroup.IpPermissions) < len(group.IpPermissions) {
				group = viGroup
			}
		}

		resultingSet[*group.GroupId] = member
	}

	result := make([]string, len(resultingSet))
	i := 0
	for k := range resultingSet {
		result[i] = k
		i++
	}

	return result
}

func elementAlreadyFound(resultingSet map[string]void, v []graph.Node, idToSg map[int]types.SecurityGroup) bool {
	for k := range resultingSet {
		for _, vi := range v {
			viGroupID := *idToSg[int(vi.ID())].GroupId
			if k == viGroupID {
				return true
			}
		}
	}
	return false
}

func (g *SecurityGenerator) InitResources() error {
	config, err := g.generateConfig()
	if err != nil {
		return err
	}
	svc := ec2.NewFromConfig(config)
	p := ec2.NewDescribeSecurityGroupsPaginator(svc, &ec2.DescribeSecurityGroupsInput{})
	var resourcesToFilter []types.SecurityGroup
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		resourcesToFilter = append(resourcesToFilter, page.SecurityGroups...)
	}
	sort.Slice(resourcesToFilter, func(i, j int) bool {
		return *resourcesToFilter[i].GroupId < *resourcesToFilter[j].GroupId
	})
	g.Resources = g.createResources(resourcesToFilter)

	return nil
}

func (g *SecurityGenerator) PostConvertHook() error {
	for _, resource := range g.Resources {
		if resource.InstanceInfo.Type == "aws_security_group_rule" {
			if resource.Item["self"] == "true" {
				delete(resource.Item, "source_security_group_id")
			}
		} else if resource.InstanceInfo.Type == "aws_security_group" {
			if resource.Item["clearRules"] == true {
				delete(resource.Item, "ingress")
				delete(resource.Item, "egress")
				delete(resource.Item, "clearRules")
				continue
			}

			if val, ok := resource.Item["ingress"]; ok {
				g.sortRules(val.([]interface{}))
			}
			if val, ok := resource.Item["egress"]; ok {
				g.sortRules(val.([]interface{}))
			}
		}
	}
	return nil
}

func (g *SecurityGenerator) sortRules(rules []interface{}) {
	for _, rule := range rules {
		ruleMap := rule.(map[string]interface{})
		g.sortIfExist("cidr_blocks", ruleMap)
		g.sortIfExist("ipv6_cidr_blocks", ruleMap)
		g.sortIfExist("security_groups", ruleMap)
	}
	sort.Slice(rules, func(i, j int) bool {
		return fmt.Sprintf("%v", rules[i]) < fmt.Sprintf("%v", rules[j])
	})
}

func (g *SecurityGenerator) sortIfExist(attribute string, ruleMap map[string]interface{}) {
	if val, ok := ruleMap[attribute]; ok {
		sort.Slice(val.([]interface{}), func(i, j int) bool {
			return val.([]interface{})[i].(string) < val.([]interface{})[j].(string)
		})
	}
}

func permissionID(sgID, ruleType, groupID string, ip types.IpPermission) string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s_%s_%s_%d_%d_", sgID, ruleType, *ip.IpProtocol, fromPort(ip), toPort(ip)))

	if len(ip.IpRanges) > 0 {
		s := make([]string, len(ip.IpRanges))
		for i, r := range ip.IpRanges {
			s[i] = *r.CidrIp
		}
		sort.Strings(s)

		for _, v := range s {
			buf.WriteString(fmt.Sprintf("%s_", v))
		}
	}

	if len(ip.Ipv6Ranges) > 0 {
		s := make([]string, len(ip.Ipv6Ranges))
		for i, r := range ip.Ipv6Ranges {
			s[i] = *r.CidrIpv6
		}
		sort.Strings(s)

		for _, v := range s {
			buf.WriteString(fmt.Sprintf("%s_", v))
		}
	}

	if len(ip.PrefixListIds) > 0 {
		s := make([]string, len(ip.PrefixListIds))
		for i, pl := range ip.PrefixListIds {
			s[i] = *pl.PrefixListId
		}
		sort.Strings(s)

		for _, v := range s {
			buf.WriteString(fmt.Sprintf("%s_", v))
		}
	}

	if groupID != "" {
		buf.WriteString(fmt.Sprintf("%s_", groupID))
	}

	idPreformatted := buf.String()
	return idPreformatted[:len(idPreformatted)-1]
}

func fromPort(ip types.IpPermission) int {
	switch {
	case *ip.IpProtocol == "icmp":
		return -1
	case *ip.IpProtocol == "-1":
		return -1
	case ip.FromPort != nil && *ip.FromPort > 0:
		return int(*ip.FromPort)
	default:
		return 0
	}
}

func toPort(ip types.IpPermission) int {
	switch {
	case *ip.IpProtocol == "icmp":
		return -1
	case *ip.IpProtocol == "-1":
		return -1
	case ip.ToPort != nil && *ip.ToPort > 0:
		return int(*ip.ToPort)
	default:
		return 65536
	}
}

func ipRange(rule types.IpPermission) []string {
	result := make([]string, len(rule.IpRanges))
	for idx, rule := range rule.IpRanges {
		result[idx] = *rule.CidrIp
	}
	return result
}

func ip6Range(rule types.IpPermission) []string {
	result := make([]string, len(rule.Ipv6Ranges))
	for idx, rule := range rule.Ipv6Ranges {
		result[idx] = *rule.CidrIpv6
	}
	return result
}

func prefixes(rule types.IpPermission) []string {
	result := make([]string, len(rule.PrefixListIds))
	for idx, rule := range rule.PrefixListIds {
		result[idx] = *rule.PrefixListId
	}
	return result
}
