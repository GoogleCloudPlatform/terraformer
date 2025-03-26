package cmd

import (
	myrasec_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/myrasec"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

//
// newCmdMyrasecImporter
//
func newCmdMyrasecImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "myrasec",
		Short: "Import current state to Terraform configuration from Myra Security",
		Long:  "Import current state to Terraform configuration from Myra Security",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newMyrasecProvider()
			err := Import(provider, options, []string{})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newMyrasecProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "domain", "")
	return cmd
}

//
// newMyrasecProvider
//
func newMyrasecProvider() terraformutils.ProviderGenerator {
	return &myrasec_terraforming.MyrasecProvider{}
}
