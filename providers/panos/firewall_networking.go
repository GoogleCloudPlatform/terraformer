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

package panos

import (
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/PaloAltoNetworks/pango"
	"github.com/PaloAltoNetworks/pango/netw/interface/eth"
	"github.com/PaloAltoNetworks/pango/netw/interface/subinterface/layer2"
	"github.com/PaloAltoNetworks/pango/netw/interface/subinterface/layer3"
	"github.com/PaloAltoNetworks/pango/util"
)

type FirewallNetworkingGenerator struct {
	PanosService
}

func (g *FirewallNetworkingGenerator) createResourcesFromList(o getGeneric, idPrefix string, useIDForResourceName bool, terraformResourceName string, checkIfIsVsys bool, checkType string) (resources []terraformutils.Resource) {
	var l []string
	var err error

	switch f := o.i.(type) {
	case getListWithoutArg:
		l, err = f.GetList()
	case getListWithOneArg:
		l, err = f.GetList(o.params[0])
	case getListWithTwoArgs:
		l, err = f.GetList(o.params[0], o.params[1])
	case getListWithThreeArgs:
		l, err = f.GetList(o.params[0], o.params[1], o.params[2])
	default:
		err = fmt.Errorf("not supported")
	}
	if err != nil || len(l) == 0 {
		return []terraformutils.Resource{}
	}

	for _, r := range l {
		if checkIfIsVsys {
			rv, err := g.client.(*pango.Firewall).IsImported(checkType, "", "", g.vsys, r)
			if err != nil || !rv {
				continue
			}
		}

		id := idPrefix + r
		resources = append(resources, terraformutils.NewSimpleResource(
			id,
			normalizeResourceName(func() string {
				if useIDForResourceName {
					return id
				}

				return r
			}()),
			terraformResourceName,
			"panos",
			[]string{},
		))
	}

	return resources
}

func (g *FirewallNetworkingGenerator) createAggregateInterfaceResources() (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Firewall).Network.AggregateInterface.GetList()
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, aggregateInterface := range l {
		rv, err := g.client.(*pango.Firewall).IsImported(util.InterfaceImport, "", "", g.vsys, aggregateInterface)
		if err != nil || !rv {
			continue
		}

		id := g.vsys + ":" + aggregateInterface
		resources = append(resources, terraformutils.NewSimpleResource(
			id,
			normalizeResourceName(aggregateInterface),
			"panos_aggregate_interface",
			"panos",
			[]string{},
		))

		e, err := g.client.(*pango.Firewall).Network.AggregateInterface.Get(aggregateInterface)
		if err != nil {
			continue
		}

		if e.Mode == eth.ModeLayer2 || e.Mode == eth.ModeVirtualWire {
			g.Resources = append(g.Resources, g.createLayer2SubInterfaceResources(layer2.AggregateInterface, aggregateInterface, e.Mode)...)
		}

		if e.Mode == eth.ModeLayer3 {
			g.Resources = append(g.Resources, g.createLayer3SubInterfaceResources(layer3.AggregateInterface, aggregateInterface)...)
		}
	}

	return resources
}

func (g *FirewallNetworkingGenerator) createBFDProfileResources() []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Network.BfdProfile, []string{}},
		"", false, "panos_bfd_profile", false, "",
	)
}

func (g *FirewallNetworkingGenerator) createBGPResource(virtualRouter string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		virtualRouter,
		normalizeResourceName(virtualRouter),
		"panos_bgp",
		"panos",
		[]string{},
	)
}

func (g *FirewallNetworkingGenerator) createBGPAggregateResources(virtualRouter string) (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Firewall).Network.BgpAggregate.GetList(virtualRouter)
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, bgpAggregate := range l {
		id := virtualRouter + ":" + bgpAggregate
		resources = append(resources, terraformutils.NewSimpleResource(
			id,
			normalizeResourceName(id),
			"panos_bgp_aggregate",
			"panos",
			[]string{},
		))

		resources = append(resources, g.createBGPAggregateAdvertiseFilterResources(virtualRouter, bgpAggregate)...)
		resources = append(resources, g.createBGPAggregateSuppressFilterResources(virtualRouter, bgpAggregate)...)
	}

	return resources
}

