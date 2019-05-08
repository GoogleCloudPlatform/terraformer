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

import (
	"bytes"
	"encoding/json"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

const pathForGenerateFiles = "/gcp_terraforming/"
const serviceTemplate = `
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

// AUTO-GENERATED CODE. DO NOT EDIT.
package gcp_terraforming

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

var {{.resource}}AllowEmptyValues = []string{"{{join .allowEmptyValues "\",\"" }}"}

var {{.resource}}AdditionalFields = map[string]string{
	{{ range $key,$value := .additionalFields}}
	"{{$key}}":			"{{$value}}",{{end}}
}

type {{.titleResourceName}}Generator struct {
	GCPService
}

// Run on {{.resource}}List and create for each TerraformResource
func (g {{.titleResourceName}}Generator) createResources({{.resource}}List *compute.{{.titleResourceName}}ListCall, ctx context.Context) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	if err := {{.resource}}List.Pages(ctx, func(page *compute.{{.responseName}}) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewResource(
				{{ if .idWithZone  }}g.GetArgs()["zone"]+"/"+obj.Name,{{else}}obj.Name,{{end}}
				obj.Name,
				"{{.terraformName}}",
				"google",
				map[string]string{
					"name":    obj.Name,
					"project": g.GetArgs()["project"],
					{{ if .needRegion}}"region":  g.GetArgs()["region"],{{end}}
					{{ if .byZone  }}"zone":    g.GetArgs()["zone"],{{end}}
					{{ range $key, $value := .additionalFieldsForRefresh}}
					"{{$key}}":			"{{$value}}",{{end}}
				},
				{{.resource}}AllowEmptyValues,
				{{.resource}}AdditionalFields,
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each {{.resource}} create 1 TerraformResource
// Need {{.resource}} name as ID for terraform resource
func (g *{{.titleResourceName}}Generator) InitResources() error {
	ctx := context.Background()
	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	{{.resource}}List := computeService.{{.titleResourceName}}.List({{.parameterOrder}})

	g.Resources = g.createResources({{.resource}}List, ctx)
	g.PopulateIgnoreKeys()
	return nil

}

`
const computeTemplate = `
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

// AUTO-GENERATED CODE. DO NOT EDIT.
package gcp_terraforming

import (
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
)

// Map of supported GCP compute service with code generate
var ComputeServices = map[string]terraform_utils.ServiceGenerator{
{{ range $key, $value := .services }}
	"{{$key}}":                   &{{title $key}}Generator{},{{ end }}

}

`

func main() {
	computeAPIData, err := ioutil.ReadFile("vendor/google.golang.org/api/compute/v1/compute-api.json") //TODO delete this hack
	if err != nil {
		log.Fatal(err)

	}
	computeAPI := map[string]interface{}{}
	json.Unmarshal(computeAPIData, &computeAPI)
	funcMap := template.FuncMap{
		"title":   strings.Title,
		"toLower": strings.ToLower,
		"join":    strings.Join,
	}
	for resource, v := range computeAPI["resources"].(map[string]interface{}) {
		if _, exist := terraformResources[resource]; !exist {
			continue
		}
		if value, exist := v.(map[string]interface{})["methods"].(map[string]interface{})["list"]; exist {
			parameters := []string{}
			for _, param := range value.(map[string]interface{})["parameterOrder"].([]interface{}) {
				parameters = append(parameters, `g.GetArgs()["`+param.(string)+`"]`)
			}
			parameterOrder := strings.Join(parameters, ", ")
			var tpl bytes.Buffer
			t := template.Must(template.New("resource.go").Funcs(funcMap).Parse(serviceTemplate))
			err := t.Execute(&tpl, map[string]interface{}{
				"titleResourceName":          strings.Title(resource),
				"resource":                   resource,
				"responseName":               value.(map[string]interface{})["response"].(map[string]interface{})["$ref"].(string),
				"terraformName":              terraformResources[resource].getTerraformName(),
				"additionalFields":           terraformResources[resource].getAdditionalFields(),
				"additionalFieldsForRefresh": terraformResources[resource].getAdditionalFieldsForRefresh(),
				"allowEmptyValues":           terraformResources[resource].getAllowEmptyValues(),
				"needRegion":                 terraformResources[resource].ifNeedRegion(),
				"resourcePackageName":        resource,
				"parameterOrder":             parameterOrder,
				"byZone":                     terraformResources[resource].ifNeedZone(strings.Contains(parameterOrder, "zone")),
				"idWithZone":                 terraformResources[resource].ifIDWithZone(strings.Contains(parameterOrder, "zone")),
			})
			if err != nil {
				log.Print(resource, err)
				continue
			}
			rootPath, _ := os.Getwd()
			currentPath := rootPath + pathForGenerateFiles
			err = os.MkdirAll(currentPath, os.ModePerm)
			if err != nil {
				log.Print(resource, err)
				continue
			}
			err = ioutil.WriteFile(currentPath+"/"+resource+".go", codeFormat(tpl.Bytes()), os.ModePerm)
			if err != nil {
				log.Print(resource, err)
				continue
			}
		} else {
			log.Println(resource)
		}
	}
	var tpl bytes.Buffer
	t := template.Must(template.New("compute.go").Funcs(funcMap).Parse(computeTemplate))
	err = t.Execute(&tpl, map[string]interface{}{
		"services": terraformResources,
	})
	if err != nil {
		log.Print(err)
	}
	rootPath, _ := os.Getwd()
	err = ioutil.WriteFile(rootPath+pathForGenerateFiles+"compute.go", codeFormat(tpl.Bytes()), os.ModePerm)
	if err != nil {
		log.Println(err)
	}
}

func codeFormat(src []byte) []byte {
	code, err := format.Source(src)
	if err != nil {
		log.Println(err)
	}
	return code
}
