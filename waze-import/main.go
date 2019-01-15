package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"waze/terraformer/terraform_utils"

	"github.com/hashicorp/terraform/terraform"
)

/*
├── infra
│   ├── aws
│   │   ├── waze
│   │   │   ├── iam
│   │   │   │   └── global
│   │   │   └── sg
│   │   │       └── us-east1
│   │   └── waze-mapreduce
│   │       ├── iam
│   │       │   └── global
│   │       └── sg
│   │           └── us-east1
│   └── gcp
│       ├── waze-ci
│       │   ├── firewall
│       │   │   ├── europe-west1
│       │   │   └── us-east1
│       │   ├── iam
│       │   │   └── global
│       │   └── subnets
│       │       ├── europe-west1
│       │       └── us-east1
│       ├── waze-development
│       │   ├── firewall
│       │   │   ├── europe-west1
│       │   │   └── us-east1
│       │   ├── iam
│       │   │   └── global
│       │   └── subnets
│       │       ├── europe-west1
│       │       └── us-east1
│       └── waze-prod
*/

type importedService struct {
	tfResources []importedResource
	project     string
	region      string
}

type importedResource struct {
	tfResource  terraform_utils.Resource
	region      string
	cloud       string
	serviceName string
}

// services list from spinnaker applications
var microserviceNameList = []string{
	"abtests",
	"adapt",
	"admanage",
	"adpacing",
	"adreports",
	"ads",
	"adsmanagementapi",
	"advendors",
	"advenuesync",
	"adview",
	"alerts",
	"ap",
	"artifactory",
	"ateam",
	"bake",
	"bash",
	"biscripts",
	"brandsserver",
	"canaryconfigs",
	"carpool",
	"carpoolactivity",
	"carpooladapter",
	"carpoolcommute",
	"carpoolcommutedaily",
	"carpooldispatcher",
	"carpoolgateway",
	"carpoolgroups",
	"carpoolindex",
	"carpoolindexrewrite",
	"carpoolmanager",
	"carpoolmatching",
	"carpoolmatchingdispatcher",
	"carpoolpayments",
	"carpoolpricing",
	"carpoolranking",
	"carpoolreviews",
	"carpoolrouting",
	"carpoolserver",
	"carpooltesting",
	"cars",
	"cartool",
	"chaosmonkeytest1",
	"client",
	"clienttiles",
	"closures",
	"cloudprober",
	"columbus",
	"configuration",
	"containertest",
	"copy",
	"crosstimes",
	"dataflow",
	"dataproc",
	"datasetmanager",
	"dategiver",
	"dbeventwriter",
	"default",
	"deployment",
	"descartes",
	"dockermanager",
	"donald",
	"driveplanner",
	"dtb",
	"elasticsearch",
	"elton",
	"emailserver",
	"engagement",
	"es",
	"esp",
	"eu.static.proxy",
	"event",
	"example",
	"externalpoi",
	"feed",
	"feedelserver",
	"feedserver",
	"fluentd",
	"forum",
	"gagentproxy",
	"gamingserver",
	"gasolina",
	"gateway",
	"gcp",
	"general",
	"geocodingserver",
	"geocodingserveronline",
	"geoindex",
	"georegistry",
	"gke",
	"grpc",
	"heapster",
	"hitchhikes",
	"ifs",
	"ifsimagecreator",
	"il.map.benchmark",
	"inbox",
	"incidents",
	"indexserver",
	"internal",
	"internaltools",
	"issh",
	"jobflowscheduler",
	"jumpserver",
	"kube",
	"kubernetes",
	"l7",
	"loganalysis",
	"login",
	"logserveranalysis",
	"maintjobs",
	"mapcomments",
	"mapnik",
	"mapnikserver",
	"mapproblems",
	"memcached",
	"memcachetest",
	"merger",
	"mergerpii",
	"metadata",
	"metrics",
	"microservice",
	"monitoring",
	"mte",
	"mtevendors",
	"nvidia",
	"osupgrades",
	"p2pproxy",
	"packages",
	"parkingserver",
	"permits",
	"pickups",
	"piitool",
	"pointsprocessor",
	"pointsserver",
	"preferences",
	"prompto",
	"pushserver",
	"realtime",
	"realtimeproxy",
	"redis",
	"redit",
	"regressionchecker",
	"remoteproxy",
	"repository",
	"repositoryprocessor",
	"reverseproxy",
	"routingserver",
	"sa",
	"sampleapp1",
	"saw",
	"scheduler",
	"searchserver",
	"servermanager",
	"sessionprocessor",
	"socialmediaserver",
	"sparta",
	"spartatrends",
	"spin",
	"spinnaker",
	"staging",
	"storagegateway",
	"stress",
	"supportool",
	"test",
	"testtest",
	"testtest2",
	"tile",
	"tilesbuilder",
	"topic",
	"topiceventsdistributor",
	"topicforwarder",
	"topicforwardereceiver",
	"topicforwardersender",
	"topicmonitoring",
	"topicserver",
	"topicwriter",
	"traceutility",
	"trip",
	"tts",
	"ttsgateway",
	"usage",
	"useractions",
	"usercredibility",
	"userjourney",
	"userrouting",
	"usersprofile",
	"usersserver",
	"usertracking",
	"v1",
	"v2",
	"venues",
	"voiceprompts",
	"was",
	"wazebischeduler",
	"weave",
	"weblogin",
	"wiki",
	"wudprocessor",
	"youyinglcodelab",
	"tiles",
}
var runOnService = ""
var runOnRegion = ""

