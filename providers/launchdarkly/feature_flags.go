// Copyright 2021 The Terraformer Authors.
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

package launchdarkly

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	launchdarkly "github.com/launchdarkly/api-client-go"
)

var featureFlagsAllowEmptyValues = []string{"variations.*.value"}

type FeatureFlagsGenerator struct {
	LaunchDarklyService
}

func (g *FeatureFlagsGenerator) loadFeatureFlagEnv(ctx context.Context, client *launchdarkly.APIClient, projectKey, flagKey string) error {
	ff, _, err := client.FeatureFlagsApi.GetFeatureFlag(ctx, projectKey, flagKey, &launchdarkly.FeatureFlagsApiGetFeatureFlagOpts{})
	if err != nil {
		return err
	}
	for envKey := range ff.Environments {
		resource := terraformutils.NewResource(
			projectKey+"/"+envKey+"/"+flagKey,
			projectKey+"-"+envKey+"-"+flagKey,
			"launchdarkly_feature_flag_environment",
			"launchdarkly",
			map[string]string{
				"env_key": envKey,
				"flag_id": projectKey + "/" + flagKey,
			},
			featureFlagsAllowEmptyValues,
			map[string]interface{}{})
		g.Resources = append(g.Resources, resource)
	}
	return nil
}

func (g *FeatureFlagsGenerator) loadFeatureFlags(ctx context.Context, client *launchdarkly.APIClient, project string) error {
	featureFlags, _, err := client.FeatureFlagsApi.GetFeatureFlags(ctx, project, &launchdarkly.FeatureFlagsApiGetFeatureFlagsOpts{})
	if err != nil {
		return err
	}
	for _, featureFlag := range featureFlags.Items {
		resource := terraformutils.NewResource(
			featureFlag.Key,
			project+"-"+featureFlag.Name,
			"launchdarkly_feature_flag",
			"launchdarkly",
			map[string]string{
				"key":         featureFlag.Key,
				"project_key": project,
			},
			featureFlagsAllowEmptyValues,
			map[string]interface{}{})
		resource.IgnoreKeys = append(resource.IgnoreKeys, "include_in_snippet")
		err = g.loadFeatureFlagEnv(ctx, client, project, featureFlag.Key)
		if err != nil {
			return err
		}
		g.Resources = append(g.Resources, resource)
	}
	return nil
}

func (g *FeatureFlagsGenerator) InitResources() error {
	projects, err := getProjects(g.GetArgs()["ctx"].(context.Context), g.GetArgs()["client"].(*launchdarkly.APIClient))
	if err != nil {
		return err
	}
	for _, project := range projects.Items {
		if err := g.loadFeatureFlags(g.GetArgs()["ctx"].(context.Context), g.GetArgs()["client"].(*launchdarkly.APIClient), project.Key); err != nil {
			return err
		}
	}
	return nil
}
