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

package ibm

import (
	"fmt"
	"os"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/catalog"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev2/controllerv2"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM/go-sdk-core/v3/core"
	dns "github.com/IBM/networking-go-sdk/dnssvcsv1"
)

// privateDNSTemplateGenerator ...
type privateDNSTemplateGenerator struct {
	IBMService
}

// loadPrivateDNS ...
func (g privateDNSTemplateGenerator) loadPrivateDNS() func(pDNSID, pDNSName, resGrpID string) terraformutils.Resource {
	names := make(map[string]struct{})
	random := true
	return func(pDNSID, pDNSName, resGrpID string) terraformutils.Resource {
		names, random = getRandom(names, pDNSName, random)
		resource := terraformutils.NewResource(
			pDNSID,
			normalizeResourceName(pDNSName, random),
			"ibm_resource_instance",
			"ibm",
			map[string]string{
				"resource_group_id": resGrpID,
			},
			[]string{},
			map[string]interface{}{})
		return resource
	}
}

// loadPrivateDNSZone ...
func (g privateDNSTemplateGenerator) loadPrivateDNSZone(pDNSGuid string, zoneID string, dependsOn []string) terraformutils.Resource {
	resources := terraformutils.NewResource(
		fmt.Sprintf("%s/%s", pDNSGuid, zoneID),
		normalizeResourceName("ibm_dns_zone", true),
		"ibm_dns_zone",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

// loadPrivateDNSPermittedNetwork ...
func (g privateDNSTemplateGenerator) loadPrivateDNSPermittedNetwork(pDNSGuid string, zoneID string, permittedNetworkID string, dependsOn []string) terraformutils.Resource {
	resources := terraformutils.NewResource(
		fmt.Sprintf("%s/%s/%s", pDNSGuid, zoneID, permittedNetworkID),
		normalizeResourceName("ibm_dns_permitted_network", true),
		"ibm_dns_permitted_network",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

// loadPrivateDNSResourceRecord ...
func (g privateDNSTemplateGenerator) loadPrivateDNSResourceRecord() func(pDNSGuid, zoneID, recordID, recordName string, dependsOn []string) terraformutils.Resource {
	names := make(map[string]struct{})
	random := false
	return func(pDNSGuid, zoneID, recordID, recordName string, dependsOn []string) terraformutils.Resource {
		names, random = getRandom(names, recordName, random)
		resources := terraformutils.NewResource(
			fmt.Sprintf("%s/%s/%s", pDNSGuid, zoneID, recordID),
			normalizeResourceName(recordName, random),
			"ibm_dns_resource_record",
			"ibm",
			map[string]string{},
			[]string{},
			map[string]interface{}{
				"depends_on": dependsOn,
			})
		return resources
	}
}

// loadPrivateDNSGLBMonitor ...
func (g privateDNSTemplateGenerator) loadPrivateDNSGLBMonitor() func(pDNSGuid, monitorID, monitorName string, dependsOn []string) terraformutils.Resource {
	names := make(map[string]struct{})
	random := false
	return func(pDNSGuid, monitorID, monitorName string, dependsOn []string) terraformutils.Resource {
		names, random = getRandom(names, monitorName, random)
		resources := terraformutils.NewResource(
			fmt.Sprintf("%s/%s", pDNSGuid, monitorID),
			normalizeResourceName(monitorName, random),
			"ibm_dns_glb_monitor",
			"ibm",
			map[string]string{},
			[]string{},
			map[string]interface{}{
				"depends_on": dependsOn,
			})
		return resources
	}
}

// loadPrivateDNSGLBPool ...
func (g privateDNSTemplateGenerator) loadPrivateDNSGLBPool() func(pDNSGuid, poolID, poolName string, dependsOn []string) terraformutils.Resource {
	names := make(map[string]struct{})
	random := false
	return func(pDNSGuid, poolID, poolName string, dependsOn []string) terraformutils.Resource {
		names, random = getRandom(names, poolName, random)
		resources := terraformutils.NewResource(
			fmt.Sprintf("%s/%s", pDNSGuid, poolID),
			normalizeResourceName(poolName, random),
			"ibm_dns_glb_pool",
			"ibm",
			map[string]string{},
			[]string{},
			map[string]interface{}{
				"depends_on": dependsOn,
			})
		return resources
	}
}

// loadPrivateDNSGLB ...
func (g privateDNSTemplateGenerator) loadPrivateDNSGLB() func(pDNSGuid, zoneID, lbID, lbName string, dependsOn []string) terraformutils.Resource {
	names := make(map[string]struct{})
	random := false
	return func(pDNSGuid, zoneID, lbID, lbName string, dependsOn []string) terraformutils.Resource {
		names, random = getRandom(names, lbName, random)
		resources := terraformutils.NewResource(
			fmt.Sprintf("%s/%s/%s", pDNSGuid, zoneID, lbID),
			normalizeResourceName(lbName, random),
			"ibm_dns_glb",
			"ibm",
			map[string]string{},
			[]string{},
			map[string]interface{}{
				"depends_on": dependsOn,
			})
		return resources
	}
}

// InitResources ...
func (g *privateDNSTemplateGenerator) InitResources() error {

	region := g.Args["region"].(string)
	bmxConfig := &bluemix.Config{
		BluemixAPIKey: os.Getenv("IC_API_KEY"),
		Region:        region,
	}
	sess, err := session.New(bmxConfig)
	if err != nil {
		return err
	}
	defaultDNSURL := "https://api.dns-svcs.cloud.ibm.com/v1"

	err = authenticateAPIKey(sess)
	if err != nil {
		return err
	}

	bluemixToken := ""
	if strings.HasPrefix(sess.Config.IAMAccessToken, "Bearer") {
		bluemixToken = sess.Config.IAMAccessToken[7:len(sess.Config.IAMAccessToken)]
	} else {
		bluemixToken = sess.Config.IAMAccessToken
	}

	catalogClient, err := catalog.New(sess)
	if err != nil {
		return err
	}

	controllerClient, err := controllerv2.New(sess)
	if err != nil {
		return err
	}

	serviceID, err := catalogClient.ResourceCatalog().FindByName("dns-svcs", true)
	if err != nil {
		return err
	}
	query := controllerv2.ServiceInstanceQuery{
		ServiceID: serviceID[0].ID,
	}
	pDNSInstances, err := controllerClient.ResourceServiceInstanceV2().ListInstances(query)
	if err != nil {
		return err
	}

	for _, instance := range pDNSInstances {
		instanceID := instance.ID
		instanceGUID := instance.Guid
		// Instance
		fnObjt := g.loadPrivateDNS()
		g.Resources = append(g.Resources, fnObjt(instanceID, instance.Name, instance.ResourceGroupID))
		resourceName := g.Resources[len(g.Resources)-1:][0].ResourceName
		var pDNSDependsOn []string
		pDNSDependsOn = append(pDNSDependsOn,
			"ibm_resource_instance."+resourceName)

		// Zones
		zoneOpts := &dns.DnsSvcsV1Options{
			URL: defaultDNSURL,
			Authenticator: &core.BearerTokenAuthenticator{
				BearerToken: bluemixToken,
			},
		}

		zService, err := dns.NewDnsSvcsV1(zoneOpts)
		if err != nil {
			return err
		}
		zoneOpt := dns.ListDnszonesOptions{
			InstanceID: &instanceGUID,
		}
		zoneList, _, err := zService.ListDnszones(&zoneOpt)
		if err != nil {
			return fmt.Errorf("error Listing Zones %s", err)
		}
		for _, zone := range zoneList.Dnszones {
			zoneID := *zone.ID
			g.Resources = append(g.Resources, g.loadPrivateDNSZone(instanceGUID, zoneID, pDNSDependsOn))
			domainResourceName := g.Resources[len(g.Resources)-1:][0].ResourceName
			domainDependsOn := makeDependsOn(pDNSDependsOn, "ibm_dns_zone."+domainResourceName)

			// Permitted Network Records
			permittedNetworkOpt := dns.ListPermittedNetworksOptions{
				InstanceID: &instanceGUID,
				DnszoneID:  &zoneID,
			}
			permittedNetworkList, _, err := zService.ListPermittedNetworks(&permittedNetworkOpt)
			if err != nil {
				return fmt.Errorf("error Listing Permitted Networks %s", err)
			}
			for _, permittedNetwork := range permittedNetworkList.PermittedNetworks {
				permittedNetworkID := *permittedNetwork.ID
				g.Resources = append(g.Resources, g.loadPrivateDNSPermittedNetwork(instanceGUID, zoneID, permittedNetworkID, domainDependsOn))
			}

			// Resource Records
			dnsRecordOpt := dns.ListResourceRecordsOptions{
				InstanceID: &instanceGUID,
				DnszoneID:  &zoneID,
			}
			resourceRecordList, _, err := zService.ListResourceRecords(&dnsRecordOpt)
			if err != nil {
				return fmt.Errorf("error Listing Resource Records %s", err)
			}

			pdnsFnObjt := g.loadPrivateDNSResourceRecord()
			for _, record := range resourceRecordList.ResourceRecords {
				g.Resources = append(g.Resources, pdnsFnObjt(instanceGUID, zoneID, *record.ID, *record.Name, domainDependsOn))
			}

			// GLB Records
			glbOpt := dns.ListLoadBalancersOptions{
				InstanceID: &instanceGUID,
				DnszoneID:  &zoneID,
			}
			glbOptList, _, err := zService.ListLoadBalancers(&glbOpt)
			if err != nil {
				return fmt.Errorf("error Listing GLBs %s", err)
			}
			glbFntObj := g.loadPrivateDNSGLB()
			for _, lb := range glbOptList.LoadBalancers {
				g.Resources = append(g.Resources, glbFntObj(instanceGUID, zoneID, *lb.ID, *lb.Name, domainDependsOn))
			}
		}
		// Monitor Records
		monitorOpt := dns.ListMonitorsOptions{
			InstanceID: &instanceGUID,
		}
		glbMonitorList, _, err := zService.ListMonitors(&monitorOpt)
		if err != nil {
			return fmt.Errorf("error Listing GLB Monitor %s", err)
		}

		lbMonitorObjt := g.loadPrivateDNSGLBMonitor()
		for _, monitor := range glbMonitorList.Monitors {
			g.Resources = append(g.Resources, lbMonitorObjt(instanceGUID, *monitor.ID, *monitor.Name, pDNSDependsOn))
		}

		// Pool Records
		glbPoolOpt := dns.ListPoolsOptions{
			InstanceID: &instanceGUID,
		}
		glbPoolOptList, _, err := zService.ListPools(&glbPoolOpt)
		if err != nil {
			return fmt.Errorf("error Listing GLB Pools %s", err)
		}
		dnsGlbfnObj := g.loadPrivateDNSGLBPool()
		for _, pool := range glbPoolOptList.Pools {
			g.Resources = append(g.Resources, dnsGlbfnObj(instanceGUID, *pool.ID, *pool.Name, pDNSDependsOn))
		}

	}

	return nil
}