func (g *FirewallNetworkingGenerator) createBGPAggregateAdvertiseFilterResources(virtualRouter, bgpAggregate string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Network.BgpAggAdvertiseFilter, []string{virtualRouter, bgpAggregate}},
		virtualRouter+":"+bgpAggregate+":", true, "panos_bgp_aggregate_advertise_filter", false, "",
	)
}

func (g *FirewallNetworkingGenerator) createBGPAggregateSuppressFilterResources(virtualRouter, bgpAggregate string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Network.BgpAggSuppressFilter, []string{virtualRouter, bgpAggregate}},
		virtualRouter+":"+bgpAggregate+":", true, "panos_bgp_aggregate_suppress_filter", false, "",
	)
}

// The secret argument will contain "(incorrect)", not the real value
func (g *FirewallNetworkingGenerator) createBGPAuthProfileResources(virtualRouter string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Network.BgpAuthProfile, []string{virtualRouter}},
		virtualRouter+":", true, "panos_bgp_auth_profile", false, "",
	)
}

func (g *FirewallNetworkingGenerator) createBGPConditionalAdvertisementResources(virtualRouter string) (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Firewall).Network.BgpConditionalAdv.GetList(virtualRouter)
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, bgpConditionalAdv := range l {
		id := virtualRouter + ":" + bgpConditionalAdv
		resources = append(resources, terraformutils.NewSimpleResource(
			id,
			normalizeResourceName(id),
			"panos_bgp_conditional_adv",
			"panos",
			[]string{},
		))

		resources = append(resources, g.createBGPConditionalAdvertisementAdvertiseFilterResources(virtualRouter, bgpConditionalAdv)...)
		resources = append(resources, g.createBGPConditionalAdvertisementNonExistFilterResources(virtualRouter, bgpConditionalAdv)...)
	}

	return resources
}

func (g *FirewallNetworkingGenerator) createBGPConditionalAdvertisementAdvertiseFilterResources(virtualRouter, bgpConditionalAdv string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Network.BgpConAdvAdvertiseFilter, []string{virtualRouter, bgpConditionalAdv}},
		virtualRouter+":"+bgpConditionalAdv+":", true, "panos_bgp_conditional_adv_advertise_filter", false, "",
	)
}

func (g *FirewallNetworkingGenerator) createBGPConditionalAdvertisementNonExistFilterResources(virtualRouter, bgpConditionalAdv string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Network.BgpConAdvNonExistFilter, []string{virtualRouter, bgpConditionalAdv}},
		virtualRouter+":"+bgpConditionalAdv+":", true, "panos_bgp_conditional_adv_non_exist_filter", false, "",
	)
}

func (g *FirewallNetworkingGenerator) createBGPDampeningProfileResources(virtualRouter string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Network.BgpDampeningProfile, []string{virtualRouter}},
		virtualRouter+":", true, "panos_bgp_dampening_profile", false, "",
	)
}

func (g *FirewallNetworkingGenerator) createBGPRuleGroupResourcesFromList(o getGeneric, terraformResourceName string) (resources []terraformutils.Resource) {
	l, err := o.i.(getListWithOneArg).GetList(o.params[0])
	if err != nil || len(l) == 0 {
		return []terraformutils.Resource{}
	}

	var positionReference string
	id := o.params[0] + ":" + strconv.Itoa(util.MoveTop) + "::"

	for k, r := range l {
		if k > 0 {
			id = o.params[0] + ":" + strconv.Itoa(util.MoveAfter) + ":" + positionReference + ":"
		}

		id += base64.StdEncoding.EncodeToString([]byte(r))
		positionReference = r

		resources = append(resources, terraformutils.NewSimpleResource(
			id,
			normalizeResourceName(r),
			terraformResourceName,
			"panos",
			[]string{},
		))
	}

	return resources
}

func (g *FirewallNetworkingGenerator) createBGPExportRuleGroupResources(virtualRouter string) []terraformutils.Resource {
	return g.createBGPRuleGroupResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Network.BgpExport, []string{virtualRouter}},
		"panos_bgp_export_rule_group",
	)
}

