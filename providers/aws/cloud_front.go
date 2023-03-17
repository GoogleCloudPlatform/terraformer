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

	if err := g.loadDistribution(svc); err != nil {
		return err
	}

	if err := g.loadCachePolicy(svc); err != nil {
		return err
	}

	return nil
}

func (g *CloudFrontGenerator) loadDistribution(svc *cloudfront.Client) error {
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

func (g *CloudFrontGenerator) loadCachePolicy(svc *cloudfront.Client) error {
	var marker *string
	for {
		out, err := svc.ListCachePolicies(context.TODO(), &cloudfront.ListCachePoliciesInput{
			Marker: marker,
		})
		if err != nil {
			return err
		}
		for _, cachePolicy := range out.CachePolicyList.Items {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				StringValue(cachePolicy.CachePolicy.Id),
				StringValue(cachePolicy.CachePolicy.Id),
				"aws_cloudfront_cache_policy",
				"aws",
				cloudFrontAllowEmptyValues,
			))
		}
		marker = out.CachePolicyList.NextMarker
		if marker == nil {
			break
		}
	}
	return nil
}

func (g *CloudFrontGenerator) PostConvertHook() error {
	for i, r := range g.Resources {
		if r.InstanceInfo.Type != "aws_cloudfront_distribution" {
			continue
		}

		for _, cachePolicy := range g.Resources {
			if cachePolicy.InstanceInfo.Type != "aws_cloudfront_cache_policy" {
				continue
			}

			if defaultCacheBehavior, ok := r.Item["default_cache_behavior"].([]interface{})[0].(map[string]interface{})["cache_policy_id"]; ok {
				if defaultCacheBehavior.(string) == cachePolicy.InstanceState.Attributes["id"] {
					g.Resources[i].Item["default_cache_behavior"].([]interface{})[0].(map[string]interface{})["cache_policy_id"] = "${aws_cloudfront_cache_policy." + cachePolicy.ResourceName + ".id}"
				}
			}

			if orderedCacheBehavior, ok := r.Item["ordered_cache_behavior"].([]interface{}); ok {
				for j, behavior := range orderedCacheBehavior {
					if behavior, ok := behavior.(map[string]interface{})["cache_policy_id"]; ok && behavior.(string) == cachePolicy.InstanceState.Attributes["id"] {
						g.Resources[i].Item["ordered_cache_behavior"].([]interface{})[j].(map[string]interface{})["cache_policy_id"] = "${aws_cloudfront_cache_policy." + cachePolicy.ResourceName + ".id}"
					}
				}
			}
		}

	}
	return nil
}
