// Copyright 2019 The Terraformer Authors.
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

package azure

import (
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-helpers/authentication"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type AzureService struct { //nolint
	terraformutils.Service
}

func (az *AzureService) getClientArgs() (subscriptionID string, resourceGroup string, authorizer autorest.Authorizer) {
	subs := az.Args["config"].(authentication.Config).SubscriptionID
	auth := az.Args["authorizer"].(autorest.Authorizer)
	resg := az.Args["resource_group"].(string)
	return subs, resg, auth
}

func (az *AzureService) AppendSimpleResource(id string, resourceName string, resourceType string) {
	newResource := terraformutils.NewSimpleResource(id, resourceName, resourceType, az.ProviderName, []string{})
	az.Resources = append(az.Resources, newResource)
}

func (az *AzureService) appendSimpleAssociation(id string, linkedResourceName string, resourceName *string, resourceType string, attributes map[string]string) {
	var resourceName2 string
	if resourceName != nil {
		resourceName2 = *resourceName
	} else {
		resourceName0 := strings.ReplaceAll(resourceType, "azurerm_", "")
		resourceName1 := resourceName0[strings.IndexByte(resourceName0, '_'):]
		resourceName2 = linkedResourceName + resourceName1
	}
	newResource := terraformutils.NewResource(
		id, resourceName2, resourceType, az.ProviderName, attributes,
		[]string{"name"},
		map[string]interface{}{},
	)
	az.Resources = append(az.Resources, newResource)
}
