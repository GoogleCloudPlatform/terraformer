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
	"context"
	"log"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils/terraformerstring"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go/aws"
)

var apiGatewayAllowEmptyValues = []string{"tags.", "parent_id", "path_part"}

type APIGatewayGenerator struct {
	AWSService
}

func (g *APIGatewayGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := apigateway.New(config)

	if err := g.loadRestApis(svc); err != nil {
		return err
	}
	if err := g.loadVpcLinks(svc); err != nil {
		return err
	}
	if err := g.loadUsagePlans(svc); err != nil {
		return err
	}

	return nil
}

func (g *APIGatewayGenerator) loadRestApis(svc *apigateway.Client) error {
	p := apigateway.NewGetRestApisPaginator(svc.GetRestApisRequest(&apigateway.GetRestApisInput{}))
	for p.Next(context.Background()) {
		for _, restAPI := range p.CurrentPage().Items {
			if g.shouldFilterRestAPI(restAPI.Tags) {
				continue
			}
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				*restAPI.Id,
				*restAPI.Name,
				"aws_api_gateway_rest_api",
				"aws",
				apiGatewayAllowEmptyValues))
			if err := g.loadStages(svc, restAPI.Id); err != nil {
				return err
			}
			if err := g.loadResources(svc, restAPI.Id); err != nil {
				return err
			}
			if err := g.loadModels(svc, restAPI.Id); err != nil {
				return err
			}
			if err := g.loadResponses(svc, restAPI.Id); err != nil {
				return err
			}
			if err := g.loadDocumentationParts(svc, restAPI.Id); err != nil {
				return err
			}
			if err := g.loadAuthorizers(svc, restAPI.Id); err != nil {
				return err
			}
		}
	}
	return p.Err()
}

func (g *APIGatewayGenerator) shouldFilterRestAPI(tags map[string]string) bool {
	for _, filter := range g.Filter {
		if strings.HasPrefix(filter.FieldPath, "tags.") && filter.IsApplicable("api_gateway_rest_api") {
			tagName := strings.Replace(filter.FieldPath, "tags.", "", 1)
			if val, ok := tags[tagName]; ok {
				return !terraformerstring.ContainsString(filter.AcceptableValues, val)
			}
			return true
		}
	}
	return false
}

