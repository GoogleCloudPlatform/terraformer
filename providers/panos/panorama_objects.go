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

type PanoramaObjectsGenerator struct {
	PanosService
}

func (g *PanoramaObjectsGenerator) createResourcesFromList(o getGeneric, dg string, terraformResourceName string) (resources []terraformutils.Resource) {
	l, err := o.i.(getListWithOneArg).GetList(o.params[0])
	if err != nil || len(l) == 0 {
		return []terraformutils.Resource{}
	}

	for _, r := range l {
		id := dg + ":" + r
		resources = append(resources, terraformutils.NewResource(
			id,
			normalizeResourceName(id),
			terraformResourceName,
			"panos",
			map[string]string{
				"device_group": dg,
			},
			[]string{},
			map[string]interface{}{},
		))
	}

	return resources
}

func (g *PanoramaObjectsGenerator) createAddressGroupResources(dg string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Objects.AddressGroup, []string{dg}},
		dg, "panos_panorama_address_group",
	)
}

func (g *PanoramaObjectsGenerator) createAdministrativeTagResources(dg string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Objects.Tags, []string{dg}},
		dg, "panos_panorama_administrative_tag",
	)
}

func (g *PanoramaObjectsGenerator) createApplicationGroupResources(dg string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Objects.AppGroup, []string{dg}},
		dg, "panos_panorama_application_group",
	)
}

func (g *PanoramaObjectsGenerator) createApplicationObjectResources(dg string) (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Panorama).Objects.Application.GetList(dg)
	if err != nil {
		return []terraformutils.Resource{}
	}

	for _, r := range l {
		id := dg + ":" + r
		resources = append(resources, terraformutils.NewSimpleResource(
			id,
			normalizeResourceName(id),
			"panos_panorama_application_object",
			"panos",
			[]string{},
		))

		// TODO
		// resources = append(resources, g.createApplicationSignatureResources(dg, r)...)
	}

	return resources
}

func (g *PanoramaObjectsGenerator) createEDLResources(dg string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Objects.Edl, []string{dg}},
		dg, "panos_panorama_edl",
	)
}

func (g *PanoramaObjectsGenerator) createLogForwardingResources(dg string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Objects.LogForwardingProfile, []string{dg}},
		dg, "panos_panorama_log_forwarding_profile",
	)
}

func (g *PanoramaObjectsGenerator) createServiceGroupResources(dg string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Objects.ServiceGroup, []string{dg}},
		dg, "panos_panorama_service_group",
	)
}

func (g *PanoramaObjectsGenerator) createServiceObjectResources(dg string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Objects.Services, []string{dg}},
		dg, "panos_panorama_service_object",
	)
}

func (g *PanoramaObjectsGenerator) createAddressObjectResources(dg string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Objects.Address, []string{dg}},
		dg, "panos_address_object",
	)
}

func (g *PanoramaObjectsGenerator) createAntiSpywareSecurityProfileResources(dg string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Objects.AntiSpywareProfile, []string{dg}},
		dg, "panos_anti_spyware_security_profile",
	)
}

func (g *PanoramaObjectsGenerator) createAntivirusSecurityProfileResources(dg string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Objects.AntivirusProfile, []string{dg}},
		dg, "panos_antivirus_security_profile",
	)
}

func (g *PanoramaObjectsGenerator) createCustomDataPatternObjectResources(dg string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Objects.DataPattern, []string{dg}},
		dg, "panos_custom_data_pattern_object",
	)
}

func (g *PanoramaObjectsGenerator) createDataFilteringSecurityProfileResources(dg string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Objects.DataFilteringProfile, []string{dg}},
		dg, "panos_data_filtering_security_profile",
	)
}

func (g *PanoramaObjectsGenerator) createDOSProtectionProfileResources(dg string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Objects.DosProtectionProfile, []string{dg}},
		dg, "panos_dos_protection_profile",
	)
}

func (g *PanoramaObjectsGenerator) createDynamicUserGroupResources(dg string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Objects.DynamicUserGroup, []string{dg}},
		dg, "panos_dynamic_user_group",
	)
}

