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
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/PaloAltoNetworks/pango"
	"github.com/PaloAltoNetworks/pango/netw/interface/eth"
	"github.com/PaloAltoNetworks/pango/netw/interface/subinterface/layer2"
	"github.com/PaloAltoNetworks/pango/netw/interface/subinterface/layer3"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/vsys"
)

type PanoramaNetworkingGenerator struct {
	PanosService
}

func (g *PanoramaNetworkingGenerator) createResourcesFromList(
	o getGeneric,
	idPrefix string,
	useIDForResourceName bool,
	terraformResourceName string,
) (resources []terraformutils.Resource) {
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
	case getListWithFourArgs:
		l, err = f.GetList(o.params[0], o.params[1], o.params[2], o.params[3])
	case getListWithFiveArgs:
		l, err = f.GetList(o.params[0], o.params[1], o.params[2], o.params[3], o.params[4])
	default:
		err = fmt.Errorf("not supported")
	}
	if err != nil || len(l) == 0 {
		return []terraformutils.Resource{}
	}

	for _, r := range l {
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

func (g *PanoramaNetworkingGenerator) createAggregateInterfaceResources(tmpl, ts string, v []vsys.Entry) (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Panorama).Network.AggregateInterface.GetList(tmpl, ts)
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, vsys := range v {
		for _, aggregateInterface := range l {
			if !contains(vsys.NetworkImports.Interfaces, aggregateInterface) {
				continue
			}

			rv, err := g.client.(*pango.Panorama).IsImported(util.InterfaceImport, tmpl, ts, vsys.Name, aggregateInterface)
			if err != nil || !rv {
				continue
			}

			id := tmpl + ":" + ts + ":" + vsys.Name + ":" + aggregateInterface
			resources = append(resources, terraformutils.NewSimpleResource(
				id,
				normalizeResourceName(id),
				"panos_panorama_aggregate_interface",
				"panos",
				[]string{},
			))

			e, err := g.client.(*pango.Panorama).Network.AggregateInterface.Get(tmpl, ts, aggregateInterface)
			if err != nil {
				continue
			}

			if e.Mode == eth.ModeLayer2 || e.Mode == eth.ModeVirtualWire {
				g.Resources = append(g.Resources, g.createLayer2SubInterfaceResources(tmpl, ts, vsys.Name, layer2.EthernetInterface, aggregateInterface, e.Mode)...)
			}

			if e.Mode == eth.ModeLayer3 {
				g.Resources = append(g.Resources, g.createLayer3SubInterfaceResources(tmpl, ts, vsys.Name, layer3.EthernetInterface, aggregateInterface)...)
			}
		}
	}

	return resources
}

func (g *PanoramaNetworkingGenerator) createBFDProfileResources(tmpl, ts string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Network.BfdProfile, []string{tmpl, ts}},
		tmpl+":"+ts+":", false, "panos_panorama_bfd_profile",
	)
}

func (g *PanoramaNetworkingGenerator) createBGPResource(tmpl, ts, virtualRouter string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		tmpl+":"+ts+":"+virtualRouter,
		normalizeResourceName(tmpl+":"+ts+":"+virtualRouter),
		"panos_panorama_bgp",
		"panos",
		[]string{},
	)
}

func (g *PanoramaNetworkingGenerator) createBGPAggregateResources(tmpl, ts, virtualRouter string) (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Panorama).Network.BgpAggregate.GetList(tmpl, ts, virtualRouter)
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, bgpAggregate := range l {
		id := tmpl + ":" + ts + ":" + virtualRouter + ":" + bgpAggregate
		resources = append(resources, terraformutils.NewSimpleResource(
			id,
			normalizeResourceName(id),
			"panos_panorama_bgp_aggregate",
			"panos",
			[]string{},
		))

		resources = append(resources, g.createBGPAggregateAdvertiseFilterResources(tmpl, ts, virtualRouter, bgpAggregate)...)
		resources = append(resources, g.createBGPAggregateSuppressFilterResources(tmpl, ts, virtualRouter, bgpAggregate)...)
	}

	return resources
}

