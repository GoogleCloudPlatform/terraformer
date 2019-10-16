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
	"github.com/aws/aws-sdk-go/service/eks"
)

var eksAllowEmptyValues = []string{"tags."}

type EksGenerator struct {
	AWSService
}

func (g *EksGenerator) InitResources() error {
	sess := g.generateSession()
	svc := eks.New(sess)

	err := svc.ListClustersPages(&eks.ListClustersInput{}, func(clusters *eks.ListClustersOutput, lastPage bool) bool {
		for _, clusterName := range clusters.Clusters {
			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				aws.StringValue(clusterName),
				aws.StringValue(clusterName),
				"aws_eks_cluster",
				"aws",
				eksAllowEmptyValues,
			))
		}
		return !lastPage
	})
	if err != nil {
		return err
	}

	return nil
}