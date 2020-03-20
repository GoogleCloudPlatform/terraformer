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
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils/terraformer_string"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go/aws"
)

var apiGatewayAllowEmptyValues = []string{"tags.", "parent_id", "path_part"}

type ApiGatewayGenerator struct {
	AWSService
}

func (g *ApiGatewayGenerator) InitResources() error {
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

func (g *ApiGatewayGenerator) loadRestApis(svc *apigateway.Client) error {
	p := apigateway.NewGetRestApisPaginator(svc.GetRestApisRequest(&apigateway.GetRestApisInput{}))
	for p.Next(context.Background()) {
		for _, restApi := range p.CurrentPage().Items {
			if g.shouldFilterRestApi(restApi.Tags) {
				continue
			}
			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				*restApi.Id,
				*restApi.Name,
				"aws_api_gateway_rest_api",
				"aws",
				apiGatewayAllowEmptyValues))
			if err := g.loadStages(svc, restApi.Id); err != nil {
				return err
			}
			if err := g.loadResources(svc, restApi.Id); err != nil {
				return err
			}
			if err := g.loadModels(svc, restApi.Id); err != nil {
				return err
			}
			if err := g.loadResponses(svc, restApi.Id); err != nil {
				return err
			}
			if err := g.loadDocumentationParts(svc, restApi.Id); err != nil {
				return err
			}
			if err := g.loadAuthorizers(svc, restApi.Id); err != nil {
				return err
			}
		}
	}
	return p.Err()
}

func (g *ApiGatewayGenerator) shouldFilterRestApi(tags map[string]string) bool {
	for _, filter := range g.Filter {
		if strings.HasPrefix(filter.FieldPath, "tags.") && filter.IsApplicable("aws_api_gateway_rest_api") {
			tagName := strings.Replace(filter.FieldPath, "tags.", "", 1)
			if val, ok := tags[tagName]; ok {
				return !terraformer_string.ContainsString(filter.AcceptableValues, val)
			} else {
				return true
			}
		}
	}
	return false
}

func (g *ApiGatewayGenerator) loadStages(svc *apigateway.Client, restApiId *string) error {
	output, err := svc.GetStagesRequest(&apigateway.GetStagesInput{
		RestApiId: restApiId,
	}).Send(context.Background())
	if err != nil {
		return err
	}
	for _, stage := range output.GetStagesOutput.Item {
		stageId := *restApiId + "/" + *stage.StageName
		g.Resources = append(g.Resources, terraform_utils.NewResource(
			stageId,
			stageId,
			"aws_api_gateway_stage",
			"aws",
			map[string]string{
				"rest_api_id": *restApiId,
				"stage_name":  *stage.StageName,
			},
			apiGatewayAllowEmptyValues,
			map[string]interface{}{},
		))
	}
	return nil
}

func (g *ApiGatewayGenerator) loadResources(svc *apigateway.Client, restApiId *string) error {
	p := apigateway.NewGetResourcesPaginator(svc.GetResourcesRequest(&apigateway.GetResourcesInput{
		RestApiId: restApiId,
	}))
	for p.Next(context.Background()) {
		for _, resource := range p.CurrentPage().Items {
			resourceId := *resource.Id
			g.Resources = append(g.Resources, terraform_utils.NewResource(
				resourceId,
				resourceId,
				"aws_api_gateway_resource",
				"aws",
				map[string]string{
					"path":        aws.StringValue(resource.Path),
					"path_part":   aws.StringValue(resource.PathPart),
					"partent_id":  aws.StringValue(resource.ParentId),
					"rest_api_id": aws.StringValue(restApiId),
				},
				apiGatewayAllowEmptyValues,
				map[string]interface{}{},
			))
			g.loadResourceMethods(svc, restApiId, resource)
		}
	}
	return p.Err()
}

func (g *ApiGatewayGenerator) loadModels(svc *apigateway.Client, restApiId *string) error {
	p := apigateway.NewGetModelsPaginator(svc.GetModelsRequest(&apigateway.GetModelsInput{
		RestApiId: restApiId,
	}))
	for p.Next(context.Background()) {
		for _, model := range p.CurrentPage().Items {
			resourceId := *model.Id
			g.Resources = append(g.Resources, terraform_utils.NewResource(
				resourceId,
				resourceId,
				"aws_api_gateway_model",
				"aws",
				map[string]string{
					"name":         aws.StringValue(model.Name),
					"content_type": aws.StringValue(model.ContentType),
					"schema":       aws.StringValue(model.Schema),
					"rest_api_id":  aws.StringValue(restApiId),
				},
				apiGatewayAllowEmptyValues,
				map[string]interface{}{},
			))
		}
	}
	return p.Err()
}

