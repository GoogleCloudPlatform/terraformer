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

package scaleway

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

type ScalewayService struct { //nolint
	terraformutils.Service
}

func (s *ScalewayService) generateClient() *scw.Client {
	region := scw.Region(s.Args["region"].(string))
	client, err := scw.NewClient(
		scw.WithAuth(s.Args["accesskey"].(string), s.Args["secretkey"].(string)),
		scw.WithDefaultOrganizationID(s.Args["organization"].(string)),
		scw.WithDefaultRegion(region),
	)
	if err != nil {
		panic(err)
	}
	return client
}
