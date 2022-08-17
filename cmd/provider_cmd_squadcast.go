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
	var refreshToken string
	var region string
	var teamName string

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
			err := Import(provider, options, []string{refreshToken, region, teamName})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newSquadcastProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "user", "")
	cmd.PersistentFlags().StringVarP(&refreshToken, "refresh-token", "", "", "YOUR_SQUADCAST_REFRESH_TOKEN or env param SQUADCAST_REFRESH_TOKEN")
	cmd.PersistentFlags().StringVarP(&region, "region", "", "", "")
	cmd.PersistentFlags().StringVarP(&teamName, "team-name", "", "", "")
	return cmd
}

func newSquadcastProvider() terraformutils.ProviderGenerator {
	return &squadcast_terraforming.SquadcastProvider{}
}
