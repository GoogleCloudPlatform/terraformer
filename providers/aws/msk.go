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
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kafka"
)

var mskAllowEmptyValues = []string{"tags."}

type MskGenerator struct {
	AWSService
}

func (g *MskGenerator) InitResources() error {
	sess := g.generateSession()
	svc := kafka.New(sess)

	err := svc.ListClustersPages(&kafka.ListClustersInput{}, func(clusters *kafka.ListClustersOutput, lastPage bool) bool {
		for _, clusterInfo := range clusters.ClusterInfoList {
			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				aws.StringValue(clusterInfo.ClusterArn),
				aws.StringValue(clusterInfo.ClusterName),
				"aws_msk_cluster",
				"aws",
				mskAllowEmptyValues,
			))
		}
		return !lastPage
	})
	if err != nil {
		return err
	}

	return nil
}

func (g *MskGenerator) PostConvertHook() error {
	for _, r := range g.Resources {
		if r.InstanceInfo.Type != "aws_msk_cluster" {
			continue
		}
		if r.InstanceState.Attributes["configuration_info.0.revision"] == "0" {
			delete(r.Item, "configuration_info")
		}
	}
	return nil
}
