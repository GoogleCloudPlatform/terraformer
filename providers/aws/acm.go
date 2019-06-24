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
	"log"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/acm"
)

var acmAllowEmptyValues = []string{}

var acmAdditionalFields = map[string]string{}

type ACMGenerator struct {
	AWSService
}

func (g ACMGenerator) createCertificatesResources(svc *acm.ACM) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	err := svc.ListCertificatesPages(
		&acm.ListCertificatesInput{},
		func(certs *acm.ListCertificatesOutput, lastPage bool) bool {
			for _, cert := range certs.CertificateSummaryList {
				certArn := aws.StringValue(cert.CertificateArn)
				certID := extractCertificateUUID(certArn)
				resources = append(resources, terraform_utils.NewResource(
					certArn,
					certID+"_"+strings.TrimSuffix(aws.StringValue(cert.DomainName), "."),
					"aws_acm_certificate",
					"aws",
					map[string]string{
						"domain_name": aws.StringValue(cert.DomainName),
					},
					acmAllowEmptyValues,
					acmAdditionalFields,
				))
			}
			return true
		},
	)

	if err != nil {
		log.Println(err)
		return resources
	}

	return resources
}

// Generate TerraformResources from AWS API,
// create terraform resource for each certificates
func (g *ACMGenerator) InitResources() error {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(g.GetArgs()["region"].(string))})
	svc := acm.New(sess)

	g.Resources = g.createCertificatesResources(svc)
	g.PopulateIgnoreKeys()
	return nil
}

// extractCertificateUUID extracts UUID from ARN
func extractCertificateUUID(arn string) string {
	if i := strings.Index(arn, "/"); i != -1 {
		return arn[i+1:]
	}
	return arn
}
