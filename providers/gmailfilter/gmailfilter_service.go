// Copyright 2020 The Terraformer Authors.
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

package gmailfilter

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/hashicorp/terraform/helper/pathorcontents"
	"golang.org/x/oauth2"
	googleoauth "golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

const gmailUser = "me"

var gmailAPIScopes = []string{
	gmail.GmailLabelsScope,
	gmail.GmailSettingsBasicScope,
}

type GmailfilterService struct { //nolint
	terraformutils.Service
}

func (s *GmailfilterService) gmailService(ctx context.Context) (*gmail.Service, error) {
	creds := s.GetArgs()["credentials"].(string)
	impersonatedEmailAddr := s.GetArgs()["impersonatedUserEmail"].(string)

	tokenSource, err := s.getTokenSource(creds, impersonatedEmailAddr)
	if err != nil {
		return nil, err
	}

	client := oauth2.NewClient(ctx, tokenSource)
	client.Timeout = 30 * time.Second

	svc, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}
	return svc, nil
}

func (s *GmailfilterService) validateCredentials(creds string) error {
	if _, err := os.Stat(creds); err == nil {
		return nil
	}
	if _, err := googleoauth.CredentialsFromJSON(context.Background(), []byte(creds)); err != nil {
		return fmt.Errorf("JSON credentials in %q are not valid: %s", creds, err)
	}
	return nil
}

func (s *GmailfilterService) getTokenSource(creds string, impersonatedEmailAddr string) (oauth2.TokenSource, error) {
	if creds != "" && impersonatedEmailAddr != "" {
		if err := s.validateCredentials(creds); err != nil {
			return nil, err
		}
		contents, _, err := pathorcontents.Read(creds)
		if err != nil {
			return nil, fmt.Errorf("Error loading credentials: %s", err)
		}

		var serviceAccount serviceAccountFile
		if err := parseJSON(&serviceAccount, contents); err != nil {
			return nil, fmt.Errorf("error parsing credentials %q: %s", contents, err)
		}

		conf := jwt.Config{
			Email:      serviceAccount.ClientEmail,
			PrivateKey: []byte(serviceAccount.PrivateKey),
			Scopes:     gmailAPIScopes,
			TokenURL:   "https://oauth2.googleapis.com/token",
		}
		conf.Subject = impersonatedEmailAddr
		return conf.TokenSource(context.Background()), nil
	}

	return googleoauth.DefaultTokenSource(context.Background(), gmailAPIScopes...)
}

type serviceAccountFile struct {
	PrivateKeyID string `json:"private_key_id"`
	PrivateKey   string `json:"private_key"`
	ClientEmail  string `json:"client_email"`
	ClientID     string `json:"client_id"`
}

func parseJSON(result interface{}, contents string) error {
	r := strings.NewReader(contents)
	dec := json.NewDecoder(r)

	return dec.Decode(result)
}
