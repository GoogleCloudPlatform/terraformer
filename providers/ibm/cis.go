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
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/catalog"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev2/controllerv2"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/IBM/networking-go-sdk/custompagesv1"
	"github.com/IBM/networking-go-sdk/dnsrecordsv1"
	"github.com/IBM/networking-go-sdk/edgefunctionsapiv1"
	"github.com/IBM/networking-go-sdk/filtersv1"
	"github.com/IBM/networking-go-sdk/globalloadbalancermonitorv1"
	"github.com/IBM/networking-go-sdk/globalloadbalancerpoolsv0"
	"github.com/IBM/networking-go-sdk/globalloadbalancerv1"
	"github.com/IBM/networking-go-sdk/pageruleapiv1"
	"github.com/IBM/networking-go-sdk/rangeapplicationsv1"
	"github.com/IBM/networking-go-sdk/routingv1"
	"github.com/IBM/networking-go-sdk/sslcertificateapiv1"
	"github.com/IBM/networking-go-sdk/useragentblockingrulesv1"
	"github.com/IBM/networking-go-sdk/wafrulegroupsapiv1"
	"github.com/IBM/networking-go-sdk/wafrulepackagesapiv1"
	"github.com/IBM/networking-go-sdk/zonefirewallaccessrulesv1"
	"github.com/IBM/networking-go-sdk/zonelockdownv1"
	"github.com/IBM/networking-go-sdk/zoneratelimitsv1"
	"github.com/IBM/networking-go-sdk/zonesv1"
)

// CISGenerator ..
type CISGenerator struct {
	IBMService
}

func (g CISGenerator) loadInstances(crn, name, resGrpID string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		crn,
		normalizeResourceName(name, false),
		"ibm_cis",
		"ibm",
		map[string]string{
			"resource_group_id": resGrpID,
		},
		[]string{},
		map[string]interface{}{})
	return resource
}

func (g CISGenerator) loadDomains(crn, domainID string, dependsOn []string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s:%s", domainID, crn),
		normalizeResourceName("ibm_cis_domain", true),
		"ibm_cis_domain",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resource
}

func (g CISGenerator) loadDNSRecords(crn, domainID, dnsRecordID string, dependsOn []string) terraformutils.Resource {
	resources := terraformutils.NewResource(
		fmt.Sprintf("%s:%s:%s", dnsRecordID, domainID, crn),
		normalizeResourceName("ibm_cis_dns_record", true),
		"ibm_cis_dns_record",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

func (g CISGenerator) loadFirewall(crn, domainID, fID, fType string, dependsOn []string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s:%s:%s:%s", fType, fID, domainID, crn),
		normalizeResourceName("ibm_cis_firewall", true),
		"ibm_cis_firewall",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resource
}

func (g CISGenerator) loadDomainSettings(crn, dID string, dependsOn []string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s:%s", dID, crn),
		normalizeResourceName("ibm_cis_domain_settings", true),
		"ibm_cis_domain_settings",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resource
}

func (g CISGenerator) loadGlobalBalancer(crn, dID, gID string, dependsOn []string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s:%s:%s", gID, dID, crn),
		normalizeResourceName("ibm_cis_global_load_balancer", true),
		"ibm_cis_global_load_balancer",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})

	// Conflicts with proxied attribute
	resource.IgnoreKeys = append(resource.IgnoreKeys,
		"^ttl$",
	)
	return resource
}

func (g CISGenerator) loadGlobalBalancerPool(crn, pID, pName string, dependsOn []string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s:%s", pID, crn),
		normalizeResourceName(pName, true),
		"ibm_cis_origin_pool",
		g.ProviderName,
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resource
}

func (g CISGenerator) loadGlobalBalancerMonitor(crn, gblmID, port string, dependsOn []string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s:%s", gblmID, crn),
		normalizeResourceName("ibm_cis_healthcheck", true),
		"ibm_cis_healthcheck",
		"ibm",
		map[string]string{
			"port": port,
		},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resource
}

func (g CISGenerator) loadRateLimit(crn, dID, rID string, dependsOn []string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s:%s:%s", rID, dID, crn),
		normalizeResourceName("ibm_cis_rate_limit", true),
		"ibm_cis_rate_limit",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resource
}

func (g CISGenerator) loadEdgeFunctionAction(crn, dID, actionID string, dependsOn []string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s:%s:%s", actionID, dID, crn),
		normalizeResourceName("ibm_cis_edge_functions_action", true),
		"ibm_cis_edge_functions_action",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resource
}

