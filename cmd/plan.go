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
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

type ImportPlan struct {
	Version          string
	Provider         string
	Options          ImportOptions
	Args             []string
	ImportedResource map[string][]terraformutils.Resource
}

func newPlanCmd() *cobra.Command {
	options := ImportOptions{
		Plan: true,
	}
	cmd := &cobra.Command{
		Use:           "plan",
		Short:         "Plan to import current state to Terraform configuration",
		Long:          "Plan to import current state to Terraform configuration",
		SilenceUsage:  true,
		SilenceErrors: false,
		//Version:       version.String(),
	}

	for _, subcommand := range providerImporterSubcommands() {
		cmd.AddCommand(subcommand(options))
	}
	return cmd
}

func newCmdPlanImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plan",
		Short: "Import planned state to Terraform configuration",
		Long:  "Import planned state to Terraform configuration",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			plan, err := LoadPlanfile(args[0])
			if err != nil {
				return err
			}

			var provider terraformutils.ProviderGenerator
			if providerGen, ok := providerGenerators()[plan.Provider]; ok {
				provider = providerGen()
			} else {
				return fmt.Errorf("unsupported provider: %s", plan.Provider)
			}

			if err = provider.Init(plan.Args); err != nil {
				return err
			}

			for _, service := range plan.Options.Resources {
				if err = provider.InitService(service, options.Verbose); err != nil {
					return err
				}
			}

			return ImportFromPlan(provider, plan)
		},
	}
	return cmd
}

func LoadPlanfile(path string) (*ImportPlan, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	plan := &ImportPlan{}
	dec := json.NewDecoder(f)
	dec.DisallowUnknownFields()
	if err := dec.Decode(plan); err != nil {
		return nil, err
	}

	if plan.Version != version {
		return nil, fmt.Errorf("planfile version did not match. expected: %s, actual: %s", version, plan.Version)
	}

	return plan, nil
}

func ExportPlanFile(plan *ImportPlan, path, filename string) error {
	plan.Version = version

	planfilePath := filepath.Join(path, filename)
	log.Println("Saving planfile to", planfilePath)

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}

	f, err := os.OpenFile(planfilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "\t")
	return enc.Encode(plan)
}
