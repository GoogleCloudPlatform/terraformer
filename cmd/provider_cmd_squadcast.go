package cmd

import (
	"log"

	squadcast_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/squadcast"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdSquadcastImporter(options ImportOptions) *cobra.Command {
	var refreshToken string
	var teamName string

	cmd := &cobra.Command{
		Use:   "squadcast",
		Short: "Import current state to Terraform configuration from Squadcast",
		Long:  "Import current state to Terraform configuration from Squadcast",
		RunE: func(cmd *cobra.Command, args []string) error {
			originalPathPattern := options.PathPattern
			for _, region := range options.Regions {
				provider := newSquadcastProvider()
				options.PathPattern = originalPathPattern
				options.PathPattern += region + "/"
				log.Println(provider.GetName() + " importing region " + region)
				err := Import(provider, options, []string{refreshToken, region, teamName})
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newSquadcastProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "user", "")
	cmd.PersistentFlags().StringVarP(&refreshToken, "refresh-token", "", "", "YOUR_SQUADCAST_REFRESH_TOKEN or env param SQUADCAST_REFRESH_TOKEN")
	cmd.PersistentFlags().StringSliceVarP(&options.Regions, "region", "", []string{}, "eu or us")
	cmd.PersistentFlags().StringVarP(&teamName, "team-name", "", "", "Squadcast team name")
	return cmd
}

func newSquadcastProvider() terraformutils.ProviderGenerator {
	return &squadcast_terraforming.SquadcastProvider{}
}
