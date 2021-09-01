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
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

// DNSGenerator Struct for generating AliCloud Elastic Compute Service
type DNSGenerator struct {
	AliCloudService
}

func resourceFromDomain(domain alidns.DomainInDescribeDomains) terraformutils.Resource {
	return terraformutils.NewResource(
		domain.DomainName,                      // id
		domain.DomainId+"__"+domain.DomainName, // nolint
		"alicloud_dns",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}

func resourceFromDomainRecord(record alidns.Record) terraformutils.Resource {
	return terraformutils.NewResource(
		record.RecordId, // nolint
		record.RecordId+"__"+record.DomainName, // nolint
		"alicloud_dns_record",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}

func initDomains(client *connectivity.AliyunClient) ([]alidns.DomainInDescribeDomains, error) {
	remaining := 1
	pageNumber := 1
	pageSize := 10

	allDomains := make([]alidns.DomainInDescribeDomains, 0)

	for remaining > 0 {
		raw, err := client.WithDNSClient(func(alidnsClient *alidns.Client) (interface{}, error) {
			request := alidns.CreateDescribeDomainsRequest()
			request.RegionId = client.RegionID
			request.PageSize = requests.NewInteger(pageSize)
			request.PageNumber = requests.NewInteger(pageNumber)
			return alidnsClient.DescribeDomains(request)
		})
		if err != nil {
			return nil, err
		}

		response := raw.(*alidns.DescribeDomainsResponse)
		allDomains = append(allDomains, response.Domains.Domain...)
		remaining = int(response.TotalCount) - pageNumber*pageSize
		pageNumber++
	}

	return allDomains, nil
}

func initDomainRecords(client *connectivity.AliyunClient, allDomains []alidns.DomainInDescribeDomains) ([]alidns.Record, error) {
	allDomainRecords := make([]alidns.Record, 0)

	for _, domain := range allDomains {
		remaining := 1
		pageNumber := 1
		pageSize := 10

		for remaining > 0 {
			raw, err := client.WithDNSClient(func(alidnsClient *alidns.Client) (interface{}, error) {
				request := alidns.CreateDescribeDomainRecordsRequest()
				request.RegionId = client.RegionID
				request.DomainName = domain.DomainName
				request.PageSize = requests.NewInteger(pageSize)
				request.PageNumber = requests.NewInteger(pageNumber)
				return alidnsClient.DescribeDomainRecords(request)
			})
			if err != nil {
				return nil, err
			}

			response := raw.(*alidns.DescribeDomainRecordsResponse)
			allDomainRecords = append(allDomainRecords, response.DomainRecords.Record...)
			remaining = int(response.TotalCount) - pageNumber*pageSize
			pageNumber++
		}
	}

	return allDomainRecords, nil
}

// InitResources Gets the list of all alidns domain ids and generates resources
func (g *DNSGenerator) InitResources() error {
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
