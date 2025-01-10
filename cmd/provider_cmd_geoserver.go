// Copyright 2018 The Terraformer Authors.
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
	"log"
	"strconv"

	geoserver_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/geoserver"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdGeoServerImporter(options ImportOptions) *cobra.Command {
	geoserverURL := ""
	geowebcacheURL := ""
	user := ""
	password := ""
	insecure := false
	targetWorkspace := ""
	targetDatastore := ""
	cmd := &cobra.Command{
		Use:   "geoserver",
		Short: "Import current state to Terraform configuration from GeoServer",
		Long:  "Import current state to Terraform configuration from GeoServer",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newGeoServerProvider()
			log.Println(provider.GetName() + " importing configuration from " + geoserverURL)
			err := Import(provider, options, []string{geoserverURL, geowebcacheURL, user, password, strconv.FormatBool(insecure), targetWorkspace, targetDatastore})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newGeoServerProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "workspaces", "workspace=id1:id2:id4")
	cmd.PersistentFlags().StringVarP(&geoserverURL, "gs-url", "", "http://localhost:8080/geoserver/rest", "REST API base URL for GeoServer")
	cmd.PersistentFlags().StringVarP(&geowebcacheURL, "gwc-url", "", "https://localhost:8080/geoserver/gwc/rest", "REST API base URL for GeoWebCache")
	cmd.PersistentFlags().StringVarP(&user, "user", "", "admin", "A login which can use the REST API")
	cmd.PersistentFlags().StringVarP(&password, "password", "", "geoserver", "The password for the user")
	cmd.PersistentFlags().BoolVarP(&insecure, "insecure", "", false, "Use insecure TLS")
	cmd.PersistentFlags().StringVarP(&targetWorkspace, "target-workspace", "", "", "The target workspace for workspace dependant resources")
	cmd.PersistentFlags().StringVarP(&targetDatastore, "target-datastore", "", "", "The target datastore for datastore dependant resources")
	return cmd
}

func newGeoServerProvider() terraformutils.ProviderGenerator {
	return &geoserver_terraforming.GeoServerProvider{}
}
