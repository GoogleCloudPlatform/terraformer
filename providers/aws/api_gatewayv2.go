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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
)

var apiGatewayV2AllowEmptyValues = []string{"tags.", "parent_id", "path_part"}

type APIGatewayV2Generator struct {
	AWSService
}

func (g *APIGatewayV2Generator) InitResources() error {

	sess, e := session.NewSession(&aws.Config{})
	if e != nil {
		return e
	}
	svc := apigatewayv2.New(sess)

	if err := g.loadRestApis(svc); err != nil {
		return err
	}
	if err := g.loadVpcLinks(svc); err != nil {
		return err
	}
	if err := g.loadUsagePlans(svc); err != nil {
		return err
	}
	if err := g.loadAPIKeys(svc); err != nil {
		return err
	}

	return nil
}

func (g *APIGatewayV2Generator) loadRestApis(svc *apigatewayv2.ApiGatewayV2) error {
	result, err := svc.GetApis(&apigatewayv2.GetApisInput{})
	if err != nil {
		fmt.Println("Failed to list APIs:", err)
		return err
	}
	for _, restAPI := range result.Items {
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
		if err := g.loadDocumentationParts(svc, restAPI.ApiId); err != nil {
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

	for _, stage := range output.Items {
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

	for _, model := range output.Items {
		resourceID := *model.ModelId
		g.Resources = append(g.Resources, terraformutils.NewResource(
			resourceID,
			resourceID,
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

	for _, route := range output.Items {
		resourceID := *route.RouteId

		g.Resources = append(g.Resources, terraformutils.NewResource(
			resourceID,
			resourceID,
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

	for _, response := range output.Items {

		resourceID := *response.RouteResponseId

		g.Resources = append(g.Resources, terraformutils.NewResource(
			resourceID,
			resourceID,
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

func (g *APIGatewayV2Generator) loadDocumentationParts(svc *apigatewayv2.ApiGatewayV2, restAPIID *string) error {

	return nil
}

func (g *APIGatewayV2Generator) loadAuthorizers(svc *apigatewayv2.ApiGatewayV2, restAPIID *string) error {

	return nil
}

func (g *APIGatewayV2Generator) loadVpcLinks(svc *apigatewayv2.ApiGatewayV2) error {
	return nil
}

func (g *APIGatewayV2Generator) loadUsagePlans(svc *apigatewayv2.ApiGatewayV2) error {
	return nil
}

func (g *APIGatewayV2Generator) loadAPIKeys(svc *apigatewayv2.ApiGatewayV2) error {
	return nil
}
