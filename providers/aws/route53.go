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
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
)

var route53AllowEmptyValues = []string{}

var route53AdditionalFields = map[string]interface{}{}

type Route53Generator struct {
	AWSService
}

func (g *Route53Generator) createZonesResources(svc *route53.Client) []terraformutils.Resource {
	var resources []terraformutils.Resource
	p := route53.NewListHostedZonesPaginator(svc, &route53.ListHostedZonesInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			log.Println(err)
			return resources
		}
		for _, zone := range page.HostedZones {
			zoneID := cleanZoneID(StringValue(zone.Id))
			resources = append(resources, terraformutils.NewResource(
				zoneID,
				zoneID+"_"+strings.TrimSuffix(StringValue(zone.Name), "."),
				"aws_route53_zone",
				"aws",
				map[string]string{
					"name":          StringValue(zone.Name),
					"force_destroy": "false",
				},
				route53AllowEmptyValues,
				route53AdditionalFields,
			))
			records := g.createRecordsResources(svc, zoneID)
			resources = append(resources, records...)
		}
	}
	return resources
}

func (Route53Generator) createRecordsResources(svc *route53.Client, zoneID string) []terraformutils.Resource {
	var resources []terraformutils.Resource
	var sets *route53.ListResourceRecordSetsOutput
	var err error
	listParams := &route53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(zoneID),
	}

	for {
		sets, err = svc.ListResourceRecordSets(context.TODO(), listParams)
		if err != nil {
			log.Println(err)
			return resources
		}
		for _, record := range sets.ResourceRecordSets {
			recordName := wildcardUnescape(StringValue(record.Name))
			typeString := string(record.Type)
			resources = append(resources, terraformutils.NewResource(
				fmt.Sprintf("%s_%s_%s_%s", zoneID, recordName, typeString, StringValue(record.SetIdentifier)),
				fmt.Sprintf("%s_%s_%s_%s", zoneID, recordName, typeString, StringValue(record.SetIdentifier)),
				"aws_route53_record",
				"aws",
				map[string]string{
					"name":           strings.TrimSuffix(recordName, "."),
					"zone_id":        zoneID,
					"type":           typeString,
					"set_identifier": StringValue(record.SetIdentifier),
				},
				route53AllowEmptyValues,
				route53AdditionalFields,
			))
		}

		if sets.IsTruncated {
			listParams.StartRecordName = sets.NextRecordName
			listParams.StartRecordType = sets.NextRecordType
			listParams.StartRecordIdentifier = sets.NextRecordIdentifier
		} else {
			break
		}
	}
	return resources
}

func (Route53Generator) createHealthChecksResources(svc *route53.Client) []terraformutils.Resource {
	var resources []terraformutils.Resource

	p := route53.NewListHealthChecksPaginator(svc, &route53.ListHealthChecksInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			log.Println(err)
			return resources
		}
		for _, healthCheck := range page.HealthChecks {
			healthCheckStringType := string(healthCheck.HealthCheckConfig.Type)

			resources = append(resources, terraformutils.NewSimpleResource(
				StringValue(healthCheck.Id),
				fmt.Sprintf("%s_%s", StringValue(healthCheck.Id), healthCheckStringType),
				"aws_route53_health_check",
				"aws",
				route53AllowEmptyValues,
			))
		}
	}
	return resources
}

// Generate TerraformResources from AWS API,
// create terraform resource for each zone + each record
func (g *Route53Generator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := route53.NewFromConfig(config)

	g.Resources = g.createZonesResources(svc)
	healthCheckResources := g.createHealthChecksResources(svc)
	g.Resources = append(g.Resources, healthCheckResources...)

	return nil
}

func (g *Route53Generator) PostConvertHook() error {
	for i, resource := range g.Resources {
		resourceType := resource.InstanceInfo.Type
		if resourceType == "aws_route53_zone" {
			continue
		}

		if resourceType == "aws_route53_health_check" {
			if _, childHealthChecksExist := resource.Item["child_healthchecks"]; !childHealthChecksExist {
				if _, childHealthCheckThreshholdExist := resource.Item["child_health_threshold"]; childHealthCheckThreshholdExist {
					delete(g.Resources[i].Item, "child_health_threshold")
				}
			}
			continue
		}

		item := resource.Item
		zoneID := item["zone_id"].(string)
		for _, resourceZone := range g.Resources {
			if resourceZone.InstanceInfo.Type != "aws_route53_zone" {
				continue
			}
			if zoneID == resourceZone.InstanceState.ID {
				g.Resources[i].Item["zone_id"] = "${aws_route53_zone." + resourceZone.ResourceName + ".zone_id}"
			}
		}
		if _, aliasExist := resource.Item["alias"]; aliasExist {
			if _, ttlExist := resource.Item["ttl"]; ttlExist {
				delete(g.Resources[i].Item, "ttl")
			}
		}
	}
	return nil
}

func wildcardUnescape(s string) string {
	if strings.Contains(s, "\\052") {
		s = strings.Replace(s, "\\052", "*", 1)
	}
	return s
}

// cleanZoneID is used to remove the leading /hostedzone/
func cleanZoneID(id string) string {
	return cleanPrefix(id, "/hostedzone/")
}

// cleanPrefix removes a string prefix from an ID
func cleanPrefix(id, prefix string) string {
	return strings.TrimPrefix(id, prefix)
}
