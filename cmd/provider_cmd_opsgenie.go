package cmd

import (
	"github.com/spf13/cobra"

	opsgenie_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/opsgenie"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

func newCmdOpsgenieImporter(options ImportOptions) *cobra.Command {
	var apiKey string
	cmd := &cobra.Command{
		Use:   "opsgenie",
		Short: "Import current state to Terraform configuration from Opsgenie",
		Long:  "Import current state to Terraform configuration from Opsgenie",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newOpsgenieProvider()
			err := Import(provider, options, []string{apiKey})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newOpsgenieProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "user,team", "")
	cmd.PersistentFlags().StringVarP(&apiKey, "api-key", "", "", "YOUR_OPSGENIE_API_KEY or env param OPSGENIE_API_KEY")
	return cmd
}

func newOpsgenieProvider() terraformutils.ProviderGenerator {
	return &opsgenie_terraforming.OpsgenieProvider{}
}