func (g *FirewallNetworkingGenerator) createBGPImportRuleGroupResources(virtualRouter string) []terraformutils.Resource {
	return g.createBGPRuleGroupResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Network.BgpImport, []string{virtualRouter}},
		"panos_bgp_import_rule_group",
	)
}

func (g *FirewallNetworkingGenerator) createBGPPeerGroupResources(virtualRouter string) (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Firewall).Network.BgpPeerGroup.GetList(virtualRouter)
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, bgpPeerGroup := range l {
		id := virtualRouter + ":" + bgpPeerGroup
		resources = append(resources, terraformutils.NewSimpleResource(
			id,
			normalizeResourceName(id),
			"panos_bgp_peer_group",
			"panos",
			[]string{},
		))

		resources = append(resources, g.createBGPPeerResources(virtualRouter, bgpPeerGroup)...)
	}

	return resources
}

func (g *FirewallNetworkingGenerator) createBGPPeerResources(virtualRouter, bgpPeerGroup string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Network.BgpPeer, []string{virtualRouter, bgpPeerGroup}},
		virtualRouter+":"+bgpPeerGroup+":", true, "panos_bgp_peer", false, "",
	)
}

func (g *FirewallNetworkingGenerator) createBGPRedistResources(virtualRouter string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Network.BgpRedistRule, []string{virtualRouter}},
		virtualRouter+":", true, "panos_bgp_redist_rule", false, "",
	)
}

func (g *FirewallNetworkingGenerator) createEthernetInterfaceResources() (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Firewall).Network.EthernetInterface.GetList()
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, ethernetInterface := range l {
		rv, err := g.client.(*pango.Firewall).IsImported(util.InterfaceImport, "", "", g.vsys, ethernetInterface)
		if err != nil || !rv {
			continue
		}

		id := g.vsys + ":" + ethernetInterface
		resources = append(resources, terraformutils.NewSimpleResource(
			id,
			normalizeResourceName(ethernetInterface),
			"panos_ethernet_interface",
			"panos",
			[]string{},
		))

		e, err := g.client.(*pango.Firewall).Network.EthernetInterface.Get(ethernetInterface)
		if err != nil {
			continue
		}

		if e.Mode == eth.ModeLayer2 || e.Mode == eth.ModeVirtualWire {
			g.Resources = append(g.Resources, g.createLayer2SubInterfaceResources(layer2.EthernetInterface, ethernetInterface, e.Mode)...)
		}

		if e.Mode == eth.ModeLayer3 {
			g.Resources = append(g.Resources, g.createLayer3SubInterfaceResources(layer3.EthernetInterface, ethernetInterface)...)
		}
	}

	return resources
}

func (g *FirewallNetworkingGenerator) createGRETunnelResources() []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Network.GreTunnel, []string{}},
		"", false, "panos_gre_tunnel", false, "",
	)
}

func (g *FirewallNetworkingGenerator) createIKECryptoProfileResources() (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Firewall).Network.IkeCryptoProfile.GetList()
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, ikeCryptoProfile := range l {
		resources = append(resources, terraformutils.NewResource(
			ikeCryptoProfile,
			normalizeResourceName(ikeCryptoProfile),
			"panos_ike_crypto_profile",
			"panos",
			map[string]string{
				"name": ikeCryptoProfile,
			},
			[]string{},
			map[string]interface{}{},
		))
	}

	return resources
}

func (g *FirewallNetworkingGenerator) createIKEGatewayResources() (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Firewall).Network.IkeGateway.GetList()
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, ikeGateway := range l {
		resources = append(resources, terraformutils.NewResource(
			ikeGateway,
			normalizeResourceName(ikeGateway),
			"panos_ike_gateway",
			"panos",
			map[string]string{
				"name": ikeGateway,
			},
			[]string{},
			map[string]interface{}{},
		))
	}

	return resources
}

func (g *FirewallNetworkingGenerator) createIPSECCryptoProfileResources() (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Firewall).Network.IpsecCryptoProfile.GetList()
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, ipsecCryptoProfile := range l {
		resources = append(resources, terraformutils.NewResource(
			ipsecCryptoProfile,
			normalizeResourceName(ipsecCryptoProfile),
			"panos_ipsec_crypto_profile",
			"panos",
			map[string]string{
				"name": ipsecCryptoProfile,
			},
			[]string{},
			map[string]interface{}{},
		))
	}

	return resources
}

