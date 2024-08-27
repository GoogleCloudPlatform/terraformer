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

package bizflycloud

import (
	"context"
	"log"
	"time"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/bizflycloud/gobizfly"
)

type BizflyCloudService struct { //nolint
	terraformutils.Service
}

func (s *BizflyCloudService) generateClient() *gobizfly.Client {
	client, err := gobizfly.NewClient(gobizfly.WithProjectID(s.Args["project_id"].(string)),
		gobizfly.WithRegionName(s.Args["region_name"].(string))) // nolint

	if err != nil {
		log.Fatal(err)
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()
	var tok *gobizfly.Token
	if s.Args["auth_method"] == "password" {
		tok, err = client.Token.Create(ctx, &gobizfly.TokenCreateRequest{
			AuthMethod:    s.Args["auth_method"].(string),
			Username:      s.Args["email"].(string),
			Password:      s.Args["password"].(string),
			AppCredID:     "",
			AppCredSecret: "",
			ProjectID:     s.Args["project_id"].(string),
		},
		)
	} else {
		tok, err = client.Token.Create(ctx, &gobizfly.TokenCreateRequest{
			AuthMethod:    s.Args["auth_method"].(string),
			Username:      "",
			Password:      "",
			AppCredID:     s.Args["app_credential_id"].(string),
			AppCredSecret: s.Args["app_credential_secret"].(string),
			ProjectID:     s.Args["project_id"].(string),
		},
		)
	}

	if err != nil {
		log.Fatal(err)
	}
	client.SetKeystoneToken(tok)

	return client
}
