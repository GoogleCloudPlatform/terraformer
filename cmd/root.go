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

package cmd

import (
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       version,
	}
	cmd.AddCommand(newImportCmd())
	cmd.AddCommand(newPlanCmd())
	cmd.AddCommand(versionCmd)
	return cmd
}

func Execute() error {
	cmd := NewCmdRoot()
	return cmd.Execute()
}

func providerImporterSubcommands() []func(options ImportOptions) *cobra.Command {
	return []func(options ImportOptions) *cobra.Command{
		// Major Cloud
		newCmdGoogleImporter,
		newCmdAwsImporter,
		newCmdAzureImporter,
		newCmdAliCloudImporter,
		// Cloud
		newCmdDigitalOceanImporter,
		newCmdFastlyImporter,
		newCmdHerokuImporter,
		newCmdLinodeImporter,
		newCmdOpenStackImporter,
		newCmdVultrImporter,
		// Infrastructure Software
		newCmdKubernetesImporter,
		newCmdOctopusDeployImporter,
		newCmdRabbitMQImporter,
		// Network
		newCmdCloudflareImporter,
		// VCS
		newCmdGithubImporter,
		// Monitoring & System Management
		newCmdDatadogImporter,
		newCmdNewRelicImporter,
		// Community
		newCmdKeycloakImporter,
		newCmdLogzioImporter,
		newCmdCommercetoolsImporter,
		newCmdMikrotikImporter,
	}
}

func providerGenerators() map[string]func() terraform_utils.ProviderGenerator {
	list := make(map[string]func() terraform_utils.ProviderGenerator)
	for _, providerGen := range []func() terraform_utils.ProviderGenerator{
		// Major Cloud
		newGoogleProvider,
		newAWSProvider,
		newAzureProvider,
		newAliCloudProvider,
		// Cloud
		newDigitalOceanProvider,
		newFastlyProvider,
		newHerokuProvider,
		newLinodeProvider,
		newOpenStackProvider,
		newVultrProvider,
		// Infrastructure Software
		newKubernetesProvider,
		newOctopusDeployProvider,
		newRabbitMQProvider,
		// Network
		newCloudflareProvider,
		// VCS
		newGitHubProvider,
		// Monitoring & System Management
		newDataDogProvider,
		newNewRelicProvider,
		// Community
		newKeycloakProvider,
		newLogzioProvider,
		newCommercetoolsProvider,
		newMikrotikProvider,
	} {
		list[providerGen().GetName()] = providerGen
	}
	return list
}
