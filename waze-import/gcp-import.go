package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"waze/terraformer/gcp_terraforming"
	"waze/terraformer/terraform_utils"

	"github.com/hashicorp/terraform/terraform"

	"golang.org/x/oauth2/google"

	"github.com/deckarep/golang-set"
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
	//"addresses",
	"cloudsql",
)

var notInfraServiceGcp = mapset.NewSetWith(
	"backendServices",
	//"regionBackendServices",
	"urlMaps",
	"targetHttpProxies",
	"targetHttpsProxies",
	"targetSslProxies",
	"targetTcpProxies",
	"globalForwardingRules",
	"forwardingRules",
	//"healthChecks",
	//"httpHealthChecks",
	//"httpsHealthChecks",
)

var resourceConnections = map[string]map[string][]string{
	"firewalls":             {"networks": []string{"network", "self_link"}},
	"routes":                {"networks": []string{"network", "self_link"}},
	"regionBackendServices": {"healthChecks": []string{"health_checks", "self_link"}},
	"backendBuckets":        {"gcs": []string{"bucket_name", "name"}},
}

func importGCP() {
	importResources := map[string]importedService{}

	for _, project := range GCPProjects {
		resources := []importedResource{}
		for _, service := range getGCPService() {
			zones := []*compute.Zone{{Name: "europe-west1-b", Region: "europe-west1"}} //dummy region
			if regionServicesGcp.Contains(service) {
				zones = getGCPZone()
			}
			for _, zone := range zones {
				provider := &gcp_terraforming.GCPProvider{}
				for _, r := range importResource(provider, service, zone.Name, project) {
					delete(r.Item, "project")
					resources = append(resources, importedResource{
						region:      zone.Name,
						tfResource:  r,
						serviceName: service,
					})
				}
			}
		}

		for _, service := range getGCPService() {
			ir := importedService{}
			serviceRegion := ""
			for _, r := range resources {
				if r.serviceName == service {
					serviceRegion = r.region
					ir.tfResources = append(ir.tfResources, r)
				}
			}
			ir.region = "global"
			if regionServicesGcp.Contains(service) {
				regionPath := strings.Split(serviceRegion, "/")
				ir.region = regionPath[len(regionPath)-1]
			}
			importResources[service] = ir
		}

		for resource, connection := range resourceConnections {
			if _, exist := importResources[resource]; exist {
				for k, v := range connection {
					if cc, ok := importResources[k]; ok {
						for _, ccc := range cc.tfResources {
							for i := range importResources[resource].tfResources {
								key := v[1]
								if v[1] == "self_link" || v[1] == "id" {
									key = ccc.tfResource.GetIDKey()
								}
								keyValue := ccc.tfResource.InstanceInfo.Type + "_" + ccc.tfResource.ResourceName + "_" + key
								linkValue := "${data.terraform_remote_state." + k + "." + keyValue + "}"

								tfResource := importResources[resource].tfResources[i].tfResource
								if ccc.tfResource.InstanceState.Attributes[key] == tfResource.InstanceState.Attributes[v[0]] {
									importResources[resource].tfResources[i].tfResource.InstanceState.Attributes[v[0]] = linkValue
									importResources[resource].tfResources[i].tfResource.Item[v[0]] = linkValue
								} else {
									for keyAttributes, j := range tfResource.InstanceState.Attributes {
										match, err := regexp.MatchString(v[0]+".\\d+$", keyAttributes)
										if match && err == nil {
											if j == ccc.tfResource.InstanceState.Attributes[key] {
												importResources[resource].tfResources[i].tfResource.InstanceState.Attributes[keyAttributes] = linkValue
												switch ar := tfResource.Item[v[0]].(type) {
												case []interface{}:
													for j, l := range ar {
														if l == ccc.tfResource.InstanceState.Attributes[key] {
															importResources[resource].tfResources[i].tfResource.Item[v[0]].([]interface{})[j] = linkValue
														}
													}
												default:
													log.Println("type not supported", ar)
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

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

		for serviceName, r := range importResources {
			rootPath, _ := os.Getwd()
			path := ""
			if notInfraServiceGcp.Contains(serviceName) {
				continue
				//path = fmt.Sprintf("%s/imported/microservices/%s/", rootPath, serviceName)
			} else {
				path = fmt.Sprintf("%s/imported/infra/gcp/%s/%s/%s", rootPath, project, serviceName, r.region)
			}
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				log.Fatal(err)
				return
			}
			resources := []terraform_utils.Resource{}
			for _, resource := range r.tfResources {
				resources = append(resources, resource.tfResource)
			}
			printHclFiles(resources, path, project, serviceName)
			tfStateFile, err := terraform_utils.PrintTfState(resources)
			if err != nil {
				log.Fatal(err)
				return
			}
			err = bucketUpload(path, tfStateFile)
			if err != nil {
				log.Fatal(err)
				return
			}
		}
	}
}

func printHclFiles(resources []terraform_utils.Resource, path, projectName, serviceName string) {
	// create provider file
	providerData := map[string]interface{}{
		"provider": map[string]interface{}{
			"google": map[string]interface{}{
				"project": projectName,
			},
		},
	}
	providerDataFile, err := terraform_utils.HclPrint(providerData)
	if err != nil {
		log.Fatal(err)
		return
	}
	printFile(path+"/provider.tf", providerDataFile)

	// create bucket file
	bucketStateDataFile, err := terraform_utils.HclPrint(bucketGetTfData(path))
	printFile(path+"/bucket.tf", bucketStateDataFile)
	// create outputs files
	outputs := map[string]interface{}{}
	outputsByResource := map[string]map[string]interface{}{}

	for i, r := range resources {
		outputState := map[string]*terraform.OutputState{}
		outputsByResource[r.InstanceInfo.Type+"_"+r.ResourceName+"_"+r.GetIDKey()] = map[string]interface{}{
			"value": "${" + r.InstanceInfo.Type + "." + r.ResourceName + "." + r.GetIDKey() + "}",
		}
		outputState[r.InstanceInfo.Type+"_"+r.ResourceName+"_"+r.GetIDKey()] = &terraform.OutputState{
			Type:  "string",
			Value: r.InstanceState.Attributes[r.GetIDKey()],
		}
		for _, v := range resourceConnections {
			for k, ids := range v {
				if k == serviceName {
					if _, exist := r.InstanceState.Attributes[ids[1]]; exist {
						key := ids[1]
						if ids[1] == "self_link" || ids[1] == "id" {
							key = r.GetIDKey()
						}
						linkKey := r.InstanceInfo.Type + "_" + r.ResourceName + "_" + key
						outputsByResource[linkKey] = map[string]interface{}{
							"value": "${" + r.InstanceInfo.Type + "." + r.ResourceName + "." + key + "}",
						}
						outputState[linkKey] = &terraform.OutputState{
							Type:  "string",
							Value: r.InstanceState.Attributes[ids[1]],
						}
					}
				}
			}
		}
		resources[i].Outputs = outputState
	}
	if len(outputsByResource) > 0 {
		outputs["output"] = outputsByResource
		outputsFile, err := terraform_utils.HclPrint(outputs)
		if err != nil {
			log.Fatal(err)
			return
		}
		printFile(path+"/outputs.tf", outputsFile)
	}

	// create variables file
	if len(resourceConnections[serviceName]) > 0 {
		variables := map[string]interface{}{}
		variablesByResource := map[string]map[string]interface{}{}
		for k := range resourceConnections[serviceName] {
			variablesByResource["terraform_remote_state"] = map[string]interface{}{
				k: map[string]interface{}{
					"backend": "gcs",
					"config": map[string]interface{}{
						"bucket": bucketStateName,
						"prefix": bucketPrefix(strings.Replace(path, serviceName, k, -1)),
					},
				},
			}
		}
		variables["data"] = variablesByResource
		variablesFile, err := terraform_utils.HclPrint(variables)
		if err != nil {
			log.Fatal(err)
			return
		}
		printFile(path+"/variables.tf", variablesFile)
	}
	// group by resource by type
	typeOfServices := map[string][]terraform_utils.Resource{}
	for _, r := range resources {
		typeOfServices[r.InstanceInfo.Type] = append(typeOfServices[r.InstanceInfo.Type], r)
	}
	for k, v := range typeOfServices {
		tfFile, err := terraform_utils.HclPrintResource(v, map[string]interface{}{})
		if err != nil {
			log.Fatal(err)
			return
		}
		fileName := strings.Replace(k, strings.Split(k, "_")[0]+"_", "", -1)
		err = ioutil.WriteFile(path+"/"+fileName+".tf", tfFile, os.ModePerm)
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

func printFile(path string, data []byte) {
	err := ioutil.WriteFile(path, data, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}
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
