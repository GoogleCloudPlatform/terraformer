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

package ibm

import (
	"fmt"
	"os"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/container-services-go-sdk/kubernetesserviceapiv1"

	"github.com/IBM/go-sdk-core/v3/core"
)

type SatelliteGenerator struct {
	IBMService
}

func (g SatelliteGenerator) loadLocations(locID, locName string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		locID,
		normalizeResourceName(locName, false),
		"ibm_satellite_location",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{})

	// Remove parameters
	resource.IgnoreKeys = append(resource.IgnoreKeys,
		"^labels$",
	)

	return resource
}

func (g SatelliteGenerator) loadAssignHostControlPlane(locID, hostID string, labels []string, dependsOn []string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s/%s", locID, hostID),
		normalizeResourceName("ibm_satellite_host", true),
		"ibm_satellite_host",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"labels":     labels,
			"depends_on": dependsOn,
		})
	return resource
}

func (g *SatelliteGenerator) InitResources() error {
	bmxConfig := &bluemix.Config{
		BluemixAPIKey: os.Getenv("IC_API_KEY"),
	}

	sess, err := session.New(bmxConfig)
	if err != nil {
		return err
	}

	err = authenticateAPIKey(sess)
	if err != nil {
		return err
	}

	bluemixToken := ""
	if strings.HasPrefix(sess.Config.IAMAccessToken, "Bearer") {
		bluemixToken = sess.Config.IAMAccessToken[7:len(sess.Config.IAMAccessToken)]
	} else {
		bluemixToken = sess.Config.IAMAccessToken
	}

	containerEndpoint := kubernetesserviceapiv1.DefaultServiceURL
	kubernetesServiceV1Options := &kubernetesserviceapiv1.KubernetesServiceApiV1Options{
		URL: envFallBack([]string{"IBMCLOUD_SATELLITE_API_ENDPOINT"}, containerEndpoint),
		Authenticator: &core.BearerTokenAuthenticator{
			BearerToken: bluemixToken,
		},
	}

	satelliteClient, err := kubernetesserviceapiv1.NewKubernetesServiceApiV1(kubernetesServiceV1Options)
	if err != nil {
		return err
	}

	getSatLocOpts := &kubernetesserviceapiv1.GetSatelliteLocationsOptions{}
	locations, _, err := satelliteClient.GetSatelliteLocations(getSatLocOpts)
	if err != nil {
		return err
	}

	for _, loc := range locations {
		var locDependsOn []string

		// Location
		if loc.Deployments != nil && !strings.Contains(*loc.Deployments.Message, "R0037") {
			g.Resources = append(g.Resources, g.loadLocations(*loc.ID, *loc.Name))
			resourceName := g.Resources[len(g.Resources)-1:][0].ResourceName
			locDependsOn = append(locDependsOn,
				"ibm_satellite_location."+resourceName)

			// Assign Host - Control plane
			getSatHostOpts := &kubernetesserviceapiv1.GetSatelliteHostsOptions{
				Controller: loc.ID,
			}
			hosts, resp, err := satelliteClient.GetSatelliteHosts(getSatHostOpts)
			if err != nil {
				return fmt.Errorf("Error getting satellite control plane hosts %s\n%s", err, resp)
			}

			for _, host := range hosts {
				if *host.Assignment.ClusterName == "infrastructure" {
					hostLabels := []string{}
					for key, value := range host.Labels {
						hostLabels = append(hostLabels, fmt.Sprintf("%s=%s", key, value))
					}
					g.Resources = append(g.Resources, g.loadAssignHostControlPlane(*loc.ID, *host.ID, hostLabels, locDependsOn))
				}
			}

		}
	}

	return nil
}
