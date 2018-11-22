package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var camel = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")

const pathForGenerateFiles = "/gcp_terraforming/compute_code_gen/"
const serviceTemplate = `
/*
AUTO GENERATE CODE - NOT FOR EDIT

*/

package {{.resourcePackageName}}

import (
	"context"
	"strings"
	"log"
	"waze/terraform/gcp_terraforming/gcp_generator"
	"waze/terraform/terraform_utils"

	"golang.org/x/oauth2/google"

	"google.golang.org/api/compute/v1"
)

var ignoreKey = map[string]bool{
	//"url":                true,
	"id":                 	true,
	"self_link":          	true,
	"fingerprint": 			true,
	"label_fingerprint": 	true,
	"creation_timestamp": 	true,
	{{ range $value := .attributesReference}}
	"{{$value}}":			true,{{end}}
}

var allowEmptyValues = map[string]bool{}

var additionalFields = map[string]string{
	"project": "waze-development",
}

type {{.resource}}Generator struct {
	gcp_generator.BasicGenerator
}

func ({{.resource}}Generator) createResources({{.resource}}List *compute.{{.resource}}ListCall, ctx context.Context, region string) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	if err := {{.resource}}List.Pages(ctx, func(page *compute.{{.responseName}}) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewTerraformResource(
				obj.Name,
				obj.Name,
				"{{.terraformName}}",
				"google",
				nil,
				map[string]string{"name": obj.Name, "project": "waze-development", "region": region},
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

func (g {{.resource}}Generator) Generate(zone string) error {
	region := strings.Join(strings.Split(zone, "-")[:len(strings.Split(zone, "-"))-1], "-")
	project := "waze-development" //os.Getenv("GOOGLE_CLOUD_PROJECT")
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	{{.resource}}List := computeService.{{.resource}}.List({{.parameterOrder}})

	resources := g.createResources({{.resource}}List, ctx, region)
	err = terraform_utils.GenerateTfState(resources)
	if err != nil {
		return err
	}
	converter := terraform_utils.TfstateConverter{}
	metadata := terraform_utils.NewResourcesMetaData(resources, ignoreKey, allowEmptyValues, additionalFields)
	resources, err = converter.Convert("terraform.tfstate", metadata)
	if err != nil {
		return err
	}
	err = terraform_utils.GenerateTf(resources, "{{.resource}}", region, "google")
	if err != nil {
		return err
	}
	return nil

}

`
const computeTemplate = `
package compute_code_gen

import (
	{{ range $key, $value := .services }}
	"waze/terraform/gcp_terraforming/compute_code_gen/{{$key}}"{{ end }}
	"waze/terraform/gcp_terraforming/gcp_generator"
)

var ComputeService = map[string]gcp_generator.Generator{
{{ range $key, $value := .services }}
	"{{$key}}":                   {{$key}}.{{title $key}}Generator{},{{ end }}

}

`

type TerraformResource struct {
	TerraformName       string
	AttributesReference []string
}