func (g *FirewallNetworkingGenerator) createIPSECTunnelProxyIDIPv4Resources(ipsecTunnel string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Network.IpsecTunnelProxyId, []string{ipsecTunnel}},
		ipsecTunnel+":", false, "panos_ipsec_tunnel_proxy_id_ipv4", false, "",
	)
}

func (g *FirewallNetworkingGenerator) createIPSECTunnelResources() (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Firewall).Network.IpsecTunnel.GetList()
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, ipsecTunnel := range l {
		resources = append(resources, terraformutils.NewSimpleResource(
			ipsecTunnel,
			normalizeResourceName(ipsecTunnel),
			"panos_ipsec_tunnel",
			"panos",
			[]string{},
		))

		resources = append(resources, g.createIPSECTunnelProxyIDIPv4Resources(ipsecTunnel)...)
	}

	return resources
}

func (g *FirewallNetworkingGenerator) createLayer2SubInterfaceResources(interfaceType, parentInterface, parentMode string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Network.Layer2Subinterface, []string{interfaceType, parentInterface, parentMode}},
		interfaceType+":"+parentInterface+":"+parentMode+":"+g.vsys+":", false, "panos_layer2_subinterface", true, util.InterfaceImport,
	)
}

func (g *FirewallNetworkingGenerator) createLayer3SubInterfaceResources(interfaceType, parentInterface string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Network.Layer3Subinterface, []string{interfaceType, parentInterface}},
		interfaceType+":"+parentInterface+":"+g.vsys+":", false, "panos_layer3_subinterface", true, util.InterfaceImport,
	)
}

func (g *FirewallNetworkingGenerator) createLoopbackInterfaceResources() []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Network.LoopbackInterface, []string{}},
		g.vsys+":", false, "panos_loopback_interface", true, util.InterfaceImport,
	)
}

func (g *FirewallNetworkingGenerator) createManagementProfileResources() (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Firewall).Network.ManagementProfile.GetList()
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, managementProfile := range l {
		resources = append(resources, terraformutils.NewResource(
			managementProfile,
			normalizeResourceName(managementProfile),
			"panos_management_profile",
			"panos",
			map[string]string{
				"name": managementProfile,
			},
			[]string{},
			map[string]interface{}{},
		))
	}

	return resources
}

func (g *FirewallNetworkingGenerator) createMonitorProfileResources() []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Network.MonitorProfile, []string{}},
		"", false, "panos_monitor_profile", false, "",
	)
}

func (g *FirewallNetworkingGenerator) createRedistributionProfileResources(virtualRouter string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Network.RedistributionProfile, []string{virtualRouter}},
		virtualRouter+":", true, "panos_redistribution_profile_ipv4", false, "",
	)
}

func (g *FirewallNetworkingGenerator) createStaticRouteIpv4Resources(virtualRouter string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Network.StaticRoute, []string{virtualRouter}},
		virtualRouter+":", true, "panos_static_route_ipv4", false, "",
	)
}

func (g *FirewallNetworkingGenerator) createTunnelInterfaceResources() []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Network.TunnelInterface, []string{}},
		g.vsys+":", false, "panos_tunnel_interface", true, util.InterfaceImport,
	)
}

