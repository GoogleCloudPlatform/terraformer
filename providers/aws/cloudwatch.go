// Copyright 2020 The Terraformer Authors.
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
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
)

var cloudwatchAllowEmptyValues = []string{"tags."}

type CloudWatchGenerator struct {
	AWSService
}

func (g *CloudWatchGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}

	cloudwatchSvc := cloudwatch.New(config)
	err := g.createMetricAlarms(cloudwatchSvc)
	if err != nil {
		return err
	}
	err = g.createDashboards(cloudwatchSvc)
	if err != nil {
		return err
	}

	cloudwatcheventsSvc := cloudwatchevents.New(config)
	err = g.createRules(cloudwatcheventsSvc)
	if err != nil {
		return err
	}

	return nil
}

func (g *CloudWatchGenerator) createMetricAlarms(cloudwatchSvc *cloudwatch.Client) error {
	output, err := cloudwatchSvc.DescribeAlarmsRequest(&cloudwatch.DescribeAlarmsInput{}).Send(context.Background())
	if err != nil {
		return err
	}
	for _, metricAlarm := range output.MetricAlarms {
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			*metricAlarm.AlarmName,
			*metricAlarm.AlarmName,
			"aws_cloudwatch_metric_alarm",
			"aws",
			cloudwatchAllowEmptyValues))
	}
	return nil
}

func (g *CloudWatchGenerator) createDashboards(cloudwatchSvc *cloudwatch.Client) error {
	output, err := cloudwatchSvc.ListDashboardsRequest(&cloudwatch.ListDashboardsInput{}).Send(context.Background())
	if err != nil {
		return err
	}
	for _, dashboardEntry := range output.DashboardEntries {
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			*dashboardEntry.DashboardName,
			*dashboardEntry.DashboardName,
			"aws_cloudwatch_dashboard",
			"aws",
			cloudwatchAllowEmptyValues))
	}
	return nil
}

func (g *CloudWatchGenerator) createRules(cloudwatcheventsSvc *cloudwatchevents.Client) error {
	output, err := cloudwatcheventsSvc.ListRulesRequest(&cloudwatchevents.ListRulesInput{}).Send(context.Background())
	if err != nil {
		return err
	}
	for _, rule := range output.Rules {
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			*rule.Name,
			*rule.Name,
			"aws_cloudwatch_event_rule",
			"aws",
			cloudwatchAllowEmptyValues))

		targetResponse, err := cloudwatcheventsSvc.ListTargetsByRuleRequest(&cloudwatchevents.ListTargetsByRuleInput{
			Rule: rule.Name,
		}).Send(context.Background())
		if err != nil {
			return err
		}
		for _, target := range targetResponse.Targets {
			targetRef := *rule.Name + "/" + *target.Id
			g.Resources = append(g.Resources, terraform_utils.NewResource(
				targetRef,
				targetRef,
				"aws_cloudwatch_event_target",
				"aws",
				map[string]string{
					"rule": *rule.Name,
					"target_id": *target.Id,
				},
				cloudwatchAllowEmptyValues,
				map[string]interface{}{}))
		}
	}

	return nil
}
