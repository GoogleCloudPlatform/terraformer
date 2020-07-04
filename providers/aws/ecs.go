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
	"fmt"
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

var ecsAllowEmptyValues = []string{"tags."}

type EcsGenerator struct {
	AWSService
}

func (g *EcsGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := ecs.New(config)

	p := ecs.NewListClustersPaginator(svc.ListClustersRequest(&ecs.ListClustersInput{}))
	for p.Next(context.Background()) {
		for _, clusterArn := range p.CurrentPage().ClusterArns {
			arnParts := strings.Split(clusterArn, "/")
			clusterName := arnParts[len(arnParts)-1]

			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				clusterArn,
				clusterName,
				"aws_ecs_cluster",
				"aws",
				ecsAllowEmptyValues,
			))

			servicePage := ecs.NewListServicesPaginator(svc.ListServicesRequest(&ecs.ListServicesInput{
				Cluster: aws.String(clusterArn),
			}))
			for servicePage.Next(context.Background()) {
				for _, serviceArn := range servicePage.CurrentPage().ServiceArns {
					arnParts := strings.Split(serviceArn, "/")
					serviceName := arnParts[len(arnParts)-1]

					serResp, err := svc.DescribeServicesRequest(&ecs.DescribeServicesInput{
						Services: []string{
							serviceName,
						},
						Cluster: aws.String(clusterArn),
					}).Send(context.Background())
					if err != nil {
						fmt.Println(err.Error())
						continue
					}
					serviceDetails := serResp.Services[0]

					g.Resources = append(g.Resources, terraformutils.NewResource(
						serviceArn,
						clusterName+"_"+serviceName,
						"aws_ecs_service",
						"aws",
						map[string]string{
							"task_definition": aws.StringValue(serviceDetails.TaskDefinition),
							"cluster":         clusterName,
							"name":            serviceName,
							"id":              serviceArn,
						},
						ecsAllowEmptyValues,
						map[string]interface{}{},
					))
				}
			}
			if err := servicePage.Err(); err != nil {
				return err
			}
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	taskDefinitionsMap := map[string]terraformutils.Resource{}
	taskDefinitionsPage := ecs.NewListTaskDefinitionsPaginator(svc.ListTaskDefinitionsRequest(&ecs.ListTaskDefinitionsInput{}))
	for taskDefinitionsPage.Next(context.Background()) {
		for _, taskDefinitionArn := range taskDefinitionsPage.CurrentPage().TaskDefinitionArns {
			arnParts := strings.Split(taskDefinitionArn, ":")
			definitionWithFamily := arnParts[len(arnParts)-2]
			revision, _ := strconv.Atoi(arnParts[len(arnParts)-1])

			// fetch only latest revision of task definitions
			if val, ok := taskDefinitionsMap[definitionWithFamily]; !ok || val.AdditionalFields["revision"].(int) < revision {
				taskDefinitionsMap[definitionWithFamily] = terraformutils.NewResource(
					taskDefinitionArn,
					definitionWithFamily,
					"aws_ecs_task_definition",
					"aws",
					map[string]string{
						"task_definition":       taskDefinitionArn,
						"container_definitions": "{}",
						"family":                "test-task",
						"arn":                   taskDefinitionArn,
					},
					[]string{},
					map[string]interface{}{
						"revision": revision,
					},
				)
			}
		}
	}
	for _, v := range taskDefinitionsMap {
		delete(v.AdditionalFields, "revision")
		g.Resources = append(g.Resources, v)
	}

	return taskDefinitionsPage.Err()
}

func (g *EcsGenerator) PostConvertHook() error {
	for _, r := range g.Resources {
		if r.InstanceInfo.Type != "aws_ecs_service" {
			continue
		}
		if r.InstanceState.Attributes["propagate_tags"] == "NONE" {
			delete(r.Item, "propagate_tags")
		}
		delete(r.Item, "iam_role")
	}

	return nil
}
