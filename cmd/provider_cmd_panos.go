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
	"reflect"
	"strings"

	panos_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/panos"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdPanosImporter(options ImportOptions) *cobra.Command {
	vsys := []string{}
	cmd := &cobra.Command{
		Use:   "panos",
		Short: "Import current state to Terraform configuration from a PAN-OS",
		Long:  "Import current state to Terraform configuration from a PAN-OS",
		RunE: func(cmd *cobra.Command, args []string) error {
			var t interface{}

			if len(vsys) == 0 {
				var err error

				vsys, t, err = panos_terraforming.GetVsysList()
				if err != nil {
					return err
				}
			} else {
				c, err := panos_terraforming.Initialize()
				if err != nil {
					return err
				}

				t = reflect.TypeOf(c)
			}

			resources := panos_terraforming.FilterCallableResources(t, options.Resources)
			options.Resources = resources

			originalPathPattern := options.PathPattern
			for _, v := range vsys {
				provider := newPanosProvider()
				log.Println(provider.GetName() + " importing VSYS " + v)
				options.PathPattern = originalPathPattern
				options.PathPattern = strings.ReplaceAll(options.PathPattern, "{provider}", "{provider}/"+v)

				err := Import(provider, options, []string{v})
				if err != nil {
					return err
				}
			}

			return nil
		},
	}

	cmd.AddCommand(listCmd(newPanosProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "firewall_device_config,firewall_networking,firewall_objects,firewall_policy", "")
	cmd.PersistentFlags().StringSliceVarP(&vsys, "vsys", "", []string{}, "")

	return cmd
}

func newPanosProvider() terraformutils.ProviderGenerator {

	return &panos_terraforming.PanosProvider{}
}
