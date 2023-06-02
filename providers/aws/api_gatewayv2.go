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