func (g *APIGatewayGenerator) loadStages(svc *apigateway.Client, restAPIID *string) error {
	output, err := svc.GetStagesRequest(&apigateway.GetStagesInput{
		RestApiId: restAPIID,
	}).Send(context.Background())
	if err != nil {
		return err
	}
	for _, stage := range output.GetStagesOutput.Item {
		stageID := *restAPIID + "/" + *stage.StageName
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

func (g *APIGatewayGenerator) loadResources(svc *apigateway.Client, restAPIID *string) error {
	p := apigateway.NewGetResourcesPaginator(svc.GetResourcesRequest(&apigateway.GetResourcesInput{
		RestApiId: restAPIID,
	}))
	for p.Next(context.Background()) {
		for _, resource := range p.CurrentPage().Items {
			resourceID := *resource.Id
			g.Resources = append(g.Resources, terraformutils.NewResource(
				resourceID,
				resourceID,
				"aws_api_gateway_resource",
				"aws",
				map[string]string{
					"path":        aws.StringValue(resource.Path),
					"path_part":   aws.StringValue(resource.PathPart),
					"partent_id":  aws.StringValue(resource.ParentId),
					"rest_api_id": aws.StringValue(restAPIID),
				},
				apiGatewayAllowEmptyValues,
				map[string]interface{}{},
			))
			err := g.loadResourceMethods(svc, restAPIID, resource)
			if err != nil {
				log.Println(err)
			}
		}
	}
	return p.Err()
}

func (g *APIGatewayGenerator) loadModels(svc *apigateway.Client, restAPIID *string) error {
	p := apigateway.NewGetModelsPaginator(svc.GetModelsRequest(&apigateway.GetModelsInput{
		RestApiId: restAPIID,
	}))
	for p.Next(context.Background()) {
		for _, model := range p.CurrentPage().Items {
			resourceID := *model.Id
			g.Resources = append(g.Resources, terraformutils.NewResource(
				resourceID,
				resourceID,
				"aws_api_gateway_model",
				"aws",
				map[string]string{
					"name":         aws.StringValue(model.Name),
					"content_type": aws.StringValue(model.ContentType),
					"schema":       aws.StringValue(model.Schema),
					"rest_api_id":  aws.StringValue(restAPIID),
				},
				apiGatewayAllowEmptyValues,
				map[string]interface{}{},
			))
		}
	}
	return p.Err()
}

func (g *APIGatewayGenerator) loadResourceMethods(svc *apigateway.Client, restAPIID *string, resource apigateway.Resource) error {
	for httpMethod, method := range resource.ResourceMethods {
		methodID := *restAPIID + "/" + *resource.Id + "/" + httpMethod
		authorizationType := "NONE"
		if method.AuthorizationType != nil {
			authorizationType = *method.AuthorizationType
		}

		g.Resources = append(g.Resources, terraformutils.NewResource(
			methodID,
			methodID,
			"aws_api_gateway_method",
			"aws",
			map[string]string{
				"rest_api_id":   *restAPIID,
				"resource_id":   *resource.Id,
				"http_method":   httpMethod,
				"authorization": authorizationType,
			},
			apiGatewayAllowEmptyValues,
			map[string]interface{}{},
		))

		methodDetails, err := svc.GetMethodRequest(&apigateway.GetMethodInput{
			HttpMethod: &httpMethod,
			ResourceId: resource.Id,
			RestApiId:  restAPIID,
		}).Send(context.Background())
		if err != nil {
			return err
		}

		if methodDetails.MethodIntegration != nil {
			typeString, _ := methodDetails.MethodIntegration.Type.MarshalValue()
			g.Resources = append(g.Resources, terraformutils.NewResource(
				methodID,
				methodID,
				"aws_api_gateway_integration",
				"aws",
				map[string]string{
					"rest_api_id": *restAPIID,
					"resource_id": *resource.Id,
					"http_method": httpMethod,
					"type":        typeString,
				},
				apiGatewayAllowEmptyValues,
				map[string]interface{}{},
			))
			integrationDetails, err := svc.GetIntegrationRequest(&apigateway.GetIntegrationInput{
				HttpMethod: &httpMethod,
				ResourceId: resource.Id,
				RestApiId:  restAPIID,
			}).Send(context.Background())
			if err != nil {
				return err
			}

			for responseCode := range integrationDetails.IntegrationResponses {
				integrationResponseID := *restAPIID + "/" + *resource.Id + "/" + httpMethod + "/" + responseCode
				g.Resources = append(g.Resources, terraformutils.NewResource(
					integrationResponseID,
					integrationResponseID,
					"aws_api_gateway_integration_response",
					"aws",
					map[string]string{
						"rest_api_id": *restAPIID,
						"resource_id": *resource.Id,
						"http_method": httpMethod,
						"status_code": responseCode,
					},
					apiGatewayAllowEmptyValues,
					map[string]interface{}{},
				))
			}
		}
		for responseCode := range methodDetails.MethodResponses {
			responseID := *restAPIID + "/" + *resource.Id + "/" + httpMethod + "/" + responseCode

			g.Resources = append(g.Resources, terraformutils.NewResource(
				responseID,
				responseID,
				"aws_api_gateway_method_response",
				"aws",
				map[string]string{
					"rest_api_id": *restAPIID,
					"resource_id": *resource.Id,
					"http_method": httpMethod,
					"status_code": responseCode,
				},
				apiGatewayAllowEmptyValues,
				map[string]interface{}{},
			))
		}
	}
	return nil
}

func (g *APIGatewayGenerator) loadResponses(svc *apigateway.Client, restAPIID *string) error {
	var position *string
	for {
		response, err := svc.GetGatewayResponsesRequest(&apigateway.GetGatewayResponsesInput{
			RestApiId: restAPIID,
			Position:  position,
		}).Send(context.Background())
		if err != nil {
			return err
		}
		for _, response := range response.Items {
			if aws.BoolValue(response.DefaultResponse) {
				continue
			}
			responseTypeString, _ := response.ResponseType.MarshalValue()
			responseID := *restAPIID + "/" + responseTypeString
			g.Resources = append(g.Resources, terraformutils.NewResource(
				responseID,
				responseID,
				"aws_api_gateway_gateway_response",
				"aws",
				map[string]string{
					"rest_api_id":   *restAPIID,
					"response_type": responseTypeString,
				},
				apiGatewayAllowEmptyValues,
				map[string]interface{}{},
			))
		}
		position = response.Position
		if position == nil {
			break
		}
	}
	return nil
}

func (g *APIGatewayGenerator) loadDocumentationParts(svc *apigateway.Client, restAPIID *string) error {
	var position *string
	for {
		response, err := svc.GetDocumentationPartsRequest(&apigateway.GetDocumentationPartsInput{
			RestApiId: restAPIID,
			Position:  position,
		}).Send(context.Background())
		if err != nil {
			return err
		}
		for _, documentationPart := range response.Items {
			documentationPartID := *restAPIID + "/" + *documentationPart.Id
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				documentationPartID,
				documentationPartID,
				"aws_api_gateway_documentation_part",
				"aws",
				apiGatewayAllowEmptyValues,
			))
		}
		position = response.Position
		if position == nil {
			break
		}
	}
	return nil
}

func (g *APIGatewayGenerator) loadAuthorizers(svc *apigateway.Client, restAPIID *string) error {
	var position *string
	for {
		response, err := svc.GetAuthorizersRequest(&apigateway.GetAuthorizersInput{
			RestApiId: restAPIID,
			Position:  position,
		}).Send(context.Background())
		if err != nil {
			return err
		}
		for _, authorizer := range response.Items {
			g.Resources = append(g.Resources, terraformutils.NewResource(
				*authorizer.Id,
				*authorizer.Id,
				"aws_api_gateway_authorizer",
				"aws",
				map[string]string{
					"rest_api_id": *restAPIID,
					"name":        aws.StringValue(authorizer.Name),
				},
				apiGatewayAllowEmptyValues,
				map[string]interface{}{},
			))
		}
		position = response.Position
		if position == nil {
			break
		}
	}
	return nil
}

func (g *APIGatewayGenerator) loadVpcLinks(svc *apigateway.Client) error {
	p := apigateway.NewGetVpcLinksPaginator(svc.GetVpcLinksRequest(&apigateway.GetVpcLinksInput{}))
	for p.Next(context.Background()) {
		for _, vpcLink := range p.CurrentPage().Items {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				*vpcLink.Id,
				*vpcLink.Name,
				"aws_api_gateway_vpc_link",
				"aws",
				apiGatewayAllowEmptyValues))
		}
	}
	return p.Err()
}

func (g *APIGatewayGenerator) loadUsagePlans(svc *apigateway.Client) error {
	p := apigateway.NewGetUsagePlansPaginator(svc.GetUsagePlansRequest(&apigateway.GetUsagePlansInput{}))
	for p.Next(context.Background()) {
		for _, usagePlan := range p.CurrentPage().Items {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				*usagePlan.Id,
				*usagePlan.Name,
				"aws_api_gateway_usage_plan",
				"aws",
				apiGatewayAllowEmptyValues))
		}
	}
	return p.Err()
}
