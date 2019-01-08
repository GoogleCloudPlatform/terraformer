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

	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

//var GCPProjects = []string{"waze-development", "waze-prod"}
var GCPProjects = []string{"waze-development"}

type importedResource struct {
	tfResources []terraform_utils.Resource
	tfState     []byte
	project     string
	region      string
	serviceName string
}

func importGCP() {
	resources := []importedResource{}
	for _, project := range GCPProjects {
		for _, zone := range getGCPZone() {
			for _, service := range getGCPService() {
				provider := &gcp_terraforming.GCPProvider{}
				err := provider.Init([]string{zone.Name, project})
				if err != nil {
					log.Fatal(err)
					return
				}

				err = provider.InitService(service)
				if err != nil {
					log.Fatal(err)
					return
				}

				err = provider.GetService().InitResources()
				if err != nil {
					log.Fatal(err)
					return
				}
				refreshedResources, err := terraform_utils.RefreshResources(provider.GetService().GetResources(), provider.GetName())
				if err != nil {
					log.Fatal(err)
					return
				}
				provider.GetService().SetResources(refreshedResources)

				// create tfstate
				tfStateFile, err := terraform_utils.PrintTfState(provider.GetService().GetResources())
				if err != nil {
					log.Fatal(err)
					return
				}
				// convert InstanceState to go struct for hcl print
				for i := range provider.GetService().GetResources() {
					provider.GetService().GetResources()[i].ConvertTFstate()
				}
				// change structs with additional data for each resource
				err = provider.GetService().PostConvertHook()
				if err != nil {
					log.Fatal(err)
					return
				}
				regionPath := strings.Split(zone.Region, "/")
				regionName := regionPath[len(regionPath)-1]
				if true { // check if service global
					regionName = "global" // TODO for debug
				}
				resources = append(resources, importedResource{
					tfResources: provider.GetService().GetResources(),
					tfState:     tfStateFile,
					project:     project,
					region:      regionName,
					serviceName: service,
				})
			}
		}
	}
	provider := &gcp_terraforming.GCPProvider{}
	for _, r := range resources {
		rootPath, _ := os.Getwd()
		path := fmt.Sprintf("%s/imported/infra/gcp/%s/%s/%s", rootPath, r.project, r.serviceName, r.region)
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
				if strings.Contains(zone.Region, "europe-west1") { // TODO for debug
					zones = append(zones, zone)
				}

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
		if service != "healthChecks" { // TODO for debug
			continue
		}
		services = append(services, service)
	}
	sort.Strings(services)
	return services
}
