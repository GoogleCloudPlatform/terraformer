package s3

import (
	"log"
	"waze/terraform/aws_terraforming/awsGenerator"
	"waze/terraform/terraform_utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var ignoreKey = map[string]bool{
	"bucket_domain_name":          true,
	"bucket_regional_domain_name": true,
	"id":                          true,
	"acceleration_status":         true,
}

var allowEmptyValues = map[string]bool{
	"tags.": true,
}

var additionalFields = map[string]string{
	"force_destroy": "false",
}

type S3Generator struct {
	awsGenerator.BasicGenerator
}

func (S3Generator) createResources(buckets *s3.ListBucketsOutput, region string) []terraform_utils.TerraformResource {
	resoures := []terraform_utils.TerraformResource{}
	for _, bucket := range buckets.Buckets {
		resourceName := aws.StringValue(bucket.Name)
		sess, _ := session.NewSession(&aws.Config{Region: aws.String(region)})
		svc := s3.New(sess)
		location, err := svc.GetBucketLocation(&s3.GetBucketLocationInput{Bucket: bucket.Name})
		if err != nil {
			log.Println(err)
		}
		if aws.StringValue(location.LocationConstraint) == region {
			resoures = append(resoures, terraform_utils.NewTerraformResource(resourceName, resourceName, "aws_s3_bucket", "aws", nil))
		}
	}
	return resoures
}

func (g S3Generator) Generate(region string) error {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(region)})
	svc := s3.New(sess)
	buckets, err := svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return err
	}

	resources := g.createResources(buckets, region)
	err = terraform_utils.GenerateTfState(resources)
	if err != nil {
		return err
	}
	converter := terraform_utils.TfstateConverter{}
	metadata := terraform_utils.NewResourcesMetaData(resources, ignoreKey, allowEmptyValues, additionalFields)
	resources, err = converter.Convert("terraform.tfstate", metadata)
	if err != nil {
		return err
	}
	err = terraform_utils.GenerateTf(resources, "buckets", region)
	if err != nil {
		return err
	}
	return nil

}
