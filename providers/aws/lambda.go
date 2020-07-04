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

package aws

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

var lambdaAllowEmptyValues = []string{"tags."}

type LambdaGenerator struct {
	AWSService
}

func (g *LambdaGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := lambda.New(config)

	err := g.addFunctions(svc)
	if err != nil {
		return err
	}
	err = g.addEventSourceMappings(svc)
	if err != nil {
		return err
	}
	err = g.addLayerVersions(svc)
	return err
}

func (g *LambdaGenerator) PostConvertHook() error {
	for i, r := range g.Resources {
		if _, exist := r.Item["environment"]; !exist {
			continue
		}
		variables := g.Resources[i].Item["environment"].([]interface{})[0].(map[string]interface{})["variables"]
		g.Resources[i].Item["environment"] = []interface{}{
			map[string]interface{}{
				"variables": []map[string]interface{}{variables.(map[string]interface{})},
			},
		}
	}
	for _, r := range g.Resources {
		if r.InstanceInfo.Type != "aws_lambda_function_event_invoke_config" {
			continue
		}
		if r.InstanceState.Attributes["maximum_event_age_in_seconds"] == "0" {
			delete(r.Item, "maximum_event_age_in_seconds")
		}
	}
	return nil
}

func (g *LambdaGenerator) addFunctions(svc *lambda.Client) error {
	p := lambda.NewListFunctionsPaginator(svc.ListFunctionsRequest(&lambda.ListFunctionsInput{}))
	for p.Next(context.Background()) {
		for _, function := range p.CurrentPage().Functions {
			g.Resources = append(g.Resources, terraformutils.NewResource(
				*function.FunctionArn,
				*function.FunctionName,
				"aws_lambda_function",
				"aws",
				map[string]string{
					"function_name": *function.FunctionName,
				},
				lambdaAllowEmptyValues,
				map[string]interface{}{},
			))

			pi := lambda.NewListFunctionEventInvokeConfigsPaginator(svc.ListFunctionEventInvokeConfigsRequest(
				&lambda.ListFunctionEventInvokeConfigsInput{
					FunctionName: function.FunctionName,
				}))
			for pi.Next(context.Background()) {
				for _, functionEventInvokeConfig := range pi.CurrentPage().FunctionEventInvokeConfigs {
					g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
						*function.FunctionArn,
						"feic_"+*functionEventInvokeConfig.FunctionArn,
						"aws_lambda_function_event_invoke_config",
						"aws",
						lambdaAllowEmptyValues,
					))
				}
			}
			if err := pi.Err(); err != nil {
				return err
			}
		}
	}
	return p.Err()
}

func (g *LambdaGenerator) addEventSourceMappings(svc *lambda.Client) error {
	p := lambda.NewListEventSourceMappingsPaginator(svc.ListEventSourceMappingsRequest(&lambda.ListEventSourceMappingsInput{}))
	for p.Next(context.Background()) {
		for _, mapping := range p.CurrentPage().EventSourceMappings {
			g.Resources = append(g.Resources, terraformutils.NewResource(
				*mapping.UUID,
				*mapping.UUID,
				"aws_lambda_event_source_mapping",
				"aws",
				map[string]string{
					"event_source_arn": *mapping.EventSourceArn,
					"function_name":    *mapping.FunctionArn,
				},
				lambdaAllowEmptyValues,
				map[string]interface{}{},
			))
		}
	}
	return p.Err()
}

func (g *LambdaGenerator) addLayerVersions(svc *lambda.Client) error {
	pl := lambda.NewListLayersPaginator(svc.ListLayersRequest(&lambda.ListLayersInput{}))
	for pl.Next(context.Background()) {
		for _, layer := range pl.CurrentPage().Layers {
			pv := lambda.NewListLayerVersionsPaginator(svc.ListLayerVersionsRequest(&lambda.ListLayerVersionsInput{
				LayerName: layer.LayerName,
			}))
			for pv.Next(context.Background()) {
				for _, layerVersion := range pv.CurrentPage().LayerVersions {
					g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
						*layerVersion.LayerVersionArn,
						*layerVersion.LayerVersionArn,
						"aws_lambda_layer_version",
						"aws",
						lambdaAllowEmptyValues,
					))
				}
			}
			if err := pv.Err(); err != nil {
				return err
			}
		}
	}
	return pl.Err()
}
