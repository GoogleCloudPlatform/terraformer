package sg

import (
	"strings"
	"waze/terraform/terraform_utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const maxResults = 1000

var ignoreKey = map[string]bool{
	"arn":      true,
	"owner_id": true,
	"id":       true,
}

var allowEmptyValues = map[string]bool{
	"tags.": true,
}

func createResources(securityGroups []*ec2.SecurityGroup) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	for _, sg := range securityGroups {
		if sg.VpcId == nil {
			continue
		}
		resources = append(resources, terraform_utils.TerraformResource{
			ResourceType: "aws_security_group",
			ResourceName: strings.Trim(aws.StringValue(sg.GroupName), " "),
			Item:         nil,
			ID:           aws.StringValue(sg.GroupId),
			Provider:     "aws",
		})
	}
	return resources
}

func Generate(region string) error {
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
				return err
			}
		} else {
			break
		}
	}
	resources := createResources(securityGroups)
	err = terraform_utils.GenerateTfState(resources)
	if err != nil {
		return err
	}
	converter := terraform_utils.TfstateConverter{
		Provider:        "aws",
		IgnoreKeys:      ignoreKey,
		AllowEmptyValue: allowEmptyValues,
	}
	resources, err = converter.Convert("terraform.tfstate")
	if err != nil {
		return err
	}
	//resources = replaceIDToName(resources)
	err = terraform_utils.GenerateTf(resources, "security_group", region)
	if err != nil {
		return err
	}
	return nil
}

func replaceIDToName(resources []terraform_utils.TerraformResource) []terraform_utils.TerraformResource {
	for _, resource := range resources {
		item := resource.Item.(map[string]interface{})
		if _, exist := item["ingress"]; !exist {
			continue
		}
		ingresses := item["ingress"].([]map[string]interface{})
		for _, ingress := range ingresses {
			if _, exist := ingress["security_groups"]; !exist {
				continue
			}
			security_groups := ingress["security_groups"].([]string)
			renamedSecurity_groups := []string{}
			for _, security_group := range security_groups {
				found := false
				for _, i := range resources {
					if i.ID == security_group {
						renamedSecurity_groups = append(renamedSecurity_groups, "${"+i.ResourceType+"."+i.ResourceName+".id}")
						found = true
						break
					}
				}
				if !found {
					renamedSecurity_groups = append(renamedSecurity_groups, security_group)
				}
			}
			ingress["security_groups"] = renamedSecurity_groups
		}
	}
	return resources
}
