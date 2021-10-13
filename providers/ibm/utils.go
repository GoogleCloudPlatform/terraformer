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

package ibm

import (
	"fmt"
	gohttp "net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/authentication"
	"github.com/IBM-Cloud/bluemix-go/http"
	"github.com/IBM-Cloud/bluemix-go/rest"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/dgrijalva/jwt-go"

	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev2/managementv2"
)

// UserConfig ...
type UserConfig struct {
	userID      string
	userEmail   string
	userAccount string
	cloudName   string `default:"bluemix"`
	cloudType   string `default:"public"`
	generation  int    `default:"2"`
}

// EnvFallBack ...
func envFallBack(envs []string, defaultValue string) string {
	for _, k := range envs {
		if v := os.Getenv(k); v != "" {
			return v
		}
	}
	return defaultValue
}

func fetchUserDetails(sess *session.Session, generation int) (*UserConfig, error) {
	config := sess.Config
	user := UserConfig{}
	var bluemixToken string

	if strings.HasPrefix(config.IAMAccessToken, "Bearer") {
		bluemixToken = config.IAMAccessToken[7:len(config.IAMAccessToken)]
	} else {
		bluemixToken = config.IAMAccessToken
	}

	token, err := jwt.Parse(bluemixToken, func(token *jwt.Token) (interface{}, error) {
		return "", nil
	})
	// TODO validate with key
	if err != nil && !strings.Contains(err.Error(), "key is of invalid type") {
		return &user, err
	}
	claims := token.Claims.(jwt.MapClaims)
	if email, ok := claims["email"]; ok {
		user.userEmail = email.(string)
	}
	user.userID = claims["id"].(string)
	user.userAccount = claims["account"].(map[string]interface{})["bss"].(string)
	iss := claims["iss"].(string)
	if strings.Contains(iss, "https://iam.cloud.ibm.com") {
		user.cloudName = "bluemix"
	} else {
		user.cloudName = "staging"
	}
	user.cloudType = "public"

	user.generation = generation
	return &user, nil
}

func authenticateAPIKey(sess *session.Session) error {
	config := sess.Config
	tokenRefresher, err := authentication.NewIAMAuthRepository(config, &rest.Client{
		DefaultHeader: gohttp.Header{
			"User-Agent": []string{http.UserAgent()},
		},
	})
	if err != nil {
		return err
	}
	return tokenRefresher.AuthenticateAPIKey(config.BluemixAPIKey)
}

func authenticateCF(sess *session.Session) error {
	config := sess.Config
	tokenRefresher, err := authentication.NewUAARepository(config, &rest.Client{
		DefaultHeader: gohttp.Header{
			"User-Agent": []string{http.UserAgent()},
		},
	})
	if err != nil {
		return err
	}
	return tokenRefresher.AuthenticateAPIKey(config.BluemixAPIKey)
}

func GetNext(next interface{}) string {
	if reflect.ValueOf(next).IsNil() {
		return ""
	}

	u, err := url.Parse(reflect.ValueOf(next).Elem().FieldByName("Href").Elem().String())
	if err != nil {
		return ""
	}

	q := u.Query()
	return q.Get("start")
}

// GetNextIAM ...
func GetNextIAM(next interface{}) string {
	if reflect.ValueOf(next).IsNil() {
		return ""
	}

	u, err := url.Parse(reflect.ValueOf(next).Elem().String())
	if err != nil {
		return ""
	}
	q := u.Query()
	return q.Get("pagetoken")
}

func GetResourceGroupID(apiKey, name, region string) (string, error) {
	bmxConfig := &bluemix.Config{
		BluemixAPIKey: apiKey,
		Region:        region,
	}

	sess, err := session.New(bmxConfig)
	if err != nil {
		return "", err
	}

	err = authenticateAPIKey(sess)
	if err != nil {
		return "", err
	}

	generation := envFallBack([]string{"Generation"}, "2")
	gen, err := strconv.Atoi(generation)
	if err != nil {
		return "", err
	}
	userInfo, err := fetchUserDetails(sess, gen)
	if err != nil {
		return "", err
	}

	accountID := userInfo.userAccount
	rsManagementAPI, err := managementv2.New(sess)
	if err != nil {
		return "", err
	}

	rsGroup := rsManagementAPI.ResourceGroup()
	resourceGroupQuery := &managementv2.ResourceGroupQuery{
		AccountID: accountID,
	}
	grp, err := rsGroup.FindByName(resourceGroupQuery, name)
	if err != nil {
		return "", err
	}
	if len(grp) > 0 {
		return grp[0].ID, nil
	}

	return "", fmt.Errorf("Unable to get ID of resource group")
}
