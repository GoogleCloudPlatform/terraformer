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
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/PaloAltoNetworks/pango"
)

type FirewallObjectsGenerator struct {
	PanosService
}

func (g *FirewallObjectsGenerator) createResourcesFromList(o getGeneric, idPrefix string, terraformResourceName string) (resources []terraformutils.Resource) {
	l, err := o.i.(getListWithOneArg).GetList(o.params[0])
	if err != nil || len(l) == 0 {
		return []terraformutils.Resource{}
	}

	for _, r := range l {
		id := idPrefix + r
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

func (g *FirewallObjectsGenerator) createResourcesFromListWithVsys(o getGeneric, idPrefix string, terraformResourceName string) (resources []terraformutils.Resource) {
	l, err := o.i.(getListWithOneArg).GetList(o.params[0])
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, r := range l {
		id := idPrefix + r
		resources = append(resources, terraformutils.NewResource(
			id,
			normalizeResourceName(r),
			terraformResourceName,
			"panos",
			map[string]string{
				"vsys":         g.vsys,
				"device_group": "shared",
			},
			[]string{},
			map[string]interface{}{},
		))
	}

	return resources
}

func (g *FirewallObjectsGenerator) createAddressGroupResources() []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Objects.AddressGroup, []string{g.vsys}},
		g.vsys+":", "panos_address_group",
	)
}

func (g *FirewallObjectsGenerator) createAdministrativeTagResources() []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Objects.Tags, []string{g.vsys}},
		g.vsys+":", "panos_administrative_tag",
	)
}

func (g *FirewallObjectsGenerator) createApplicationGroupResources() []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Objects.AppGroup, []string{g.vsys}},
		g.vsys+":", "panos_application_group",
	)
}

func (g *FirewallObjectsGenerator) createApplicationObjectResources() (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Firewall).Objects.Application.GetList(g.vsys)
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, r := range l {
		id := g.vsys + ":" + r
		resources = append(resources, terraformutils.NewSimpleResource(
			id,
			normalizeResourceName(r),
			"panos_application_object",
			"panos",
			[]string{},
		))

		// TODO: fix
		// resources = append(resources, g.createApplicationSignatureResources(r)...)
	}

	return resources
}

// func (g *FirewallObjectsGenerator) createApplicationSignatureResources(applicationObject string) []terraformutils.Resource {
// 	return g.createResourcesFromList(
// 		getGeneric{g.client.(*pango.Firewall).Objects.AppSignature, []string{g.vsys, applicationObject}},
// 		g.vsys+":"+applicationObject+":", "panos_application_signature",
// 	)
// }

func (g *FirewallObjectsGenerator) createEDLResources() []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Objects.Edl, []string{g.vsys}},
		g.vsys+":", "panos_edl",
	)
}

func (g *FirewallObjectsGenerator) createLogForwardingResources() []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Objects.LogForwardingProfile, []string{g.vsys}},
		g.vsys+":", "panos_log_forwarding_profile",
	)
}

func (g *FirewallObjectsGenerator) createServiceGroupResources() []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Objects.ServiceGroup, []string{g.vsys}},
		g.vsys+":", "panos_service_group",
	)
}

func (g *FirewallObjectsGenerator) createServiceObjectResources() []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Objects.Services, []string{g.vsys}},
		g.vsys+":", "panos_service_object",
	)
}

func (g *FirewallObjectsGenerator) createAddressObjectResources() []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Firewall).Objects.Address, []string{g.vsys}},
		g.vsys+":", "panos_address_object",
	)
}

func (g *FirewallObjectsGenerator) createAntiSpywareSecurityProfileResources() []terraformutils.Resource {
	return g.createResourcesFromListWithVsys(
		getGeneric{g.client.(*pango.Firewall).Objects.AntiSpywareProfile, []string{g.vsys}},
		g.vsys+":", "panos_anti_spyware_security_profile",
	)
}

func (g *FirewallObjectsGenerator) createAntivirusSecurityProfileResources() []terraformutils.Resource {
	return g.createResourcesFromListWithVsys(
		getGeneric{g.client.(*pango.Firewall).Objects.AntivirusProfile, []string{g.vsys}},
		g.vsys+":", "panos_antivirus_security_profile",
	)
}

func (g *FirewallObjectsGenerator) createCustomDataPatternObjectResources() []terraformutils.Resource {
	return g.createResourcesFromListWithVsys(
		getGeneric{g.client.(*pango.Firewall).Objects.DataPattern, []string{g.vsys}},
		g.vsys+":", "panos_custom_data_pattern_object",
	)
}

func (g *FirewallObjectsGenerator) createDataFilteringSecurityProfileResources() []terraformutils.Resource {
	return g.createResourcesFromListWithVsys(
		getGeneric{g.client.(*pango.Firewall).Objects.DataFilteringProfile, []string{g.vsys}},
		g.vsys+":", "panos_data_filtering_security_profile",
	)
}

func (g *FirewallObjectsGenerator) createDOSProtectionProfileResources() []terraformutils.Resource {
	return g.createResourcesFromListWithVsys(
		getGeneric{g.client.(*pango.Firewall).Objects.DosProtectionProfile, []string{g.vsys}},
		g.vsys+":", "panos_dos_protection_profile",
	)
}

func (g *FirewallObjectsGenerator) createDynamicUserGroupResources() []terraformutils.Resource {
	return g.createResourcesFromListWithVsys(
		getGeneric{g.client.(*pango.Firewall).Objects.DynamicUserGroup, []string{g.vsys}},
		g.vsys+":", "panos_dynamic_user_group",
	)
}

