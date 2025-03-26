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
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
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
	svc := apigateway.NewFromConfig(config)

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

func (g *APIGatewayGenerator) loadRestApis(svc *apigateway.Client) error {
	p := apigateway.NewGetRestApisPaginator(svc, &apigateway.GetRestApisInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, restAPI := range page.Items {
			if g.shouldFilterRestAPI(restAPI.Tags) {
				continue
			}
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				*restAPI.Id,
				*restAPI.Id+"_"+*restAPI.Name,
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
	return nil
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
	output, err := svc.GetStages(context.TODO(), &apigateway.GetStagesInput{
		RestApiId: restAPIID,
	})
	if err != nil {
		return err
	}
	for _, stage := range output.Item {
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

func (g *APIGatewayGenerator) loadResources(svc *apigateway.Client, restAPIID *string) error {
	p := apigateway.NewGetResourcesPaginator(svc, &apigateway.GetResourcesInput{
		RestApiId: restAPIID,
	})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, resource := range page.Items {
			resourceID := *restAPIID + "/" + *resource.Id
			g.Resources = append(g.Resources, terraformutils.NewResource(
				resourceID,
				resourceID,
				"aws_api_gateway_resource",
				"aws",
				map[string]string{
					"path":        StringValue(resource.Path),
					"path_part":   StringValue(resource.PathPart),
					"partent_id":  StringValue(resource.ParentId),
					"rest_api_id": StringValue(restAPIID),
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
	return nil
}

func (g *APIGatewayGenerator) loadModels(svc *apigateway.Client, restAPIID *string) error {
	p := apigateway.NewGetModelsPaginator(svc, &apigateway.GetModelsInput{
		RestApiId: restAPIID,
	})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return nil
		}
		for _, model := range page.Items {
			resourceID := *restAPIID + "/" + *model.Id
			g.Resources = append(g.Resources, terraformutils.NewResource(
				resourceID,
				resourceID,
				"aws_api_gateway_model",
				"aws",
				map[string]string{
					"name":         StringValue(model.Name),
					"content_type": StringValue(model.ContentType),
					"schema":       StringValue(model.Schema),
					"rest_api_id":  StringValue(restAPIID),
				},
				apiGatewayAllowEmptyValues,
				map[string]interface{}{},
			))
		}
	}
	return nil
}

func (g *APIGatewayGenerator) loadResourceMethods(svc *apigateway.Client, restAPIID *string, resource types.Resource) error {
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

		methodDetails, err := svc.GetMethod(context.TODO(), &apigateway.GetMethodInput{
			HttpMethod: &httpMethod,
			ResourceId: resource.Id,
			RestApiId:  restAPIID,
		})
		if err != nil {
			return err
		}

		if methodDetails.MethodIntegration != nil {
			typeString := string(methodDetails.MethodIntegration.Type)
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
			integrationDetails, err := svc.GetIntegration(context.TODO(), &apigateway.GetIntegrationInput{
				HttpMethod: &httpMethod,
				ResourceId: resource.Id,
				RestApiId:  restAPIID,
			})
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
		response, err := svc.GetGatewayResponses(context.TODO(), &apigateway.GetGatewayResponsesInput{
			RestApiId: restAPIID,
			Position:  position,
		})
		if err != nil {
			return err
		}
		for _, response := range response.Items {
			if response.DefaultResponse {
				continue
			}
			responseTypeString := string(response.ResponseType)
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
		response, err := svc.GetDocumentationParts(context.TODO(), &apigateway.GetDocumentationPartsInput{
			RestApiId: restAPIID,
			Position:  position,
		})
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
		response, err := svc.GetAuthorizers(context.TODO(), &apigateway.GetAuthorizersInput{
			RestApiId: restAPIID,
			Position:  position,
		})
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
					"name":        StringValue(authorizer.Name),
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
	p := apigateway.NewGetVpcLinksPaginator(svc, &apigateway.GetVpcLinksInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, vpcLink := range page.Items {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				*vpcLink.Id,
				*vpcLink.Name,
				"aws_api_gateway_vpc_link",
				"aws",
				apiGatewayAllowEmptyValues))
		}
	}
	return nil
}

func (g *APIGatewayGenerator) loadUsagePlans(svc *apigateway.Client) error {
	p := apigateway.NewGetUsagePlansPaginator(svc, &apigateway.GetUsagePlansInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, usagePlan := range page.Items {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				*usagePlan.Id,
				*usagePlan.Name,
				"aws_api_gateway_usage_plan",
				"aws",
				apiGatewayAllowEmptyValues))
		}
	}
	return nil
}

func (g *APIGatewayGenerator) loadAPIKeys(svc *apigateway.Client) error {
	p := apigateway.NewGetApiKeysPaginator(svc, &apigateway.GetApiKeysInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, apiKey := range page.Items {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				*apiKey.Id,
				*apiKey.Name,
				"aws_api_gateway_api_key",
				"aws",
				apiGatewayAllowEmptyValues))
		}
	}

	return nil
}
