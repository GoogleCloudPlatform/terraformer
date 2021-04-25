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

package linode

import (
	"context"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/linode/linodego"
)

type DomainGenerator struct {
	LinodeService
}

func (g *DomainGenerator) loadDomains(client linodego.Client) ([]linodego.Domain, error) {
	domainList, err := client.ListDomains(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	for _, domain := range domainList {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			strconv.Itoa(domain.ID),
			strconv.Itoa(domain.ID),
			"linode_domain",
			"linode",
			[]string{}))
	}
	return domainList, nil
}

func (g *DomainGenerator) loadDomainRecords(client linodego.Client, domainID int) error {
	domainRecordList, err := client.ListDomainRecords(context.Background(), domainID, nil)
	if err != nil {
		return err
	}
	for _, domainRecord := range domainRecordList {
		g.Resources = append(g.Resources, terraformutils.NewResource(
			strconv.Itoa(domainRecord.ID),
			strconv.Itoa(domainRecord.ID),
			"linode_domain_record",
			"linode",
			map[string]string{"domain_id": strconv.Itoa(domainID)},
			[]string{},
			map[string]interface{}{}))
	}
	return nil
}

func (g *DomainGenerator) InitResources() error {
	client := g.generateClient()
	domainList, err := g.loadDomains(client)
	if err != nil {
		return err
	}
	for _, domain := range domainList {
		err := g.loadDomainRecords(client, domain.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
