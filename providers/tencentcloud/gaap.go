// Copyright 2022 The Terraformer Authors.
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

package tencentcloud

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	gaap "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gaap/v20180529"
)

type GaapGenerator struct {
	TencentCloudService
}

func (g *GaapGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := gaap.NewClient(&credential, region, profile)
	if err != nil {
		return err
	}

	if err := g.loadProxy(client); err != nil {
		return err
	}
	if err := g.loadRealServer(client); err != nil {
		return err
	}
	if err := g.loadCertificate(client); err != nil {
		return err
	}

	return nil
}

func (g *GaapGenerator) loadProxy(client *gaap.Client) error {
	request := gaap.NewDescribeProxiesRequest()
	filters := make([]string, 0)
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("tencentcloud_gaap_proxy") {
			filters = append(filters, filter.AcceptableValues...)
		}
	}
	for i := range filters {
		request.InstanceIds = append(request.InstanceIds, &filters[i])
	}

	var offset uint64
	var pageSize uint64 = 50
	allInstances := make([]*gaap.ProxyInfo, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		response, err := client.DescribeProxies(request)
		if err != nil {
			return err
		}

		allInstances = append(allInstances, response.Response.ProxySet...)
		if len(response.Response.ProxySet) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.ProxyId,
			*instance.ProxyName+"_"+*instance.ProxyId,
			"tencentcloud_gaap_proxy",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)

		if len(g.Filter) > 0 {
			match := false
			for _, filter := range g.Filter {
				if filter.Filter(resource) {
					match = true
					break
				}
			}
			if !match {
				continue
			}
		}
		if err := g.loadHTTPListener(client, *instance.ProxyId, resource.ResourceName); err != nil {
			return err
		}
		if err := g.loadHTTPSListener(client, *instance.ProxyId, resource.ResourceName); err != nil {
			return err
		}
		if err := g.loadTCPListener(client, *instance.ProxyId, resource.ResourceName); err != nil {
			return err
		}
		if err := g.loadUDPListener(client, *instance.ProxyId, resource.ResourceName); err != nil {
			return err
		}
	}

	return nil
}

func (g *GaapGenerator) matchFilter(resource *terraformutils.Resource) bool {

	return false
}

func (g *GaapGenerator) loadRealServer(client *gaap.Client) error {
	request := gaap.NewDescribeRealServersRequest()
	var projectID int64 = -1
	request.ProjectId = &projectID

	var offset uint64
	var pageSize uint64 = 50
	allInstances := make([]*gaap.BindRealServerInfo, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		response, err := client.DescribeRealServers(request)
		if err != nil {
			return err
		}

		allInstances = append(allInstances, response.Response.RealServerSet...)
		if len(response.Response.RealServerSet) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.RealServerId,
			*instance.RealServerName+"_"+*instance.RealServerId,
			"tencentcloud_gaap_realserver",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}

func (g *GaapGenerator) loadHTTPListener(client *gaap.Client, proxyID, resourceName string) error {
	request := gaap.NewDescribeHTTPListenersRequest()
	request.ProxyId = &proxyID
	var offset uint64
	var pageSize uint64 = 50
	allInstances := make([]*gaap.HTTPListener, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		response, err := client.DescribeHTTPListeners(request)
		if err != nil {
			return err
		}

		allInstances = append(allInstances, response.Response.ListenerSet...)
		if len(response.Response.ListenerSet) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.ListenerId,
			*instance.ListenerName+"_"+*instance.ListenerId,
			"tencentcloud_gaap_layer7_listener",
			"tencentcloud",
			map[string]string{"proxy_id": proxyID},
			[]string{},
			map[string]interface{}{},
		)
		resource.AdditionalFields["proxy_id"] = "${tencentcloud_gaap_proxy." + resourceName + ".id}"
		g.Resources = append(g.Resources, resource)

		if err := g.loadDomain(client, *instance.ListenerId, "HTTP", resource.ResourceName); err != nil {
			return err
		}
	}
	return nil
}

func (g *GaapGenerator) loadHTTPSListener(client *gaap.Client, proxyID, resourceName string) error {
	request := gaap.NewDescribeHTTPSListenersRequest()
	request.ProxyId = &proxyID
	var offset uint64
	var pageSize uint64 = 50
	allInstances := make([]*gaap.HTTPSListener, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		response, err := client.DescribeHTTPSListeners(request)
		if err != nil {
			return err
		}

		allInstances = append(allInstances, response.Response.ListenerSet...)
		if len(response.Response.ListenerSet) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.ListenerId,
			*instance.ListenerName+"_"+*instance.ListenerId,
			"tencentcloud_gaap_layer7_listener",
			"tencentcloud",
			map[string]string{"proxy_id": proxyID},
			[]string{},
			map[string]interface{}{},
		)
		resource.AdditionalFields["proxy_id"] = "${tencentcloud_gaap_proxy." + resourceName + ".id}"
		g.Resources = append(g.Resources, resource)

		if err := g.loadDomain(client, *instance.ListenerId, "HTTPS", resource.ResourceName); err != nil {
			return err
		}
	}
	return nil
}