func (g *PanoramaNetworkingGenerator) createBGPAggregateAdvertiseFilterResources(tmpl, ts, virtualRouter, bgpAggregate string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Network.BgpAggAdvertiseFilter, []string{tmpl, ts, virtualRouter, bgpAggregate}},
		tmpl+":"+ts+":"+virtualRouter+":"+bgpAggregate+":", true, "panos_panorama_bgp_aggregate_advertise_filter",
	)
}

func (g *PanoramaNetworkingGenerator) createBGPAggregateSuppressFilterResources(tmpl, ts, virtualRouter, bgpAggregate string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Network.BgpAggSuppressFilter, []string{tmpl, ts, virtualRouter, bgpAggregate}},
		tmpl+":"+ts+":"+virtualRouter+":"+bgpAggregate+":", true, "panos_panorama_bgp_aggregate_suppress_filter",
	)
}

// The secret argument will contain "(incorrect)", not the real value
func (g *PanoramaNetworkingGenerator) createBGPAuthProfileResources(tmpl, ts, virtualRouter string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Network.BgpAuthProfile, []string{tmpl, ts, virtualRouter}},
		tmpl+":"+ts+":"+virtualRouter+":", true, "panos_panorama_bgp_auth_profile",
	)
}

func (g *PanoramaNetworkingGenerator) createBGPConditionalAdvertisementResources(tmpl, ts, virtualRouter string) (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Panorama).Network.BgpConditionalAdv.GetList(tmpl, ts, virtualRouter)
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, bgpConditionalAdv := range l {
		id := tmpl + ":" + ts + ":" + virtualRouter + ":" + bgpConditionalAdv
		resources = append(resources, terraformutils.NewSimpleResource(
			id,
			normalizeResourceName(id),
			"panos_panorama_bgp_conditional_adv",
			"panos",
			[]string{},
		))

		resources = append(resources, g.createBGPConditionalAdvertisementAdvertiseFilterResources(tmpl, ts, virtualRouter, bgpConditionalAdv)...)
		resources = append(resources, g.createBGPConditionalAdvertisementNonExistFilterResources(tmpl, ts, virtualRouter, bgpConditionalAdv)...)
	}

	return resources
}

func (g *PanoramaNetworkingGenerator) createBGPConditionalAdvertisementAdvertiseFilterResources(tmpl, ts, virtualRouter, bgpConditionalAdv string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Network.BgpConAdvAdvertiseFilter, []string{tmpl, ts, virtualRouter, bgpConditionalAdv}},
		tmpl+":"+ts+":"+virtualRouter+":"+bgpConditionalAdv+":", true, "panos_panorama_bgp_conditional_adv_advertise_filter",
	)
}

func (g *PanoramaNetworkingGenerator) createBGPConditionalAdvertisementNonExistFilterResources(tmpl, ts, virtualRouter, bgpConditionalAdv string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Network.BgpConAdvNonExistFilter, []string{tmpl, ts, virtualRouter, bgpConditionalAdv}},
		tmpl+":"+ts+":"+virtualRouter+":"+bgpConditionalAdv+":", true, "panos_panorama_bgp_conditional_adv_non_exist_filter",
	)
}

func (g *PanoramaNetworkingGenerator) createBGPDampeningProfileResources(tmpl, ts, virtualRouter string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Network.BgpDampeningProfile, []string{tmpl, ts, virtualRouter}},
		tmpl+":"+ts+":"+virtualRouter+":", true, "panos_panorama_bgp_dampening_profile",
	)
}

func (g *PanoramaNetworkingGenerator) createBGPExportRuleGroupResources(tmpl, ts, virtualRouter string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Network.BgpExport, []string{tmpl, ts, virtualRouter}},
		tmpl+":"+ts+":"+virtualRouter+":", true, "panos_panorama_bgp_export_rule_group",
	)
}

func (g *PanoramaNetworkingGenerator) createBGPImportRuleGroupResources(tmpl, ts, virtualRouter string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Network.BgpImport, []string{tmpl, ts, virtualRouter}},
		tmpl+":"+ts+":"+virtualRouter+":", true, "panos_panorama_bgp_import_rule_group",
	)
}

