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
	"github.com/IBM/networking-go-sdk/dnsrecordsv1"
	"github.com/IBM/networking-go-sdk/edgefunctionsapiv1"
	"github.com/IBM/networking-go-sdk/globalloadbalancermonitorv1"
	"github.com/IBM/networking-go-sdk/globalloadbalancerv1"
	"github.com/IBM/networking-go-sdk/pageruleapiv1"
	"github.com/IBM/networking-go-sdk/zonelockdownv1"
	"github.com/IBM/networking-go-sdk/zoneratelimitsv1"
	"github.com/IBM/networking-go-sdk/zonesv1"
)

// CISGenerator ..
type CISGenerator struct {
	IBMService
}

func (g CISGenerator) loadInstances(crn, name, resGrpID string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		crn,
		name,
		"ibm_cis",
		"ibm",
		map[string]string{
			"resource_group_id": resGrpID,
		},
		[]string{},
		map[string]interface{}{})
	return resources
}

func (g CISGenerator) loadDomains(crn, domainID string, dependsOn []string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		fmt.Sprintf("%s:%s", domainID, crn),
		domainID,
		"ibm_cis_domain",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

func (g CISGenerator) loadDNSRecords(crn, domainID, dnsRecordID string, dependsOn []string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		fmt.Sprintf("%s:%s:%s", dnsRecordID, domainID, crn),
		dnsRecordID,
		"ibm_cis_dns_record",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

func (g CISGenerator) loadFirewallLockdown(resourceName, crn, domainID, fID, fType string, dependsOn []string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		fmt.Sprintf("%s:%s:%s:%s", fType, fID, domainID, crn),
		resourceName,
		"ibm_cis_firewall",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

func (g CISGenerator) loadDomainSettings(crn, dID string, dependsOn []string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		fmt.Sprintf("%s:%s", dID, crn),
		dID,
		"ibm_cis_domain_settings",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

func (g CISGenerator) loadGlobalBalancer(crn, dID, gID string, dependsOn []string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		fmt.Sprintf("%s:%s:%s", gID, dID, crn),
		fmt.Sprintf("%s:%s:%s", gID, dID, crn),
		"ibm_cis_global_load_balancer",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

func (g CISGenerator) loadGlobalBalancerPool(crn, pID string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		fmt.Sprintf("%s:%s", pID, crn),
		pID,
		"ibm_cis_origin_pool",
		g.ProviderName,
		[]string{})
	return resources
}

func (g CISGenerator) loadGlobalBalancerMonitor(crn, gblmID string, dependsOn []string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		fmt.Sprintf("%s:%s", gblmID, crn),
		gblmID,
		"ibm_cis_healthcheck",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

func (g CISGenerator) loadRateLimit(resourceName, crn, dID, rID string, dependsOn []string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		fmt.Sprintf("%s:%s:%s", rID, dID, crn),
		resourceName,
		"ibm_cis_rate_limit",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

func (g CISGenerator) loadEdgeFunctionAction(crn, dID, actionID string, dependsOn []string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		fmt.Sprintf("%s:%s:%s", actionID, dID, crn),
		actionID,
		"ibm_cis_edge_functions_action",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

func (g CISGenerator) loadEdgeFunctionTrigger(crn, dID, triggerID string, dependsOn []string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		fmt.Sprintf("%s:%s:%s", triggerID, dID, crn),
		triggerID,
		"ibm_cis_edge_functions_trigger",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

func (g CISGenerator) loadWafRulePackage(crn, dID, pkgID string, dependsOn []string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		fmt.Sprintf("%s:%s:%s", pkgID, dID, crn),
		pkgID,
		"ibm_cis_waf_package",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{})
	return resources
}

func (g CISGenerator) loadPageRule(crn, dID, ruleID string, dependsOn []string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		fmt.Sprintf("%s:%s:%s", ruleID, dID, crn),
		ruleID,
		"ibm_cis_page_rule",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{})
	return resources
}

func (g CISGenerator) loadCustomPage(crn, dID, cpID string, dependsOn []string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		fmt.Sprintf("%s:%s:%s", cpID, dID, crn),
		cpID,
		"ibm_cis_custom_page",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{})
	return resources
}

// InitResources ..
func (g *CISGenerator) InitResources() error {
	DefaultCisURL := "https://api.cis.cloud.ibm.com"

	bmxConfig := &bluemix.Config{
		BluemixAPIKey: os.Getenv("IC_API_KEY"),
	}

	sess, err := session.New(bmxConfig)
	if err != nil {
		return err
	}

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

	serviceID, err := catalogClient.ResourceCatalog().FindByName("internet-svcs", true)
	if err != nil {
		return err
	}
	query := controllerv2.ServiceInstanceQuery{
		ServiceID: serviceID[0].ID,
	}
	cisInstances, err := controllerClient.ResourceServiceInstanceV2().ListInstances(query)
	if err != nil {
		return err
	}

	for _, c := range cisInstances {
		//Instance
		crn := c.Crn.String()
		g.Resources = append(g.Resources, g.loadInstances(crn, c.Name, c.ResourceGroupID))

		var cisDependsOn []string
		cisDependsOn = append(cisDependsOn,
			"ibm_cis."+terraformutils.TfSanitize(c.Name))

		//Domain
		zoneOpts := &zonesv1.ZonesV1Options{
			URL: DefaultCisURL,
			Authenticator: &core.BearerTokenAuthenticator{
				BearerToken: bluemixToken,
			},
			Crn: &crn,
		}

		zService, err := zonesv1.NewZonesV1(zoneOpts)
		if err != nil {
			return err
		}

		domainOpts := &zonesv1.ListZonesOptions{}
		zoneList, _, err := zService.ListZones(domainOpts)
		if err != nil {
			return err
		}

		//Health Monitor
		gblmOpts := &globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
			URL:           DefaultCisURL,
			Authenticator: &core.BearerTokenAuthenticator{BearerToken: bluemixToken},
			Crn:           &crn,
		}

		gblmService, _ := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1(gblmOpts)
		gblmList, _, err := gblmService.ListAllLoadBalancerMonitors(&globalloadbalancermonitorv1.ListAllLoadBalancerMonitorsOptions{})
		if err != nil {
			return err
		}

		for _, gblm := range gblmList.Result {
			g.Resources = append(g.Resources, g.loadGlobalBalancerMonitor(crn, *gblm.ID, cisDependsOn))
		}

		for _, z := range zoneList.Result {
			var domainDependsOn []string
			domainDependsOn = append(domainDependsOn,
				"ibm_cis."+terraformutils.TfSanitize(c.Name))
			domainDependsOn = append(domainDependsOn,
				"ibm_cis_domain."+terraformutils.TfSanitize(*z.ID))

			g.Resources = append(g.Resources, g.loadDomains(crn, *z.ID, domainDependsOn))

			//DNS Record
			zoneID := *z.ID
			dnsOpts := &dnsrecordsv1.DnsRecordsV1Options{
				URL: DefaultCisURL,
				Authenticator: &core.BearerTokenAuthenticator{
					BearerToken: bluemixToken,
				},
				Crn:            &crn,
				ZoneIdentifier: &zoneID,
			}

			//Domain Setting
			g.Resources = append(g.Resources, g.loadDomainSettings(crn, *z.ID, domainDependsOn))

			dnsService, err := dnsrecordsv1.NewDnsRecordsV1(dnsOpts)
			if err != nil {
				return err
			}

			dOpts := &dnsrecordsv1.ListAllDnsRecordsOptions{}
			dnsList, _, err := dnsService.ListAllDnsRecords(dOpts)
			if err != nil {
				return err
			}

			//IBM Network CIS WAF Package
			// cisWAFPackageOpt := &wafrulepackagesapiv1.WafRulePackagesApiV1Options{
			// 	URL: DefaultCisURL,
			// 	Authenticator: &core.BearerTokenAuthenticator{
			// 		BearerToken: bluemixToken,
			// 	},
			// 	Crn:    &crn,
			// 	ZoneID: &zoneID,
			// }
			// cisWAFPackageClient, _ := wafrulepackagesapiv1.NewWafRulePackagesApiV1(cisWAFPackageOpt)
			// wasPkgList, _, err := cisWAFPackageClient.ListWafPackages(&wafrulepackagesapiv1.ListWafPackagesOptions{})
			// if err != nil {
			// 	return err
			// }

			// for _, wafPkg := range wasPkgList.Result {
			// 	fmt.Println("*wfpackge.ID ::", *wafPkg.Name)
			// 	//g.Resources = append(g.Resources, g.loadWafRulePackage(crn, *z.ID, *wafPkg.ID, domainDependsOn))
			// }

			//IBM Network CIS Page Rules
			cisPageRuleOpt := &pageruleapiv1.PageRuleApiV1Options{
				URL: DefaultCisURL,
				Authenticator: &core.BearerTokenAuthenticator{
					BearerToken: bluemixToken,
				},
				Crn:    &crn,
				ZoneID: &zoneID,
			}
			cisPageRuleClient, _ := pageruleapiv1.NewPageRuleApiV1(cisPageRuleOpt)
			cisPgList, _, err := cisPageRuleClient.ListPageRules(&pageruleapiv1.ListPageRulesOptions{})
			if err != nil {
				return err
			}

			for _, pg := range cisPgList.Result {
				g.Resources = append(g.Resources, g.loadPageRule(crn, *z.ID, *pg.ID, domainDependsOn))
			}

			//Rate Limit
			rateLimitPoolOpts := &zoneratelimitsv1.ZoneRateLimitsV1Options{
				URL:            DefaultCisURL,
				Authenticator:  &core.BearerTokenAuthenticator{BearerToken: bluemixToken},
				Crn:            &crn,
				ZoneIdentifier: &zoneID,
			}

			rateLimitService, _ := zoneratelimitsv1.NewZoneRateLimitsV1(rateLimitPoolOpts)
			rateLimitList, _, err := rateLimitService.ListAllZoneRateLimits(&zoneratelimitsv1.ListAllZoneRateLimitsOptions{})
			if err != nil {
				return err
			}

			for _, rl := range rateLimitList.Result {
				resourceName := fmt.Sprintf("%s:%s", "ibm_cis_rate_limit", *z.ID)
				g.Resources = append(g.Resources, g.loadRateLimit(resourceName, crn, *z.ID, *rl.ID, domainDependsOn))
			}

			//Firewall Lockdown
			firewallOpts := &zonelockdownv1.ZoneLockdownV1Options{
				URL: DefaultCisURL,
				Authenticator: &core.BearerTokenAuthenticator{
					BearerToken: bluemixToken,
				},
				Crn:            &crn,
				ZoneIdentifier: &zoneID,
			}

			fService, err := zonelockdownv1.NewZoneLockdownV1(firewallOpts)
			if err != nil {
				return err
			}

			firewallList, _, err := fService.ListAllZoneLockownRules(&zonelockdownv1.ListAllZoneLockownRulesOptions{})
			if err != nil {
				return err
			}

			// IBM Network CIS Edge Function Triggers
			cisEdgeFunctionOpt := &edgefunctionsapiv1.EdgeFunctionsApiV1Options{
				URL: DefaultCisURL,
				Authenticator: &core.BearerTokenAuthenticator{
					BearerToken: bluemixToken,
				},
				Crn:            &crn,
				ZoneIdentifier: &zoneID,
			}

			cisEdgeFunctionClient, _ := edgefunctionsapiv1.NewEdgeFunctionsApiV1(cisEdgeFunctionOpt)
			edgeTriggerList, _, err := cisEdgeFunctionClient.ListEdgeFunctionsTriggers(&edgefunctionsapiv1.ListEdgeFunctionsTriggersOptions{})
			if err != nil {
				return err
			}

			actionName := ""
			for _, el := range edgeTriggerList.Result {
				actionName = *el.Script
				g.Resources = append(g.Resources, g.loadEdgeFunctionTrigger(crn, *z.ID, *el.ID, domainDependsOn))
			}

			if actionName != "" {
				g.Resources = append(g.Resources, g.loadEdgeFunctionAction(crn, *z.ID, actionName, domainDependsOn))
			}

			for _, f := range firewallList.Result {
				resourceName := fmt.Sprintf("%s:%s", "ibm_cis_firewall", *f.ID)
				g.Resources = append(g.Resources, g.loadFirewallLockdown(resourceName, crn, *z.ID, *f.ID, "lockdowns", domainDependsOn))
			}

			for _, d := range dnsList.Result {
				dnsDependsOn := append(domainDependsOn,
					"ibm_cis_dns_record."+terraformutils.TfSanitize(*d.ID))

				g.Resources = append(g.Resources, g.loadDNSRecords(crn, *z.ID, *d.ID, domainDependsOn))

				//Global Load Balancer
				gblSetttingOpts := &globalloadbalancerv1.GlobalLoadBalancerV1Options{
					URL: DefaultCisURL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bluemixToken,
					},
					Crn:            &crn,
					ZoneIdentifier: &zoneID,
				}

				gblService, err := globalloadbalancerv1.NewGlobalLoadBalancerV1(gblSetttingOpts)
				if err != nil {
					return err
				}

				gblList, _, err := gblService.ListAllLoadBalancers(&globalloadbalancerv1.ListAllLoadBalancersOptions{})
				if err != nil {
					return err
				}

				for _, gb := range gblList.Result {
					g.Resources = append(g.Resources, g.loadGlobalBalancer(crn, *z.ID, *gb.ID, dnsDependsOn))
				}
			}
		}
	}

	return nil
}
