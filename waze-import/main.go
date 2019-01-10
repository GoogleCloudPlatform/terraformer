package main

import (
	"log"
	"waze/terraformer/terraform_utils"
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

func main() {
	//importAWS()
	importGCP()
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