func (g *PanoramaNetworkingGenerator) createBGPPeerGroupResources(tmpl, ts, virtualRouter string) (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Panorama).Network.BgpPeerGroup.GetList(tmpl, ts, virtualRouter)
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, bgpPeerGroup := range l {
		id := tmpl + ":" + ts + ":" + virtualRouter + ":" + bgpPeerGroup
		resources = append(resources, terraformutils.NewSimpleResource(
			id,
			normalizeResourceName(id),
			"panos_panorama_bgp_peer_group",
			"panos",
			[]string{},
		))

		resources = append(resources, g.createBGPPeerResources(tmpl, ts, virtualRouter, bgpPeerGroup)...)
	}

	return resources
}

func (g *PanoramaNetworkingGenerator) createBGPPeerResources(tmpl, ts, virtualRouter, bgpPeerGroup string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Network.BgpPeer, []string{tmpl, ts, virtualRouter, bgpPeerGroup}},
		tmpl+":"+ts+":"+virtualRouter+":"+bgpPeerGroup+":", true, "panos_panorama_bgp_peer",
	)
}

func (g *PanoramaNetworkingGenerator) createBGPRedistResources(tmpl, ts, virtualRouter string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Network.BgpRedistRule, []string{tmpl, ts, virtualRouter}},
		tmpl+":"+ts+":"+virtualRouter+":", true, "panos_panorama_bgp_redist_rule",
	)
}

func (g *PanoramaNetworkingGenerator) createEthernetInterfaceResources(tmpl, ts string, v []vsys.Entry) (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Panorama).Network.EthernetInterface.GetList(tmpl, ts)
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, vsys := range v {
		for _, ethernetInterface := range l {
			if !contains(vsys.NetworkImports.Interfaces, ethernetInterface) {
				continue
			}

			rv, err := g.client.(*pango.Panorama).IsImported(util.InterfaceImport, tmpl, ts, vsys.Name, ethernetInterface)
			if err != nil || !rv {
				continue
			}

			id := tmpl + ":" + ts + ":" + vsys.Name + ":" + ethernetInterface
			resources = append(resources, terraformutils.NewSimpleResource(
				id,
				normalizeResourceName(id),
				"panos_panorama_ethernet_interface",
				"panos",
				[]string{},
			))

			e, err := g.client.(*pango.Panorama).Network.EthernetInterface.Get(tmpl, ts, ethernetInterface)
			if err != nil {
				continue
			}

			if e.Mode == eth.ModeLayer2 || e.Mode == eth.ModeVirtualWire {
				g.Resources = append(g.Resources, g.createLayer2SubInterfaceResources(tmpl, ts, vsys.Name, layer2.EthernetInterface, ethernetInterface, e.Mode)...)
			}

			if e.Mode == eth.ModeLayer3 {
				g.Resources = append(g.Resources, g.createLayer3SubInterfaceResources(tmpl, ts, vsys.Name, layer3.EthernetInterface, ethernetInterface)...)
			}
		}
	}

	return resources
}

func (g *PanoramaNetworkingGenerator) createGRETunnelResources(tmpl, ts string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Network.GreTunnel, []string{tmpl, ts}},
		tmpl+":"+ts+":", false, "panos_panorama_gre_tunnel",
	)
}

