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

package ns1

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	ns1 "gopkg.in/ns1/ns1-go.v2/rest"
	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
	"net/http"
	"time"
)

type ZoneGenerator struct {
	Ns1Service
}

func (g *ZoneGenerator) createZoneRecordResources(client *ns1.Client, zone_name string) error {

	zone, _, err := client.Zones.Get(zone_name)
	if err != nil {
		return err
	}

	for _, record := range zone.Records {
		r, _, err := client.Records.Get(zone_name, record.Domain, record.Type)
		if err != nil {
			return err
		}

		g.Resources = append(g.Resources, terraformutils.NewResource(
			r.ID,
			r.ID,
			"ns1_record",
			"ns1",
			map[string]string{"zone": r.Zone, "domain": r.Domain, "type": r.Type},
			[]string{},
			map[string]interface{}{},
		))

	}

	return nil
}

func (g *ZoneGenerator) createZoneResources(client *ns1.Client, includeZones []string) error {

	var zones []*dns.Zone

	if len(includeZones) > 0 {
		for _, filter := range includeZones {
			var z *dns.Zone
			z, _, err := client.Zones.Get(filter)
			if err != nil {
				return err
			}
			zones = append(zones, z)
		}
	} else {
		var err error
		zones, _, err = client.Zones.List()
		if err != nil {
			return err
		}
	}

	for _, zone := range zones {
		g.Resources = append(g.Resources, terraformutils.NewResource(
			zone.ID,
			zone.Zone,
			"ns1_zone",
			"ns1",
			map[string]string{"zone": zone.Zone},
			[]string{},
			map[string]interface{}{},
		))

		g.createZoneRecordResources(client, zone.Zone)
	}

	return nil
}

func (g *ZoneGenerator) InitResources() error {
	filters := make([]string, 0)
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("zone") {
			filters = append(filters, filter.AcceptableValues...)
		}
	}
	httpClient := &http.Client{Timeout: time.Second * 10}
	client := ns1.NewClient(httpClient, ns1.SetAPIKey(g.Args["api_key"].(string)))

	if err := g.createZoneResources(client, filters); err != nil {
		return err
	}

	return nil
}
