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

package main

type gcpResourceRenderable interface {
	getTerraformName() string
	getAdditionalFields() map[string]string
	getAllowEmptyValues() []string
	ifNeedRegion() bool
	ifNeedZone(zoneInParameters bool) bool
	ifIDWithZone(zoneInParameters bool) bool
	getAdditionalFieldsForRefresh() map[string]string
}

type basicGCPResource struct {
	terraformName              string
	allowEmptyValues           []string
	additionalFields           map[string]string
	additionalFieldsForRefresh map[string]string
}

func (b basicGCPResource) getTerraformName() string {
	return b.terraformName
}

func (b basicGCPResource) getAdditionalFields() map[string]string {
	return b.additionalFields
}

func (b basicGCPResource) getAdditionalFieldsForRefresh() map[string]string {
	return b.additionalFieldsForRefresh
}

func (b basicGCPResource) getAllowEmptyValues() []string {
	return b.allowEmptyValues
}
func (b basicGCPResource) ifNeedRegion() bool {
	return true
}

func (b basicGCPResource) ifNeedZone(zoneInParameters bool) bool {
	return zoneInParameters
}

func (b basicGCPResource) ifIDWithZone(zoneInParameters bool) bool {
	return zoneInParameters
}