func (g *PanoramaNetworkingGenerator) createIKECryptoProfileResources(tmpl, ts string) (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Panorama).Network.IkeCryptoProfile.GetList(tmpl, ts)
	if err != nil {
		return []terraformutils.Resource{}
	}

	idPrefix := tmpl + ":" + ts + ":"
	for _, ikeCryptoProfile := range l {
		id := idPrefix + ikeCryptoProfile
		resources = append(resources, terraformutils.NewResource(
			id,
			normalizeResourceName(id),
			"panos_panorama_ike_crypto_profile",
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

func (g *PanoramaNetworkingGenerator) createIKEGatewayResources(tmpl, ts string) (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Panorama).Network.IkeGateway.GetList(tmpl, ts)
	if err != nil {
		return []terraformutils.Resource{}
	}

	idPrefix := tmpl + ":" + ts + ":"
	for _, ikeGateway := range l {
		id := idPrefix + ikeGateway
		resources = append(resources, terraformutils.NewResource(
			id,
			normalizeResourceName(id),
			"panos_panorama_ike_gateway",
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

func (g *PanoramaNetworkingGenerator) createIPSECCryptoProfileResources(tmpl, ts string) (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Panorama).Network.IpsecCryptoProfile.GetList(tmpl, ts)
	if err != nil {
		return []terraformutils.Resource{}
	}

	idPrefix := tmpl + ":" + ts + ":"
	for _, ipsecCryptoProfile := range l {
		id := idPrefix + ipsecCryptoProfile
		resources = append(resources, terraformutils.NewResource(
			id,
			normalizeResourceName(id),
			"panos_panorama_ipsec_crypto_profile",
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

func (g *PanoramaNetworkingGenerator) createIPSECTunnelProxyIDIPv4Resources(tmpl, ts, ipsecTunnel string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Network.IpsecTunnelProxyId, []string{tmpl, ts, ipsecTunnel}},
		tmpl+":"+ts+":"+ipsecTunnel+":", true, "panos_panorama_ipsec_tunnel_proxy_id_ipv4",
	)
}

func (g *PanoramaNetworkingGenerator) createIPSECTunnelResources(tmpl, ts string) (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Panorama).Network.IpsecTunnel.GetList(tmpl, ts)
	if err != nil {
		return []terraformutils.Resource{}
	}

	idPrefix := tmpl + "::"
	for _, ipsecTunnel := range l {
		id := idPrefix + ipsecTunnel
		resources = append(resources, terraformutils.NewSimpleResource(
			id,
			normalizeResourceName(id),
			"panos_panorama_ipsec_tunnel",
			"panos",
			[]string{},
		))

		resources = append(resources, g.createIPSECTunnelProxyIDIPv4Resources(tmpl, ts, ipsecTunnel)...)
	}

	return resources
}

func (g *PanoramaNetworkingGenerator) createLayer2SubInterfaceResources(tmpl, ts, vsys, interfaceType, parentInterface, parentMode string) []terraformutils.Resource {
	// TO FIX: check disabled!
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Network.Layer2Subinterface, []string{tmpl, ts, interfaceType, parentInterface, parentMode}},
		tmpl+":"+ts+":"+interfaceType+":"+parentInterface+":"+parentMode+":"+vsys+":", true, "panos_panorama_layer2_subinterface",
	)
}

func (g *PanoramaNetworkingGenerator) createLayer3SubInterfaceResources(tmpl, ts, vsys, interfaceType, parentInterface string) []terraformutils.Resource {
	// TO FIX: check disabled!
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Network.Layer3Subinterface, []string{tmpl, ts, interfaceType, parentInterface}},
		tmpl+":"+ts+":"+interfaceType+":"+parentInterface+":"+vsys+":", true, "panos_panorama_layer3_subinterface",
	)
}

func (g *PanoramaNetworkingGenerator) createLoopbackInterfaceResources(tmpl, ts string, v []vsys.Entry) (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Panorama).Network.LoopbackInterface.GetList(tmpl, ts)
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, vsys := range v {
		for _, loopbackInterface := range l {
			if !contains(vsys.NetworkImports.Interfaces, loopbackInterface) {
				continue
			}

			rv, err := g.client.(*pango.Panorama).IsImported(util.InterfaceImport, tmpl, ts, vsys.Name, loopbackInterface)
			if err != nil || !rv {
				continue
			}

			id := tmpl + ":" + ts + ":" + vsys.Name + ":" + loopbackInterface
			resources = append(resources, terraformutils.NewSimpleResource(
				id,
				normalizeResourceName(id),
				"panos_panorama_loopback_interface",
				"panos",
				[]string{},
			))
		}
	}

	return resources
}

func (g *PanoramaNetworkingGenerator) createManagementProfileResources(tmpl, ts string) (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Panorama).Network.ManagementProfile.GetList(tmpl, ts)
	if err != nil {
		return []terraformutils.Resource{}
	}

	idPrefix := tmpl + ":" + ts + ":"
	for _, managementProfile := range l {
		id := idPrefix + managementProfile
		resources = append(resources, terraformutils.NewResource(
			id,
			normalizeResourceName(id),
			"panos_panorama_management_profile",
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

func (g *PanoramaNetworkingGenerator) createMonitorProfileResources(tmpl, ts string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Network.MonitorProfile, []string{tmpl, ts}},
		tmpl+":"+ts+":", true, "panos_panorama_monitor_profile",
	)
}

func (g *PanoramaNetworkingGenerator) createRedistributionProfileResources(tmpl, ts, virtualRouter string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Network.RedistributionProfile, []string{tmpl, ts, virtualRouter}},
		tmpl+":"+ts+":"+virtualRouter+":", true, "panos_panorama_redistribution_profile_ipv4",
	)
}

