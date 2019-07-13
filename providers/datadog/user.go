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

package datadog

import (
	"fmt"

	datadog "github.com/zorkian/go-datadog-api"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
)

var (
	// UserAllowEmptyValues ...
	UserAllowEmptyValues = []string{}
	// UserAttributes ...
	UserAttributes = map[string]string{}
	// UserAdditionalFields ...
	UserAdditionalFields = map[string]string{}
)

// UserGenerator ...
type UserGenerator struct {
	DatadogService
}

func (UserGenerator) createResources(users []datadog.User) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	for _, user := range users {
		resourceName := user.GetHandle()
		resources = append(resources, terraform_utils.NewResource(
			resourceName,
			fmt.Sprintf("user_%s", resourceName),
			"datadog_user",
			"datadog",
			UserAttributes,
			UserAllowEmptyValues,
			UserAdditionalFields,
		))
	}

	return resources
}

// InitResources Generate TerraformResources from Datadog API,
// from each user create 1 TerraformResource.
// Need User ID as ID for terraform resource
func (g *UserGenerator) InitResources() error {
	client := datadog.NewClient(g.Args["api-key"].(string), g.Args["app-key"].(string))
	_, err := client.Validate()
	if err != nil {
		return err
	}
	users, err := client.GetUsers()
	if err != nil {
		return err
	}
	g.Resources = g.createResources(users)
	g.PopulateIgnoreKeys()
	return nil
}
