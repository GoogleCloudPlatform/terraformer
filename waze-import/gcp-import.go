package main

import (
	"context"
	"log"
	"sort"
	"strings"
	"waze/terraformer/gcp_terraforming"

	"golang.org/x/oauth2/google"

	"github.com/deckarep/golang-set"
	"google.golang.org/api/compute/v1"
)

//var GCPProjects = []string{"waze-development", "waze-prod"}
var GCPProjects = []string{ "waze-development"}

const gcpProviderVersion = "~>2.0.0"

type gcpImporter struct {
	project string
	name    string
}

func (gcpImporter) getIgnoreService() mapset.Set {
	return mapset.NewSetWith(
		"disks",
		"iam",
		"autoscalers",
		"instanceGroupManagers",
		"instances",
		"instanceGroups",
		"regionAutoscalers",
		"regionDisks",
		"regionInstanceGroupManagers",
		"regionAutoscalers",
		"instanceTemplates",
		"images",
		"addresses",
		"regionBackendServices",
		"backendServices",
		"healthChecks", //google_compute_http_health_check is a legacy health check https://www.terraform.io/docs/providers/google/r/compute_http_health_check.html
	)
}
func (gcpImporter) getRegionServices() mapset.Set {
	return mapset.NewSetWith(
		"disks",
		"autoscalers",
		"instanceGroupManagers",
		"instances",
		"instanceGroups",
	)
}

func (gcpImporter) getNotInfraService() mapset.Set {
	return mapset.NewSetWith(
		"urlMaps",
		"targetHttpProxies",
		"targetHttpsProxies",
		"targetSslProxies",
		"targetTcpProxies",
		"globalForwardingRules",
		"forwardingRules",
		"httpHealthChecks",
		"httpsHealthChecks",
	)
}

func (g gcpImporter) getProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			g.name: map[string]interface{}{
				"project": g.project,
				"version": gcpProviderVersion,
			},
		},
	}
}

func (gcpImporter) getResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		"firewalls":             {"networks": []string{"network", "self_link"}},
		"routes":                {"networks": []string{"network", "self_link"}},
		"regionBackendServices": {"healthChecks": []string{"health_checks", "self_link"}},
		"backendBuckets":        {"gcs": []string{"bucket_name", "name"}},
	}
}

func (g gcpImporter) getAccount() string {
	return g.project
}

func (g gcpImporter) getName() string {
	return g.name
}

func (g gcpImporter) getGcpZonesForService(service string) []*compute.Zone {
	zones := []*compute.Zone{{Name: "europe-west1-b", Region: "europe-west1"}} //dummy region
	if g.getRegionServices().Contains(service) {
		zones = g.getZone()
	}
	return zones
}

func importGCP() {
	importResources := map[string]importedService{}
	for _, project := range GCPProjects {
		importer := gcpImporter{
			name:    "google",
			project: project,
		}
		resources := []importedResource{}
		for _, service := range importer.getService() {
			zones := importer.getGcpZonesForService(service)
			for _, zone := range zones {
				provider := &gcp_terraforming.GCPProvider{}
				for _, r := range importResource(provider, service, zone.Name, project) {
					if strings.Contains(r.ResourceName, filters) {
						continue
					}
					delete(r.Item, "project")
					resources = append(resources, importedResource{
						region:      zone.Name,
						tfResource:  r,
						serviceName: service,
					})
				}
			}
		}

		for _, service := range importer.getService() {
			ir := importedService{}
			for _, r := range resources {
				if r.serviceName == service {
					if importer.getRegionServices().Contains(service) {
						regionPath := strings.Split(r.region, "/")
						ir.region = regionPath[len(regionPath)-1]
					} else {
						ir.region = "global"
						r.region = "global"
					}
					ir.tfResources = append(ir.tfResources, r)
				}
			}
			importResources[service] = ir
		}

		importResources = connectServices(importResources, importer.getResourceConnections())

		/*for _, microserviceName := range microserviceNameList {
			for cloudServiceName, value := range importResources {
				if notInfraServiceGcp.Contains(cloudServiceName) {
					continue
				}
				for _, obj := range value.tfResources {
					resourceName := strings.Replace(obj.tfResource.ResourceName, "_", "-", -1)
					ObjNamePrefix := strings.Split(resourceName, "-")[0]
					if ObjNamePrefix == microserviceName {
						log.Println(microserviceName, cloudServiceName)
					}
				}
			}
		}*/

		generateFilesAndUploadState(importResources, importer)

	}
}

func (g gcpImporter) getZone() []*compute.Zone {
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

func (g gcpImporter) getService() []string {
	services := []string{}
	provider := &gcp_terraforming.GCPProvider{}
	for service := range provider.GetGCPSupportService() {
		if !g.getIgnoreService().Contains(service) {
			services = append(services, service)
		}
	}
	sort.Strings(services)
	return services
}
