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

package kubernetes

import (
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	tfschema "github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	tfk8s "github.com/terraform-providers/terraform-provider-kubernetes/kubernetes"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/discovery"
	"k8s.io/kubectl/pkg/pluginutils"
)

type KubernetesProvider struct {
	terraform_utils.Provider
	region string
}

const k8sProviderVersion = ">=1.4.0"

func (p KubernetesProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p KubernetesProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"kubernetes": map[string]interface{}{
				"version": k8sProviderVersion,
			},
		},
	}
}

func (p *KubernetesProvider) Init(args []string) error {
	return nil
}

func (p *KubernetesProvider) GetName() string {
	return "kubernetes"
}

func (p *KubernetesProvider) InitService(serviceName string) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("kubernetes: " + serviceName + " not supported resource")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetProviderName(p.GetName())
	return nil
}

// GetSupportService return map of supported resource for Kubernetes
func (p *KubernetesProvider) GetSupportedService() map[string]terraform_utils.ServiceGenerator {
	resources := make(map[string]terraform_utils.ServiceGenerator)

	config, _, err := pluginutils.InitClientAndConfig()
	if err != nil {
		return resources
	}

	dc, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		return resources
	}

	lists, err := dc.ServerPreferredResources()
	if err != nil {
		return resources
	}

	for _, list := range lists {
		if len(list.APIResources) == 0 {
			continue
		}

		gv, err := schema.ParseGroupVersion(list.GroupVersion)
		if err != nil {
			continue
		}

		for _, resource := range list.APIResources {
			if len(resource.Verbs) == 0 {
				continue
			}

			// filter to resources that support list
			if len(resource.Verbs) > 0 && !sets.NewString(resource.Verbs...).Has("list") {
				continue
			}

			// filter to resource that are supported by terraform kubernetes provider
			if _, ok := tfk8s.Provider().(*tfschema.Provider).ResourcesMap[extractTfResourceName(resource.Kind)]; !ok {
				continue
			}

			resources[resource.Name] = &Kind{
				Group:      gv.Group,
				Version:    gv.Version,
				Name:       resource.Kind,
				Namespaced: resource.Namespaced,
			}
		}
	}
	return resources
}
