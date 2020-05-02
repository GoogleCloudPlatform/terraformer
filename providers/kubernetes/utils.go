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
	"strings"

	"github.com/iancoleman/strcase"
)

func extractClientSetFuncGroupName(group, version string) string {
	v := strings.Title(version)
	if len(group) > 0 {
		return strings.Title(strings.Split(group, ".")[0]) + v
	}
	return "Core" + v
}

func extractClientSetFuncTypeName(kind string) string {
	switch string(kind[len(kind)-1]) {
	case "s":
		return kind + "es"
	case "y":
		return strings.TrimSuffix(kind, "y") + "ies"
	}
	return kind + "s"
}

func extractTfResourceName(kind string) string {
	return "kubernetes_" + strcase.ToSnake(kind)
}
