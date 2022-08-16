package cmd

import (
	"os"

	squadcast_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/squadcast"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

const (
	defaultSquadcastEndpoint = "http://api.squadcast.com/v3/"
)

func newCmdSquadcastImporter(options ImportOptions) *cobra.Command {
	var accessToken string
	// accessToken := ""
	cmd := &cobra.Command{
		Use:   "squadcast",
		Short: "Import current state to Terraform configuration from SquadCast",
		Long:  "Import current state to Terraform configuration from SquadCast",
		RunE: func(cmd *cobra.Command, args []string) error {
			endpoint := os.Getenv("SQUADCAST_SERVER_URL")
			if len(endpoint) == 0 {
				endpoint = defaultSquadcastEndpoint
			}
			provider := newSquadcastProvider()
			err := Import(provider, options, []string{accessToken})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newSquadcastProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "user", "")
	cmd.PersistentFlags().StringVarP(&accessToken, "access-token", "", "", "YOUR_SQUADCAST_ACCESS_TOKEN or env param SQUADCAST_ACCESS_TOKEN")
	return cmd
}

func newSquadcastProvider() terraformutils.ProviderGenerator {
	return &squadcast_terraforming.SquadcastProvider{}
}
