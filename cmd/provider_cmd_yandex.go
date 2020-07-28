package cmd

import (
	"log"
	"strings"

	yandex_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/yandex"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdYandexImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "yandex",
		Short: "Import current state to Terraform configuration from Yandex Cloud",
		Long:  "Import current state to Terraform configuration from Yandex Cloud",
		RunE: func(cmd *cobra.Command, args []string) error {

			originalPathPattern := options.PathPattern
			// iterate over provided folder_ids
			for _, folderID := range options.Projects {
				provider := newYandexProvider()
				options.PathPattern = originalPathPattern
				options.PathPattern = strings.Replace(options.PathPattern, "{provider}/{service}", "{provider}/"+folderID+"/{service}", -1)
				log.Println(provider.GetName() + " importing folder id " + folderID)
				err := Import(provider, options, []string{folderID})
				if err != nil {
					return err
				}
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newYandexProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "instance,disk", "")
	cmd.Flags().StringSliceVarP(&options.Projects, "folder_ids", "", []string{}, "folder_id_1,folder_id_2")
	_ = cmd.MarkFlagRequired("folder_ids")
	return cmd
}

func newYandexProvider() terraformutils.ProviderGenerator {
	return &yandex_terraforming.YandexProvider{}
}
