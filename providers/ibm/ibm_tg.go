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
	"time"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/IBM/go-sdk-core/v4/core"
	tg "github.com/IBM/networking-go-sdk/transitgatewayapisv1"
)

// TGGenerator ...
type TGGenerator struct {
	IBMService
}

func (g TGGenerator) createTransitGatewayResources(gatewayID, gatewayName string) terraformutils.Resource {
	resource := terraformutils.NewSimpleResource(
		gatewayID,
		normalizeResourceName(gatewayName, false),
		"ibm_tg_gateway",
		"ibm",
		[]string{})
	return resource
}

func (g TGGenerator) createTransitGatewayConnectionResources(gatewayID, connectionID, connectionName string, dependsOn []string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s/%s", gatewayID, connectionID),
		normalizeResourceName(connectionName, false),
		"ibm_tg_connection",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resource
}

// CreateVersionDate requires mandatory version attribute. Any date from 2019-12-13 up to the currentdate may be provided. Specify the current date to request the latest version.
func CreateVersionDate() *string {
	version := time.Now().Format("2006-01-02")
	return &version
}

// InitResources ...
func (g *TGGenerator) InitResources() error {
	apiKey := os.Getenv("IC_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("no API key set")
	}
	tgURL := "https://transit.cloud.ibm.com/v1"
	transitgatewayOptions := &tg.TransitGatewayApisV1Options{
		URL: envFallBack([]string{"IBMCLOUD_TG_API_ENDPOINT"}, tgURL),
		Authenticator: &core.IamAuthenticator{
			ApiKey: apiKey,
		},
		Version: CreateVersionDate(),
	}

	tgclient, err := tg.NewTransitGatewayApisV1(transitgatewayOptions)
	if err != nil {
		return err
	}
	start := ""
	allrecs := []tg.TransitGateway{}
	for {
		listTransitGatewaysOptions := &tg.ListTransitGatewaysOptions{}
		if start != "" {
			listTransitGatewaysOptions.Start = &start
		}

		gateways, resp, err := tgclient.ListTransitGateways(listTransitGatewaysOptions)
		if err != nil {
			return fmt.Errorf("Error Listing Transit Gateways %s\n%s", err, resp)
		}
		start = GetNext(gateways.Next)
		allrecs = append(allrecs, gateways.TransitGateways...)
		if start == "" {
			break
		}
	}
	for _, gateway := range allrecs {
		g.Resources = append(g.Resources, g.createTransitGatewayResources(*gateway.ID, *gateway.Name))
		resourceName := g.Resources[len(g.Resources)-1:][0].ResourceName
		var dependsOn []string
		dependsOn = append(dependsOn,
			"ibm_tg_gateway."+resourceName)
		listTransitGatewayConnectionsOptions := &tg.ListTransitGatewayConnectionsOptions{
			TransitGatewayID: gateway.ID,
		}
		connections, response, err := tgclient.ListTransitGatewayConnections(listTransitGatewayConnectionsOptions)
		if err != nil {
			return fmt.Errorf("Error Listing Transit Gateway connections %s\n%s", err, response)
		}
		for _, connection := range connections.Connections {
			g.Resources = append(g.Resources, g.createTransitGatewayConnectionResources(*gateway.ID, *connection.ID, *connection.Name, dependsOn))
		}
	}
	return nil
}
