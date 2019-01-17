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

const awsProviderVersion = "~>1.56.0"
const terraformTagName = "terraform"

var awsAccount = []string{"waze"}

type awsImporter struct {
	name    string
	account string
	region  string
}

func (g awsImporter) getIgnoreService() mapset.Set {
	return mapset.NewSetWith(
		"auto_scaling",
	)
}

func (g awsImporter) getGlobalService() mapset.Set {
	return mapset.NewSetWith(
		"iam",
		"route53",
	)
}

func (awsImporter) getNotInfraService() mapset.Set {
	return mapset.NewSetWith(
		"elb",
		//"elasticache",
	)
}

func (g awsImporter) getProviderData(arg ...string) map[string]interface{} {
	d := map[string]interface{}{
		"provider": map[string]interface{}{
			"aws": map[string]interface{}{
				"version": awsProviderVersion,
			},
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
			if runOnRegion != "" && region != runOnRegion {
				continue
			}
			for _, service := range importer.getService() {
				if importer.getGlobalService().Contains(service) {
					continue
				}
				provider := &aws_terraforming.AWSProvider{}
				for _, r := range importResource(provider, service, region, account) {
					if strings.Contains(r.ResourceName, filters) {
						continue
					}
					resources = append(resources, importedResource{
						region:      region,
						tfResource:  r,
						serviceName: service,
					})
				}
			}
		}
		for _, service := range importer.getService() {
			if !importer.getGlobalService().Contains(service) {
				continue
			}
			provider := &aws_terraforming.AWSProvider{}
			for _, r := range importResource(provider, service, "eu-west-1", account) {
				if strings.Contains(r.ResourceName, filters) {
					continue
				}
				resources = append(resources, importedResource{
					region:      "eu-west-1",
					tfResource:  r,
					serviceName: service,
				})
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
				if _, exist := r.tfResource.Item["tags"]; exist {
					r.tfResource.Item["tags"].(map[string]interface{})[terraformTagName] = "true"
				}
				r.tfResource.Item["lifecycle"] = map[string]interface{}{
					"prevent_destroy": true,
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
			if runOnService == "" || service == runOnService {
				services = append(services, service)
			}

		}
	}
	sort.Strings(services)
	return services
}
