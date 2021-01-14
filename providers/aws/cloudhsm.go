// Copyright 2021 The Terraformer Authors.
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
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/aws"
)

var cloudHsmAllowEmptyValues = []string{"tags."}

type CloudHsmGenerator struct {
	AWSService
}

func (g *CloudHsmGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := cloudhsmv2.New(config)

	p := cloudhsmv2.NewDescribeClustersPaginator(svc.DescribeClustersRequest(&cloudhsmv2.DescribeClustersInput{}))
	for p.Next(context.Background()) {
		for _, cluster := range p.CurrentPage().Clusters {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				aws.StringValue(cluster.ClusterId),
				aws.StringValue(cluster.ClusterId),
				"aws_cloudhsm_v2_cluster",
				"aws",
				cloudHsmAllowEmptyValues,
			))

			for _, hsm := range cluster.Hsms {
				g.Resources = append(g.Resources, terraformutils.NewResource(
					aws.StringValue(hsm.HsmId),
					aws.StringValue(hsm.HsmId),
					"aws_cloudhsm_v2_hsm",
					"aws",
					map[string]string{
						"cluster_id": aws.StringValue(hsm.ClusterId),
					},
					cloudHsmAllowEmptyValues,
					map[string]interface{}{},
				))

			}
		}
	}
	return p.Err()
}
