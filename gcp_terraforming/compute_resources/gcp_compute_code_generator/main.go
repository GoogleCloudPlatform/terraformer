package main

import (
	"bytes"
	"encoding/json"
	"go/format"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const pathForGenerateFiles = "/gcp_terraforming/compute_resources/"
const serviceTemplate = `
// AUTO-GENERATED CODE. DO NOT EDIT.
package computeTerrforming

import (
	"context"
	"strings"
	"log"
	"os"

	"waze/terraform/gcp_terraforming/gcp_generator"
	"waze/terraform/terraform_utils"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

var {{.resource}}IgnoreKey = map[string]bool{
	"id":                 	true,
	"self_link":          	true,
	"fingerprint": 			true,
	"label_fingerprint": 	true,
	"creation_timestamp": 	true,
	{{ range $value := .ignoreKeys}}
	"{{$value}}":			true,{{end}}
}

var {{.resource}}AllowEmptyValues = map[string]bool{
{{ range $value := .allowEmptyValues }}
	"{{$value}}":		true,
{{end}}
}

var {{.resource}}AdditionalFields = map[string]string{
	"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
	{{ range $key,$value := .additionalFields}}
	"{{$key}}":			"{{$value}}",{{end}}
}

type {{.titleResourceName}}Generator struct {
	gcp_generator.BasicGenerator
}

func ({{.titleResourceName}}Generator) createResources({{.resource}}List *compute.{{.titleResourceName}}ListCall, ctx context.Context, region, zone string) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	if err := {{.resource}}List.Pages(ctx, func(page *compute.{{.responseName}}) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewTerraformResource(
				{{ if .idWithZone  }}zone+"/"+obj.Name,{{else}}obj.Name,{{end}}
				obj.Name,
				"{{.terraformName}}",
				"google",
				nil,
				map[string]string{
					"name":    obj.Name,
					"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
					{{ if .needRegion}}"region":  region,{{end}}
					{{ if .byZone  }}"zone":    zone,{{end}}
				},
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

func (g {{.titleResourceName}}Generator) Generate(zone string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
	region := strings.Join(strings.Split(zone, "-")[:len(strings.Split(zone, "-"))-1], "-")
	project := os.Getenv("GOOGLE_CLOUD_PROJECT")
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

	resources := g.createResources({{.resource}}List, ctx, region, zone)
	metadata := terraform_utils.NewResourcesMetaData(resources, {{.resource}}IgnoreKey, {{.resource}}AllowEmptyValues, {{.resource}}AdditionalFields)
	return resources, metadata, nil

}

`
const computeTemplate = `
// AUTO-GENERATED CODE. DO NOT EDIT.
package computeTerrforming

import (
	"waze/terraform/gcp_terraforming/gcp_generator"
)

var ComputeService = map[string]gcp_generator.Generator{
{{ range $key, $value := .services }}
	"{{$key}}":                   {{title $key}}Generator{},{{ end }}

}

`

func main() {
	computeAPIData, err := ioutil.ReadFile(os.Getenv("GOPATH") + "/src/google.golang.org/api/compute/v1/compute-api.json")
	if err != nil {
		log.Fatal(err)

	}
	computeAPI := map[string]interface{}{}
	json.Unmarshal(computeAPIData, &computeAPI)
	funcMap := template.FuncMap{
		"title":   strings.Title,
		"toLower": strings.ToLower,
	}
	for resource, v := range computeAPI["resources"].(map[string]interface{}) {
		if _, exist := terraformResources[resource]; !exist {
			continue
		}
		if value, exist := v.(map[string]interface{})["methods"].(map[string]interface{})["list"]; exist {
			parameters := []string{}
			for _, param := range value.(map[string]interface{})["parameterOrder"].([]interface{}) {
				parameters = append(parameters, param.(string))
			}
			parameterOrder := strings.Join(parameters, ", ")
			var tpl bytes.Buffer
			t := template.Must(template.New("resource.go").Funcs(funcMap).Parse(serviceTemplate))
			err := t.Execute(&tpl, map[string]interface{}{
				"titleResourceName":   strings.Title(resource),
				"resource":            resource,
				"responseName":        value.(map[string]interface{})["response"].(map[string]interface{})["$ref"].(string),
				"terraformName":       terraformResources[resource].getTerraformName(),
				"ignoreKeys":          terraformResources[resource].getIgnoreKeys(),
				"additionalFields":    terraformResources[resource].getAdditionalFields(),
				"allowEmptyValues":    terraformResources[resource].getAllowEmptyValues(),
				"needRegion":          terraformResources[resource].ifNeedRegion(),
				"resourcePackageName": resource,
				"parameterOrder":      parameterOrder,
				"byZone":              terraformResources[resource].ifNeedZone(strings.Contains(parameterOrder, "zone")),
				"idWithZone":          terraformResources[resource].ifIDWithZone(strings.Contains(parameterOrder, "zone")),
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
	ioutil.WriteFile(rootPath+pathForGenerateFiles+"compute.go", codeFormat(tpl.Bytes()), os.ModePerm)

}

func codeFormat(src []byte) []byte {
	code, err := format.Source(src)
	if err != nil {
		log.Println(err)
	}
	return code
}
