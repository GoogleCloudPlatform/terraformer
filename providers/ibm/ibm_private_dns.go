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
	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/catalog"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev2/controllerv2"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM/go-sdk-core/v3/core"
	dns "github.com/IBM/networking-go-sdk/dnssvcsv1"
)

//privateDNSTemplateGenerator ...
type privateDNSTemplateGenerator struct {
	IBMService
}

// loadPrivateDNS ...
func (g privateDNSTemplateGenerator) loadPrivateDNS(pDNSID string, pDNSName string, resGrpID string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		pDNSID,
		pDNSName,
		"ibm_resource_instance",
		"ibm",
		map[string]string{
			"resource_group_id": resGrpID,
		},
		[]string{},
		map[string]interface{}{})
	return resources
}

// loadPrivateDNSZone ...
func (g privateDNSTemplateGenerator) loadPrivateDNSZone(pDNSGuid string, zoneID string, dependsOn []string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		fmt.Sprintf("%s/%s", pDNSGuid, zoneID),
		zoneID,
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
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		fmt.Sprintf("%s/%s/%s", pDNSGuid, zoneID, permittedNetworkID),
		permittedNetworkID,
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
func (g privateDNSTemplateGenerator) loadPrivateDNSResourceRecord(pDNSGuid string, zoneID string, recordID string, dependsOn []string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		fmt.Sprintf("%s/%s/%s", pDNSGuid, zoneID, recordID),
		recordID,
		"ibm_dns_resource_record",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

// loadPrivateDNSGLBMonitor ...
func (g privateDNSTemplateGenerator) loadPrivateDNSGLBMonitor(pDNSGuid string, monitorID string, dependsOn []string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		fmt.Sprintf("%s/%s", pDNSGuid, monitorID),
		monitorID,
		"ibm_dns_glb_monitor",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

// loadPrivateDNSGLBPool ...
func (g privateDNSTemplateGenerator) loadPrivateDNSGLBPool(pDNSGuid string, poolID string, dependsOn []string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		fmt.Sprintf("%s/%s", pDNSGuid, poolID),
		poolID,
		"ibm_dns_glb_pool",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

// loadPrivateDNSGLB ...
func (g privateDNSTemplateGenerator) loadPrivateDNSGLB(pDNSGuid string, zoneID string, lbID string, dependsOn []string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		fmt.Sprintf("%s/%s/%s", pDNSGuid, zoneID, lbID),
		lbID,
		"ibm_dns_glb",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

//InitResources ...
func (g *privateDNSTemplateGenerator) InitResources() error {

	region := os.Getenv("IC_REGION")
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
		g.Resources = append(g.Resources, g.loadPrivateDNS(instanceID, instance.Name, instance.ResourceGroupID))
		var pDNSDependsOn []string
		pDNSDependsOn = append(pDNSDependsOn,
			"ibm_resource_instance."+terraformutils.TfSanitize(instance.Name))

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
			return fmt.Errorf("Error Listing Zones %s", err)
		}
		for _, zone := range zoneList.Dnszones {
			zoneID := *zone.ID
			g.Resources = append(g.Resources, g.loadPrivateDNSZone(instanceGUID, zoneID, pDNSDependsOn))

			var domainDependsOn []string
			domainDependsOn = append(pDNSDependsOn,
				"ibm_dns_zone."+terraformutils.TfSanitize(zoneID))

			// Permitted Network Records
			permittedNetworkOpt := dns.ListPermittedNetworksOptions{
				InstanceID: &instanceGUID,
				DnszoneID:  &zoneID,
			}
			permittedNetworkList, _, err := zService.ListPermittedNetworks(&permittedNetworkOpt)
			if err != nil {
				return fmt.Errorf("Error Listing Permitted Networks %s", err)
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
				return fmt.Errorf("Error Listing Resource Records %s", err)
			}
			for _, record := range resourceRecordList.ResourceRecords {
				recordID := *record.ID
				g.Resources = append(g.Resources, g.loadPrivateDNSResourceRecord(instanceGUID, zoneID, recordID, domainDependsOn))
			}

			// GLB Records
			glbOpt := dns.ListLoadBalancersOptions{
				InstanceID: &instanceGUID,
				DnszoneID:  &zoneID,
			}
			glbOptList, _, err := zService.ListLoadBalancers(&glbOpt)
			if err != nil {
				return fmt.Errorf("Error Listing GLBs %s", err)
			}
			for _, lb := range glbOptList.LoadBalancers {
				lbID := *lb.ID
				g.Resources = append(g.Resources, g.loadPrivateDNSGLB(instanceGUID, zoneID, lbID, domainDependsOn))
			}
		}
		// Monitor Records
		monitorOpt := dns.ListMonitorsOptions{
			InstanceID: &instanceGUID,
		}
		glbMonitorList, _, err := zService.ListMonitors(&monitorOpt)
		if err != nil {
			return fmt.Errorf("Error Listing GLB Monitor %s", err)
		}
		for _, monitor := range glbMonitorList.Monitors {
			monitorID := *monitor.ID
			g.Resources = append(g.Resources, g.loadPrivateDNSGLBMonitor(instanceGUID, monitorID, pDNSDependsOn))
		}

		// Pool Records
		glbPoolOpt := dns.ListPoolsOptions{
			InstanceID: &instanceGUID,
		}
		glbPoolOptList, _, err := zService.ListPools(&glbPoolOpt)
		if err != nil {
			return fmt.Errorf("Error Listing GLB Pools %s", err)
		}
		for _, pool := range glbPoolOptList.Pools {
			poolID := *pool.ID
			g.Resources = append(g.Resources, g.loadPrivateDNSGLBPool(instanceGUID, poolID, pDNSDependsOn))
		}

	}

	return nil
}
