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
	"log"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"

	"github.com/aws/aws-sdk-go-v2/service/acm"
)

var acmAllowEmptyValues = []string{}

var acmAdditionalFields = map[string]interface{}{}

type ACMGenerator struct {
	AWSService
}

func (g *ACMGenerator) createCertificatesResources(svc *acm.Client) []terraformutils.Resource {
	var resources []terraformutils.Resource
	p := acm.NewListCertificatesPaginator(svc, &acm.ListCertificatesInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			log.Println(err)
			return resources
		}
		for _, cert := range page.CertificateSummaryList {
			certArn := *cert.CertificateArn
			certID := extractCertificateUUID(certArn)
			resources = append(resources, terraformutils.NewResource(
				certArn,
				certID+"_"+strings.TrimSuffix(*cert.DomainName, "."),
				"aws_acm_certificate",
				"aws",
				map[string]string{
					"domain_name": *cert.DomainName,
				},
				acmAllowEmptyValues,
				acmAdditionalFields,
			))
		}
	}
	return resources
}

// Generate TerraformResources from AWS API,
// create terraform resource for each certificates
func (g *ACMGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := acm.NewFromConfig(config)

	g.Resources = g.createCertificatesResources(svc)
	return nil
}

// extractCertificateUUID extracts UUID from ARN
func extractCertificateUUID(arn string) string {
	if i := strings.Index(arn, "/"); i != -1 {
		return arn[i+1:]
	}
	return arn
}