func (g *PanoramaNetworkingGenerator) createStaticRouteIpv4Resources(tmpl, ts, virtualRouter string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Network.StaticRoute, []string{tmpl, ts, virtualRouter}},
		tmpl+":"+ts+":"+virtualRouter+":", true, "panos_panorama_static_route_ipv4",
	)
}

func (g *PanoramaNetworkingGenerator) createTunnelInterfaceResources(tmpl, ts string, v []vsys.Entry) (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Panorama).Network.TunnelInterface.GetList(tmpl, ts)
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, vsys := range v {
		for _, tunnelInterface := range l {
			if !contains(vsys.NetworkImports.Interfaces, tunnelInterface) {
				continue
			}

			rv, err := g.client.(*pango.Panorama).IsImported(util.InterfaceImport, tmpl, ts, vsys.Name, tunnelInterface)
			if err != nil || !rv {
				continue
			}

			id := tmpl + ":" + ts + ":" + vsys.Name + ":" + tunnelInterface
			resources = append(resources, terraformutils.NewSimpleResource(
				id,
				normalizeResourceName(id),
				"panos_panorama_tunnel_interface",
				"panos",
				[]string{},
			))
		}
	}

	return resources
}

func (g *PanoramaNetworkingGenerator) createVirtualRouterResources(tmpl, ts string, v []vsys.Entry) (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Panorama).Network.VirtualRouter.GetList(tmpl, ts)
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, vsys := range v {
		for _, virtualRouter := range l {
			if !contains(vsys.NetworkImports.VirtualRouters, virtualRouter) {
				continue
			}

			// TODO: doesn't work!!?
			// rv, err := g.client.(*pango.Panorama).IsImported(util.InterfaceImport, tmpl, ts, vsys.Name, virtualRouter)
			// if err != nil || !rv {
			// 	continue
			// }

			id := tmpl + ":" + ts + ":" + vsys.Name + ":" + virtualRouter
			resources = append(resources, terraformutils.NewSimpleResource(
				id,
				normalizeResourceName(id),
				"panos_panorama_virtual_router",
				"panos",
				[]string{},
			))

			resources = append(resources, g.createBGPResource(tmpl, ts, virtualRouter))
			resources = append(resources, g.createBGPAggregateResources(tmpl, ts, virtualRouter)...)
			resources = append(resources, g.createBGPAuthProfileResources(tmpl, ts, virtualRouter)...)
			resources = append(resources, g.createBGPConditionalAdvertisementResources(tmpl, ts, virtualRouter)...)
			resources = append(resources, g.createBGPDampeningProfileResources(tmpl, ts, virtualRouter)...)
			resources = append(resources, g.createBGPExportRuleGroupResources(tmpl, ts, virtualRouter)...)
			resources = append(resources, g.createBGPImportRuleGroupResources(tmpl, ts, virtualRouter)...)
			resources = append(resources, g.createBGPPeerGroupResources(tmpl, ts, virtualRouter)...)
			resources = append(resources, g.createBGPRedistResources(tmpl, ts, virtualRouter)...)
			resources = append(resources, g.createRedistributionProfileResources(tmpl, ts, virtualRouter)...)
			resources = append(resources, g.createStaticRouteIpv4Resources(tmpl, ts, virtualRouter)...)
		}
	}

	return resources
}

// FIX: get VLANs in Vsys = None
func (g *PanoramaNetworkingGenerator) createVlanResources(tmpl, ts string, v []vsys.Entry) (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Panorama).Network.Vlan.GetList(tmpl, ts)
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, vsys := range v {
		for _, vlan := range l {
			if !contains(vsys.NetworkImports.Vlans, vlan) {
				continue
			}

			rv, err := g.client.(*pango.Panorama).IsImported(util.VlanImport, tmpl, ts, vsys.Name, vlan)
			if err != nil || !rv {
				continue
			}

			id := tmpl + ":" + ts + ":" + vsys.Name + ":" + vlan
			resources = append(resources, terraformutils.NewSimpleResource(
				id,
				normalizeResourceName(id),
				"panos_panorama_vlan",
				"panos",
				[]string{},
			))
		}
	}

	return resources
}

