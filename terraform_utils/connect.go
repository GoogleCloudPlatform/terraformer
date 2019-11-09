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

func ConnectServices(importResources map[string][]Resource, isServicePath bool, resourceConnections map[string]map[string][]string) map[string][]Resource {
	for resource, connection := range resourceConnections {
		if _, exist := importResources[resource]; exist {
			for k, connectionPairs := range connection {
				if len(connectionPairs)%2 == 1 {
					continue
				}
				if cc, ok := importResources[k]; ok {
					for i := 0; i < len(connectionPairs)/2; i++ {
						connectionPair := []string{connectionPairs[i*2], connectionPairs[i*2+1]}
						for _, ccc := range cc {
							if !isServicePath {
								mapResource(importResources, resource, connectionPair, ccc, "local")
							} else {
								mapResource(importResources, resource, connectionPair, ccc, k)
							}
						}
					}
				}
			}
		}
	}
	return importResources
}

func mapResource(importResources map[string][]Resource, resource string, connectionPair []string, resourceToMap Resource, k string) {
	for i := range importResources[resource] {
		key := connectionPair[1]
		if connectionPair[1] == "self_link" || connectionPair[1] == "id" {
			key = resourceToMap.GetIDKey()
		}
		mappingResourceAttr := WalkAndGet(key, resourceToMap.InstanceState.Attributes)
		keyValue := resourceToMap.InstanceInfo.Type + "_" + resourceToMap.ResourceName + "_" + key
		linkValue := "${data.terraform_remote_state." + k + ".outputs." + keyValue + "}"

		if len(mappingResourceAttr) == 1 {
			resourceIdentifier := mappingResourceAttr[0].(string)
			WalkAndOverride(connectionPair[0], resourceIdentifier, linkValue, importResources[resource][i].Item)
		}
	}
}
