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
	"fmt"
	"log"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

var S3AllowEmptyValues = []string{"tags."}

var S3AdditionalFields = map[string]string{}

type S3Generator struct {
	AWSService
}

// createResources iterate on all buckets
// for each bucket we check region and choose only bucket from set region
// for each bucket try get bucket policy, if policy exist create additional NewTerraformResource for policy
func (g S3Generator) createResources(sess *session.Session, buckets *s3.ListBucketsOutput, region string) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	svc := s3.New(sess)
	for _, bucket := range buckets.Buckets {
		resourceName := aws.StringValue(bucket.Name)
		location, err := svc.GetBucketLocation(&s3.GetBucketLocationInput{Bucket: bucket.Name})
		if err != nil {
			log.Println(err)
			continue
		}
		// check if bucket in region
		if s3.NormalizeBucketLocation(aws.StringValue(location.LocationConstraint)) == region {
			resources = append(resources, terraform_utils.NewResource(
				resourceName,
				resourceName,
				"aws_s3_bucket",
				"aws",
				map[string]string{
					"force_destroy": "false",
					"acl":           "private",
				},
				S3AllowEmptyValues,
				S3AdditionalFields))
			// try get policy
			_, err := svc.GetBucketPolicy(&s3.GetBucketPolicyInput{
				Bucket: bucket.Name,
			})

			if err != nil {
				if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "NoSuchBucketPolicy" {
					// Bucket without policy
					continue
				}
				log.Println(err)
				continue
			}
			// if bucket policy exist create TerraformResource with bucket name as ID
			resources = append(resources, terraform_utils.NewResource(
				resourceName,
				resourceName,
				"aws_s3_bucket_policy",
				"aws",
				map[string]string{},
				S3AllowEmptyValues,
				S3AdditionalFields))
		}
	}
	return resources
}

// Generate TerraformResources from AWS API,
// from each s3 bucket create 2 TerraformResource(bucket and bucket policy)
// Need bucket name as ID for terraform resource
func (g *S3Generator) InitResources() error {
	sess := g.generateSession()
	svc := s3.New(sess)
	buckets, err := svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return err
	}
	g.Resources = g.createResources(sess, buckets, g.GetArgs()["region"])
	g.PopulateIgnoreKeys()
	return nil
}

// PostGenerateHook for add bucket policy json as heredoc
// support only bucket with policy
func (g *S3Generator) PostConvertHook() error {
	for i, resource := range g.Resources {
		if resource.InstanceInfo.Type != "aws_s3_bucket_policy" {
			continue
		}
		policy := resource.Item["policy"].(string)
		g.Resources[i].Item["policy"] = fmt.Sprintf(`<<POLICY
%s
POLICY`, policy)
	}
	return nil
}

func (g *S3Generator) ParseFilter(rawFilter []string) {
	g.Filter = map[string][]string{}
	for _, resource := range rawFilter {
		t := strings.Split(resource, "=")
		if len(t) != 2 {
			log.Println("Pattern for filter must be resource_type=id1:id2:id4")
			continue
		}
		resourceName, resourcesID := t[0], t[1]
		g.Filter[resourceName] = strings.Split(resourcesID, ":")
		if resourceName == "aws_s3_bucket" {
			g.Filter["aws_s3_bucket_policy"] = strings.Split(resourcesID, ":")
		}
	}
}