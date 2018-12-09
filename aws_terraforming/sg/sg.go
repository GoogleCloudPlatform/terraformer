package sg

import (
	"log"
	"strings"

	"waze/terraformer/aws_terraforming/aws_generator"
	"waze/terraformer/terraform_utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const maxResults = 1000

var ignoreKey = map[string]bool{
	"^arn":      true,
	"^owner_id": true,
	"^id$":      true,
}

var allowEmptyValues = map[string]bool{
	"tags.": true,
}

type SecurityGenerator struct {
	aws_generator.BasicGenerator
}

func (SecurityGenerator) createResources(securityGroups []*ec2.SecurityGroup) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	for _, sg := range securityGroups {
		if sg.VpcId == nil {
			continue
		}
		resources = append(resources, terraform_utils.NewTerraformResource(
			aws.StringValue(sg.GroupId),
			strings.Trim(aws.StringValue(sg.GroupName), " "),
			"aws_security_group",
			"aws",
			nil,
			map[string]string{}))
	}
	return resources
}

// Generate TerraformResources from AWS API,
// from each security group create 1 TerraformResource.
// Need GroupId as ID for terraform resource
// AWS support pagination with NextToken patter
func (g SecurityGenerator) Generate(region string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(region)})
	svc := ec2.New(sess)
	var securityGroups []*ec2.SecurityGroup
	var err error
	firstRun := true
	securityGroupsOutput := &ec2.DescribeSecurityGroupsOutput{}
	for {
		if firstRun || securityGroupsOutput.NextToken != nil {
			firstRun = false
			securityGroupsOutput, err = svc.DescribeSecurityGroups(&ec2.DescribeSecurityGroupsInput{
				MaxResults: aws.Int64(maxResults),
				NextToken:  securityGroupsOutput.NextToken,
			})
			securityGroups = append(securityGroups, securityGroupsOutput.SecurityGroups...)
			if err != nil {
				log.Println(err)
			}
		} else {
			break
		}
	}
	resources := g.createResources(securityGroups)
	metadata := terraform_utils.NewResourcesMetaData(resources, ignoreKey, allowEmptyValues, map[string]string{})
	return resources, metadata, nil
}

// PostGenerateHook - replace sg-xxxxx string to terraform ID in all security group
func (g SecurityGenerator) PostGenerateHook(resources []terraform_utils.TerraformResource) ([]terraform_utils.TerraformResource, error) {
	for _, resource := range resources {
		for _, typeOfRule := range []string{"ingress", "egress"} {
			item := resource.Item.(map[string]interface{})
			if _, exist := item[typeOfRule]; !exist {
				continue
			}
			for i, k := range item[typeOfRule].([]interface{}) {
				ingresses := k.(map[string]interface{})
				for key, ingress := range ingresses {
					if key != "security_groups" {
						continue
					}
					securityGroups := ingress.([]interface{})
					renamedSecurityGroups := []string{}
					for _, securityGroup := range securityGroups {
						found := false
						for _, i := range resources {
							if i.ID == securityGroup {
								renamedSecurityGroups = append(renamedSecurityGroups, "${"+i.ResourceType+"."+i.ResourceName+".id}")
								found = true
								break
							}
						}
						if !found {
							renamedSecurityGroups = append(renamedSecurityGroups, securityGroup.(string))
						}
					}
					item[typeOfRule].([]interface{})[i].(map[string]interface{})["security_groups"] = renamedSecurityGroups
				}
			}
		}
	}
	return resources, nil
}
