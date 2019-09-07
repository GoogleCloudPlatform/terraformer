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

package aws

import (
	"os"
	"encoding/json"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

type AWSService struct {
	terraform_utils.Service
}

func (s *AWSService) generateSession() *session.Session {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:                 s.GetArgs()["profile"].(string),
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
		SharedConfigState:       session.SharedConfigEnable,
	}))
	creds, _ := sess.Config.Credentials.Get()
	// terraform cannot ask for MFA token, so we need to pass STS session token
	os.Setenv("AWS_ACCESS_KEY_ID", creds.AccessKeyID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", creds.SecretAccessKey)
	os.Setenv("AWS_SESSION_TOKEN", creds.SessionToken)

	return sess
}

func (s *AWSService) PostConvertHook() error {

	for _, r := range s.Resources {
		tagsMap, resourceHasTags := r.Item["tags"]
		if resourceHasTags && len(tagsMap.(map[string]interface{})) > 0 {
			var newTags []map[string]interface{}
			for tagName, tagValue := range r.Item["tags"].(map[string]interface{}) {
				newTag := make(map[string]interface{})
				newTag[tagName] = tagValue
				newTags = append(newTags, newTag)
			}
			r.Item["tags"] = newTags
		}
	}

	dataJsonBytes, _ := json.MarshalIndent(s.Resources[0], "", "  ")
	println(string(dataJsonBytes))

	return nil
}
