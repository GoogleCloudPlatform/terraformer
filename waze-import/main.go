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

type importedResource struct {
	tfResources []terraform_utils.Resource
	tfState     []byte
	project     string
	region      string
	serviceName string
}

func main() {
	//importAWS()
	importGCP()
}

func importResource(provider terraform_utils.ProviderGenerator, regionName, service, region, project string) importedResource {
	log.Println(service, region)
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

	// create tfstate
	tfStateFile, err := terraform_utils.PrintTfState(provider.GetService().GetResources())
	if err != nil {
		log.Fatal(err)
	}
	// convert InstanceState to go struct for hcl print
	for i := range provider.GetService().GetResources() {
		provider.GetService().GetResources()[i].ConvertTFstate()
	}
	// change structs with additional data for each resource
	err = provider.GetService().PostConvertHook()
	if err != nil {
		log.Fatal(err)
	}

	return importedResource{
		tfResources: provider.GetService().GetResources(),
		tfState:     tfStateFile,
		project:     project,
		region:      regionName,
		serviceName: service,
	}
}
