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

type TerraformResource struct {
	ResourceType string
	ResourceName string
	Item         interface{}
	ID           string
	Provider     string
	Attributes   map[string]string
}

func NewTerraformResource(ID, resourceName, resourceType, provider string, item interface{}, attributes map[string]string) TerraformResource {
	return TerraformResource{
		ResourceType: resourceType,
		ResourceName: TfSanitize(resourceName),
		Item:         item,
		ID:           ID,
		Provider:     provider,
		Attributes:   attributes,
	}
}
