package cmd

import (
	octopusdeploy_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/octopusdeploy"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/spf13/cobra"
)

func newCmdOctopusDeployImporter(options ImportOptions) *cobra.Command {
	var server, apiKey string
	cmd := &cobra.Command{
		Use:   "octopusdeploy",
		Short: "Import current state to Terraform configuration from Octopus Deploy",
		Long:  "Import current state to Terraform configuration from Octopus Deploy",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newOctopusDeployProvider()
			options.PathPattern = "{output}/{provider}/"
			err := Import(provider, options, []string{server, apiKey})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newOctopusDeployProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "octopusdeploy", "tagset")
	cmd.PersistentFlags().StringVar(&server, "server", "", "Octopus Server's API endpoint or env param OCTOPUS_CLI_SERVER")
	cmd.PersistentFlags().StringVar(&apiKey, "apikey", "", "Octopus API key or env param OCTOPUS_CLI_API_KEY")
	return cmd
}

func newOctopusDeployProvider() terraform_utils.ProviderGenerator {
	return &octopusdeploy_terraforming.OctopusDeployProvider{}
}
