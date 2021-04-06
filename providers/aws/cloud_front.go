// Copyright 2018 The Terraformer Authors.
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
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
)

var cloudFrontAllowEmptyValues = []string{"tags."}

type CloudFrontGenerator struct {
	AWSService
}

func (g *CloudFrontGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := cloudfront.NewFromConfig(config)
	p := cloudfront.NewListDistributionsPaginator(svc, &cloudfront.ListDistributionsInput{})
	for p.HasMorePages() {
		page, e := p.NextPage(context.TODO())
		if e != nil {
			return e
		}
		for _, distribution := range page.DistributionList.Items {
			r := terraformutils.NewResource(
				StringValue(distribution.Id),
				StringValue(distribution.Id),
				"aws_cloudfront_distribution",
				"aws",
				map[string]string{
					"retain_on_delete": "false",
				},
				cloudFrontAllowEmptyValues,
				map[string]interface{}{},
			)
			r.IgnoreKeys = append(r.IgnoreKeys, "^active_trusted_signers.(.*)")
			g.Resources = append(g.Resources, r)
		}
	}
	return nil
}
