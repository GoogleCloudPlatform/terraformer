package cmd

import (
	"errors"
	"os"

	googleworkspace_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/googleworkspace"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdGoogleWorkspaceImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "googleworkspace",
		Short: "Import current State to terraform configuration from Google Workspace",
		Long:  "Import current State to terraform configuration from Google Workspace",
		RunE: func(cmd *cobra.Command, args []string) error {
			orgID := os.Getenv("GOOGLEWORKSPACE_CUSTOMER_ID")
			if len(orgID) == 0 {
				return errors.New("Google Workspace Org ID must be set through the `GOOGLEWORKSPACE_ORG_ID` env var")
			}
			credentialJson := os.Getenv("GOOGLEWORKSPACE_CREDENTIALS")
			if len(credentialJson) == 0 {
				return errors.New("Path to the credential JSON file for the Google Workspace Service Account must be set through the `GOOGLEWORKSPACE_CREDENTIALS` env var")
			}

			impersonatedUserEmail := os.Getenv("GOOGLEWORKSPACE_IMPERSONATED_USER_EMAIL")
			if len(impersonatedUserEmail) == 0 {
				return errors.New("Email address of a user to impersonate for Google Admin actions must be set through the `GOOGLEWORKSPACE_IMPERSONATED_USER_EMAIL` env var")
			}

			provider := newGoogleWorkspaceProvider()
			err := Import(provider, options, []string{orgID, credentialJson, impersonatedUserEmail})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newGoogleWorkspaceProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "bleh", "bleh=bleh1,i don't know what this does")
	return cmd
}

func newGoogleWorkspaceProvider() terraformutils.ProviderGenerator {
	return &googleworkspace_terraforming.GoogleWorkspaceProvider{}
}
