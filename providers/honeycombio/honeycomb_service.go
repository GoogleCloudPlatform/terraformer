// Copyright 2022 The Terraformer Authors.
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

package honeycombio

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	hnyclient "github.com/honeycombio/terraform-provider-honeycombio/client"
)

type HoneycombService struct { //nolint
	terraformutils.Service
	datasets   []hnyclient.Dataset
	datasetMap map[string]bool
}

func (s *HoneycombService) newClient() (*hnyclient.Client, error) {
	enableDebug, _ := strconv.ParseBool(os.Getenv("HONEYCOMBIO_DEBUG"))

	client, err := hnyclient.NewClient(&hnyclient.Config{
		APIKey:    s.GetArgs()["api_key"].(string),
		APIUrl:    s.GetArgs()["api_url"].(string),
		UserAgent: fmt.Sprintf("terraformer-honeycombio/%s", honeycombTerraformerProviderVersion),
		Debug:     enableDebug,
	})
	if err != nil {
		return client, err
	}

	ctx := context.TODO()
	datasets := s.GetArgs()["datasets"].([]string)
	datasetMap := make(map[string]bool)
	if len(datasets) == 0 {
		// assume all datasets
		s.datasets, err = client.Datasets.List(ctx)
		if err != nil {
			return client, fmt.Errorf("unable to list Honeycomb datasets: %v", err)
		}
		for _, ds := range s.datasets {
			datasetMap[ds.Name] = true
		}
	} else {
		// verify the provided datasets exist
		for _, d := range datasets {
			ds, err := client.Datasets.Get(ctx, d)
			if err != nil {
				return client, fmt.Errorf("unable to get Honeycomb dataset '%s': %v", d, err)
			}
			s.datasets = append(s.datasets, *ds)
			datasetMap[ds.Name] = true
		}
	}
	s.datasetMap = datasetMap

	return client, nil
}
