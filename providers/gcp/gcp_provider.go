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
	"errors"
	"os"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
)

const gcpProviderVersion = ">2.0.0"

type GCPProvider struct {
	terraform_utils.Provider
	projectName string
	zone        string
	region      string
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
	p.zone = args[0]
	p.region = strings.Join(strings.Split(p.zone, "-")[:len(strings.Split(p.zone, "-"))-1], "-")
	return nil
}

func (p *GCPProvider) GetName() string {
	return "google"
}

func (p *GCPProvider) InitService(serviceName string) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("gcp: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]string{
		"zone":    p.zone,
		"region":  p.region,
		"project": p.projectName,
	})
	return nil
}

// GetGCPSupportService return map of support service for GCP
func (p *GCPProvider) GetSupportedService() map[string]terraform_utils.ServiceGenerator {
	services := map[string]terraform_utils.ServiceGenerator{}
	services = ComputeServices
	services["gcs"] = &GcsGenerator{}
	services["monitoring"] = &MonitoringGenerator{}
	services["iam"] = &IamGenerator{}
	services["dns"] = &CloudDNSGenerator{}
	services["cloudsql"] = &CloudSQLGenerator{}
	services["gke"] = &GkeGenerator{}
	services["memoryStore"] = &MemoryStoreGenerator{}
	services["schedulerJobs"] = &SchedulerJobsGenerator{}
	services["bigQuery"] = &BigQueryGenerator{}
	services["dataProc"] = &DataprocGenerator{}
	services["cloudFunctions"] = &CloudFunctionsGenerator{}
	services["pubsub"] = &PubsubGenerator{}
	services["kms"] = &KmsGenerator{}
	services["project"] = &ProjectGenerator{}
	return services
}

func (GCPProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		"firewalls":             {"networks": []string{"network", "self_link"}},
		"routes":                {"networks": []string{"network", "self_link"}},
		"regionBackendServices": {"healthChecks": []string{"health_checks", "self_link"}},
		"backendBuckets":        {"gcs": []string{"bucket_name", "name"}},
		"instanceTemplates": {
			"networks":    []string{"network", "self_link"},
			"subnetworks": []string{"subnetworks", "self_link"},
		},
		"subnetworks": {"networks": []string{"network", "self_link"}},
		"gke": {
			"networks":    []string{"network", "self_link"},
			"subnetworks": []string{"subnetwork", "self_link"},
		},
		"regionInstanceGroupManagers": {"instanceTemplates": []string{"instance_template", "self_link"}},
	}
}

func (p GCPProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			p.GetName(): map[string]interface{}{
				"project": p.projectName,
				"version": gcpProviderVersion,
			},
		},
	}
}