func (g *PanoramaNetworkingGenerator) createVlanInterfaceResources(tmpl, ts string, v []vsys.Entry) (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Panorama).Network.VlanInterface.GetList(tmpl, ts)
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, vsys := range v {
		for _, vlanInterface := range l {
			if !contains(vsys.NetworkImports.Interfaces, vlanInterface) {
				continue
			}

			rv, err := g.client.(*pango.Panorama).IsImported(util.InterfaceImport, tmpl, ts, vsys.Name, vlanInterface)
			if err != nil || !rv {
				continue
			}

			id := tmpl + ":" + ts + ":" + vsys.Name + ":" + vlanInterface
			resources = append(resources, terraformutils.NewSimpleResource(
				id,
				normalizeResourceName(id),
				"panos_panorama_vlan_interface",
				"panos",
				[]string{},
			))
		}
	}

	return resources
}

func (g *PanoramaNetworkingGenerator) createZoneResources(tmpl, ts string, v []vsys.Entry) (resources []terraformutils.Resource) {
	for _, vsys := range v {
		l, err := g.client.(*pango.Panorama).Network.Zone.GetList(tmpl, ts, vsys.Name)
		if err != nil {
			return []terraformutils.Resource{}
		}

		for _, zone := range l {
			id := tmpl + ":" + ts + ":" + vsys.Name + ":" + zone
			resources = append(resources, terraformutils.NewSimpleResource(
				id,
				normalizeResourceName(id),
				"panos_panorama_zone",
				"panos",
				[]string{},
			))
		}
	}

	return resources
}

func (g *PanoramaNetworkingGenerator) InitResources() error {
	if err := g.Initialize(); err != nil {
		return err
	}

	ts, err := g.client.(*pango.Panorama).Panorama.TemplateStack.GetList()
	if err != nil {
		return err
	}

	for _, v := range ts {
		g.Resources = append(g.Resources, g.createBFDProfileResources("", v)...)
		g.Resources = append(g.Resources, g.createIKECryptoProfileResources("", v)...)
		g.Resources = append(g.Resources, g.createIKEGatewayResources("", v)...)
		g.Resources = append(g.Resources, g.createIPSECCryptoProfileResources("", v)...)
		g.Resources = append(g.Resources, g.createManagementProfileResources("", v)...)
		g.Resources = append(g.Resources, g.createMonitorProfileResources("", v)...)
	}

	tmpl, err := g.client.(*pango.Panorama).Panorama.Template.GetList()
	if err != nil {
		return err
	}

	for _, v := range tmpl {
		vsysAll, err := g.client.(*pango.Panorama).Vsys.GetAll(v, "")
		if err != nil {
			return err
		}

		g.Resources = append(g.Resources, g.createAggregateInterfaceResources(v, "", vsysAll)...)
		g.Resources = append(g.Resources, g.createBFDProfileResources(v, "")...)
		g.Resources = append(g.Resources, g.createEthernetInterfaceResources(v, "", vsysAll)...)
		g.Resources = append(g.Resources, g.createGRETunnelResources(v, "")...)
		g.Resources = append(g.Resources, g.createIKECryptoProfileResources(v, "")...)
		g.Resources = append(g.Resources, g.createIKEGatewayResources(v, "")...)
		g.Resources = append(g.Resources, g.createIPSECCryptoProfileResources(v, "")...)
		g.Resources = append(g.Resources, g.createIPSECTunnelResources(v, "")...)
		g.Resources = append(g.Resources, g.createLoopbackInterfaceResources(v, "", vsysAll)...)
		g.Resources = append(g.Resources, g.createManagementProfileResources(v, "")...)
		g.Resources = append(g.Resources, g.createMonitorProfileResources(v, "")...)
		g.Resources = append(g.Resources, g.createTunnelInterfaceResources(v, "", vsysAll)...)
		g.Resources = append(g.Resources, g.createVirtualRouterResources(v, "", vsysAll)...)
		g.Resources = append(g.Resources, g.createVlanResources(v, "", vsysAll)...)
		g.Resources = append(g.Resources, g.createVlanInterfaceResources(v, "", vsysAll)...)
		g.Resources = append(g.Resources, g.createZoneResources(v, "", vsysAll)...)
	}

	return nil
}