func (g *PanoramaObjectsGenerator) createFileBlockingSecurityProfileResources(dg string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Objects.FileBlockingProfile, []string{dg}},
		dg, "panos_file_blocking_security_profile",
	)
}

func (g *PanoramaObjectsGenerator) createURLFilteringSecurityProfileResources(dg string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Objects.UrlFilteringProfile, []string{dg}},
		dg, "panos_url_filtering_security_profile",
	)
}

func (g *PanoramaObjectsGenerator) createVulnerabilitySecurityProfileResources(dg string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Objects.VulnerabilityProfile, []string{dg}},
		dg, "panos_vulnerability_security_profile",
	)
}

func (g *PanoramaObjectsGenerator) createWildfireAnalysisSecurityProfileResources(dg string) []terraformutils.Resource {
	return g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Objects.WildfireAnalysisProfile, []string{dg}},
		dg, "panos_wildfire_analysis_security_profile",
	)
}
func (g *PanoramaObjectsGenerator) InitResources() error {
	if err := g.Initialize(); err != nil {
		return err
	}

	dg, err := g.client.(*pango.Panorama).Panorama.DeviceGroup.GetList()
	if err != nil {
		return err
	}

	for _, v := range dg {
		g.Resources = append(g.Resources, g.createAddressGroupResources(v)...)
		g.Resources = append(g.Resources, g.createAdministrativeTagResources(v)...)
		g.Resources = append(g.Resources, g.createApplicationGroupResources(v)...)
		g.Resources = append(g.Resources, g.createApplicationObjectResources(v)...)
		g.Resources = append(g.Resources, g.createEDLResources(v)...)
		g.Resources = append(g.Resources, g.createLogForwardingResources(v)...)
		g.Resources = append(g.Resources, g.createServiceGroupResources(v)...)
		g.Resources = append(g.Resources, g.createServiceObjectResources(v)...)
		g.Resources = append(g.Resources, g.createAddressObjectResources(v)...)
		g.Resources = append(g.Resources, g.createAntiSpywareSecurityProfileResources(v)...)
		g.Resources = append(g.Resources, g.createAntivirusSecurityProfileResources(v)...)
		g.Resources = append(g.Resources, g.createCustomDataPatternObjectResources(v)...)
		g.Resources = append(g.Resources, g.createDataFilteringSecurityProfileResources(v)...)
		g.Resources = append(g.Resources, g.createDOSProtectionProfileResources(v)...)
		g.Resources = append(g.Resources, g.createDynamicUserGroupResources(v)...)
		g.Resources = append(g.Resources, g.createFileBlockingSecurityProfileResources(v)...)
		g.Resources = append(g.Resources, g.createURLFilteringSecurityProfileResources(v)...)
		g.Resources = append(g.Resources, g.createVulnerabilitySecurityProfileResources(v)...)
		g.Resources = append(g.Resources, g.createWildfireAnalysisSecurityProfileResources(v)...)
	}

	return nil
}

func (g *PanoramaObjectsGenerator) PostConvertHook() error {
	mapAddressObjectIDs := map[string]string{}
	mapApplicationObjectIDs := map[string]string{}
	mapServiceObjectIDs := map[string]string{}

	for _, r := range g.Resources {
		if _, ok := r.Item["name"]; ok {
			if r.InstanceInfo.Type == "panos_address_object" {
				mapAddressObjectIDs[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			}

			if r.InstanceInfo.Type == "panos_panorama_application_object" {
				mapApplicationObjectIDs[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			}

			if r.InstanceInfo.Type == "panos_panorama_service_object" {
				mapServiceObjectIDs[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			}
		}
	}

	for _, r := range g.Resources {
		if r.InstanceInfo.Type == "panos_panorama_address_group" {
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

		if r.InstanceInfo.Type == "panos_panorama_application_group" {
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

		if r.InstanceInfo.Type == "panos_panorama_service_group" {
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
