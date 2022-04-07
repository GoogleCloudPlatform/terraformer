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
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
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
		newCmdIbmImporter,
		// Cloud
		newCmdDigitalOceanImporter,
		newCmdEquinixMetalImporter,
		newCmdHerokuImporter,
		newCmdLaunchDarklyImporter,
		newCmdLinodeImporter,
		newCmdOpenStackImporter,
		newCmdTencentCloudImporter,
		newCmdVultrImporter,
		newCmdYandexImporter,
		// Infrastructure Software
		newCmdKubernetesImporter,
		newCmdOctopusDeployImporter,
		newCmdRabbitMQImporter,
		// Network
		newCmdCloudflareImporter,
		newCmdFastlyImporter,
		newCmdNs1Importer,
		newCmdPanosImporter,
		// VCS
		newCmdAzureDevOpsImporter,
		newCmdAzureADImporter,
		newCmdGithubImporter,
		newCmdGitLabImporter,
		// Monitoring & System Management
		newCmdDatadogImporter,
		newCmdNewRelicImporter,
		newCmdMackerelImporter,
		newCmdGrafanaImporter,
		newCmdPagerDutyImporter,
		newCmdOpsgenieImporter,
		// Community
		newCmdKeycloakImporter,
		newCmdLogzioImporter,
		newCmdCommercetoolsImporter,
		newCmdMikrotikImporter,
		newCmdXenorchestraImporter,
		newCmdGmailfilterImporter,
		newCmdVaultImporter,
		newCmdOktaImporter,
		newCmdAuth0Importer,
	}
}

func providerGenerators() map[string]func() terraformutils.ProviderGenerator {
	list := make(map[string]func() terraformutils.ProviderGenerator)
	for _, providerGen := range []func() terraformutils.ProviderGenerator{
		// Major Cloud
		newGoogleProvider,
		newAWSProvider,
		newAzureProvider,
		newAliCloudProvider,
		newIbmProvider,
		// Cloud
		newDigitalOceanProvider,
		newEquinixMetalProvider,
		newFastlyProvider,
		newHerokuProvider,
		newLaunchDarklyProvider,
		newLinodeProvider,
		newNs1Provider,
		newOpenStackProvider,
		newTencentCloudProvider,
		newVultrProvider,
		// Infrastructure Software
		newKubernetesProvider,
		newOctopusDeployProvider,
		newRabbitMQProvider,
		// Network
		newCloudflareProvider,
		// VCS
		newAzureDevOpsProvider,
		newAzureADProvider,
		newGitHubProvider,
		newGitLabProvider,
		// Monitoring & System Management
		newDataDogProvider,
		newNewRelicProvider,
		newPagerDutyProvider,
		// Community
		newKeycloakProvider,
		newLogzioProvider,
		newCommercetoolsProvider,
		newMikrotikProvider,
		newXenorchestraProvider,
		newGmailfilterProvider,
		newVaultProvider,
		newOktaProvider,
		newAuth0Provider,
	} {
		list[providerGen().GetName()] = providerGen
	}
	return list
}