var filters = "cassandra"

func main() {
	cloud := ""
	if len(os.Args) > 1 {
		cloud = os.Args[1]
	}
	if len(os.Args) > 2 {
		runOnService = os.Args[2]
	}
	if len(os.Args) > 3 {
		runOnRegion = os.Args[3]
	}
	switch cloud {
	case "aws":
		importAWS()
	case "google":
		importGCP()
	default:
		importAWS()
		importGCP()
	}

}

func rootPath() string {
	//rootPath, _ := os.Getwd()
	rootPath := "/Users/sergeylanz/code/terraform" // for debug
	return rootPath
}

func importResource(provider terraform_utils.ProviderGenerator, service, region, project string) []terraform_utils.Resource {
	log.Println(service, region, project)
	err := provider.Init([]string{region, project})
	if err != nil {
		log.Fatal(err)
	}

	err = provider.InitService(service)
	if err != nil {
		log.Fatal(err)
	}

	err = provider.GetService().InitResources()
	if err != nil {
		log.Fatal(err)
	}
	refreshedResources, err := terraform_utils.RefreshResources(provider.GetService().GetResources(), provider.GetName())
	if err != nil {
		log.Fatal(err)
	}
	provider.GetService().SetResources(refreshedResources)
	for i := range provider.GetService().GetResources() {
		provider.GetService().GetResources()[i].ConvertTFstate()
	}
	err = provider.GetService().PostConvertHook()
	if err != nil {
		log.Fatal(err)
	}
	return provider.GetService().GetResources()
}

func connectServices(importResources map[string]importedService, resourceConnections map[string]map[string][]string) map[string]importedService {
	for resource, connection := range resourceConnections {
		if _, exist := importResources[resource]; exist {
			for k, v := range connection {
				if cc, ok := importResources[k]; ok {
					for _, ccc := range cc.tfResources {
						for i := range importResources[resource].tfResources {
							if ccc.region != importResources[resource].tfResources[i].region {
								continue
							}
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
	return importResources
}

func generateFilesAndUploadState(importResources map[string]importedService, importer providerImporter) {
	for serviceName, r := range importResources {
		rootPath := rootPath()
		path := ""
		regionsMapping := map[string][]terraform_utils.Resource{}
		for _, resource := range r.tfResources {
			if _, exist := regionsMapping[resource.region]; !exist {
				regionsMapping[resource.region] = []terraform_utils.Resource{}
			}
			regionsMapping[resource.region] = append(regionsMapping[resource.region], resource.tfResource)
		}
		for region, resources := range regionsMapping {
			if importer.getNotInfraService().Contains(serviceName) {
				continue
				//path = fmt.Sprintf("%s/imported/microservices/%s/", rootPath, serviceName)
			} else {
				path = fmt.Sprintf("%s/infra/%s/%s/%s/%s", rootPath, importer.getName(), importer.getAccount(), serviceName, region)
			}
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				log.Fatal(err)
				return
			}

			printHclFiles(resources, importer.getResourceConnections(), path, serviceName, importer.getProviderData(r.project, region))
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

func printHclFiles(resources []terraform_utils.Resource, resourceConnections map[string]map[string][]string, path, serviceName string, providerData map[string]interface{}) {
	// create provider file
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
		variables := map[string]map[string]map[string]interface{}{}
		variables["data"] = map[string]map[string]interface{}{}
		variables["data"]["terraform_remote_state"] = map[string]interface{}{}
		for k := range resourceConnections[serviceName] {
			variables["data"]["terraform_remote_state"][k] = map[string]interface{}{
				"backend": "gcs",
				"config": map[string]interface{}{
					"bucket": bucketStateName,
					"prefix": bucketPrefix(strings.Replace(path, serviceName, k, -1)),
				},
			}
		}
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

func printFile(path string, data []byte) {
	err := ioutil.WriteFile(path, data, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}
}
