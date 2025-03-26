// Copyright 2020 The Terraformer Authors.
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

package azure

import (
	"context"
	"log"
	"regexp"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-03-01/network"
	"github.com/Azure/go-autorest/autorest"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/hashicorp/go-azure-helpers/authentication"
)

type LoadBalancerGenerator struct {
	AzureService
}

func (g *LoadBalancerGenerator) listLoadBalancerProbes(resourceGroupName string, loadBalancerName string) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	resourceManagerEndpoint := g.Args["config"].(authentication.Config).CustomResourceManagerEndpoint

	LoadBalancerProbesClient := network.NewLoadBalancerProbesClientWithBaseURI(resourceManagerEndpoint, subscriptionID)
	LoadBalancerProbesClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)
	loadBalancerProbeIterator, err := LoadBalancerProbesClient.ListComplete(ctx, resourceGroupName, loadBalancerName)

	if err != nil {
		return nil, err
	}
	for loadBalancerProbeIterator.NotDone() {
		loadBalancerProbe := loadBalancerProbeIterator.Value()
		// NOTE:
		// This works out the loadBalancer resource id from current probe
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/probes/probe1
		//
		// As the related data_source in azurerm provider works by starting to look up with loadbalancer_id
		// https://github.com/terraform-providers/terraform-provider-azurerm/blob/v2.18.0/azurerm/internal/services/network/lb_probe_resource.go#L186
		re := regexp.MustCompile(`/probes/.*$`)
		loadBalancerID := re.ReplaceAllLiteralString(*loadBalancerProbe.ID, "")
		resources = append(resources, terraformutils.NewResource(
			*loadBalancerProbe.ID,
			*loadBalancerProbe.Name,
			"azurerm_lb_probe",
			g.ProviderName,
			map[string]string{
				"loadbalancer_id": loadBalancerID,
			},
			[]string{},
			map[string]interface{}{},
		))

		if err := loadBalancerProbeIterator.Next(); err != nil {
			log.Println(err)
			break
		}
	}

	return resources, nil
}

func (g *LoadBalancerGenerator) listInboundNatRules(resourceGroupName string, loadBalancerName string) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	resourceManagerEndpoint := g.Args["config"].(authentication.Config).CustomResourceManagerEndpoint

	InboundNatRulesClient := network.NewInboundNatRulesClientWithBaseURI(resourceManagerEndpoint, subscriptionID)
	InboundNatRulesClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)
	InboundNatRuleIterator, err := InboundNatRulesClient.ListComplete(ctx, resourceGroupName, loadBalancerName)

	if err != nil {
		return nil, err
	}
	for InboundNatRuleIterator.NotDone() {
		InboundNatRule := InboundNatRuleIterator.Value()
		// NOTE:
		// Similar to above explanation, work out loadbalancer_id for azurerm datasource impl
		re := regexp.MustCompile(`/inboundNatRules/.*$`)
		loadBalancerID := re.ReplaceAllLiteralString(*InboundNatRule.ID, "")
		resources = append(resources, terraformutils.NewResource(
			*InboundNatRule.ID,
			*InboundNatRule.Name,
			"azurerm_lb_nat_rule",
			g.ProviderName,
			map[string]string{
				"loadbalancer_id": loadBalancerID,
			},
			[]string{},
			map[string]interface{}{},
		))

		if err := InboundNatRuleIterator.Next(); err != nil {
			log.Println(err)
			break
		}
	}

	return resources, nil
}

func (g *LoadBalancerGenerator) listLoadBalancerBackendAddressPools(resourceGroupName string, loadBalancerName string) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	resourceManagerEndpoint := g.Args["config"].(authentication.Config).CustomResourceManagerEndpoint

	LoadBalancerBackendAddressPoolsClient := network.NewLoadBalancerBackendAddressPoolsClientWithBaseURI(resourceManagerEndpoint, subscriptionID)
	LoadBalancerBackendAddressPoolsClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)
	loadBalancerBackendAddressPoolIterator, err := LoadBalancerBackendAddressPoolsClient.ListComplete(ctx, resourceGroupName, loadBalancerName)

	if err != nil {
		return nil, err
	}
	for loadBalancerBackendAddressPoolIterator.NotDone() {
		loadBalancerBackendAddressPool := loadBalancerBackendAddressPoolIterator.Value()
		// NOTE:
		// Similar to above explanation, work out loadbalancer_id for azurerm datasource impl
		re := regexp.MustCompile(`/backendAddressPools/.*$`)
		loadBalancerID := re.ReplaceAllLiteralString(*loadBalancerBackendAddressPool.ID, "")
		resources = append(resources, terraformutils.NewResource(
			*loadBalancerBackendAddressPool.ID,
			*loadBalancerBackendAddressPool.Name,
			"azurerm_lb_backend_address_pool",
			g.ProviderName,
			map[string]string{
				"loadbalancer_id": loadBalancerID,
			},
			[]string{},
			map[string]interface{}{},
		))
		if err := loadBalancerBackendAddressPoolIterator.Next(); err != nil {
			log.Println(err)
			break
		}
	}

	return resources, nil
}

func (g *LoadBalancerGenerator) listAndAddForLoadBalancers() ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	resourceManagerEndpoint := g.Args["config"].(authentication.Config).CustomResourceManagerEndpoint

	LoadBalancersClient := network.NewLoadBalancersClientWithBaseURI(resourceManagerEndpoint, subscriptionID)
	LoadBalancersClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)

	var (
		loadBalancerIterator network.LoadBalancerListResultIterator
		err                  error
	)

	if rg := g.Args["resource_group"].(string); rg != "" {
		loadBalancerIterator, err = LoadBalancersClient.ListComplete(ctx, rg)
	} else {
		loadBalancerIterator, err = LoadBalancersClient.ListAllComplete(ctx)
	}

	if err != nil {
		return nil, err
	}
	for loadBalancerIterator.NotDone() {
		loadBalancer := loadBalancerIterator.Value()
		resources = append(resources, terraformutils.NewSimpleResource(
			*loadBalancer.ID,
			*loadBalancer.Name,
			"azurerm_lb",
			g.ProviderName,
			[]string{}))

		id, err := ParseAzureResourceID(*loadBalancer.ID)
		if err != nil {
			return nil, err
		}

		probes, err := g.listLoadBalancerProbes(id.ResourceGroup, *loadBalancer.Name)
		if err != nil {
			return nil, err
		}
		resources = append(resources, probes...)

		inboundNatRules, err := g.listInboundNatRules(id.ResourceGroup, *loadBalancer.Name)
		if err != nil {
			return nil, err
		}
		resources = append(resources, inboundNatRules...)

		backendAddressPools, err := g.listLoadBalancerBackendAddressPools(id.ResourceGroup, *loadBalancer.Name)
		if err != nil {
			return nil, err
		}
		resources = append(resources, backendAddressPools...)

		if err := loadBalancerIterator.Next(); err != nil {
			log.Println(err)
			return resources, err
		}
	}

	return resources, nil
}

func (g *LoadBalancerGenerator) InitResources() error {
	functions := []func() ([]terraformutils.Resource, error){
		g.listAndAddForLoadBalancers,
	}

	for _, f := range functions {
		resources, err := f()
		if err != nil {
			return err
		}
		g.Resources = append(g.Resources, resources...)
	}

	return nil
}
