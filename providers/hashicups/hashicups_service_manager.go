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
package hashicups

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	container "github.com/Jeffail/gabs"
)

func (client *Client) GetViaTokenAndPoint(endpoint, token string) (*container.Container, error) {
	req, err := client.MakeRequest("GET", endpoint, token, nil, true)
	if err != nil {
		return nil, err
	}

	obj, _, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if obj == nil {
		return nil, fmt.Errorf("Empty response body")
	}

	return obj, nil
}

func (c *Client) Do(req *http.Request) (*container.Container, *http.Response, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	obj, err := container.ParseJSON(bodyBytes)
	if err != nil {
		log.Printf("Error occured while json parsing %+v", err)
		return nil, resp, err
	}
	return obj, resp, err
}

func (c *Client) MakeRequest(method string, rpath string, token string, body *container.Container, authenticated bool) (*http.Request, error) {
	fURL := c.HostURL + rpath

	req, err := http.NewRequest(method, fURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", token)
	return req, nil
}
