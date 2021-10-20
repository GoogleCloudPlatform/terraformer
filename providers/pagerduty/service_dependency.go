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

package pagerduty

import (
	//"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"fmt"
	"github.com/MakeNowJust/heredoc"
	pagerduty "github.com/heimweh/go-pagerduty/pagerduty"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

type ServiceDependencyGenerator struct {
	PagerDutyService
}

/*
We are using a different approach here, as long as https://github.com/GoogleCloudPlatform/terraformer/issues/1051 is not solved
With the usual way of creating resources (using the terraform pagerduty provider) we are not able to populate the userContactMethod.ID variable in the provider properly
What this code here does as a workaround is to create the terraform config without sanity checks from the provider. We also need to import the state manually.
THIS IS BAD AND SHOULD BE CHANGED! I just needed a quick fix as I couldn't wait for someone to come up with a fix.
*/

// Comment these code lines in once you found a solution for the issue mentioned above
/*
func (g *ServiceDependencyGenerator) createBusinessServiceDependencyResources(client *pagerduty.Client) error {
  respBusinessServices, _, err := client.BusinessServices.List()
	if err != nil {
		return err
	}
	for _, businessService := range respBusinessServices.BusinessServices {
    respBusinessServiceDependencies, _, err := client.ServiceDependencies.GetServiceDependenciesForType(businessService.ID, businessService.Type)
    if err != nil {
      return err
    }
    for _, businessServiceRelationships := range respBusinessServiceDependencies.Relationships {
      g.Resources = append(g.Resources, terraformutils.NewResource(
        businessServiceRelationships.ID,
        fmt.Sprintf("%s", strings.Replace(businessService.Name, " ", "_", -1)),
        "pagerduty_service_dependency",
        g.ProviderName,
        map[string]string{},
        []string{},
        map[string]interface{}{},
      ))
    }
	}

	return nil
}

func (g *ServiceDependencyGenerator) createServiceDependencyResources(client *pagerduty.Client) error {
  var offset = 0
	options := pagerduty.ListServicesOptions{}
	for {
		options.Offset = offset
		respServices, _, err := client.Services.List(&options)
		if err != nil {
			return err
		}

		for _, service := range respServices.Services {
      respServiceDependencies, _, err := client.ServiceDependencies.GetServiceDependenciesForType(service.ID, service.Type)
      if err != nil {
        return err
      }
      for _, serviceRelationships := range respServiceDependencies.Relationships {
        g.Resources = append(g.Resources, terraformutils.NewResource(
          serviceRelationships.ID,
          fmt.Sprintf("%s", strings.Replace(service.Name, " ", "_", -1)),
          "pagerduty_service_dependency",
          g.ProviderName,
					map[string]string{},
          []string{},
					map[string]interface{}{},
        ))
      }
		}

		if !respServices.More {
			break
		}
		offset += respServices.Limit
	}
	return nil
}
*/

// Here comes the dirty workaround. Remove this once you fixed the issue above.
// Defining a struct to populate our template properly
type ServiceDependency struct {
	ResourceName          string
	DependentServiceID    string
	DependentServiceType  string
	SupportingServiceID   string
	SupportingServiceType string
	DependencyID          string
}

/*
By design of the PagerDuty API and the go-pagerduty library that builds on it,
we need to fetch the services, iterate over the fetched services to find the ones
that got service dependencies and then iterate over the list of services that
have service dependencies. We populate our template with the previously fetched data.
*/
// Fucntion that generates the terraform code for technical services.
func (g *ServiceDependencyGenerator) createServiceDependencyResources(client *pagerduty.Client) error {
	log.Println("Generating technical service dependency config...")
	// Setting up folder
	basePath := filepath.Join("./generated/pagerduty/service_dependency")
	err := os.MkdirAll(basePath, os.ModePerm)
	if err != nil {
		panic(err)
	}
	// Setting offset variable because PagerDuty only responds wit maximum 25 services per request
	// By increasing the offset we make sure that we get all services
	var offset = 0
	options := pagerduty.ListServicesOptions{}
	for {
		options.Offset = offset
		// First API call: get all services
		respServices, _, err := client.Services.List(&options)
		if err != nil {
			return err
		}
		// Second API call: iterate over previously fetched services and find all services
		// that have service dependencies
		for _, service := range respServices.Services {
			respServiceDependencies, _, err := client.ServiceDependencies.GetServiceDependenciesForType(service.ID, service.Type)
			if err != nil {
				return err
			}
			// Iterate over every service dependency found and create our template
			for _, serviceRelationships := range respServiceDependencies.Relationships {
				// Fetching service name info to create unique terraform resource name
				var dependentServiceName string
				var supportingServiceName string
				if serviceRelationships.DependentService.Type == "business_service_reference" {
					dependentService, _, err := client.BusinessServices.Get(serviceRelationships.DependentService.ID)
					dependentServiceName = dependentService.Name
					if err != nil {
						return err
					}
				} else {
					dependentService, _, err := client.Services.Get(serviceRelationships.DependentService.ID, &pagerduty.GetServiceOptions{})
					dependentServiceName = dependentService.Name
					if err != nil {
						return err
					}
				}
				if serviceRelationships.SupportingService.Type == "business_service_reference" {
					supportingService, _, err := client.BusinessServices.Get(serviceRelationships.SupportingService.ID)
					supportingServiceName = supportingService.Name
					if err != nil {
						return err
					}
				} else {
					supportingService, _, err := client.Services.Get(serviceRelationships.SupportingService.ID, &pagerduty.GetServiceOptions{})
					supportingServiceName = supportingService.Name
					if err != nil {
						return err
					}
				}
				resourceName := fmt.Sprintf("%s_to_%s", dependentServiceName, supportingServiceName)
				// It is possible to create objects in the PagerDuty webapp with whitespaces in their respective names, but terraform doesn't allow whitespaces, so we replace them
				regexpMatch := regexp.MustCompile(`[\(,\)]`)
				userContactMethod := ServiceDependency{fmt.Sprintf("%s", strings.Replace(regexpMatch.ReplaceAllString(resourceName, ""), " ", "_", -1)), serviceRelationships.DependentService.ID, serviceRelationships.DependentService.Type, serviceRelationships.SupportingService.ID, serviceRelationships.SupportingService.Type, serviceRelationships.ID}
				userContactMethodTemplate, err := template.New("terraform").Parse(heredoc.Doc(`
        resource "pagerduty_service_dependency" "{{ .ResourceName}}" {
          dependency {
            dependent_service {
              id = "{{ .DependentServiceID}}"
              type = "{{if eq .DependentServiceType "business_service_reference" }}business_service{{else}}service{{end}}"
            }
            supporting_service {
              id = "{{ .SupportingServiceID}}"
              type = "{{if eq .SupportingServiceType "business_service_reference" }}business_service{{else}}service{{end}}"
            }
          }
        }

			`))
				if err != nil {
					panic(err)
				}
				// Create terraform import shell script
				userContactMethodImportTemplate, err := template.New("importer_script").Parse(heredoc.Doc(`
			terraform import pagerduty_service_dependency.{{ .ResourceName}} {{ .SupportingServiceID}}.{{if eq .SupportingServiceType "business_service_reference" }}business_service{{else}}service{{end}}.{{ .DependencyID}}
			`))
				userContactMethodOutputTemplate, err := template.New("importer_script").Parse(heredoc.Doc(`
      output "{{ .ResourceName}}_id" {
        value = pagerduty_service_dependency.{{ .ResourceName}}.id
      }
			`))
				// Populate and write our template
				outputDirTf := filepath.Join("./generated/pagerduty/service_dependency/service_dependency.tf")
				outputDirTfOutput := filepath.Join("./generated/pagerduty/service_dependency/outputs.tf")
				outputDirSh := filepath.Join("./generated/pagerduty/service_dependency/import_service_dependency.sh")
				tfFile, err := os.OpenFile(outputDirTf, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				shFile, err := os.OpenFile(outputDirSh, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
				if err != nil {
					panic(err)
				}
				outputFile, err := os.OpenFile(outputDirTfOutput, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
				if err != nil {
					panic(err)
				}
				err = userContactMethodTemplate.Execute(tfFile, userContactMethod)
				if err != nil {
					panic(err)
				}
				err = userContactMethodImportTemplate.Execute(shFile, userContactMethod)
				if err != nil {
					panic(err)
				}
				err = userContactMethodOutputTemplate.Execute(outputFile, userContactMethod)
				if err != nil {
					panic(err)
				}
			}
		}
		// Stop iteration if there are no more services
		if !respServices.More {
			break
		}
		// Increase offset after each iteration
		offset += respServices.Limit
	}
	return nil
}

func (g *ServiceDependencyGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*pagerduty.Client) error{
		g.createServiceDependencyResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}
