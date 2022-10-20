// Copyright 2019 The Terraformer Authors.
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

package ibm

import (
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

// VPCGenerator ...
type VPCRoutingTableGenerator struct {
	IBMService
}

func (g VPCRoutingTableGenerator) loadVPCRouteTableResources(vpcID, routeTableID, routeTableName string) terraformutils.Resource {
	resources := terraformutils.NewResource(
		fmt.Sprintf("%s/%s", vpcID, routeTableID),
		normalizeResourceName(routeTableName, false),
		"ibm_is_vpc_routing_table",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{})

	return resources
}

func (g VPCRoutingTableGenerator) loadVPCRouteTableRouteResources(vpcID, routeTableID, routeTableRouteID, routeTableRouteName string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s/%s/%s", vpcID, routeTableID, routeTableRouteID),
		normalizeResourceName(routeTableRouteName, false),
		"ibm_is_vpc_routing_table_route",
		"ibm",
		map[string]string{
			"routing_table": routeTableID,
			"action":        "deliver",
		},
		[]string{},
		map[string]interface{}{})

	// Deprecated parameters
	resource.IgnoreKeys = append(resource.IgnoreKeys,
		"^action$",
	)
	return resource
}

// InitResources ...
func (g *VPCRoutingTableGenerator) InitResources() error {
	region := g.Args["region"].(string)
	apiKey := os.Getenv("IC_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("no API key set")
	}

	isURL := GetVPCEndPoint(region)
	iamURL := GetAuthEndPoint()
	vpcoptions := &vpcv1.VpcV1Options{
		URL: isURL,
		Authenticator: &core.IamAuthenticator{
			ApiKey: apiKey,
			URL:    iamURL,
		},
	}
	vpcclient, err := vpcv1.NewVpcV1(vpcoptions)
	if err != nil {
		return err
	}
	start := ""
	var allrecs []vpcv1.VPC
	for {
		listVpcsOptions := &vpcv1.ListVpcsOptions{}
		if start != "" {
			listVpcsOptions.Start = &start
		}
		if rg := g.Args["resource_group"].(string); rg != "" {
			rg, err = GetResourceGroupID(apiKey, rg, region)
			if err != nil {
				return fmt.Errorf("Error Fetching Resource Group Id %s", err)
			}
			listVpcsOptions.ResourceGroupID = &rg
		}
		vpcs, response, err := vpcclient.ListVpcs(listVpcsOptions)
		if err != nil {
			return fmt.Errorf("Error Fetching vpcs %s\n%s", err, response)
		}
		start = GetNext(vpcs.Next)
		allrecs = append(allrecs, vpcs.Vpcs...)
		if start == "" {
			break
		}
	}

	for _, vpc := range allrecs {

		// routing table
		listVPCRoutingTablesOptions := &vpcv1.ListVPCRoutingTablesOptions{
			VPCID: vpc.ID,
		}
		tables, response, err := vpcclient.ListVPCRoutingTables(listVPCRoutingTablesOptions)
		if err != nil {
			return fmt.Errorf("Error Fetching vpc routing tables %s\n%s", err, response)
		}
		for _, table := range tables.RoutingTables {
			g.Resources = append(g.Resources, g.loadVPCRouteTableResources(*vpc.ID, *table.ID, *table.Name))
			listVPCRoutingTableRoutesOptions := &vpcv1.ListVPCRoutingTableRoutesOptions{
				VPCID:          vpc.ID,
				RoutingTableID: table.ID,
			}
			tableroutes, response, err := vpcclient.ListVPCRoutingTableRoutes(listVPCRoutingTableRoutesOptions)
			if err != nil {
				return fmt.Errorf("Error Fetching vpc route table routes %s\n%s", err, response)
			}
			for _, tableroute := range tableroutes.Routes {
				g.Resources = append(g.Resources, g.loadVPCRouteTableRouteResources(*vpc.ID, *table.ID, *tableroute.ID, *tableroute.Name))
			}
		}
	}

	return nil
}

func (g *VPCRoutingTableGenerator) PostConvertHook() error {
	for i, r := range g.Resources {
		if r.InstanceInfo.Type != "ibm_is_vpc_routing_table_route" {
			continue
		}
		for _, rt := range g.Resources {
			if rt.InstanceInfo.Type != "ibm_is_vpc_routing_table" {
				continue
			}
			if r.InstanceState.Attributes["routing_table"] == rt.InstanceState.Attributes["id"] {
				g.Resources[i].Item["routing_table"] = "${ibm_is_vpc_routing_table." + rt.ResourceName + ".id}"
			}
		}
	}

	return nil
}
