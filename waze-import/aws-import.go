package main

import (
	"sort"
	"strings"
	"waze/terraformer/aws_terraforming"

	"github.com/deckarep/golang-set"
)

var awsRegions = []string{
	"us-east-2",
	"us-east-1",
	"us-west-1",
	"us-west-2",
	"ap-south-1",
	"ap-northeast-2",
	"ap-southeast-1",
	"ap-southeast-2",
	"ap-northeast-1",
	"ca-central-1",
	"eu-central-1",
	"eu-west-1",
	"eu-west-2",
	"eu-west-3",
	"sa-east-1",
}

const awsProviderVersion = "~>1.55.0"

var awsAccount = []string{"waze"}

type awsImporter struct {
	name    string
	account string
	region  string
}

func (g awsImporter) getIgnoreService() mapset.Set {
	return mapset.NewSetWith(
		"auto_scaling",
		"iam",     //for debug
		"route53", //for debug
		//"s3",          //for debug
		//"sg",  //for debug
		"elb", //for debug
		//"elasticache", //for debug
	)
}

func (g awsImporter) getGlobalService() mapset.Set {
	return mapset.NewSetWith(
		"iam",
	)
}

func (g awsImporter) getProviderData(arg ...string) map[string]interface{} {
	d := map[string]interface{}{
		"provider": map[string]interface{}{
			"aws":     map[string]interface{}{},
			"version": awsProviderVersion,
		},
	}
	if arg[1] != "global" {
		d["provider"].(map[string]interface{})["aws"].(map[string]interface{})["region"] = arg[1]
	}
	return d
}

func (awsImporter) getResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		"subnet":         {"vpc": []string{"vpc_id", "id"}},
		"vpn_gateway":    {"vpc": []string{"vpc_id", "id"}},
		"vpn_connection": {"vpn_gateway": []string{"vpn_gateway_id", "id"}},
		"rds": {
			"subnet": []string{"subnet_ids", "id"},
			"sg":     []string{"vpc_security_group_ids", "id"},
		},
		"nacl": {
			"subnet": []string{"subnet_ids", "id"},
			"vpc":    []string{"vpc_id", "id"},
		},
		"igw": {"vpc": []string{"vpc_id", "id"}},
		"elasticache": {
			"vpc":    []string{"vpc_id", "id"},
			"subnet": []string{"subnet_ids", "id"},
			"sg":     []string{"security_group_ids", "id"},
		},
	}
}

func (awsImporter) getNotInfraService() mapset.Set {
	return mapset.NewSetWith(
		"elb",
		//"elasticache",
	)
}

func (g awsImporter) getAccount() string {
	return g.account
}

func (g awsImporter) getName() string {
	return g.name
}

func importAWS() {
	importResources := map[string]importedService{}
	resources := []importedResource{}
	for _, account := range awsAccount {
		importer := awsImporter{
			name:    "aws",
			account: account,
		}
		for _, region := range awsRegions {
			for _, service := range importer.getService() {
				provider := &aws_terraforming.AWSProvider{}
				for _, r := range importResource(provider, service, region, account) {
					resources = append(resources, importedResource{
						region:      region,
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
					if !importer.getGlobalService().Contains(service) {
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

		generateFilesAndUploadState(importResources, importer)
	}

}

func (g awsImporter) getService() []string {
	services := []string{}
	provider := &aws_terraforming.AWSProvider{}
	for service := range provider.GetAWSSupportService() {
		if !g.getIgnoreService().Contains(service) {
			services = append(services, service)
		}
	}
	sort.Strings(services)
	return services
}
