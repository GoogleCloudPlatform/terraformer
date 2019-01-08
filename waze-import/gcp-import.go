package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"waze/terraformer/gcp_terraforming"
	"waze/terraformer/terraform_utils"

	"github.com/deckarep/golang-set"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

//var GCPProjects = []string{"waze-development", "waze-prod"}
var GCPProjects = []string{"waze-development"}

var regionServicesGcp = mapset.NewSetWith(
	"disks",
	"autoscalers",
	"instanceGroupManagers",
	"instances",
	"instanceGroups",
	"regionAutoscalers",
	"regionBackendServices",
	"regionDisks",
	"regionInstanceGroupManagers",
	"subnetworks",
	"addresses",
	"routers",
	"vpnTunnels",
	"forwardingRules",
)

var ignoreServicesGcp = mapset.NewSetWith(
	"disks",
	"iam",
	"autoscalers",
	"instanceGroupManagers",
	"instances",
	"instanceGroups",
	"regionAutoscalers",
	"regionDisks",
	"regionInstanceGroupManagers",
	"instanceTemplates",
	"images",
	"addresses",
)

var notInfraServiceGcp = mapset.NewSetWith(
	"backendServices",
	"urlMaps",
	"targetHttpProxies",
	"targetHttpsProxies",
	"targetSslProxies",
	"targetTcpProxies",
	"globalForwardingRules",
	"forwardingRules",
	"healthChecks",
	"httpHealthChecks",
	"httpsHealthChecks",
)

func importGCP() {
	resources := []importedResource{}
	for _, project := range GCPProjects {
		for _, service := range getGCPService() {
			zones := []*compute.Zone{{Name: "europe-west1-b", Region: "europe-west1"}} //dummy region
			if regionServicesGcp.Contains(service) {
				zones = getGCPZone()
			}
			for _, zone := range zones {
				provider := &gcp_terraforming.GCPProvider{}
				regionName := "global"
				if regionServicesGcp.Contains(service) {
					regionPath := strings.Split(zone.Name, "/")
					regionName = regionPath[len(regionPath)-1]
				}
				resources = append(resources, importResource(provider, regionName, service, zone.Name, project))
			}
		}
	}
	provider := &gcp_terraforming.GCPProvider{}
	//services := map[string]importedResource{}
	for _, r := range resources {
		rootPath, _ := os.Getwd()
		path := ""
		if notInfraServiceGcp.Contains(r.serviceName) {
			continue
			//path = fmt.Sprintf("%s/imported/microservices/%s/", rootPath, r.serviceName)
		} else {
			path = fmt.Sprintf("%s/imported/infra/gcp/%s/%s/%s", rootPath, r.project, r.serviceName, r.region)
		}
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			log.Fatal(err)
			return
		}
		tfFile, err := terraform_utils.HclPrint(r.tfResources, provider.RegionResource())
		err = ioutil.WriteFile(path+"/"+r.serviceName+".tf", tfFile, os.ModePerm)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = ioutil.WriteFile(path+"/terraform.tfstate", r.tfState, os.ModePerm) //TODO copy to bucket
		if err != nil {
			log.Fatal(err)
			return
		}
	}

}

func getGCPZone() []*compute.Zone {
	ctx := context.Background()
	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}
	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}
	zones := []*compute.Zone{}
	for _, project := range GCPProjects {
		req := computeService.Zones.List(project)
		if err := req.Pages(ctx, func(page *compute.ZoneList) error {
			for _, zone := range page.Items {
				//if strings.Contains(zone.Region, "europe-west1") { // TODO for debug
				zones = append(zones, zone)
				//}
			}
			return nil
		}); err != nil {
			log.Fatal(err)
		}
	}
	return zones
}

func getGCPService() []string {
	services := []string{}
	provider := &gcp_terraforming.GCPProvider{}
	for service := range provider.GetGCPSupportService() {
		if !ignoreServicesGcp.Contains(service) {
			services = append(services, service)
		}
	}
	sort.Strings(services)
	return services
}