func (g CISGenerator) loadEdgeFunctionTrigger(crn, dID, triggerID string, dependsOn []string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s:%s:%s", triggerID, dID, crn),
		normalizeResourceName("ibm_cis_edge_functions_trigger", true),
		"ibm_cis_edge_functions_trigger",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resource
}

func (g CISGenerator) loadWafRulePackage(crn, dID, pkgID, actionMode, sensitivity string, dependsOn []string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s:%s:%s", pkgID, dID, crn),
		normalizeResourceName("ibm_cis_waf_package", true),
		"ibm_cis_waf_package",
		"ibm",
		map[string]string{
			"action_mode": actionMode,
			"sensitivity": sensitivity,
		},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resource
}

func (g CISGenerator) loadWafGroups(crn, dID, pkgID, grpID string, dependsOn []string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s:%s:%s:%s", grpID, pkgID, dID, crn),
		normalizeResourceName("ibm_cis_waf_group", true),
		"ibm_cis_waf_group",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resource
}

func (g CISGenerator) loadPageRule(crn, dID, ruleID string, dependsOn []string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s:%s:%s", ruleID, dID, crn),
		normalizeResourceName("ibm_cis_page_rule", true),
		"ibm_cis_page_rule",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resource
}

func (g CISGenerator) loadCustomPage(crn, dID, cpID, url string, dependsOn []string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s:%s:%s", cpID, dID, crn),
		normalizeResourceName("ibm_cis_custom_page", true),
		"ibm_cis_custom_page",
		"ibm",
		map[string]string{
			"url": url,
		},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resource
}

func (g CISGenerator) loadRangeApp(crn, dID, appID string, dependsOn []string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s:%s:%s", appID, dID, crn),
		normalizeResourceName("ibm_cis_range_app", true),
		"ibm_cis_range_app",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resource
}

func (g CISGenerator) loadSSLCertificates(crn, dID, cID string, dependsOn []string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s:%s:%s", cID, dID, crn),
		normalizeResourceName("ibm_cis_certificate_order", true),
		"ibm_cis_certificate_order",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resource
}

func (g CISGenerator) loadCISRouting(crn, dID string, dependsOn []string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s:%s", dID, crn),
		normalizeResourceName("ibm_cis_routing", true),
		"ibm_cis_routing",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resource
}

func (g CISGenerator) loadCacheSettings(crn, dID string, dependsOn []string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s:%s", dID, crn),
		normalizeResourceName("ibm_cis_cache_settings", true),
		"ibm_cis_cache_settings",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resource
}

func (g CISGenerator) loadTLSSettings(crn, dID string, dependsOn []string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s:%s", dID, crn),
		normalizeResourceName("ibm_cis_tls_settings", true),
		"ibm_cis_tls_settings",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resource
}

