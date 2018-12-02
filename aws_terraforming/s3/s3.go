package s3

import (
	"fmt"
	"log"

	"waze/terraform/aws_terraforming/aws_generator"
	"waze/terraform/terraform_utils"

	"github.com/aws/aws-sdk-go/aws/awserr"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var ignoreKey = map[string]bool{
	"^bucket_domain_name$":          true,
	"^bucket_regional_domain_name$": true,
	"^id$":                          true,
	"^acceleration_status$":         true,
}

var allowEmptyValues = map[string]bool{
	"tags.": true,
}

var additionalFields = map[string]string{}

type S3Generator struct {
	aws_generator.BasicGenerator
}

func (S3Generator) createResources(buckets *s3.ListBucketsOutput, region string) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(region)})
	svc := s3.New(sess)
	for _, bucket := range buckets.Buckets {
		resourceName := aws.StringValue(bucket.Name)
		location, err := svc.GetBucketLocation(&s3.GetBucketLocationInput{Bucket: bucket.Name})
		if err != nil {
			log.Println(err)
			continue
		}

		if s3.NormalizeBucketLocation(aws.StringValue(location.LocationConstraint)) == region {
			resources = append(resources, terraform_utils.NewTerraformResource(
				resourceName,
				resourceName,
				"aws_s3_bucket",
				"aws",
				nil,
				map[string]string{}))
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
			resources = append(resources, terraform_utils.NewTerraformResource(
				resourceName,
				resourceName,
				"aws_s3_bucket_policy",
				"aws",
				nil,
				map[string]string{}))
		}

	}
	return resources
}

func (g S3Generator) Generate(region string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(region)})
	svc := s3.New(sess)
	buckets, err := svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return []terraform_utils.TerraformResource{}, map[string]terraform_utils.ResourceMetaData{}, err
	}
	resources := g.createResources(buckets, region)
	metadata := terraform_utils.NewResourcesMetaData(resources, ignoreKey, allowEmptyValues, additionalFields)
	return resources, metadata, nil
}

func (S3Generator) PostGenerateHook(resources []terraform_utils.TerraformResource) ([]terraform_utils.TerraformResource, error) {
	for _, resource := range resources {
		if resource.ResourceType != "aws_s3_bucket_policy" {
			continue
		}
		policy := resource.Item.(interface{}).(map[string]interface{})["policy"].(string)
		resource.Item.(interface{}).(map[string]interface{})["policy"] = fmt.Sprintf(`<<POLICY
%s
POLICY`, policy)
	}
	return resources, nil
}
