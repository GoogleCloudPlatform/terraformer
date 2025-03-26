// Copyright 2020 The Terraformer Authors.
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

package aws

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
)

var apiGatewayV2AllowEmptyValues = []string{"tags.", "parent_id", "path_part"}

type APIGatewayV2Generator struct {
	AWSService
}

func (g *APIGatewayV2Generator) InitResources() error {

	svc := apigatewayv2.New(session.Must(session.NewSession()))

	if err := g.loadRestApis(svc); err != nil {
		return err
	}
	if err := g.loadVpcLinks(svc); err != nil {
		return err
	}
	return nil
}

func (g *APIGatewayV2Generator) loadRestApis(svc *apigatewayv2.ApiGatewayV2) error {
	output, err := svc.GetApis(&apigatewayv2.GetApisInput{})
	if err != nil {
		fmt.Println("Failed to list APIs:", err)
		return err
	}

	err = g.processRestApis(svc, output.Items)
	if err != nil {
		fmt.Println("Failed to list APIs:", err)
		return err
	}

	for output.NextToken != nil {
		output, err = svc.GetApis(&apigatewayv2.GetApisInput{
			NextToken: output.NextToken,
		})
		if err != nil {
			fmt.Println("Failed to list APIs:", err)
			return err
		}
		if err = g.processRestApis(svc, output.Items); err != nil {
			fmt.Println("Failed to list APIs:", err)
			return err
		}
	}

	return nil
}

func (g *APIGatewayV2Generator) processRestApis(svc *apigatewayv2.ApiGatewayV2, output []*apigatewayv2.Api) error {
	for _, restAPI := range output {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			*restAPI.ApiId,
			*restAPI.ApiId+"_"+*restAPI.Name,
			"aws_apigatewayv2_api",
			"aws",
			apiGatewayV2AllowEmptyValues,
		))
		if err := g.loadStages(svc, restAPI.ApiId); err != nil {
			return err
		}
		if err := g.loadModels(svc, restAPI.ApiId); err != nil {
			return err
		}
		if err := g.loadRoutes(svc, restAPI.ApiId); err != nil {
			return err
		}

		if err := g.loadAuthorizers(svc, restAPI.ApiId); err != nil {
			return err
		}

	}
	return nil
}

