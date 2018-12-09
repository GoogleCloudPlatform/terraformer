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

package route53

import (
	"fmt"
	"log"
	"strings"

	"waze/terraformer/aws_terraforming/aws_generator"
	"waze/terraformer/terraform_utils"

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
		zoneID := cleanZoneID(aws.StringValue(zone.Id))
		resources = append(resources, terraform_utils.NewTerraformResource(
			zoneID,
			strings.TrimSuffix(aws.StringValue(zone.Name), "."),
			"aws_route53_zone",
			"aws",
			nil,
			map[string]string{
				"name": aws.StringValue(zone.Name),
			},
		))
		records := g.createRecordsResources(svc, zoneID)
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

func (Route53Generator) PostGenerateHook(resources []terraform_utils.TerraformResource) ([]terraform_utils.TerraformResource, error) {
	for i, resourceRecord := range resources {
		if resourceRecord.ResourceType == "aws_route53_zone" {
			continue
		}
		item := resourceRecord.Item.(map[string]interface{})
		zoneID := item["zone_id"].(string)
		for _, resourceZone := range resources {
			if resourceZone.ResourceType != "aws_route53_zone" {
				continue
			}
			if zoneID == resourceZone.ID {
				resources[i].Item.(map[string]interface{})["zone_id"] = "${aws_route53_zone." + resourceZone.ResourceName + ".zone_id}"
			}
		}

	}
	return resources, nil
}

// cleanZoneID is used to remove the leading /hostedzone/
func cleanZoneID(ID string) string {
	return cleanPrefix(ID, "/hostedzone/")
}

// cleanPrefix removes a string prefix from an ID
func cleanPrefix(ID, prefix string) string {
	if strings.HasPrefix(ID, prefix) {
		ID = strings.TrimPrefix(ID, prefix)
	}
	return ID
}