func (g CISGenerator) loadFilters(crn, dID, fID string, dependsOn []string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s:%s:%s", fID, dID, crn),
		normalizeResourceName("ibm_cis_filter", true),
		"ibm_cis_filter",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resource
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
		// Instance
		crn := c.Crn.String()
		g.Resources = append(g.Resources, g.loadInstances(crn, c.Name, c.ResourceGroupID))
		resourceName := g.Resources[len(g.Resources)-1:][0].ResourceName

		var cisDependsOn []string
		cisDependsOn = append(cisDependsOn,
			"ibm_cis."+resourceName)

		// Domain
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

		domainOpts := zService.NewListZonesOptions()
		domainOpts.SetPage(1)       // list all zones in one page
		domainOpts.SetPerPage(1000) // maximum allowed limit is 1000 per page

		zoneList, _, err := zService.ListZones(domainOpts)
		if err != nil {
			return err
		}

		// Origin pool
		gblOpts := &globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
			URL:           DefaultCisURL,
			Authenticator: &core.BearerTokenAuthenticator{BearerToken: bluemixToken},
			Crn:           &crn,
		}

		gblService, err := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0(gblOpts)
		if err != nil {
			return err
		}

		gblPoolList, _, err := gblService.ListAllLoadBalancerPools(&globalloadbalancerpoolsv0.ListAllLoadBalancerPoolsOptions{})
		if err != nil {
			return err
		}

		for _, gbl := range gblPoolList.Result {
			if gbl.ID != nil {
				g.Resources = append(g.Resources, g.loadGlobalBalancerPool(crn, *gbl.ID, *gbl.Name, cisDependsOn))
			}
		}

		// Health Monitor
		gblmOpts := &globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
			URL:           DefaultCisURL,
			Authenticator: &core.BearerTokenAuthenticator{BearerToken: bluemixToken},
			Crn:           &crn,
		}

		gblmService, err := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1(gblmOpts)
		if err != nil {
			return err
		}

		gblmList, _, err := gblmService.ListAllLoadBalancerMonitors(&globalloadbalancermonitorv1.ListAllLoadBalancerMonitorsOptions{})
		if err != nil {
			return err
		}
		for _, gblm := range gblmList.Result {
			if gblm.Port != nil {
				port := strconv.FormatInt(*gblm.Port, 10)
				g.Resources = append(g.Resources, g.loadGlobalBalancerMonitor(crn, *gblm.ID, port, cisDependsOn))
			}
		}

		for _, z := range zoneList.Result {
			var domainDependsOn []string
			domainDependsOn = append(domainDependsOn,
				"ibm_cis."+resourceName)

			g.Resources = append(g.Resources, g.loadDomains(crn, *z.ID, domainDependsOn))
			zoneResourceName := g.Resources[len(g.Resources)-1:][0].ResourceName
			domainDependsOn = append(domainDependsOn,
				"ibm_cis_domain."+zoneResourceName)

			// DNS Record
			zoneID := *z.ID
			dnsOpts := &dnsrecordsv1.DnsRecordsV1Options{
				URL: DefaultCisURL,
				Authenticator: &core.BearerTokenAuthenticator{
					BearerToken: bluemixToken,
				},
				Crn:            &crn,
				ZoneIdentifier: &zoneID,
			}

			// Domain Setting
			g.Resources = append(g.Resources, g.loadDomainSettings(crn, *z.ID, domainDependsOn))

			// DNS Records
			dnsService, err := dnsrecordsv1.NewDnsRecordsV1(dnsOpts)
			if err != nil {
				return err
			}

			dOpts := &dnsrecordsv1.ListAllDnsRecordsOptions{}
			dnsList, _, err := dnsService.ListAllDnsRecords(dOpts)
			if err != nil {
				return err
			}

			// IBM Network CIS WAF Package
			cisWAFPackagesOpt := &wafrulepackagesapiv1.WafRulePackagesApiV1Options{
				URL: DefaultCisURL,
				Authenticator: &core.BearerTokenAuthenticator{
					BearerToken: bluemixToken,
				},
				Crn:    &crn,
				ZoneID: &zoneID,
			}
			cisWAFPackageClient, _ := wafrulepackagesapiv1.NewWafRulePackagesApiV1(cisWAFPackagesOpt)
			wasPkgList, _, err := cisWAFPackageClient.ListWafPackages(&wafrulepackagesapiv1.ListWafPackagesOptions{})
			if err != nil {
				return err
			}

			for _, wafPkg := range wasPkgList.Result {
				cisWAFPackageOpt := &wafrulepackagesapiv1.GetWafPackageOptions{
					PackageID: wafPkg.ID,
				}
				wafPkg, _, err := cisWAFPackageClient.GetWafPackage(cisWAFPackageOpt)
				if err != nil {
					return err
				}

				if wafPkg.Result != nil && wafPkg.Result.ActionMode != nil {
					g.Resources = append(g.Resources, g.loadWafRulePackage(crn, *z.ID, *wafPkg.Result.ID, *wafPkg.Result.ActionMode, *wafPkg.Result.Sensitivity, domainDependsOn))

					// CIS waf-groups
					cisWAFGroupOpt := &wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
						URL: DefaultCisURL,
						Authenticator: &core.BearerTokenAuthenticator{
							BearerToken: bluemixToken,
						},
						Crn:    &crn,
						ZoneID: &zoneID,
					}
					cisWAFGroupClient, _ := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(cisWAFGroupOpt)
					wasGrpList, _, err := cisWAFGroupClient.ListWafRuleGroups(&wafrulegroupsapiv1.ListWafRuleGroupsOptions{
						PkgID: wafPkg.Result.ID,
					})
					if err != nil {
						return err
					}
					for _, wafGrp := range wasGrpList.Result {
						g.Resources = append(g.Resources, g.loadWafGroups(crn, *z.ID, *wafPkg.Result.ID, *wafGrp.ID, domainDependsOn))
					}
				}
			}

			// Rate Limit
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
				g.Resources = append(g.Resources, g.loadRateLimit(crn, *z.ID, *rl.ID, domainDependsOn))
			}

			// Firewall -  Lockdown
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

			for _, f := range firewallList.Result {
				g.Resources = append(g.Resources, g.loadFirewall(crn, *z.ID, *f.ID, "lockdowns", domainDependsOn))
			}

			// Firewall -  AccessRules
			firewallAccessOpts := &zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
				URL: DefaultCisURL,
				Authenticator: &core.BearerTokenAuthenticator{
					BearerToken: bluemixToken,
				},
				Crn:            &crn,
				ZoneIdentifier: &zoneID,
			}

			fAccessService, err := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1(firewallAccessOpts)
			if err != nil {
				return err
			}

			firewalAccesslList, _, err := fAccessService.ListAllZoneAccessRules(&zonefirewallaccessrulesv1.ListAllZoneAccessRulesOptions{})
			if err != nil {
				return err
			}

			for _, f := range firewalAccesslList.Result {
				if f.Configuration.Target != nil {
					g.Resources = append(g.Resources, g.loadFirewall(crn, *z.ID, *f.ID, "access_rules", domainDependsOn))
				}
			}

			// Useragent blocking rules
			firewallUAOpts := &useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
				URL: DefaultCisURL,
				Authenticator: &core.BearerTokenAuthenticator{
					BearerToken: bluemixToken,
				},
				Crn:            &crn,
				ZoneIdentifier: &zoneID,
			}

			fUAService, err := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(firewallUAOpts)
			if err != nil {
				return err
			}

			firewalUAlList, _, err := fUAService.ListAllZoneUserAgentRules(&useragentblockingrulesv1.ListAllZoneUserAgentRulesOptions{})
			if err != nil {
				return err
			}

			for _, f := range firewalUAlList.Result {
				if f.Configuration.Target != nil {
					g.Resources = append(g.Resources, g.loadFirewall(crn, *z.ID, *f.ID, "ua_rules", domainDependsOn))
				}
			}

			// IBM Network CIS Edge Function Action & Triggers
			cisEdgeFunctionOpt := &edgefunctionsapiv1.EdgeFunctionsApiV1Options{
				URL: DefaultCisURL,
				Authenticator: &core.BearerTokenAuthenticator{
					BearerToken: bluemixToken,
				},
				Crn:            &crn,
				ZoneIdentifier: &zoneID,
			}

			cisEdgeFunctionClient, _ := edgefunctionsapiv1.NewEdgeFunctionsApiV1(cisEdgeFunctionOpt)
			edgeActionResonse, _, err := cisEdgeFunctionClient.ListEdgeFunctionsActions(&edgefunctionsapiv1.ListEdgeFunctionsActionsOptions{})
			if err != nil {
				return err
			}

			for _, el := range edgeActionResonse.Result {
				if el.Routes != nil {
					for _, elT := range el.Routes {
						g.Resources = append(g.Resources, g.loadEdgeFunctionAction(crn, *z.ID, *elT.Script, domainDependsOn))
						elResourceName := g.Resources[len(g.Resources)-1:][0].ResourceName
						edgeFunctionActionDependsOn := makeDependsOn(domainDependsOn,
							"ibm_cis_edge_functions_action."+elResourceName)

						g.Resources = append(g.Resources, g.loadEdgeFunctionTrigger(crn, *z.ID, *elT.ID, edgeFunctionActionDependsOn))
					}
				}
			}

			// Range app
			rangeAppOpt := &rangeapplicationsv1.RangeApplicationsV1Options{
				URL: DefaultCisURL,
				Authenticator: &core.BearerTokenAuthenticator{
					BearerToken: bluemixToken,
				},
				Crn:            &crn,
				ZoneIdentifier: &zoneID,
			}

			rangeAppClient, _ := rangeapplicationsv1.NewRangeApplicationsV1(rangeAppOpt)
			ranegAppList, _, err := rangeAppClient.ListRangeApps(&rangeapplicationsv1.ListRangeAppsOptions{})
			if err != nil {
				return err
			}

			for _, r := range ranegAppList.Result {
				g.Resources = append(g.Resources, g.loadRangeApp(crn, *z.ID, *r.ID, domainDependsOn))
			}

			// Page Rules
			pageRueleOpt := &pageruleapiv1.PageRuleApiV1Options{
				URL: DefaultCisURL,
				Authenticator: &core.BearerTokenAuthenticator{
					BearerToken: bluemixToken,
				},
				Crn:    &crn,
				ZoneID: &zoneID,
			}

			pageRuleClient, _ := pageruleapiv1.NewPageRuleApiV1(pageRueleOpt)
			pageRuleList, _, err := pageRuleClient.ListPageRules(&pageruleapiv1.ListPageRulesOptions{})
			if err != nil {
				return err
			}

			for _, p := range pageRuleList.Result {
				g.Resources = append(g.Resources, g.loadPageRule(crn, *z.ID, *p.ID, domainDependsOn))
			}

			// Custom Page
			customPageOpt := &custompagesv1.CustomPagesV1Options{
				URL: DefaultCisURL,
				Authenticator: &core.BearerTokenAuthenticator{
					BearerToken: bluemixToken,
				},
				Crn:            &crn,
				ZoneIdentifier: &zoneID,
			}

			customPageClient, _ := custompagesv1.NewCustomPagesV1(customPageOpt)
			customPageList, _, err := customPageClient.ListInstanceCustomPages(&custompagesv1.ListInstanceCustomPagesOptions{})
			if err != nil {
				return err
			}

			for _, cp := range customPageList.Result {
				if cp.URL != nil {
					g.Resources = append(g.Resources, g.loadCustomPage(crn, *z.ID, *cp.ID, *cp.URL, domainDependsOn))
				}
			}

			// SSL Certificate - order
			sslOpt := &sslcertificateapiv1.SslCertificateApiV1Options{
				URL: DefaultCisURL,
				Authenticator: &core.BearerTokenAuthenticator{
					BearerToken: bluemixToken,
				},
				Crn:            &crn,
				ZoneIdentifier: &zoneID,
			}
			sslClient, err := sslcertificateapiv1.NewSslCertificateApiV1(sslOpt)
			if err != nil {
				return err
			}
			sslList, _, err := sslClient.ListCertificates(&sslcertificateapiv1.ListCertificatesOptions{})
			if err != nil {
				return err
			}
			for _, cert := range sslList.Result {
				g.Resources = append(g.Resources, g.loadSSLCertificates(crn, *z.ID, *cert.ID, domainDependsOn))
			}

			// routingv1
			routingOpt := &routingv1.RoutingV1Options{
				URL: DefaultCisURL,
				Authenticator: &core.BearerTokenAuthenticator{
					BearerToken: bluemixToken,
				},
				Crn:            &crn,
				ZoneIdentifier: &zoneID,
			}

			routingClient, err := routingv1.NewRoutingV1(routingOpt)
			if err != nil {
				return err
			}
			routingList, _, err := routingClient.GetSmartRouting(&routingv1.GetSmartRoutingOptions{})
			if err != nil {
				return err
			}
			if routingList != nil {
				g.Resources = append(g.Resources, g.loadCISRouting(crn, *z.ID, domainDependsOn))
			}

			// Filters
			filterOpts := &filtersv1.FiltersV1Options{
				URL: DefaultCisURL,
				Authenticator: &core.BearerTokenAuthenticator{
					BearerToken: bluemixToken,
				},
			}

			filterClient, err := filtersv1.NewFiltersV1(filterOpts)
			if err != nil {
				return err
			}

			filterList, _, err := filterClient.ListAllFilters(&filtersv1.ListAllFiltersOptions{
				Crn:            &crn,
				ZoneIdentifier: &zoneID,
				XAuthUserToken: &bluemixToken,
			})
			if err != nil {
				return err
			}

			if filterList != nil {
				for _, f := range filterList.Result {
					g.Resources = append(g.Resources, g.loadFilters(crn, *z.ID, *f.ID, domainDependsOn))
				}
			}

			// Cache Settings
			g.Resources = append(g.Resources, g.loadCacheSettings(crn, *z.ID, domainDependsOn))

			// TLS Settings
			g.Resources = append(g.Resources, g.loadTLSSettings(crn, *z.ID, domainDependsOn))

			for _, d := range dnsList.Result {
				g.Resources = append(g.Resources, g.loadDNSRecords(crn, *z.ID, *d.ID, domainDependsOn))
				dnsResourceName := g.Resources[len(g.Resources)-1:][0].ResourceName
				dnsDependsOn := makeDependsOn(domainDependsOn,
					"ibm_cis_dns_record."+dnsResourceName)

				// Global Load Balancer
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

func makeDependsOn(dependsOn []string, resource string) []string {
	return append(dependsOn, resource)
}