func (g *APIGatewayV2Generator) loadStages(svc *apigatewayv2.ApiGatewayV2, restAPIID *string) error {

	output, err := svc.GetStages(&apigatewayv2.GetStagesInput{
		ApiId: restAPIID,
	})
	if err != nil {
		return err
	}
	err = g.processStages(output.Items, restAPIID)
	if err != nil {
		return err
	}

	for output.NextToken != nil {
		output, err = svc.GetStages(&apigatewayv2.GetStagesInput{
			NextToken: output.NextToken,
		})
		if err != nil {
			return err
		}
		err = g.processStages(output.Items, restAPIID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *APIGatewayV2Generator) processStages(output []*apigatewayv2.Stage, restAPIID *string) error {
	for _, stage := range output {
		stageID := *restAPIID + "/" + StringValue(stage.StageName)
		g.Resources = append(g.Resources, terraformutils.NewResource(
			stageID,
			stageID,
			"aws_api_gateway_stage",
			"aws",
			map[string]string{
				"rest_api_id": *restAPIID,
				"stage_name":  *stage.StageName,
			},
			apiGatewayAllowEmptyValues,
			map[string]interface{}{},
		))
	}
	return nil
}

func (g *APIGatewayV2Generator) loadModels(svc *apigatewayv2.ApiGatewayV2, restAPIID *string) error {

	output, err := svc.GetModels(
		&apigatewayv2.GetModelsInput{
			ApiId: restAPIID,
		})
	if err != nil {
		return err
	}
	err = g.processModels(output.Items, restAPIID)
	if err != nil {
		return err
	}

	for output.NextToken != nil {
		output, err = svc.GetModels(
			&apigatewayv2.GetModelsInput{
				NextToken: output.NextToken,
			})
		if err != nil {
			return err
		}
		err = g.processModels(output.Items, restAPIID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *APIGatewayV2Generator) processModels(output []*apigatewayv2.Model, restAPIID *string) error {
	for _, model := range output {

		g.Resources = append(g.Resources, terraformutils.NewResource(
			*model.ModelId,
			*model.ModelId,
			"aws_apigatewayv2_model",
			"aws",
			map[string]string{
				"name":         StringValue(model.Name),
				"content_type": StringValue(model.ContentType),
				"schema":       StringValue(model.Schema),
				"api_id":       StringValue(restAPIID),
			},
			apiGatewayAllowEmptyValues,
			map[string]interface{}{},
		))
	}
	return nil
}

func (g *APIGatewayV2Generator) loadRoutes(svc *apigatewayv2.ApiGatewayV2, restAPIID *string) error {

	output, err := svc.GetRoutes(
		&apigatewayv2.GetRoutesInput{
			ApiId: restAPIID,
		})
	if err != nil {
		return err
	}

	err = g.processRoutes(svc, output.Items, restAPIID)
	if err != nil {
		return err
	}

	for output.NextToken != nil {
		output, err := svc.GetRoutes(
			&apigatewayv2.GetRoutesInput{
				NextToken: output.NextToken,
			})
		if err != nil {
			return err
		}

		err = g.processRoutes(svc, output.Items, restAPIID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *APIGatewayV2Generator) processRoutes(svc *apigatewayv2.ApiGatewayV2, output []*apigatewayv2.Route, restAPIID *string) error {
	for _, route := range output {

		g.Resources = append(g.Resources, terraformutils.NewResource(
			*route.RouteId,
			*route.RouteId,
			"aws_apigatewayv2_route",
			"aws",
			map[string]string{
				"api_id":    *restAPIID,
				"route_key": *route.RouteKey,
			},
			apiGatewayAllowEmptyValues,
			map[string]interface{}{},
		))
		if err := g.loadResponses(svc, restAPIID, route.RouteId); err != nil {
			return err
		}
	}
	return nil
}

func (g *APIGatewayV2Generator) loadResponses(svc *apigatewayv2.ApiGatewayV2, restAPIID *string, routeID *string) error {

	output, err := svc.GetRouteResponses(
		&apigatewayv2.GetRouteResponsesInput{
			ApiId:   restAPIID,
			RouteId: routeID,
		})

	if err != nil {
		return err
	}
	err = g.processResponses(output.Items, restAPIID, routeID)
	if err != nil {
		return err
	}

	for output.NextToken != nil {
		output, err = svc.GetRouteResponses(
			&apigatewayv2.GetRouteResponsesInput{
				NextToken: output.NextToken,
			})

		if err != nil {
			return err
		}
		err = g.processResponses(output.Items, restAPIID, routeID)
		if err != nil {
			return err
		}
	}

	return nil
}
func (g *APIGatewayV2Generator) processResponses(output []*apigatewayv2.RouteResponse, restAPIID *string, routeID *string) error {
	for _, response := range output {

		g.Resources = append(g.Resources, terraformutils.NewResource(
			*response.RouteResponseId,
			*response.RouteResponseId,
			"aws_apigatewayv2_route_response",
			"aws",
			map[string]string{
				"api_id":             *restAPIID,
				"route_id":           *routeID,
				"route_response_key": "$default",
			},
			apiGatewayAllowEmptyValues,
			map[string]interface{}{},
		))

	}
	return nil
}

func (g *APIGatewayV2Generator) loadAuthorizers(svc *apigatewayv2.ApiGatewayV2, restAPIID *string) error {

	output, err := svc.GetAuthorizers(
		&apigatewayv2.GetAuthorizersInput{
			ApiId: restAPIID,
		})
	if err != nil {
		return err
	}
	g.processAuthorizers(output.Items, restAPIID)

	for output.NextToken != nil {
		output, err = svc.GetAuthorizers(
			&apigatewayv2.GetAuthorizersInput{
				NextToken: output.NextToken,
			})
		if err != nil {
			return err
		}
		g.processAuthorizers(output.Items, restAPIID)
	}

	return nil
}

func (g *APIGatewayV2Generator) processAuthorizers(output []*apigatewayv2.Authorizer, restAPIID *string) error {
	for _, authoriser := range output {

		g.Resources = append(g.Resources, terraformutils.NewResource(
			*authoriser.AuthorizerId,
			*authoriser.AuthorizerId,
			"aws_apigatewayv2_authorizer",
			"aws",
			map[string]string{
				"api_id":          *restAPIID,
				"name":            StringValue(authoriser.Name),
				"authorizer_type": *authoriser.AuthorizerType,
			},
			apiGatewayAllowEmptyValues,
			map[string]interface{}{},
		))

	}
	return nil
}

func (g *APIGatewayV2Generator) loadVpcLinks(svc *apigatewayv2.ApiGatewayV2) error {

	output, err := svc.GetVpcLinks(
		&apigatewayv2.GetVpcLinksInput{})
	if err != nil {
		return err
	}
	g.processVpcLinks(output.Items)

	for output.NextToken != nil {
		output, err := svc.GetVpcLinks(
			&apigatewayv2.GetVpcLinksInput{
				NextToken: output.NextToken,
			})
		if err != nil {
			return err
		}
		g.processVpcLinks(output.Items)
	}

	return nil
}

func (g *APIGatewayV2Generator) processVpcLinks(output []*apigatewayv2.VpcLink) error {
	for _, vpcLink := range output {

		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			*vpcLink.VpcLinkId,
			*vpcLink.VpcLinkId,
			"aws_apigatewayv2_vpc_link",
			"aws",
			apiGatewayAllowEmptyValues))
	}
	return nil
}
