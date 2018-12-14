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

package terraform_utils

type ResourceMetaData struct {
	Provider         string
	IgnoreKeys       map[string][]string
	AllowEmptyValue  map[string]bool
	AdditionalFields map[string]string
}

func NewResourcesMetaData(resources []TerraformResource, ignoreKeys map[string][]string, allowEmptyValue map[string]bool, AdditionalFields map[string]string) map[string]ResourceMetaData {
	data := map[string]ResourceMetaData{}
	for _, resource := range resources {
		data[resource.ID] = ResourceMetaData{
			Provider:         resource.Provider,
			IgnoreKeys:       ignoreKeys,
			AllowEmptyValue:  allowEmptyValue,
			AdditionalFields: AdditionalFields,
		}
	}
	return data
}
