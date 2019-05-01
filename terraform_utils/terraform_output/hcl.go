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
package terraform_output

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"github.com/hashicorp/terraform/terraform"
)

func OutputHclFiles(resources []terraform_utils.Resource, provider terraform_utils.ProviderGenerator, path string, serviceName string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}
	// create provider file
	providerDataFile, err := terraform_utils.HclPrint(provider.GetProviderData())
	if err != nil {
		return err
	}
	PrintFile(path+"/provider.tf", providerDataFile)

	// create outputs files
	outputs := map[string]interface{}{}
	outputsByResource := map[string]map[string]interface{}{}

	for i, r := range resources {
		outputState := map[string]*terraform.OutputState{}
		outputsByResource[r.InstanceInfo.Type+"_"+r.ResourceName+"_"+r.GetIDKey()] = map[string]interface{}{
			"value": "${" + r.InstanceInfo.Type + "." + r.ResourceName + "." + r.GetIDKey() + "}",
		}
		outputState[r.InstanceInfo.Type+"_"+r.ResourceName+"_"+r.GetIDKey()] = &terraform.OutputState{
			Type:  "string",
			Value: r.InstanceState.Attributes[r.GetIDKey()],
		}
		for _, v := range provider.GetResourceConnections() {
			for k, ids := range v {
				if k == serviceName {
					if _, exist := r.InstanceState.Attributes[ids[1]]; exist {
						key := ids[1]
						if ids[1] == "self_link" || ids[1] == "id" {
							key = r.GetIDKey()
						}
						linkKey := r.InstanceInfo.Type + "_" + r.ResourceName + "_" + key
						outputsByResource[linkKey] = map[string]interface{}{
							"value": "${" + r.InstanceInfo.Type + "." + r.ResourceName + "." + key + "}",
						}
						outputState[linkKey] = &terraform.OutputState{
							Type:  "string",
							Value: r.InstanceState.Attributes[ids[1]],
						}
					}
				}
			}
		}
		resources[i].Outputs = outputState
	}
	if len(outputsByResource) > 0 {
		outputs["output"] = outputsByResource
		outputsFile, err := terraform_utils.HclPrint(outputs)
		if err != nil {
			return err
		}
		PrintFile(path+"/outputs.tf", outputsFile)
	}

	// group by resource by type
	typeOfServices := map[string][]terraform_utils.Resource{}
	for _, r := range resources {
		typeOfServices[r.InstanceInfo.Type] = append(typeOfServices[r.InstanceInfo.Type], r)
	}
	for k, v := range typeOfServices {
		tfFile, err := terraform_utils.HclPrintResource(v, map[string]interface{}{})
		if err != nil {
			return err
		}
		fileName := strings.Replace(k, strings.Split(k, "_")[0]+"_", "", -1)
		err = ioutil.WriteFile(path+"/"+fileName+".tf", tfFile, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func PrintFile(path string, data []byte) {
	err := ioutil.WriteFile(path, data, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}
}
