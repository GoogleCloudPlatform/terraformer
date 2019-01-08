package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"waze/terraformer/aws_terraforming"
	"waze/terraformer/terraform_utils"

	"github.com/deckarep/golang-set"
)

var awsRegions = []string{
	"us-east-2",
	"us-east-1",
	"us-west-1",
	"us-west-2",
	"ap-south-1",
	//"ap-northeast-3",
	"ap-northeast-2",
	"ap-southeast-1",
	"ap-southeast-2",
	"ap-northeast-1",
	"ca-central-1",
	//"cn-north-1",
	//"cn-northwest-1",
	"eu-central-1",
	"eu-west-1",
	"eu-west-2",
	"eu-west-3",
	//"eu-north-1",
	"sa-east-1",
}

var awsAccount = []string{"waze"}
var ignoreServicesAws = mapset.NewSetWith(
	"auto_scaling",
	"iam",         //for debug
	"route53",     //for debug
	"s3",          //for debug
	"sg",          //for debug
	"elb",         //for debug
	"elasticache", //for debug
)

var notInfraServiceAws = mapset.NewSetWith(
	"elb",
	"elasticache",
)

func importAWS() {
	resources := []importedResource{}
	for _, account := range awsAccount {
		for _, region := range awsRegions {
			for _, service := range getAWSService() {
				provider := &aws_terraforming.AWSProvider{}
				resources = append(resources, importResource(provider, region, service, region, account))
			}
		}
	}
	provider := &aws_terraforming.AWSProvider{}
	//services := map[string]importedResource{}
	for _, r := range resources {
		rootPath, _ := os.Getwd()
		path := ""
		if notInfraServiceAws.Contains(r.serviceName) {
			continue
			//path = fmt.Sprintf("%s/imported/microservices/%s/", rootPath, r.serviceName)
		} else {
			path = fmt.Sprintf("%s/imported/infra/aws/%s/%s/%s", rootPath, r.project, r.serviceName, r.region)
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

func getAWSService() []string {
	services := []string{}
	provider := &aws_terraforming.AWSProvider{}
	for service := range provider.GetAWSSupportService() {
		if !ignoreServicesAws.Contains(service) {
			services = append(services, service)
		}
	}
	sort.Strings(services)
	return services
}
