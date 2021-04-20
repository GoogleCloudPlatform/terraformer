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

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
)

var configAllowEmptyValues = []string{"tags."}

type ConfigGenerator struct {
	AWSService
}

func (g *ConfigGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	client := configservice.NewFromConfig(config)

	configurationRecorderRefs, err := g.addConfigurationRecorders(client)
	if err != nil {
		return err
	}
	err = g.addConfigRules(client, configurationRecorderRefs)
	if err != nil {
		return err
	}
	err = g.addDeliveryChannels(client, configurationRecorderRefs)
	return err
}

func (g *ConfigGenerator) addConfigurationRecorders(svc *configservice.Client) ([]string, error) {
	configurationRecorders, err := svc.DescribeConfigurationRecorders(context.TODO(),
		&configservice.DescribeConfigurationRecordersInput{})

	if err != nil {
		return nil, err
	}
	var configurationRecorderRefs []string
	for _, configurationRecorder := range configurationRecorders.ConfigurationRecorders {
		name := *configurationRecorder.Name
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			name,
			name,
			"aws_config_configuration_recorder",
			"aws",
			configAllowEmptyValues,
		))
		configurationRecorderRefs = append(configurationRecorderRefs,
			"aws_config_configuration_recorder.tfer--"+name)
	}
	return configurationRecorderRefs, nil
}

func (g *ConfigGenerator) addConfigRules(svc *configservice.Client, configurationRecorderRefs []string) error {
	var nextToken *string

	for {
		configRules, err := svc.DescribeConfigRules(
			context.TODO(),
			&configservice.DescribeConfigRulesInput{
				NextToken: nextToken,
			})

		if err != nil {
			return err
		}
		for _, configRule := range configRules.ConfigRules {
			name := *configRule.ConfigRuleName
			g.Resources = append(g.Resources, terraformutils.NewResource(
				name,
				name,
				"aws_config_config_rule",
				"aws",
				map[string]string{},
				configAllowEmptyValues,
				map[string]interface{}{
					"depends_on": configurationRecorderRefs,
				},
			))
		}
		nextToken = configRules.NextToken
		if nextToken == nil {
			break
		}
	}
	return nil
}

func (g *ConfigGenerator) addDeliveryChannels(svc *configservice.Client, configurationRecorderRefs []string) error {
	deliveryChannels, err := svc.DescribeDeliveryChannels(context.TODO(),
		&configservice.DescribeDeliveryChannelsInput{})

	if err != nil {
		return err
	}
	for _, deliveryChannel := range deliveryChannels.DeliveryChannels {
		name := *deliveryChannel.Name
		g.Resources = append(g.Resources, terraformutils.NewResource(
			name,
			name,
			"aws_config_delivery_channel",
			"aws",
			map[string]string{},
			configAllowEmptyValues,
			map[string]interface{}{
				"depends_on": configurationRecorderRefs,
			},
		))
	}
	return nil
}