/*
autoscalers
backendServices
disks
globalForwardingRules
images
instanceGroupManagers
instanceGroups
instances
instanceTemplates
networks
regionInstanceGroupManagers
routes
securityPolicies
snapshots
sslCertificates
sslPolicies-empty
targetHttpProxies


regionDisks -empty
routers- empty
targetTcpProxies-empty
vpnTunnels-empty
*/
var terraformResources = map[string]TerraformResource{
	"addresses": {
		TerraformName: "google_compute_address",
		AttributesReference: []string{
			"type",
			"users",
			"address",
		},
	},
	"autoscalers": {
		TerraformName: "google_compute_autoscaler",
	},
	"backendBuckets": {
		TerraformName: "google_compute_backend_bucket",
	},
	"backendServices": {
		TerraformName: "google_compute_backend_service",
	},
	"disks": {
		TerraformName: "google_compute_disk",
		AttributesReference: []string{
			"last_attach_timestamp",
			"last_detach_timestamp",
			"users",
			"source_image_id",
			"source_snapshot_id",
		},
	},
	"firewalls": {
		TerraformName: "google_compute_firewall",
	},
	"forwardingRules": {
		TerraformName:       "google_compute_forwarding_rule",
		AttributesReference: []string{"service_name"},
	},
	"globalAddresses": {
		TerraformName:       "google_compute_global_address",
		AttributesReference: []string{"address"},
	},
	"globalForwardingRules": {
		TerraformName: "google_compute_global_forwarding_rule",
	},
	"healthChecks": {
		TerraformName:       "google_compute_health_check",
		AttributesReference: []string{"type"},
	},
	"httpHealthChecks": {
		TerraformName: "google_compute_http_health_check",
	},
	"httpsHealthChecks": {
		TerraformName: "google_compute_https_health_check",
	},
	"images": {
		TerraformName: "google_compute_image",
	},
	"instanceGroupManagers": {
		TerraformName:       "google_compute_instance_group_manager",
		AttributesReference: []string{"instance_group"},
	},
	"instanceGroups": {
		TerraformName:       "google_compute_instance_group",
		AttributesReference: []string{"size"},
	},
	"instanceTemplates": {
		TerraformName: "google_compute_instance_template",
	},
	"instances": {
		TerraformName: "google_compute_instance",
		AttributesReference: []string{
			"instance_id",
			"metadata_fingerprint",
			"tags_fingerprint",
			"cpu_platform",
		},
	},
	"networks": {
		TerraformName:       "google_compute_network",
		AttributesReference: []string{"gateway_ipv4", "name"},
	},
	"regionAutoscalers": {
		TerraformName: "google_compute_region_autoscaler",
	},
	"regionBackendServices": {
		TerraformName: "google_compute_region_backend_service",
	},
	"regionDisks": {
		TerraformName: "google_compute_region_disk",
		AttributesReference: []string{
			"last_attach_timestamp",
			"last_detach_timestamp",
			"users",
			"source_snapshot_id",
		},
	},
	"regionInstanceGroupManagers": {
		TerraformName:       "google_compute_region_instance_group_manager",
		AttributesReference: []string{"instance_group"},
	},
	"routers": {
		TerraformName: "google_compute_router",
	},
	"routes": {
		TerraformName:       "google_compute_route",
		AttributesReference: []string{"google_compute_route"},
	},
	"securityPolicies": {
		TerraformName: "google_compute_security_policy",
	},
	"snapshots": {
		TerraformName: "google_compute_snapshot",
		AttributesReference: []string{
			"snapshot_encryption_key_sha256",
			"source_disk_encryption_key_sha256",
			"source_disk_link",
		},
	},
	"sslCertificates": {
		TerraformName:       "google_compute_ssl_certificate",
		AttributesReference: []string{"certificate_id"},
	},
	"sslPolicies": {
		TerraformName: "google_compute_ssl_policy",
		AttributesReference: []string{
			"enabled_features",
		},
	},
	"subnetworks": {
		TerraformName: "google_compute_subnetwork",
		AttributesReference: []string{
			"gateway_address",
		},
	},
	"targetHttpProxies": {
		TerraformName:       "google_compute_target_http_proxy",
		AttributesReference: []string{"proxy_id"},
	},
	"targetHttpsProxies": {
		TerraformName:       "google_compute_target_https_proxy",
		AttributesReference: []string{"proxy_id"},
	},
	"targetSslProxies": {
		TerraformName:       "google_compute_target_ssl_proxy",
		AttributesReference: []string{"proxy_id"},
	},
	"targetTcpProxies": {
		TerraformName:       "google_compute_target_tcp_proxy",
		AttributesReference: []string{"proxy_id"},
	},
	"urlMaps": {
		TerraformName:       "google_compute_url_map",
		AttributesReference: []string{"map_id"},
	},
	"vpnTunnels": {
		TerraformName: "google_compute_vpn_tunnel",
		AttributesReference: []string{
			"shared_secret_hash",
			"detailed_status",
		},
	},
}

func main() {
	computeAPIdata, err := ioutil.ReadFile(os.Getenv("GOPATH") + "/src/google.golang.org/api/compute/v1/compute-api.json")
	if err != nil {
		log.Fatal(err)

	}
	computeAPI := map[string]interface{}{}
	json.Unmarshal(computeAPIdata, &computeAPI)
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
			t := template.Must(template.New("go_file").Parse(serviceTemplate))
			t.Execute(&tpl, map[string]interface{}{
				"resource":            strings.Title(resource),
				"responseName":        value.(map[string]interface{})["response"].(map[string]interface{})["$ref"].(string),
				"terraformName":       terraformResources[resource].TerraformName,
				"attributesReference": terraformResources[resource].AttributesReference,
				"resourcePackageName": resource,
				"parameterOrder":      parameterOrder,
			})
			rootPath, _ := os.Getwd()
			currentPath := rootPath + pathForGenerateFiles + resource
			os.MkdirAll(currentPath, os.ModePerm)
			ioutil.WriteFile(currentPath+"/"+resource+".go", tpl.Bytes(), os.ModePerm)
		} else {
			log.Println(resource)
		}
	}
	var tpl bytes.Buffer
	funcMap := template.FuncMap{
		"title": strings.Title,
	}
	t := template.Must(template.New("compute.go").Funcs(funcMap).Parse(computeTemplate))
	t.Execute(&tpl, map[string]interface{}{
		"services": terraformResources,
	})
	rootPath, _ := os.Getwd()
	ioutil.WriteFile(rootPath+pathForGenerateFiles+"compute.go", tpl.Bytes(), os.ModePerm)

}
