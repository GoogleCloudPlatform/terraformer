// Copyright 2018 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gcp

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"google.golang.org/api/compute/v1"
)

type GCPProvider struct { //nolint
	terraformutils.Provider
	projectName  string
	region       compute.Region
	providerType string
}

func GetRegions(project string) []string {
	computeService, err := compute.NewService(context.Background())
	if err != nil {
		return []string{}
	}
	regionsList, err := computeService.Regions.List(project).Do()
	if err != nil {
		return []string{}
	}
	regions := []string{}
	for _, region := range regionsList.Items {
		regions = append(regions, region.Name)
	}
	return regions
}

func getRegion(project, regionName string) *compute.Region {
	computeService, err := compute.NewService(context.Background())
	if err != nil {
		log.Println(err)
		return &compute.Region{}
	}
	region, err := computeService.Regions.Get(project, regionName).Do()
	if err != nil {
		log.Println(err)
		return &compute.Region{}
	}
	return region
}

// check projectName in env params
func (p *GCPProvider) Init(args []string) error {
	projectName := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if len(args) > 1 {
		projectName = args[1]
	}
	if projectName == "" {
		return errors.New("google cloud project name must be set")
	}
	p.projectName = projectName
	p.region = *getRegion(projectName, args[0])
	p.providerType = args[2]
	return nil
}

func (p *GCPProvider) GetName() string {
	if p.providerType != "" {
		return "google-" + p.providerType
	}
	return "google"
}

func (p *GCPProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("gcp: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"region":  p.region,
		"project": p.projectName,
	})
	return nil
}

// GetGCPSupportService return map of support service for GCP
func (p *GCPProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	services := ComputeServices
	services["bigQuery"] = &GCPFacade{service: &BigQueryGenerator{}}
	services["cloudFunctions"] = &GCPFacade{service: &CloudFunctionsGenerator{}}
	services["cloudsql"] = &GCPFacade{service: &CloudSQLGenerator{}}
	services["cloudtasks"] = &GCPFacade{service: &CloudTaskGenerator{}}
	services["dataProc"] = &GCPFacade{service: &DataprocGenerator{}}
	services["dns"] = &GCPFacade{service: &CloudDNSGenerator{}}
	services["gcs"] = &GCPFacade{service: &GcsGenerator{}}
	services["gke"] = &GCPFacade{service: &GkeGenerator{}}
	services["iam"] = &GCPFacade{service: &IamGenerator{}}
	services["kms"] = &GCPFacade{service: &KmsGenerator{}}
	services["logging"] = &GCPFacade{service: &LoggingGenerator{}}
	services["memoryStore"] = &GCPFacade{service: &MemoryStoreGenerator{}}
	services["monitoring"] = &GCPFacade{service: &MonitoringGenerator{}}
	services["project"] = &GCPFacade{service: &ProjectGenerator{}}
	services["instances"] = &GCPFacade{service: &InstancesGenerator{}}
	services["pubsub"] = &GCPFacade{service: &PubsubGenerator{}}
	services["schedulerJobs"] = &GCPFacade{service: &SchedulerJobsGenerator{}}
	return services
}

func (GCPProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		"backendBuckets": {"gcs": []string{"bucket_name", "name"}},
		"firewall":       {"networks": []string{"network", "self_link"}},
		"gke": {
			"networks":    []string{"network", "self_link"},
			"subnetworks": []string{"subnetwork", "self_link"},
		},
		"instanceTemplates": {
			"networks":    []string{"network", "self_link"},
			"subnetworks": []string{"subnetworks", "self_link"},
		},
		"regionInstanceGroupManagers": {"instanceTemplates": []string{"version.instance_template", "self_link"}},
		"instanceGroups":              {"instanceTemplates": []string{"version.instance_template", "self_link"}},
		"routes":                      {"networks": []string{"network", "self_link"}},
		"subnetworks":                 {"networks": []string{"network", "self_link"}},
		"forwardingRules": {
			"regionBackendServices": []string{"backend_service", "self_link"},
			"networks":              []string{"network", "self_link"},
		},
		"globalForwardingRules": {
			"targetHttpsProxies": []string{"target", "self_link"},
			"targetHttpProxies":  []string{"target", "self_link"},
			"targetSslProxies":   []string{"target", "self_link"},
		},
		"targetHttpsProxies": {
			"urlMaps": []string{"url_map", "self_link"},
		},
		"targetHttpProxies": {
			"urlMaps": []string{"url_map", "self_link"},
		},
		"targetSslProxies": {
			"backendServices": []string{"backend_service", "self_link"},
		},
		"backendServices": {
			"regionInstanceGroupManagers": []string{"backend.group", "instance_group"},
			"instanceGroupManagers":       []string{"backend.group", "instance_group"},
			"healthChecks":                []string{"health_checks", "self_link"},
		},
		"regionBackendServices": {
			"regionInstanceGroupManagers": []string{"backend.group", "instance_group"},
			"instanceGroupManagers":       []string{"backend.group", "instance_group"},
			"healthChecks":                []string{"health_checks", "self_link"},
		},
		"urlMaps": {
			"backendServices": []string{
				"default_service", "self_link",
				"path_matcher.default_service", "self_link",
				"path_matcher.path_rule.service", "self_link",
			},
			"regionBackendServices": []string{
				"default_service", "self_link",
				"path_matcher.default_service", "self_link",
				"path_matcher.path_rule.service", "self_link",
			},
		},
	}
}
func (p GCPProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			p.GetName(): map[string]interface{}{
				"project": p.projectName,
			},
		},
	}
}
