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
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aws/aws-sdk-go/service/cloudfront"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var cloudFrontAllowEmptyValues = []string{"tags."}

type CloudFrontGenerator struct {
	AWSService
}

func (g *CloudFrontGenerator) InitResources() error {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(g.GetArgs()["region"].(string))})
	svc := cloudfront.New(sess)
	err := svc.ListDistributionsPages(&cloudfront.ListDistributionsInput{}, func(distributions *cloudfront.ListDistributionsOutput, lastPage bool) bool {
		for _, distribution := range distributions.DistributionList.Items {
			r := terraform_utils.NewResource(
				aws.StringValue(distribution.Id),
				aws.StringValue(distribution.Id),
				"aws_cloudfront_distribution",
				"aws",
				map[string]string{
					"retain_on_delete": "false",
				},
				cloudFrontAllowEmptyValues,
				map[string]string{},
			)
			r.IgnoreKeys = append(r.IgnoreKeys, "^active_trusted_signers.(.*)")
			g.Resources = append(g.Resources, r)
		}
		return true
	})
	if err != nil {
		return err
	}
	g.PopulateIgnoreKeys()
	return nil

}
