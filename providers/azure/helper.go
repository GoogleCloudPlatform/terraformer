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
	"fmt"
	"math/rand"
	"net/url"
	"strings"
)

// FROM https://github.com/terraform-providers/terraform-provider-azurerm/blob/6e006ff4e5d1fb200a6b37eb2743ff0ec8b11e0d/azurerm/helpers/azure/resourceid.go#L24

// ResourceID represents a parsed long-form Azure Resource Manager ID
// with the Subscription ID, Resource Group and the Provider as top-
// level fields, and other key-value pairs available via a map in the
// Path field.
type ResourceID struct {
	SubscriptionID string
	ResourceGroup  string
	Provider       string
	Path           map[string]string
}

// ParseAzureResourceID converts a long-form Azure Resource Manager ID
// into a ResourceID. We make assumptions about the structure of URLs,
// which is obviously not good, but the best thing available given the
// SDK.
func ParseAzureResourceID(id string) (*ResourceID, error) {
	idURL, err := url.ParseRequestURI(id)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse Azure ID: %s", err)
	}

	path := idURL.Path

	path = strings.TrimPrefix(path, "/")
	path = strings.TrimSuffix(path, "/")

	components := strings.Split(path, "/")

	// We should have an even number of key-value pairs.
	if len(components)%2 != 0 {
		return nil, fmt.Errorf("The number of path segments is not divisible by 2 in %q", path)
	}

	var subscriptionID string

	// Put the constituent key-value pairs into a map
	componentMap := make(map[string]string, len(components)/2)
	for current := 0; current < len(components); current += 2 {
		key := components[current]
		value := components[current+1]

		// Check key/value for empty strings.
		if key == "" || value == "" {
			return nil, fmt.Errorf("Key/Value cannot be empty strings. Key: '%s', Value: '%s'", key, value)
		}

		// Catch the subscriptionID before it can be overwritten by another "subscriptions"
		// value in the ID which is the case for the Service Bus subscription resource
		if key == "subscriptions" && subscriptionID == "" {
			subscriptionID = value
		} else {
			componentMap[key] = value
		}
	}

	// Build up a TargetResourceID from the map
	idObj := &ResourceID{}
	idObj.Path = componentMap

	if subscriptionID != "" {
		idObj.SubscriptionID = subscriptionID
	} else {
		return nil, fmt.Errorf("No subscription ID found in: %q", path)
	}

	if resourceGroup, ok := componentMap["resourceGroups"]; ok {
		idObj.ResourceGroup = resourceGroup
		delete(componentMap, "resourceGroups")
	} else if resourceGroup, ok := componentMap["resourcegroups"]; ok {
		// Some Azure APIs are weird and provide things in lower case...
		// However it's not clear whether the casing of other elements in the URI
		// matter, so we explicitly look for that case here.
		idObj.ResourceGroup = resourceGroup
		delete(componentMap, "resourcegroups")
	}

	// It is OK not to have a provider in the case of a resource group
	if provider, ok := componentMap["providers"]; ok {
		idObj.Provider = provider
		delete(componentMap, "providers")
	}

	return idObj, nil
}

func GenerateRandomString(strlen int) string {
	var lettersToUsed = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	RandomSlice := make([]rune, strlen)
	for index := range RandomSlice {
		RandomSlice[index] = lettersToUsed[rand.Intn(len(lettersToUsed))]
	}
	return string(RandomSlice)
}

func asHereDoc(json string) string {
	return fmt.Sprintf(`<<JSON
%s
JSON`, json)
}
