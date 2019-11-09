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

package alicloud

import (
	"github.com/GoogleCloudPlatform/terraformer/providers/alicloud/connectivity"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

// DnsGenerator Struct for generating AliCloud Elastic Compute Service
type DnsGenerator struct {
	AliCloudService
}

func resourceFromDomain(domain alidns.Domain) terraform_utils.Resource {
	return terraform_utils.NewResource(
		domain.DomainName,                      // id
		domain.DomainId+"__"+domain.DomainName, // name
		"alicloud_dns",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}

func resourceFromDomainRecord(record alidns.Record) terraform_utils.Resource {
	return terraform_utils.NewResource(
		record.RecordId, // id
		record.RecordId+"__"+record.DomainName, // name
		"alicloud_dns_record",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}

func initDomains(client *connectivity.AliyunClient) ([]alidns.Domain, error) {
	remaining := 1
	pageNumber := 1
	pageSize := 10

	allDomains := make([]alidns.Domain, 0)

	for remaining > 0 {
		raw, err := client.WithDnsClient(func(alidnsClient *alidns.Client) (interface{}, error) {
			request := alidns.CreateDescribeDomainsRequest()
			request.RegionId = client.RegionId
			request.PageSize = requests.NewInteger(pageSize)
			request.PageNumber = requests.NewInteger(pageNumber)
			return alidnsClient.DescribeDomains(request)
		})
		if err != nil {
			return nil, err
		}

		response := raw.(*alidns.DescribeDomainsResponse)
		for _, domain := range response.Domains.Domain {
			allDomains = append(allDomains, domain)

		}
		remaining = int(response.TotalCount) - pageNumber*pageSize
		pageNumber++
	}

	return allDomains, nil
}

func initDomainRecords(client *connectivity.AliyunClient, allDomains []alidns.Domain) ([]alidns.Record, error) {
	allDomainRecords := make([]alidns.Record, 0)

	for _, domain := range allDomains {
		remaining := 1
		pageNumber := 1
		pageSize := 10

		for remaining > 0 {
			raw, err := client.WithDnsClient(func(alidnsClient *alidns.Client) (interface{}, error) {
				request := alidns.CreateDescribeDomainRecordsRequest()
				request.RegionId = client.RegionId
				request.DomainName = domain.DomainName
				request.PageSize = requests.NewInteger(pageSize)
				request.PageNumber = requests.NewInteger(pageNumber)
				return alidnsClient.DescribeDomainRecords(request)
			})
			if err != nil {
				return nil, err
			}

			response := raw.(*alidns.DescribeDomainRecordsResponse)
			for _, record := range response.DomainRecords.Record {
				allDomainRecords = append(allDomainRecords, record)

			}
			remaining = int(response.TotalCount) - pageNumber*pageSize
			pageNumber++
		}
	}

	return allDomainRecords, nil
}

// InitResources Gets the list of all alidns domain ids and generates resources
func (g *DnsGenerator) InitResources() error {
	client, err := g.LoadClientFromProfile()
	if err != nil {
		return err
	}

	allDomains, err := initDomains(client)
	if err != nil {
		return err
	}

	allDomainRecords, err := initDomainRecords(client, allDomains)
	if err != nil {
		return err
	}

	for _, domain := range allDomains {
		resource := resourceFromDomain(domain)
		g.Resources = append(g.Resources, resource)
	}

	for _, record := range allDomainRecords {
		resource := resourceFromDomainRecord(record)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
