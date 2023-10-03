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

package rabbitmq

import (
	"io"
	"net/http"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type RBTService struct { //nolint
	terraformutils.Service
}

func (s *RBTService) generateRequest(uri string) ([]byte, error) {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", s.Args["endpoint"].(string)+uri, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(s.Args["username"].(string), s.Args["password"].(string))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
