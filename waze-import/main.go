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
	err = provider.GetService().PostConvertHook()
	if err != nil {
		log.Fatal(err)
	}
	return provider.GetService().GetResources()
}
