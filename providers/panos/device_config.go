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
)

type DeviceConfigGenerator struct {
	PanosService
}

func (g *DeviceConfigGenerator) createResourcesFromList(o getGeneric, idPrefix, terraformResourceName string) (resources []terraformutils.Resource) {
	l, err := o.i.(getListWithOneArg).GetList(o.params[0])
	if err != nil {
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

func (g *DeviceConfigGenerator) createGeneralSettingsResource(hostname string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		hostname,
		normalizeResourceName(hostname),
		"panos_general_settings",
		"panos",
		[]string{},
	)
}

func (g *DeviceConfigGenerator) createTelemetryResource(ipAddress, hostname string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		ipAddress,
		normalizeResourceName(hostname),
		"panos_telemetry",
		"panos",
		[]string{},
	)
}

func (g *DeviceConfigGenerator) createEmailServerProfileResources() []terraformutils.Resource {
	return g.createResourcesFromList(getGeneric{g.client.Device.EmailServerProfile, []string{g.vsys}},
		g.vsys+":", "panos_email_server_profile",
	)
}

func (g *DeviceConfigGenerator) createHTTPServerProfileResources() []terraformutils.Resource {
	return g.createResourcesFromList(getGeneric{g.client.Device.HttpServerProfile, []string{g.vsys}},
		g.vsys+":", "panos_http_server_profile",
	)
}

func (g *DeviceConfigGenerator) createSNMPTrapServerProfileResources() []terraformutils.Resource {
	return g.createResourcesFromList(getGeneric{g.client.Device.SnmpServerProfile, []string{g.vsys}},
		g.vsys+":", "panos_snmptrap_server_profile",
	)
}

func (g *DeviceConfigGenerator) createSyslogServerProfileResources() []terraformutils.Resource {
	return g.createResourcesFromList(getGeneric{g.client.Device.SyslogServerProfile, []string{g.vsys}},
		g.vsys+":", "panos_syslog_server_profile",
	)
}

func (g *DeviceConfigGenerator) InitResources() error {
	if err := g.Initialize(); err != nil {
		return err
	}

	if g.vsys == "vsys1" {
		g.vsys = "shared"
	}

	generalConfig, err := g.client.Device.GeneralSettings.Get()
	if err != nil {
		return err
	}

	g.Resources = append(g.Resources, g.createGeneralSettingsResource(generalConfig.Hostname))
	g.Resources = append(g.Resources, g.createTelemetryResource(generalConfig.IpAddress, generalConfig.Hostname))
	g.Resources = append(g.Resources, g.createEmailServerProfileResources()...)
	g.Resources = append(g.Resources, g.createHTTPServerProfileResources()...)
	g.Resources = append(g.Resources, g.createSNMPTrapServerProfileResources()...)
	g.Resources = append(g.Resources, g.createSyslogServerProfileResources()...)

	return nil
}
