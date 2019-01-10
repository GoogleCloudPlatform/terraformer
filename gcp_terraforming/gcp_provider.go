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

package gcp_terraforming

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"waze/terraformer/terraform_utils"
)

const GenerateFilesFolderPath = "%s/generated/gcp/%s/%s/%s/"

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

func (p *GCPProvider) GenerateOutputPath() error {
	if err := os.MkdirAll(p.CurrentPath(), os.ModePerm); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func (p *GCPProvider) CurrentPath() string {
	rootPath, _ := os.Getwd()
	return fmt.Sprintf(GenerateFilesFolderPath, rootPath, p.projectName, p.zone, p.Service.GetName())
}

func (p *GCPProvider) InitService(serviceName string) error {
	var isSupported bool
	if _, isSupported = p.GetGCPSupportService()[serviceName]; !isSupported {
		return errors.New("gcp: " + serviceName + " not supported service")
	}
	p.Service = p.GetGCPSupportService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]string{
		"zone":    p.zone,
		"region":  p.region,
		"project": p.projectName,
	})
	return nil
}

func (p *GCPProvider) RegionResource() map[string]interface{} {
	return map[string]interface{}{
		"google": map[string]interface{}{
			"region":  p.region,
			"project": p.projectName,
		},
	}
}

// GetGCPSupportService return map of support service for GCP
func (p *GCPProvider) GetGCPSupportService() map[string]terraform_utils.ServiceGenerator {
	services := map[string]terraform_utils.ServiceGenerator{}
	services = ComputeServices
	services["gcs"] = &GcsGenerator{}
	services["monitoring"] = &MonitoringGenerator{}
	services["iam"] = &IamGenerator{}
	services["dns"] = &CloudDNSGenerator{}
	services["cloudsql"] = &CloudSQLGenerator{}
	return services
}
