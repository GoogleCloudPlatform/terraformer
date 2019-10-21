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

package alicloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/GoogleCloudPlatform/terraformer/providers/alicloud/connectivity"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
)

// AliCloudService Service struct for AliCloud
type AliCloudService struct {
	terraform_utils.Service
}

// ConfigFile go struct for ~/.aliyun/config.json
type ConfigFile struct {
	Current  string `json:"current"`
	MetaPath string `json:"meta_path"`
	Profiles []struct {
		AccessKeyID     string `json:"access_key_id"`
		AccessKeySecret string `json:"access_key_secret"`
		ExpiredSeconds  int    `json:"expired_seconds"`
		KeyPairName     string `json:"key_pair_name"`
		Language        string `json:"language"`
		Mode            string `json:"mode"`
		Name            string `json:"name"`
		OutputFormat    string `json:"output_format"`
		PrivateKey      string `json:"private_key"`
		RAMRoleArn      string `json:"ram_role_arn"`
		RAMRoleName     string `json:"ram_role_name"`
		RAMSessionName  string `json:"ram_session_name"`
		RegionID        string `json:"region_id"`
		RetryCount      int    `json:"retry_count"`
		RetryTimeout    int    `json:"retry_timeout"`
		Site            string `json:"site"`
		StsToken        string `json:"sts_token"`
		Verified        string `json:"verified"`
	} `json:"profiles"`
}

// LoadClientFromProfile Loads profile from ~/.aliyun/config.json and then applies the region from cmd line
func (s *AliCloudService) LoadClientFromProfile() (*connectivity.AliyunClient, error) {
	args := s.GetArgs()
	region := args["region"].(string)
	profileName := args["profile"].(string)

	config, err := LoadConfigFromProfile(profileName)
	if err != nil {
		return nil, err
	}

	config.RegionId = region
	config.Region = connectivity.Region(config.RegionId)

	return config.Client()
}

// LoadConfigFromProfile Loads profile from ~/.aliyun/config.json
func LoadConfigFromProfile(profileName string) (*connectivity.Config, error) {
	// Set the path depending on OS from where to pull the config.json
	profilePath := ""
	if runtime.GOOS == "windows" {
		profilePath = fmt.Sprintf("%s/.aliyun/config.json", os.Getenv("USERPROFILE"))
	} else {
		profilePath = fmt.Sprintf("%s/.aliyun/config.json", os.Getenv("HOME"))
	}

	// Make sure the profile exists
	_, err := os.Stat(profilePath)
	if os.IsNotExist(err) {
		return nil, err
	}

	// Try to parse JSON
	data, err := ioutil.ReadFile(profilePath)
	if err != nil {
		return nil, err
	}
	var configFile ConfigFile
	err = json.Unmarshal(data, &configFile)
	if err != nil {
		return nil, err
	}

	// If profile argument is missing then use the config file
	currentProfile := profileName
	if currentProfile == "" {
		currentProfile = configFile.Current
	}

	// Default to loading the first profile
	config := configFile.Profiles[0]

	// Set profile as current profile
	found := false
	for _, profile := range configFile.Profiles {
		if currentProfile == profile.Name {
			config = profile
			found = true
		}
	}

	if !found {
		fmt.Printf("ERROR: Profile %s not found. Using profile %s\n", profileName, config.Name)
	}

	conf := connectivity.Config{
		AccessKey:          config.AccessKeyID,
		SecretKey:          config.AccessKeySecret,
		EcsRoleName:        config.Name,
		Region:             connectivity.Region(config.RegionID),
		RegionId:           config.RegionID,
		SecurityToken:      config.StsToken,
		RamRoleArn:         config.RAMRoleArn,
		RamRoleSessionName: config.RAMSessionName,
		// OtsInstanceName:    "", // TODO: Figure out what to do with this
		// AccountId:          "", // TODO: Figure out what to do with this
		// RamRolePolicy:      "", // TODO: Figure out what to do with this
	}

	return &conf, nil
}