func (g *PanoramaNetworkingGenerator) PostConvertHook() error {
	mapInterfaceNames := map[string]string{}
	mapInterfaceModes := map[string]string{}
	mapIKECryptoProfileNames := map[string]string{}
	mapIKEGatewayNames := map[string]string{}
	mapIPSECCryptoProfileNames := map[string]string{}

	for _, r := range g.Resources {
		if _, ok := r.Item["name"]; ok {
			if r.InstanceInfo.Type == "panos_panorama_aggregate_interface" {
				mapInterfaceNames[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
				mapInterfaceModes[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".mode}"
			}

			if r.InstanceInfo.Type == "panos_panorama_ethernet_interface" {
				mapInterfaceNames[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
				mapInterfaceModes[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".mode}"
			}

			if r.InstanceInfo.Type == "panos_panorama_layer2_subinterface" {
				mapInterfaceNames[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			}

			if r.InstanceInfo.Type == "panos_panorama_layer3_subinterface" {
				mapInterfaceNames[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			}

			if r.InstanceInfo.Type == "panos_panorama_loopback_interface" {
				mapInterfaceNames[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			}

			if r.InstanceInfo.Type == "panos_panorama_tunnel_interface" {
				mapInterfaceNames[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			}

			if r.InstanceInfo.Type == "panos_panorama_vlan_interface" {
				mapInterfaceNames[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			}

			if r.InstanceInfo.Type == "panos_panorama_ike_crypto_profile" {
				mapIKECryptoProfileNames[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			}

			if r.InstanceInfo.Type == "panos_panorama_ike_gateway" {
				mapIKEGatewayNames[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			}

			if r.InstanceInfo.Type == "panos_panorama_ipsec_crypto_profile" {
				mapIPSECCryptoProfileNames[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			}
		}
	}

	for _, r := range g.Resources {
		if r.InstanceInfo.Type == "panos_panorama_bgp" ||
			r.InstanceInfo.Type == "panos_panorama_redistribution_profile_ipv4" ||
			r.InstanceInfo.Type == "panos_panorama_static_route_ipv4" {
			if _, ok := r.Item["virtual_router"]; ok {
				if r.Item["virtual_router"].(string) != "default" {
					r.Item["virtual_router"] = "${panos_panorama_virtual_router." + normalizeResourceName(r.Item["virtual_router"].(string)) + ".name}"
				}
			}
		}

		if r.InstanceInfo.Type == "panos_panorama_bgp_aggregate" ||
			r.InstanceInfo.Type == "panos_panorama_bgp_auth_profile" ||
			r.InstanceInfo.Type == "panos_panorama_bgp_conditional_adv" ||
			r.InstanceInfo.Type == "panos_panorama_bgp_dampening_profile" ||
			r.InstanceInfo.Type == "panos_panorama_bgp_export_rule_group" ||
			r.InstanceInfo.Type == "panos_panorama_bgp_import_rule_group" ||
			r.InstanceInfo.Type == "panos_panorama_bgp_peer_group" ||
			r.InstanceInfo.Type == "panos_panorama_bgp_redist_rule" {
			if _, ok := r.Item["virtual_router"]; ok {
				if r.Item["virtual_router"].(string) != "default" {
					r.Item["virtual_router"] = "${panos_panorama_bgp." + normalizeResourceName(r.Item["virtual_router"].(string)) + ".virtual_router}"
				}
			}
		}

		if r.InstanceInfo.Type == "panos_panorama_bgp_aggregate_advertise_filter" ||
			r.InstanceInfo.Type == "panos_panorama_bgp_aggregate_suppress_filter" {
			if _, ok := r.Item["virtual_router"]; ok {
				if r.Item["virtual_router"].(string) != "default" {
					r.Item["virtual_router"] = "${panos_panorama_bgp_aggregate." + normalizeResourceName(r.Item["virtual_router"].(string)) + ".virtual_router}"
				}
			}
			if _, ok := r.Item["bgp_aggregate"]; ok {
				r.Item["bgp_aggregate"] = "${panos_panorama_bgp_aggregate." + normalizeResourceName(r.Item["bgp_aggregate"].(string)) + ".name}"
			}
		}

		if r.InstanceInfo.Type == "panos_panorama_bgp_peer" {
			if _, ok := r.Item["virtual_router"]; ok {
				if r.Item["virtual_router"].(string) != "default" {
					r.Item["virtual_router"] = "${panos_panorama_bgp." + normalizeResourceName(r.Item["virtual_router"].(string)) + ".virtual_router}"
					r.Item["peer_as"] = "${panos_panorama_bgp." + normalizeResourceName(r.Item["virtual_router"].(string)) + ".as_number}"
				}
			}
			if _, ok := r.Item["panos_bgp_peer_group"]; ok {
				r.Item["bgp_peer_group"] = "${panos_panorama_bgp_peer_group." + normalizeResourceName(r.Item["panos_bgp_peer_group"].(string)) + ".name}"
			}
		}

		if r.InstanceInfo.Type == "panos_panorama_bgp_conditional_adv_advertise_filter" ||
			r.InstanceInfo.Type == "panos_panorama_bgp_conditional_adv_non_exist_filter" {
			if _, ok := r.Item["virtual_router"]; ok {
				if r.Item["virtual_router"].(string) != "default" {
					r.Item["virtual_router"] = "${panos_panorama_bgp." + normalizeResourceName(r.Item["virtual_router"].(string)) + ".virtual_router}"
				}
			}
			if _, ok := r.Item["panos_bgp_conditional_adv"]; ok {
				r.Item["bgp_conditional_adv"] = "${panos_panorama_bgp_conditional_adv." + normalizeResourceName(r.Item["panos_bgp_conditional_adv"].(string)) + ".name}"
			}
		}

		if r.InstanceInfo.Type == "panos_panorama_gre_tunnel" {
			if mapExists(mapInterfaceNames, r.Item, "interface") {
				r.Item["interface"] = mapInterfaceNames[r.Item["interface"].(string)]
			}
			if mapExists(mapInterfaceNames, r.Item, "tunnel_interface") {
				r.Item["tunnel_interface"] = mapInterfaceNames[r.Item["tunnel_interface"].(string)]
			}
		}

		if r.InstanceInfo.Type == "panos_panorama_ike_gateway" {
			if mapExists(mapIKECryptoProfileNames, r.Item, "ikev1_crypto_profile") {
				r.Item["ikev1_crypto_profile"] = mapIKECryptoProfileNames[r.Item["ikev1_crypto_profile"].(string)]
			}
		}

		if r.InstanceInfo.Type == "panos_panorama_ipsec_tunnel" {
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

		if r.InstanceInfo.Type == "panos_panorama_ipsec_tunnel_proxy_id_ipv4" {
			if mapExists(mapInterfaceNames, r.Item, "tunnel_interface") {
				r.Item["tunnel_interface"] = mapInterfaceNames[r.Item["tunnel_interface"].(string)]
			}
		}

		if r.InstanceInfo.Type == "panos_panorama_layer2_subinterface" {
			if mapExists(mapInterfaceModes, r.Item, "parent_interface") {
				r.Item["parent_mode"] = mapInterfaceModes[r.Item["parent_interface"].(string)]
			}
		}

		if r.InstanceInfo.Type == "panos_panorama_layer2_subinterface" ||
			r.InstanceInfo.Type == "panos_panorama_layer3_subinterface" {
			if mapExists(mapInterfaceNames, r.Item, "parent_interface") {
				r.Item["parent_interface"] = mapInterfaceNames[r.Item["parent_interface"].(string)]
			}
		}

		if r.InstanceInfo.Type == "panos_panorama_virtual_router" {
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

		if r.InstanceInfo.Type == "panos_panorama_virtual_router" ||
			r.InstanceInfo.Type == "panos_panorama_zone" {
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

		if r.InstanceInfo.Type == "panos_panorama_vlan" {
			if mapExists(mapInterfaceNames, r.Item, "vlan_interface") {
				r.Item["vlan_interface"] = mapInterfaceNames[r.Item["vlan_interface"].(string)]
			}
		}
	}

	return nil
}
