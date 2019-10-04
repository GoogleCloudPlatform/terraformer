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

package alicloud

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aws/aws-sdk-go/aws/session"

	// "github.com/alicloud/alicloud-sdk-go/alicloud/aliclouderr"
	// "github.com/alicloud/alicloud-sdk-go/alicloud/session"
	// "github.com/alicloud/alicloud-sdk-go/alicloud"
	oss "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var OSSAllowEmptyValues = []string{"tags."}

var OSSAdditionalFields = map[string]string{}

type OSSGenerator struct {
	AlicloudService
}

// createResources iterate on all buckets
// for each bucket we check region and choose only bucket from set region
// for each bucket try get bucket policy, if policy exist create additional NewTerraformResource for policy
func (g OSSGenerator) createResources(sess *session.Session, buckets interface{}, region string) []terraform_utils.Resource {
	fmt.Println("createResources: Not implemented")
	resources := []terraform_utils.Resource{}
	// svc := oss.New(sess)
	// for _, bucket := range buckets.Buckets {
	// 	resourceName := alicloud.StringValue(bucket.Name)
	// 	location, err := svc.GetBucketLocation(&oss.GetBucketLocationInput{Bucket: bucket.Name})
	// 	if err != nil {
	// 		log.Println(err)
	// 		continue
	// 	}
	// 	// check if bucket in region
	// 	if oss.NormalizeBucketLocation(alicloud.StringValue(location.LocationConstraint)) == region {
	// 		resources = append(resources, terraform_utils.NewResource(
	// 			resourceName,
	// 			resourceName,
	// 			"alicloud_oss_bucket",
	// 			"alicloud",
	// 			map[string]string{
	// 				"force_destroy": "false",
	// 				"acl":           "private",
	// 			},
	// 			OSSAllowEmptyValues,
	// 			OSSAdditionalFields))
	// 		// try get policy
	// 		_, err := svc.GetBucketPolicy(&oss.GetBucketPolicyInput{
	// 			Bucket: bucket.Name,
	// 		})

	// 		if err != nil {
	// 			if alicloudErr, ok := err.(aliclouderr.Error); ok && alicloudErr.Code() == "NoSuchBucketPolicy" {
	// 				// Bucket without policy
	// 				continue
	// 			}
	// 			log.Println(err)
	// 			continue
	// 		}
	// 		// if bucket policy exist create TerraformResource with bucket name as ID
	// 		resources = append(resources, terraform_utils.NewResource(
	// 			resourceName,
	// 			resourceName,
	// 			"alicloud_oss_bucket_policy",
	// 			"alicloud",
	// 			map[string]string{},
	// 			OSSAllowEmptyValues,
	// 			OSSAdditionalFields))
	// 	}
	// }
	return resources
}

// Generate TerraformResources from Alicloud API,
// from each oss bucket create 2 TerraformResource(bucket and bucket policy)
// Need bucket name as ID for terraform resource
func (g *OSSGenerator) InitResources() error {
	fmt.Println("WARNING: AssumeRole not supported by upstream SDK (github.com/aliyun/aliyun-oss-go-sdk/oss)")
	client, err := LoadClientFromProfile()
	if err != nil {
		return err
	}
	raw, err := client.WithOssClient(func(ossClient *oss.Client) (interface{}, error) {
		return ossClient.ListBuckets()
	})

	bucketResult := (raw).(oss.ListBucketsResult)
	if err != nil {
		return err
	}
	fmt.Println("number of bucketResult")
	fmt.Println(len(bucketResult.Buckets))
	for _, bucket := range bucketResult.Buckets {
		fmt.Println(bucket.Name)
	}
	fmt.Println("TODO: Resource generation not implemented")
	// g.Resources = g.createResources(sess, buckets, g.GetArgs()["region"].(string))
	// g.PopulateIgnoreKeys()
	return nil
}

// PostGenerateHook for add bucket policy json as heredoc
// support only bucket with policy
func (g *OSSGenerator) PostConvertHook() error {
	fmt.Println("PostConvertHook: Not implemented")
	return nil
}

func (g *OSSGenerator) ParseFilter(rawFilter []string) {
	fmt.Println("ParseFilter: Not implemented")
}