func (g *ApiGatewayGenerator) loadResourceMethods(svc *apigateway.Client, restApiId *string, resource apigateway.Resource) error {
	for httpMethod, method := range resource.ResourceMethods {
		methodId := *restApiId + "/" + *resource.Id + "/" + httpMethod
		authorizationType := "NONE"
		if method.AuthorizationType != nil {
			authorizationType = *method.AuthorizationType
		}

		g.Resources = append(g.Resources, terraform_utils.NewResource(
			methodId,
			methodId,
			"aws_api_gateway_method",
			"aws",
			map[string]string{
				"rest_api_id":   *restApiId,
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
			RestApiId:  restApiId,
		}).Send(context.Background())
		if err != nil {
			return err
		}

		if methodDetails.MethodIntegration != nil {
			typeString, _ := methodDetails.MethodIntegration.Type.MarshalValue()
			g.Resources = append(g.Resources, terraform_utils.NewResource(
				methodId,
				methodId,
				"aws_api_gateway_integration",
				"aws",
				map[string]string{
					"rest_api_id": *restApiId,
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
				RestApiId:  restApiId,
			}).Send(context.Background())
			if err != nil {
				return err
			}

			for responseCode := range integrationDetails.IntegrationResponses {
				integrationResponseId := *restApiId + "/" + *resource.Id + "/" + httpMethod + "/" + responseCode
				g.Resources = append(g.Resources, terraform_utils.NewResource(
					integrationResponseId,
					integrationResponseId,
					"aws_api_gateway_integration_response",
					"aws",
					map[string]string{
						"rest_api_id": *restApiId,
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
			responseId := *restApiId + "/" + *resource.Id + "/" + httpMethod + "/" + responseCode

			g.Resources = append(g.Resources, terraform_utils.NewResource(
				responseId,
				responseId,
				"aws_api_gateway_method_response",
				"aws",
				map[string]string{
					"rest_api_id": *restApiId,
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

func (g *ApiGatewayGenerator) loadResponses(svc *apigateway.Client, restApiId *string) error {
	response, err := svc.GetGatewayResponsesRequest(&apigateway.GetGatewayResponsesInput{
		RestApiId: restApiId,
	}).Send(context.Background())
	if err != nil {
		return err
	}
	for _, response := range response.Items {
		if aws.BoolValue(response.DefaultResponse) {
			continue
		}
		responseTypeString, _ := response.ResponseType.MarshalValue()
		responseId := *restApiId + "/" + responseTypeString
		g.Resources = append(g.Resources, terraform_utils.NewResource(
			responseId,
			responseId,
			"aws_api_gateway_gateway_response",
			"aws",
			map[string]string{
				"rest_api_id":   *restApiId,
				"response_type": responseTypeString,
			},
			apiGatewayAllowEmptyValues,
			map[string]interface{}{},
		))
	}
	return nil
}

func (g *ApiGatewayGenerator) loadDocumentationParts(svc *apigateway.Client, restApiId *string) error {
	response, err := svc.GetDocumentationPartsRequest(&apigateway.GetDocumentationPartsInput{
		RestApiId: restApiId,
	}).Send(context.Background())
	if err != nil {
		return err
	}
	for _, documentationPart := range response.Items {
		documentationPartId := *restApiId + "/" + *documentationPart.Id
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			documentationPartId,
			documentationPartId,
			"aws_api_gateway_documentation_part",
			"aws",
			apiGatewayAllowEmptyValues,
		))
	}
	return nil
}

func (g *ApiGatewayGenerator) loadAuthorizers(svc *apigateway.Client, restApiId *string) error {
	response, err := svc.GetAuthorizersRequest(&apigateway.GetAuthorizersInput{
		RestApiId: restApiId,
	}).Send(context.Background())
	if err != nil {
		return err
	}
	for _, authorizer := range response.Items {
		g.Resources = append(g.Resources, terraform_utils.NewResource(
			*authorizer.Id,
			*authorizer.Id,
			"aws_api_gateway_authorizer",
			"aws",
			map[string]string{
				"rest_api_id": *restApiId,
				"name":        aws.StringValue(authorizer.Name),
			},
			apiGatewayAllowEmptyValues,
			map[string]interface{}{},
		))
	}
	return nil
}

func (g *ApiGatewayGenerator) loadVpcLinks(svc *apigateway.Client) error {
	p := apigateway.NewGetVpcLinksPaginator(svc.GetVpcLinksRequest(&apigateway.GetVpcLinksInput{}))
	for p.Next(context.Background()) {
		for _, vpcLink := range p.CurrentPage().Items {
			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				*vpcLink.Id,
				*vpcLink.Name,
				"aws_api_gateway_vpc_link",
				"aws",
				apiGatewayAllowEmptyValues))
		}
	}
	return p.Err()
}

func (g *ApiGatewayGenerator) loadUsagePlans(svc *apigateway.Client) error {
	p := apigateway.NewGetUsagePlansPaginator(svc.GetUsagePlansRequest(&apigateway.GetUsagePlansInput{}))
	for p.Next(context.Background()) {
		for _, usagePlan := range p.CurrentPage().Items {
			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				*usagePlan.Id,
				*usagePlan.Name,
				"aws_api_gateway_usage_plan",
				"aws",
				apiGatewayAllowEmptyValues))
		}
	}
	return p.Err()
}
