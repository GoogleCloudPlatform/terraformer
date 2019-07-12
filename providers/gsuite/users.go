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

package gsuite

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/oauth2"

	"google.golang.org/api/option"

	directory "google.golang.org/api/admin/directory/v1"
)

var (
	usersAllowEmptyValues = []string{"tags."}
	usersAttributes       = map[string]string{}
	usersAdditionalFields = map[string]string{}
)

// MonitorGenerator ...
type UsersGenerator struct {
	GSuiteService
}

/*

func (UsersGenerator) createResources(monitors []datadog.Monitor) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	for _, monitor := range monitors {
		resourceName := strconv.Itoa(monitor.GetId())
		resources = append(resources, terraform_utils.NewResource(
			resourceName,
			fmt.Sprintf("monitor_%s", resourceName),
			"gsuite_user",
			"gsuite",
			usersAttributes,
			usersAllowEmptyValues,
			usersAdditionalFields,
		))
	}

	return resources
}
*/

var defaultOauthScopes = []string{
	directory.AdminDirectoryGroupScope,
	directory.AdminDirectoryUserScope,
	directory.AdminDirectoryUserschemaScope,
}

func (g *UsersGenerator) InitResources() error {
	ctx := context.Background()
	j, _ := ioutil.ReadFile(g.GetArgs()["credentials"])
	tok := &oauth2.Token{}
	err1 := json.Unmarshal(j, tok)
	log.Println(err1)
	service, err := directory.NewService(ctx, option.WithTokenSource(oauth2.StaticTokenSource(tok)))
	if err != nil {
		log.Println("NewService", err)
		return err
	}
	d := directory.NewUsersService(service)
	err = d.List().Pages(ctx, func(users *directory.Users) error {
		for _, user := range users.Users {
			log.Println(user)
		}
		return nil
	})
	if err != nil {
		log.Println("NewUsersService", err)
		return err
	}
	os.Exit(0)
	//g.Resources = g.createResources(monitors)
	g.PopulateIgnoreKeys()
	return nil
}
