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
	"github.com/IBM/networking-go-sdk/globalloadbalancermonitorv1"
	"github.com/IBM/networking-go-sdk/globalloadbalancerpoolsv0"
	"github.com/IBM/networking-go-sdk/globalloadbalancerv1"
	"github.com/IBM/networking-go-sdk/zonelockdownv1"
	"github.com/IBM/networking-go-sdk/zoneratelimitsv1"
	"github.com/IBM/networking-go-sdk/zonesv1"
)

// CISGenerator ..
type CISGenerator struct {
	IBMService
}

func (g CISGenerator) loadInstances(crn, name string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		crn,
		name,
		"ibm_cis",
		g.ProviderName,
		[]string{})
	return resources
}

func (g CISGenerator) loadDomains(crn, domainID string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		fmt.Sprintf("%s:%s", domainID, crn),
		domainID,
		"ibm_cis_domain",
		g.ProviderName,
		[]string{})
	return resources
}

func (g CISGenerator) loadDNSRecords(crn, domainID, dnsRecordID string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		fmt.Sprintf("%s:%s:%s", dnsRecordID, domainID, crn),
		dnsRecordID,
		"ibm_cis_dns_record",
		g.ProviderName,
		[]string{})
	return resources
}

func (g CISGenerator) loadFirewallLockdown(crn, domainID, fID, fType string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		fmt.Sprintf("%s:%s:%s:%s", fType, fID, domainID, crn),
		fID,
		"ibm_cis_firewall",
		g.ProviderName,
		[]string{})
	return resources
}

func (g CISGenerator) loadDomainSettings(crn, dID string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		fmt.Sprintf("%s:%s", dID, crn),
		dID,
		"ibm_cis_domain_settings",
		g.ProviderName,
		[]string{})
	return resources
}

func (g CISGenerator) loadGlobalBalancer(crn, dID, gID string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		fmt.Sprintf("%s:%s:%s", gID, dID, crn),
		dID,
		"ibm_cis_global_load_balancer",
		g.ProviderName,
		[]string{})
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

func (g CISGenerator) loadGlobalBalancerMonitor(crn, gblmID string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		fmt.Sprintf("%s:%s", gblmID, crn),
		gblmID,
		"ibm_cis_healthcheck",
		g.ProviderName,
		[]string{})
	return resources
}

func (g CISGenerator) loadRateLimit(crn, dID, rID string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		fmt.Sprintf("%s:%s:%s", rID, dID, crn),
		rID,
		"ibm_cis_rate_limit",
		g.ProviderName,
		[]string{})
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
		g.Resources = append(g.Resources, g.loadInstances(crn, c.Name))

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

		for _, z := range zoneList.Result {
			g.Resources = append(g.Resources, g.loadDomains(crn, *z.ID))

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

			dnsService, err := dnsrecordsv1.NewDnsRecordsV1(dnsOpts)
			if err != nil {
				return err
			}

			dOpts := &dnsrecordsv1.ListAllDnsRecordsOptions{}
			dnsList, _, err := dnsService.ListAllDnsRecords(dOpts)
			if err != nil {
				return err
			}

			for _, d := range dnsList.Result {
				g.Resources = append(g.Resources, g.loadDNSRecords(crn, *z.ID, *d.ID))
			}

			//Domain Setting
			g.Resources = append(g.Resources, g.loadDomainSettings(crn, *z.ID))

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
				g.Resources = append(g.Resources, g.loadGlobalBalancer(crn, *z.ID, *gb.ID))
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

			for _, f := range firewallList.Result {
				g.Resources = append(g.Resources, g.loadFirewallLockdown(crn, *z.ID, *f.ID, "lockdowns"))
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
				g.Resources = append(g.Resources, g.loadGlobalBalancerMonitor(crn, *gblm.ID))
			}

			//GlobaloadBalancer Pool
			gblPoolOpts := &globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
				URL:           DefaultCisURL,
				Authenticator: &core.BearerTokenAuthenticator{BearerToken: bluemixToken},
				Crn:           &crn,
			}

			gblpService, _ := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0(gblPoolOpts)
			gblpList, _, err := gblpService.ListAllLoadBalancerPools(&globalloadbalancerpoolsv0.ListAllLoadBalancerPoolsOptions{})
			if err != nil {
				return err
			}

			for _, gblp := range gblpList.Result {
				g.Resources = append(g.Resources, g.loadGlobalBalancerPool(crn, *gblp.ID))
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
				g.Resources = append(g.Resources, g.loadRateLimit(crn, *z.ID, *rl.ID))
			}
		}

	}

	return nil
}