func (g *GaapGenerator) loadTCPListener(client *gaap.Client, proxyID, resourceName string) error {
	request := gaap.NewDescribeTCPListenersRequest()
	request.ProxyId = &proxyID
	var offset uint64
	var pageSize uint64 = 50
	allInstances := make([]*gaap.TCPListener, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		response, err := client.DescribeTCPListeners(request)
		if err != nil {
			return err
		}

		allInstances = append(allInstances, response.Response.ListenerSet...)
		if len(response.Response.ListenerSet) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.ListenerId,
			*instance.ListenerName+"_"+*instance.ListenerId,
			"tencentcloud_gaap_layer4_listener",
			"tencentcloud",
			map[string]string{"proxy_id": proxyID},
			[]string{},
			map[string]interface{}{},
		)
		resource.AdditionalFields["proxy_id"] = "${tencentcloud_gaap_proxy." + resourceName + ".id}"
		g.Resources = append(g.Resources, resource)
	}
	return nil
}

func (g *GaapGenerator) loadUDPListener(client *gaap.Client, proxyID, resourceName string) error {
	request := gaap.NewDescribeUDPListenersRequest()
	request.ProxyId = &proxyID
	var offset uint64
	var pageSize uint64 = 50
	allInstances := make([]*gaap.UDPListener, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		response, err := client.DescribeUDPListeners(request)
		if err != nil {
			return err
		}

		allInstances = append(allInstances, response.Response.ListenerSet...)
		if len(response.Response.ListenerSet) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.ListenerId,
			*instance.ListenerName+"_"+*instance.ListenerId,
			"tencentcloud_gaap_layer4_listener",
			"tencentcloud",
			map[string]string{"proxy_id": proxyID},
			[]string{},
			map[string]interface{}{},
		)
		resource.AdditionalFields["proxy_id"] = "${tencentcloud_gaap_proxy." + resourceName + ".id}"
		g.Resources = append(g.Resources, resource)
	}
	return nil
}

func (g *GaapGenerator) loadDomain(client *gaap.Client, listenerID, protocol, resourceName string) error {
	request := gaap.NewDescribeRulesRequest()
	request.ListenerId = &listenerID
	response, err := client.DescribeRules(request)
	if err != nil {
		return err
	}

	for _, domain := range response.Response.DomainRuleSet {
		resource := terraformutils.NewResource(
			fmt.Sprintf("%s+%s+%s", listenerID, protocol, *domain.Domain),
			fmt.Sprintf("%s+%s+%s", listenerID, protocol, *domain.Domain),
			"tencentcloud_gaap_http_domain",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		resource.AdditionalFields["listener_id"] = "${tencentcloud_gaap_layer7_listener." + resourceName + ".id}"
		g.Resources = append(g.Resources, resource)

		for _, rule := range domain.RuleSet {
			ruleResource := terraformutils.NewResource(
				*rule.RuleId,
				*rule.RuleId,
				"tencentcloud_gaap_http_rule",
				"tencentcloud",
				map[string]string{},
				[]string{},
				map[string]interface{}{},
			)
			ruleResource.AdditionalFields["listener_id"] = "${tencentcloud_gaap_layer7_listener." + resourceName + ".id}"
			ruleResource.AdditionalFields["domain"] = "${tencentcloud_gaap_http_domain." + resource.ResourceName + ".domain}"
			g.Resources = append(g.Resources, ruleResource)
		}
	}
	return nil
}

func (g *GaapGenerator) loadCertificate(client *gaap.Client) error {
	request := gaap.NewDescribeCertificatesRequest()
	var offset uint64
	var pageSize uint64 = 50
	allInstances := make([]*gaap.Certificate, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		response, err := client.DescribeCertificates(request)
		if err != nil {
			return err
		}

		allInstances = append(allInstances, response.Response.CertificateSet...)
		if len(response.Response.CertificateSet) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.CertificateId,
			*instance.CertificateAlias+"_"+*instance.CertificateId,
			"tencentcloud_gaap_certificate",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		resource.Item["content"] = ""
		g.Resources = append(g.Resources, resource)
	}

	return nil
}

func (g *GaapGenerator) PostConvertHook() error {
	for _, resource := range g.Resources {
		if resource.InstanceInfo.Type == "tencentcloud_gaap_http_domain" {
			delete(resource.Item, "client_certificate_id")
			delete(resource.Item, "realserver_certificate_id")
		} else if resource.InstanceInfo.Type == "tencentcloud_gaap_layer7_listener" {
			delete(resource.Item, "client_certificate_id")
		}
	}

	return nil
}
