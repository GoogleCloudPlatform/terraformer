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
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"strconv"
	"strings"
)

var ecsAllowEmptyValues = []string{"tags."}

type EcsGenerator struct {
	AWSService
}

func (g *EcsGenerator) InitResources() error {
	sess := g.generateSession()
	svc := ecs.New(sess)

	err := svc.ListClustersPages(&ecs.ListClustersInput{}, func(clusters *ecs.ListClustersOutput, lastPage bool) bool {
		for _, clusterArn := range clusters.ClusterArns {
			arnParts := strings.Split(aws.StringValue(clusterArn), "/")
			clusterName := arnParts[len(arnParts)-1]

			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				aws.StringValue(clusterArn),
				clusterName,
				"aws_ecs_cluster",
				"aws",
				ecsAllowEmptyValues,
			))

			_ = svc.ListServicesPages(&ecs.ListServicesInput{
				Cluster: clusterArn,
			}, func(services *ecs.ListServicesOutput, lastPage bool) bool {
				for _, serviceArn := range services.ServiceArns {
					arnParts := strings.Split(aws.StringValue(serviceArn), "/")
					serviceName := arnParts[len(arnParts)-1]

					serResp, err := svc.DescribeServices(&ecs.DescribeServicesInput{
						Services: []*string{
							aws.String(serviceName),
						},
						Cluster: clusterArn,
					})
					if err != nil {
						fmt.Println(err.Error())
						continue
					}
					serviceDetails := serResp.Services[0]

					g.Resources = append(g.Resources, terraform_utils.NewResource(
						aws.StringValue(serviceArn),
						serviceName,
						"aws_ecs_service",
						"aws",
						map[string]string{
							"task_definition": aws.StringValue(serviceDetails.TaskDefinition),
							"cluster":         clusterName,
							"name":            serviceName,
							"id":              aws.StringValue(serviceArn),
						},
						ecsAllowEmptyValues,
						map[string]interface{}{},
					))
				}
				return !lastPage
			})
		}
		return !lastPage
	})
	if err != nil {
		return err
	}

	taskDefinitionsMap := map[string]terraform_utils.Resource{}
	err = svc.ListTaskDefinitionsPages(&ecs.ListTaskDefinitionsInput{}, func(taskDefinitions *ecs.ListTaskDefinitionsOutput, lastPage bool) bool {
		for _, taskDefinitionArn := range taskDefinitions.TaskDefinitionArns {
			arnParts := strings.Split(aws.StringValue(taskDefinitionArn), ":")
			definitionWithFamily := arnParts[len(arnParts)-2]
			revision, _ := strconv.Atoi(arnParts[len(arnParts)-1])

			// fetch only latest revision of task definitions
			if val, ok := taskDefinitionsMap[definitionWithFamily]; !ok || val.AdditionalFields["revision"].(int) < revision {
				taskDefinitionsMap[definitionWithFamily] = terraform_utils.NewResource(
					aws.StringValue(taskDefinitionArn),
					definitionWithFamily,
					"aws_ecs_task_definition",
					"aws",
					map[string]string{
						"task_definition":       aws.StringValue(taskDefinitionArn),
						"container_definitions": "{}",
						"family":                "test-task",
						"arn":                   aws.StringValue(taskDefinitionArn),
					},
					[]string{},
					map[string]interface{}{
						"revision": revision,
					},
				)
			}
		}

		return !lastPage
	})
	for _, v := range taskDefinitionsMap {
		g.Resources = append(g.Resources, v)
	}
	if err != nil {
		return err
	}

	g.PopulateIgnoreKeys()
	return nil
}
