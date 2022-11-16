package cmd

import (
	squadcast_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/squadcast"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdSquadcastImporter(options ImportOptions) *cobra.Command {
	var refreshToken, teamName, region, serviceName string

	cmd := &cobra.Command{
		Use:   "squadcast",
		Short: "Import current state to Terraform configuration from Squadcast",
		Long:  "Import current state to Terraform configuration from Squadcast",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newSquadcastProvider()
			options.PathPattern += region + "/"
			err := Import(provider, options, []string{refreshToken, region, teamName, serviceName})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newSquadcastProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "user", "")
	cmd.PersistentFlags().StringVarP(&refreshToken, "refresh-token", "", "", "YOUR_SQUADCAST_REFRESH_TOKEN or env variable SQUADCAST_REFRESH_TOKEN")
	cmd.PersistentFlags().StringVarP(&region, "region", "", "", "eu or us")
	cmd.PersistentFlags().StringVarP(&teamName, "team-name", "", "", "Squadcast team name")
	cmd.PersistentFlags().StringVarP(&serviceName, "service-name", "", "", "Squadcast service name")
	return cmd
}

func newSquadcastProvider() terraformutils.ProviderGenerator {
	return &squadcast_terraforming.SCProvider{}
}