func (g *FirewallObjectsGenerator) createFileBlockingSecurityProfileResources() []terraformutils.Resource {
	return g.createResourcesFromListWithVsys(
		getGeneric{g.client.(*pango.Firewall).Objects.FileBlockingProfile, []string{g.vsys}},
		g.vsys+":", "panos_file_blocking_security_profile",
	)
}

func (g *FirewallObjectsGenerator) createURLFilteringSecurityProfileResources() []terraformutils.Resource {
	return g.createResourcesFromListWithVsys(
		getGeneric{g.client.(*pango.Firewall).Objects.UrlFilteringProfile, []string{g.vsys}},
		g.vsys+":", "panos_url_filtering_security_profile",
	)
}

func (g *FirewallObjectsGenerator) createVulnerabilitySecurityProfileResources() []terraformutils.Resource {
	return g.createResourcesFromListWithVsys(
		getGeneric{g.client.(*pango.Firewall).Objects.VulnerabilityProfile, []string{g.vsys}},
		g.vsys+":", "panos_vulnerability_security_profile",
	)
}

func (g *FirewallObjectsGenerator) createWildfireAnalysisSecurityProfileResources() []terraformutils.Resource {
	return g.createResourcesFromListWithVsys(
		getGeneric{g.client.(*pango.Firewall).Objects.WildfireAnalysisProfile, []string{g.vsys}},
		g.vsys+":", "panos_wildfire_analysis_security_profile",
	)
}

func (g *FirewallObjectsGenerator) InitResources() error {
	if err := g.Initialize(); err != nil {
		return err
	}

	g.Resources = append(g.Resources, g.createAddressGroupResources()...)
	g.Resources = append(g.Resources, g.createAdministrativeTagResources()...)
	g.Resources = append(g.Resources, g.createApplicationGroupResources()...)
	g.Resources = append(g.Resources, g.createApplicationObjectResources()...)
	g.Resources = append(g.Resources, g.createEDLResources()...)
	g.Resources = append(g.Resources, g.createLogForwardingResources()...)
	g.Resources = append(g.Resources, g.createServiceGroupResources()...)
	g.Resources = append(g.Resources, g.createServiceObjectResources()...)

	g.Resources = append(g.Resources, g.createAddressObjectResources()...)
	g.Resources = append(g.Resources, g.createAntiSpywareSecurityProfileResources()...)
	g.Resources = append(g.Resources, g.createAntivirusSecurityProfileResources()...)
	g.Resources = append(g.Resources, g.createCustomDataPatternObjectResources()...)
	g.Resources = append(g.Resources, g.createDataFilteringSecurityProfileResources()...)
	g.Resources = append(g.Resources, g.createDOSProtectionProfileResources()...)
	g.Resources = append(g.Resources, g.createDynamicUserGroupResources()...)
	g.Resources = append(g.Resources, g.createFileBlockingSecurityProfileResources()...)
	g.Resources = append(g.Resources, g.createURLFilteringSecurityProfileResources()...)
	g.Resources = append(g.Resources, g.createVulnerabilitySecurityProfileResources()...)
	g.Resources = append(g.Resources, g.createWildfireAnalysisSecurityProfileResources()...)

	return nil
}

func (g *FirewallObjectsGenerator) PostConvertHook() error {
	mapAddressObjectIDs := map[string]string{}
	mapApplicationObjectIDs := map[string]string{}
	mapServiceObjectIDs := map[string]string{}

	for _, r := range g.Resources {
		if _, ok := r.Item["name"]; ok {
			if r.InstanceInfo.Type == "panos_address_object" {
				mapAddressObjectIDs[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			}

			if r.InstanceInfo.Type == "panos_application_object" {
				mapApplicationObjectIDs[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			}

			if r.InstanceInfo.Type == "panos_service_object" {
				mapServiceObjectIDs[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			}
		}
	}

	for _, r := range g.Resources {
		if r.InstanceInfo.Type == "panos_address_group" {
			if _, ok := r.Item["static_addresses"]; ok {
				staticAddresses := make([]string, len(r.Item["static_addresses"].([]interface{})))
				for k, staticAddress := range r.Item["static_addresses"].([]interface{}) {
					if _, ok2 := mapAddressObjectIDs[staticAddress.(string)]; ok2 {
						staticAddresses[k] = mapAddressObjectIDs[staticAddress.(string)]
						continue
					}
					staticAddresses[k] = staticAddress.(string)
				}

				r.Item["static_addresses"] = staticAddresses
			}
		}

		if r.InstanceInfo.Type == "panos_application_group" {
			if _, ok := r.Item["applications"]; ok {
				applications := make([]string, len(r.Item["applications"].([]interface{})))
				for k, application := range r.Item["applications"].([]interface{}) {
					if _, ok2 := mapApplicationObjectIDs[application.(string)]; ok2 {
						applications[k] = mapApplicationObjectIDs[application.(string)]
						continue
					}
					applications[k] = application.(string)
				}

				r.Item["applications"] = applications
			}
		}

		if r.InstanceInfo.Type == "panos_service_group" {
			if _, ok := r.Item["services"]; ok {
				services := make([]string, len(r.Item["services"].([]interface{})))
				for k, service := range r.Item["services"].([]interface{}) {
					if _, ok2 := mapServiceObjectIDs[service.(string)]; ok2 {
						services[k] = mapServiceObjectIDs[service.(string)]
						continue
					}
					services[k] = service.(string)
				}

				r.Item["services"] = services
			}
		}
	}

	return nil
}
