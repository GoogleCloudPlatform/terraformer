package route53

import (
	"fmt"
	"log"
	"strings"

	"waze/terraform/aws_terraforming/aws_generator"
	"waze/terraform/terraform_utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

var route53IgnoreKey = map[string]bool{
	"^id$":          true,
	"^name_servers": true,
	"^zone_id$":     true,
	"^vpc_id":       true,
	"^vpc_region$":  true,
}

var route53AllowEmptyValues = map[string]bool{}

var route53AdditionalFields = map[string]string{}

type Route53Generator struct {
	aws_generator.BasicGenerator
}

func (g Route53Generator) createZonesResources(svc *route53.Route53) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	zones, err := svc.ListHostedZones(&route53.ListHostedZonesInput{})

	if err != nil {
		log.Println(err)
		return resources
	}
	for _, zone := range zones.HostedZones {
		resources = append(resources, terraform_utils.NewTerraformResource(
			aws.StringValue(zone.Id),
			strings.TrimSuffix(aws.StringValue(zone.Name), "."),
			"aws_route53_zone",
			"aws",
			nil,
			map[string]string{
				"name": aws.StringValue(zone.Name),
			},
		))
		records := g.createRecordsResources(svc, aws.StringValue(zone.Id))
		resources = append(resources, records...)
	}
	if err != nil {
		log.Println(err)
		return []terraform_utils.TerraformResource{}
	}
	return resources
}
func (Route53Generator) createRecordsResources(svc *route53.Route53, zoneID string) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	listParams := &route53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(zoneID),
	}
	recordSet, err := svc.ListResourceRecordSets(listParams)
	for _, record := range recordSet.ResourceRecordSets {
		resources = append(resources, terraform_utils.NewTerraformResource(
			fmt.Sprintf("%s_%s_%s", zoneID, aws.StringValue(record.Name), aws.StringValue(record.Type)),
			strings.TrimSuffix(aws.StringValue(record.Name), "."),
			"aws_route53_record",
			"aws",
			nil,
			map[string]string{
				"name":    aws.StringValue(record.Name),
				"zone_id": zoneID,
				"type":    aws.StringValue(record.Type),
			},
		))
	}
	if err != nil {
		log.Println(err)
		return []terraform_utils.TerraformResource{}
	}
	return resources
}

// Generate TerraformResources from AWS API,
// create terraform resource for each zone + each record
func (g Route53Generator) Generate(region string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(region)})
	svc := route53.New(sess)

	resources := g.createZonesResources(svc)
	metadata := map[string]terraform_utils.ResourceMetaData{}
	for _, resource := range resources {
		resourceMeta := terraform_utils.ResourceMetaData{}
		resourceMeta.AllowEmptyValue = route53AllowEmptyValues
		resourceMeta.AdditionalFields = route53AdditionalFields
		switch resource.ResourceType {
		case "aws_route53_record":
			resourceMeta.IgnoreKeys = map[string]bool{
				"^id$":          true,
				"^name_servers": true,
				"^fqdn$":        true,
			}
		case "aws_route53_zone":
			resourceMeta.IgnoreKeys = route53IgnoreKey
		}
		metadata[resource.ID] = resourceMeta
	}
	return resources, metadata, nil
}
