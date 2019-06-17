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

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

var route53AllowEmptyValues = []string{}

var route53AdditionalFields = map[string]string{}

type Route53Generator struct {
	AWSService
}

func (g Route53Generator) createZonesResources(svc *route53.Route53) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	err := svc.ListHostedZonesPages(
		&route53.ListHostedZonesInput{},
		func(zones *route53.ListHostedZonesOutput, lastPage bool) bool {
			for _, zone := range zones.HostedZones {
				zoneID := cleanZoneID(aws.StringValue(zone.Id))
				resources = append(resources, terraform_utils.NewResource(
					zoneID,
					zoneID+"_"+strings.TrimSuffix(aws.StringValue(zone.Name), "."),
					"aws_route53_zone",
					"aws",
					map[string]string{
						"name":          aws.StringValue(zone.Name),
						"force_destroy": "false",
					},
					route53AllowEmptyValues,
					route53AdditionalFields,
				))
				records := g.createRecordsResources(svc, zoneID)
				resources = append(resources, records...)
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

func (Route53Generator) createRecordsResources(svc *route53.Route53, zoneID string) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	listParams := &route53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(zoneID),
	}

	err := svc.ListResourceRecordSetsPages(
		listParams,
		func(recordSet *route53.ListResourceRecordSetsOutput, lastPage bool) bool {
			for _, record := range recordSet.ResourceRecordSets {
				resources = append(resources, terraform_utils.NewResource(
					fmt.Sprintf("%s_%s_%s_%s", zoneID, aws.StringValue(record.Name), aws.StringValue(record.Type), aws.StringValue(record.SetIdentifier)),
					fmt.Sprintf("%s_%s_%s_%s", zoneID, aws.StringValue(record.Name), aws.StringValue(record.Type), aws.StringValue(record.SetIdentifier)),
					"aws_route53_record",
					"aws",
					map[string]string{
						"name":    aws.StringValue(record.Name),
						"zone_id": zoneID,
						"type":    aws.StringValue(record.Type),
						"set_identifier": aws.StringValue(record.SetIdentifier),
					},
					route53AllowEmptyValues,
					route53AdditionalFields,
				))
			}
			return true
		},
	)

	if err != nil {
		log.Println(err)
		return []terraform_utils.Resource{}
	}
	return resources
}

// Generate TerraformResources from AWS API,
// create terraform resource for each zone + each record
func (g *Route53Generator) InitResources() error {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(g.GetArgs()["region"])})
	svc := route53.New(sess)

	g.Resources = g.createZonesResources(svc)
	g.PopulateIgnoreKeys()
	return nil
}

func (g *Route53Generator) PostConvertHook() error {
	for i, resourceRecord := range g.Resources {
		if resourceRecord.InstanceInfo.Type == "aws_route53_zone" {
			continue
		}
		item := resourceRecord.Item
		zoneID := item["zone_id"].(string)
		for _, resourceZone := range g.Resources {
			if resourceZone.InstanceInfo.Type != "aws_route53_zone" {
				continue
			}
			if zoneID == resourceZone.InstanceState.ID {
				g.Resources[i].Item["zone_id"] = "${aws_route53_zone." + resourceZone.ResourceName + ".zone_id}"
			}
		}
		if _, aliasExist := resourceRecord.Item["alias"]; aliasExist {
			if _, ttlExist := resourceRecord.Item["ttl"]; ttlExist {
				delete(g.Resources[i].Item, "ttl")
			}
		}
	}
	return nil
}

// cleanZoneID is used to remove the leading /hostedzone/
func cleanZoneID(ID string) string {
	return cleanPrefix(ID, "/hostedzone/")
}

// cleanPrefix removes a string prefix from an ID
func cleanPrefix(ID, prefix string) string {
	return strings.TrimPrefix(ID, prefix)
}
