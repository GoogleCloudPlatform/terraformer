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
	"github.com/PaloAltoNetworks/pango/util"
)

type PanoramaDeviceConfigGenerator struct {
	PanosService
}

func (g *PanoramaDeviceConfigGenerator) createResourcesFromList(o getGeneric, idPrefix string, useIDForResourceName bool, terraformResourceName string) (resources []terraformutils.Resource) {
	var l []string
	var err error

	switch f := o.i.(type) {
	case getListWithoutArg:
		l, err = f.GetList()
	case getListWithTwoArgs:
		l, err = f.GetList(o.params[0], o.params[1])
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

func (g *PanoramaDeviceConfigGenerator) createDeviceGroupResources() []terraformutils.Resource {
	return g.createResourcesFromList(getGeneric{g.client.(*pango.Panorama).Panorama.DeviceGroup, []string{}},
		"", false, "panos_panorama_device_group",
	)
}

func (g *PanoramaDeviceConfigGenerator) createDeviceGroupParentResources() (resources []terraformutils.Resource) {
	p, err := g.client.(*pango.Panorama).Panorama.DeviceGroup.GetParents()
	if err != nil {
		return resources
	}

	for dg, parent := range p {
		if parent != "" {
			resources = append(resources, terraformutils.NewResource(
				dg,
				normalizeResourceName(dg),
				"panos_device_group_parent",
				"panos",
				map[string]string{
					"device_group": dg,
					"parent":       parent,
				},
				[]string{},
				map[string]interface{}{},
			))
		}
	}

	return resources
}

func createAttributes(tmpl, ts, dg string) map[string]string {
	attributes := make(map[string]string)

	if tmpl != "" {
		attributes["template"] = tmpl
	}
	if ts != "" {
		attributes["template_stack"] = ts
	}
	if dg != "" {
		attributes["device_group"] = dg
	}

	return attributes
}

func createServerProfileResources(tmpl, ts, vsys, dg, terraformResourceName string, l []string) (resources []terraformutils.Resource) {
	attributes := createAttributes(tmpl, ts, dg)

	for _, r := range l {
		id := tmpl + ":" + ts + ":" + vsys + ":" + dg + ":" + r
		resources = append(resources, terraformutils.NewResource(
			id,
			normalizeResourceName(id),
			terraformResourceName,
			"panos",
			attributes,
			[]string{},
			map[string]interface{}{},
		))
	}

	return resources
}

func (g *PanoramaDeviceConfigGenerator) createEmailServerProfileResources(tmpl, ts, vsys, dg string) (resources []terraformutils.Resource) {
	l := make([]string, 0)
	var err error

	if tmpl != "" || ts != "" {
		l, err = g.client.(*pango.Panorama).Device.EmailServerProfile.GetList(tmpl, ts, vsys)
	}
	if dg != "" {
		ans := make([]string, 0, 7)
		ans = append(ans, util.DeviceGroupXpathPrefix(dg)...)
		ans = append(ans, []string{"log-settings", "email"}...)
		l, err = g.client.(util.XapiClient).EntryListUsing(g.client.(util.XapiClient).Get, ans)
	}

	if err != nil || len(l) == 0 {
		return resources
	}

	return createServerProfileResources(tmpl, ts, vsys, dg, "panos_panorama_email_server_profile", l)
}

func (g *PanoramaDeviceConfigGenerator) createHTTPServerProfileResources(tmpl, ts, vsys, dg string) (resources []terraformutils.Resource) {
	l := make([]string, 0)
	var err error

	if tmpl != "" || ts != "" {
		l, err = g.client.(*pango.Panorama).Device.HttpServerProfile.GetList(tmpl, ts, vsys)
	}
	if dg != "" {
		ans := make([]string, 0, 7)
		ans = append(ans, util.DeviceGroupXpathPrefix(dg)...)
		ans = append(ans, []string{"log-settings", "http"}...)
		l, err = g.client.(util.XapiClient).EntryListUsing(g.client.(util.XapiClient).Get, ans)
	}

	if err != nil || len(l) == 0 {
		return resources
	}

	return createServerProfileResources(tmpl, ts, vsys, dg, "panos_panorama_http_server_profile", l)
}

func (g *PanoramaDeviceConfigGenerator) createSNMPTrapServerProfileResources(tmpl, ts, vsys, dg string) (resources []terraformutils.Resource) {
	l := make([]string, 0)
	var err error

	if tmpl != "" || ts != "" {
		l, err = g.client.(*pango.Panorama).Device.SnmpServerProfile.GetList(tmpl, ts, vsys)
	}
	if dg != "" {
		ans := make([]string, 0, 7)
		ans = append(ans, util.DeviceGroupXpathPrefix(dg)...)
		ans = append(ans, []string{"log-settings", "snmptrap"}...)
		l, err = g.client.(util.XapiClient).EntryListUsing(g.client.(util.XapiClient).Get, ans)
	}

	if err != nil || len(l) == 0 {
		return resources
	}

	return createServerProfileResources(tmpl, ts, vsys, dg, "panos_panorama_snmptrap_server_profile", l)
}

func (g *PanoramaDeviceConfigGenerator) createSyslogServerProfileResources(tmpl, ts, vsys, dg string) (resources []terraformutils.Resource) {
	l := make([]string, 0)
	var err error

	if tmpl != "" || ts != "" {
		l, err = g.client.(*pango.Panorama).Device.SyslogServerProfile.GetList(tmpl, ts, vsys)
	}
	if dg != "" {
		ans := make([]string, 0, 7)
		ans = append(ans, util.DeviceGroupXpathPrefix(dg)...)
		ans = append(ans, []string{"log-settings", "syslog"}...)
		l, err = g.client.(util.XapiClient).EntryListUsing(g.client.(util.XapiClient).Get, ans)
	}

	if err != nil || len(l) == 0 {
		return resources
	}

	return createServerProfileResources(tmpl, ts, vsys, dg, "panos_panorama_syslog_server_profile", l)
}

func (g *PanoramaDeviceConfigGenerator) createTemplateResources() []terraformutils.Resource {
	return g.createResourcesFromList(getGeneric{g.client.(*pango.Panorama).Panorama.Template, []string{}},
		"", false, "panos_panorama_template",
	)
}

func (g *PanoramaDeviceConfigGenerator) createTemplateStackResources() []terraformutils.Resource {
	return g.createResourcesFromList(getGeneric{g.client.(*pango.Panorama).Panorama.TemplateStack, []string{}},
		"", false, "panos_panorama_template_stack",
	)
}

func (g *PanoramaDeviceConfigGenerator) createTemplateVariableResources(tmpl, ts string) []terraformutils.Resource {
	return g.createResourcesFromList(getGeneric{g.client.(*pango.Panorama).Panorama.TemplateVariable, []string{tmpl, ts}},
		tmpl+":"+ts+":", true, "panos_panorama_template_variable",
	)
}

func (g *PanoramaDeviceConfigGenerator) InitResources() error {
	if err := g.Initialize(); err != nil {
		return err
	}

	g.Resources = append(g.Resources, g.createTemplateStackResources()...)
	g.Resources = append(g.Resources, g.createTemplateResources()...)
	g.Resources = append(g.Resources, g.createDeviceGroupResources()...)
	g.Resources = append(g.Resources, g.createDeviceGroupParentResources()...)

	ts, err := g.client.(*pango.Panorama).Panorama.TemplateStack.GetList()
	if err != nil {
		return err
	}

	for _, v := range ts {
		g.Resources = append(g.Resources, g.createTemplateVariableResources("", v)...)

		vsysList, err := g.client.(*pango.Panorama).Vsys.GetList("", v)
		if err != nil {
			continue
		}

		vsysList = append(vsysList, "shared")

		for _, vsys := range vsysList {
			g.Resources = append(g.Resources, g.createEmailServerProfileResources("", v, vsys, "")...)
			g.Resources = append(g.Resources, g.createHTTPServerProfileResources("", v, vsys, "")...)
			g.Resources = append(g.Resources, g.createSNMPTrapServerProfileResources("", v, vsys, "")...)
			g.Resources = append(g.Resources, g.createSyslogServerProfileResources("", v, vsys, "")...)
		}
	}

	tmpl, err := g.client.(*pango.Panorama).Panorama.Template.GetList()
	if err != nil {
		return err
	}

	for _, v := range tmpl {
		g.Resources = append(g.Resources, g.createTemplateVariableResources(v, "")...)

		vsysList, err := g.client.(*pango.Panorama).Vsys.GetList(v, "")
		if err != nil {
			continue
		}
		if err != nil {
			continue
		}

		vsysList = append(vsysList, "shared")

		for _, vsys := range vsysList {
			g.Resources = append(g.Resources, g.createEmailServerProfileResources(v, "", vsys, "")...)
			g.Resources = append(g.Resources, g.createHTTPServerProfileResources(v, "", vsys, "")...)
			g.Resources = append(g.Resources, g.createSNMPTrapServerProfileResources(v, "", vsys, "")...)
			g.Resources = append(g.Resources, g.createSyslogServerProfileResources(v, "", vsys, "")...)
		}
	}

	dg, err := g.client.(*pango.Panorama).Panorama.DeviceGroup.GetList()
	if err != nil {
		return err
	}

	for _, v := range dg {
		g.Resources = append(g.Resources, g.createEmailServerProfileResources("", "", "", v)...)
		g.Resources = append(g.Resources, g.createHTTPServerProfileResources("", "", "", v)...)
		g.Resources = append(g.Resources, g.createSNMPTrapServerProfileResources("", "", "", v)...)
		g.Resources = append(g.Resources, g.createSyslogServerProfileResources("", "", "", v)...)
	}

	// TODO: Panorama's own profiles are not yet supported by the Terraform provider

	return nil
}