func (g *FirewallNetworkingGenerator) createVirtualRouterResources() (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Firewall).Network.VirtualRouter.GetList()
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, virtualRouter := range l {
		// TODO: doesn't work!!?
		// rv, err := g.client.(*pango.Firewall).IsImported(util.VirtualRouterImport, "", "", g.vsys, virtualRouter)
		// if err != nil || !rv {
		// 	continue
		// }

		id := g.vsys + ":" + virtualRouter
		resources = append(resources, terraformutils.NewSimpleResource(
			id,
			normalizeResourceName(virtualRouter),
			"panos_virtual_router",
			"panos",
			[]string{},
		))

		resources = append(resources, g.createBGPResource(virtualRouter))
		resources = append(resources, g.createBGPAggregateResources(virtualRouter)...)
		resources = append(resources, g.createBGPAuthProfileResources(virtualRouter)...)
		resources = append(resources, g.createBGPConditionalAdvertisementResources(virtualRouter)...)
		resources = append(resources, g.createBGPDampeningProfileResources(virtualRouter)...)
		resources = append(resources, g.createBGPExportRuleGroupResources(virtualRouter)...)
		resources = append(resources, g.createBGPImportRuleGroupResources(virtualRouter)...)
		resources = append(resources, g.createBGPPeerGroupResources(virtualRouter)...)
		resources = append(resources, g.createBGPRedistResources(virtualRouter)...)
		resources = append(resources, g.createRedistributionProfileResources(virtualRouter)...)
		resources = append(resources, g.createStaticRouteIpv4Resources(virtualRouter)...)
	}

	return resources
}

func (g *FirewallNetworkingGenerator) createVlanResources() []terraformutils.Resource {
	// TODO: should activate check with util.VlanImport, but doesn't work?
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Network.Vlan, []string{}},
		g.vsys+":", false, "panos_vlan", false, "",
	)
}

func (g *FirewallNetworkingGenerator) createVlanInterfaceResources() []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Network.VlanInterface, []string{}},
		g.vsys+":", false, "panos_vlan_interface", true, util.InterfaceImport,
	)
}

func (g *FirewallNetworkingGenerator) createZoneResources() []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Network.Zone, []string{g.vsys}},
		g.vsys+":", false, "panos_zone", false, "",
	)
}

func (g *FirewallNetworkingGenerator) InitResources() error {
	if err := g.Initialize(); err != nil {
		return err
	}

	g.Resources = append(g.Resources, g.createAggregateInterfaceResources()...)
	g.Resources = append(g.Resources, g.createBFDProfileResources()...)
	g.Resources = append(g.Resources, g.createEthernetInterfaceResources()...)
	g.Resources = append(g.Resources, g.createGRETunnelResources()...)
	g.Resources = append(g.Resources, g.createIKECryptoProfileResources()...)
	g.Resources = append(g.Resources, g.createIKEGatewayResources()...)
	g.Resources = append(g.Resources, g.createIPSECCryptoProfileResources()...)
	g.Resources = append(g.Resources, g.createIPSECTunnelResources()...)
	g.Resources = append(g.Resources, g.createLoopbackInterfaceResources()...)
	g.Resources = append(g.Resources, g.createManagementProfileResources()...)
	g.Resources = append(g.Resources, g.createMonitorProfileResources()...)
	g.Resources = append(g.Resources, g.createTunnelInterfaceResources()...)
	g.Resources = append(g.Resources, g.createVirtualRouterResources()...)
	g.Resources = append(g.Resources, g.createVlanResources()...)
	g.Resources = append(g.Resources, g.createVlanInterfaceResources()...)
	g.Resources = append(g.Resources, g.createZoneResources()...)

	return nil
}

