// Copyright 2019 The Terraformer Authors.
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
	mongodbatlas_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/mongodbatlas"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdMongoDBAtlasImporter(options ImportOptions) *cobra.Command {
	var publicKey, privateKey, orgID string
	cmd := &cobra.Command{
		Use:   "mongodbatlas",
		Short: "Import current state to Terraform configuration from MongoDB Atlas",
		Long:  "Import current state to Terraform configuration from MongoDB Atlas",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newMongoDBAtlasProvider()
			err := Import(provider, options, []string{publicKey, privateKey, orgID})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newMongoDBAtlasProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "project,cluster", "project=name1:name2:name3")
	cmd.PersistentFlags().StringVarP(&privateKey, "private-key", "", "", "YOUR_MONGODBATLAS_PRIVATE_KEY or env param MONGODB_ATLAS_PRIVATE_KEY")
	cmd.PersistentFlags().StringVarP(&publicKey, "public-key", "", "", "YOUR_MONGODBATLAS_PUBLIC_KEY or env param MONGODB_ATLAS_PUBLIC_KEY")
	cmd.PersistentFlags().StringVarP(&orgID, "org-id", "", "", "YOUR_MONGODBATLAS_ORG_ID or env param MONGODB_ATLAS_ORG_ID")
	return cmd
}

func newMongoDBAtlasProvider() terraformutils.ProviderGenerator {
	return &mongodbatlas_terraforming.MongoDBAtlasProvider{}
}
