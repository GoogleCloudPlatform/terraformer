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

	"github.com/aws/aws-sdk-go/service/glue"
)

var glueAllowEmptyValues = []string{"tags."}

type GlueGenerator struct {
	AWSService
}

func (g *GlueGenerator) loadGlueCrawlers(svc *glue.Glue) error {
	var GlueCrawlerAllowEmptyValues = []string{"tags."}
	var GlueCrawlerAdditionalFields = map[string]string{}
	crawlers, err := svc.GetCrawlers(&glue.GetCrawlersInput{})
	if err != nil {
		return err
	}

	var resources []terraform_utils.Resource
	for _, crawler := range crawlers.Crawlers {
		resource := terraform_utils.NewResource(*crawler.Name, *crawler.Name,
			"aws_glue_crawler",
			"aws",
			map[string]string{},
			GlueCrawlerAllowEmptyValues, GlueCrawlerAdditionalFields)
		resources = append(resources, resource)
	}
	g.Resources = resources
	return nil
}

// Generate TerraformResources from AWS API,
// from each database create 1 TerraformResource.
// Need only database name as ID for terraform resource
// AWS api support paging
func (g *GlueGenerator) InitResources() error {
	sess := g.generateSession()
	svc := glue.New(sess)

	if err := g.loadGlueCrawlers(svc); err != nil {
		return err
	}
	g.PopulateIgnoreKeys()
	return nil

}