func (g *FirewallNetworkingGenerator) PostConvertHook() error {
	mapInterfaceNames := map[string]string{}
	mapInterfaceModes := map[string]string{}
	mapIKECryptoProfileNames := map[string]string{}
	mapIKEGatewayNames := map[string]string{}
	mapIPSECCryptoProfileNames := map[string]string{}

	for _, r := range g.Resources {
		if _, ok := r.Item["name"]; ok {
			if r.InstanceInfo.Type == "panos_aggregate_interface" {
				mapInterfaceNames[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
				mapInterfaceModes[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".mode}"
			}

			if r.InstanceInfo.Type == "panos_ethernet_interface" {
				mapInterfaceNames[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
				mapInterfaceModes[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".mode}"
			}

			if r.InstanceInfo.Type == "panos_layer2_subinterface" {
				mapInterfaceNames[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			}

			if r.InstanceInfo.Type == "panos_layer3_subinterface" {
				mapInterfaceNames[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			}

			if r.InstanceInfo.Type == "panos_loopback_interface" {
				mapInterfaceNames[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			}

			if r.InstanceInfo.Type == "panos_tunnel_interface" {
				mapInterfaceNames[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			}

			if r.InstanceInfo.Type == "panos_vlan_interface" {
				mapInterfaceNames[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			}

			if r.InstanceInfo.Type == "panos_ike_crypto_profile" {
				mapIKECryptoProfileNames[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			}

			if r.InstanceInfo.Type == "panos_ike_gateway" {
				mapIKEGatewayNames[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			}

			if r.InstanceInfo.Type == "panos_ipsec_crypto_profile" {
				mapIPSECCryptoProfileNames[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			}
		}
	}

	for _, r := range g.Resources {
		if r.InstanceInfo.Type == "panos_bgp" ||
			r.InstanceInfo.Type == "panos_redistribution_profile_ipv4" ||
			r.InstanceInfo.Type == "panos_static_route_ipv4" {
			if _, ok := r.Item["virtual_router"]; ok {
				r.Item["virtual_router"] = "${panos_virtual_router." + normalizeResourceName(r.Item["virtual_router"].(string)) + ".name}"
			}
		}

		if r.InstanceInfo.Type == "panos_bgp_aggregate" ||
			r.InstanceInfo.Type == "panos_bgp_auth_profile" ||
			r.InstanceInfo.Type == "panos_bgp_conditional_adv" ||
			r.InstanceInfo.Type == "panos_bgp_dampening_profile" ||
			r.InstanceInfo.Type == "panos_bgp_export_rule_group" ||
			r.InstanceInfo.Type == "panos_bgp_import_rule_group" ||
			r.InstanceInfo.Type == "panos_bgp_peer_group" ||
			r.InstanceInfo.Type == "panos_bgp_redist_rule" {
			if _, ok := r.Item["virtual_router"]; ok {
				r.Item["virtual_router"] = "${panos_bgp." + normalizeResourceName(r.Item["virtual_router"].(string)) + ".virtual_router}"
			}
		}

		if r.InstanceInfo.Type == "panos_bgp_aggregate_advertise_filter" ||
			r.InstanceInfo.Type == "panos_bgp_aggregate_suppress_filter" {
			if _, ok := r.Item["virtual_router"]; ok {
				r.Item["virtual_router"] = "${panos_bgp_aggregate." + normalizeResourceName(r.Item["virtual_router"].(string)) + ".virtual_router}"
			}
			if _, ok := r.Item["bgp_aggregate"]; ok {
				r.Item["bgp_aggregate"] = "${panos_bgp_aggregate." + normalizeResourceName(r.Item["bgp_aggregate"].(string)) + ".name}"
			}
		}

		if r.InstanceInfo.Type == "panos_bgp_peer" {
			if _, ok := r.Item["virtual_router"]; ok {
				r.Item["virtual_router"] = "${panos_bgp." + normalizeResourceName(r.Item["virtual_router"].(string)) + ".virtual_router}"
				r.Item["peer_as"] = "${panos_bgp." + normalizeResourceName(r.Item["virtual_router"].(string)) + ".as_number}"
			}
		}

		if r.InstanceInfo.Type == "panos_bgp_conditional_adv_advertise_filter" ||
			r.InstanceInfo.Type == "panos_bgp_conditional_adv_non_exist_filter" {
			if _, ok := r.Item["virtual_router"]; ok {
				r.Item["virtual_router"] = "${panos_bgp." + normalizeResourceName(r.Item["virtual_router"].(string)) + ".virtual_router}"
			}
			if _, ok := r.Item["panos_bgp_conditional_adv"]; ok {
				r.Item["bgp_conditional_adv"] = "${panos_bgp_conditional_adv." + normalizeResourceName(r.Item["panos_bgp_conditional_adv"].(string)) + ".name}"
			}
		}

		if r.InstanceInfo.Type == "panos_gre_tunnel" {
			if mapExists(mapInterfaceNames, r.Item, "interface") {
				r.Item["interface"] = mapInterfaceNames[r.Item["interface"].(string)]
			}
			if mapExists(mapInterfaceNames, r.Item, "tunnel_interface") {
				r.Item["tunnel_interface"] = mapInterfaceNames[r.Item["tunnel_interface"].(string)]
			}
		}

		if r.InstanceInfo.Type == "panos_ike_gateway" {
			if mapExists(mapIKECryptoProfileNames, r.Item, "ikev1_crypto_profile") {
				r.Item["ikev1_crypto_profile"] = mapIKECryptoProfileNames[r.Item["ikev1_crypto_profile"].(string)]
			}
		}

		if r.InstanceInfo.Type == "panos_ipsec_tunnel" {
			if mapExists(mapInterfaceNames, r.Item, "tunnel_interface") {
				r.Item["tunnel_interface"] = mapInterfaceNames[r.Item["tunnel_interface"].(string)]
			}
			if mapExists(mapIKEGatewayNames, r.Item, "ak_ike_gateway") {
				r.Item["ak_ike_gateway"] = mapIKEGatewayNames[r.Item["ak_ike_gateway"].(string)]
			}
			if mapExists(mapIPSECCryptoProfileNames, r.Item, "ak_ipsec_crypto_profile") {
				r.Item["ak_ipsec_crypto_profile"] = mapIPSECCryptoProfileNames[r.Item["ak_ipsec_crypto_profile"].(string)]
			}
		}

		if r.InstanceInfo.Type == "panos_ipsec_tunnel_proxy_id_ipv4" {
			if mapExists(mapInterfaceNames, r.Item, "ipsec_tunnel") {
				r.Item["ipsec_tunnel"] = mapInterfaceNames[r.Item["ipsec_tunnel"].(string)]
			}
		}

		if r.InstanceInfo.Type == "panos_layer2_subinterface" {
			if mapExists(mapInterfaceModes, r.Item, "parent_interface") {
				r.Item["parent_mode"] = mapInterfaceModes[r.Item["parent_interface"].(string)]
			}
		}

		if r.InstanceInfo.Type == "panos_layer2_subinterface" ||
			r.InstanceInfo.Type == "panos_layer3_subinterface" {
			if mapExists(mapInterfaceNames, r.Item, "parent_interface") {
				r.Item["parent_interface"] = mapInterfaceNames[r.Item["parent_interface"].(string)]
			}
		}

		if r.InstanceInfo.Type == "panos_virtual_router" {
			if r.Item["ospfv3_ext_dist"].(string) == "0" {
				r.Item["ospfv3_ext_dist"] = "110"
			}

			if r.Item["ebgp_dist"].(string) == "0" {
				r.Item["ebgp_dist"] = "20"
			}

			if r.Item["rip_dist"].(string) == "0" {
				r.Item["rip_dist"] = "120"
			}

			if r.Item["ibgp_dist"].(string) == "0" {
				r.Item["ibgp_dist"] = "200"
			}

			if r.Item["static_dist"].(string) == "0" {
				r.Item["static_dist"] = "10"
			}

			if r.Item["ospf_int_dist"].(string) == "0" {
				r.Item["ospf_int_dist"] = "30"
			}

			if r.Item["static_ipv6_dist"].(string) == "0" {
				r.Item["static_ipv6_dist"] = "10"
			}

			if r.Item["ospf_ext_dist"].(string) == "0" {
				r.Item["ospf_ext_dist"] = "110"
			}

			if r.Item["ospfv3_int_dist"].(string) == "0" {
				r.Item["ospfv3_int_dist"] = "30"
			}
		}

		if r.InstanceInfo.Type == "panos_virtual_router" ||
			r.InstanceInfo.Type == "panos_zone" {
			if _, ok := r.Item["interfaces"]; ok {
				interfaces := make([]string, len(r.Item["interfaces"].([]interface{})))
				for k, eth := range r.Item["interfaces"].([]interface{}) {
					if name, ok2 := mapInterfaceNames[eth.(string)]; ok2 {
						interfaces[k] = name
						continue
					}
					interfaces[k] = eth.(string)
				}

				r.Item["interfaces"] = interfaces
			}
		}

		if r.InstanceInfo.Type == "panos_vlan" {
			if mapExists(mapInterfaceNames, r.Item, "vlan_interface") {
				r.Item["vlan_interface"] = mapInterfaceNames[r.Item["vlan_interface"].(string)]
			}
		}
	}

	return nil
}
