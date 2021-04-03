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
	svc := cloudhsmv2.NewFromConfig(config)

	p := cloudhsmv2.NewDescribeClustersPaginator(svc, &cloudhsmv2.DescribeClustersInput{})
	for p.HasMorePages() {
		page, e := p.NextPage(context.TODO())
		if e != nil {
			return e
		}
		for _, cluster := range page.Clusters {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				StringValue(cluster.ClusterId),
				StringValue(cluster.ClusterId),
				"aws_cloudhsm_v2_cluster",
				"aws",
				cloudHsmAllowEmptyValues,
			))

			for _, hsm := range cluster.Hsms {
				g.Resources = append(g.Resources, terraformutils.NewResource(
					StringValue(hsm.HsmId),
					StringValue(hsm.HsmId),
					"aws_cloudhsm_v2_hsm",
					"aws",
					map[string]string{
						"cluster_id": StringValue(hsm.ClusterId),
					},
					cloudHsmAllowEmptyValues,
					map[string]interface{}{},
				))

			}
		}
	}
	return nil
}
