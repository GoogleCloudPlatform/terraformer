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
	"context"
	"os"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/service/sts"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type AWSService struct { //nolint
	terraformutils.Service
}

var awsVariable = regexp.MustCompile(`(\${[0-9A-Za-z:]+})`)

func (s *AWSService) generateConfig() (aws.Config, error) {
	config, e := s.buildBaseConfig()

	if e != nil {
		return config, e
	}
	if s.Verbose {
		config.ClientLogMode = aws.LogRequestWithBody & aws.LogResponseWithBody & aws.LogRetries
	}

	creds, e := config.Credentials.Retrieve(context.TODO())

	if e != nil {
		return config, e
	}

	// terraform cannot ask for MFA token, so we need to pass STS session token, which might contain credentials with MFA requirement
	accessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if accessKey == "" {
		os.Setenv("AWS_ACCESS_KEY_ID", creds.AccessKeyID)
		os.Setenv("AWS_SECRET_ACCESS_KEY", creds.SecretAccessKey)

		if creds.SessionToken != "" {
			os.Setenv("AWS_SESSION_TOKEN", creds.SessionToken)
		}
	}

	return config, nil
}

func (s *AWSService) buildBaseConfig() (aws.Config, error) {
	if s.GetArgs()["profile"].(string) != "" {
		os.Setenv("AWS_PROFILE", s.GetArgs()["profile"].(string))
	}
	if s.GetArgs()["region"].(string) != "" {
		return config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	}
	return config.LoadDefaultConfig(context.TODO())
}

// for CF interpolation and IAM Policy variables
func (*AWSService) escapeAwsInterpolation(str string) string {
	return awsVariable.ReplaceAllString(str, "$$$1")
}

func (s *AWSService) getAccountNumber(config aws.Config) (*string, error) {
	stsSvc := sts.NewFromConfig(config)
	identity, err := stsSvc.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		return nil, err
	}
	return identity.Account, nil
}
