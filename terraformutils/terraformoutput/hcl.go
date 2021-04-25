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
package terraformoutput

import (
	"github.com/hashicorp/terraform/addrs"
	"github.com/hashicorp/terraform/states"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils/providerwrapper"
)

func OutputHclFiles(resources []terraformutils.Resource, provider terraformutils.ProviderGenerator, path string, serviceName string, isCompact bool, output string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}
	// create provider file
	providerData := provider.GetProviderData()
	providerData["terraform"] = map[string]interface{}{
		"required_providers": []map[string]interface{}{{
			provider.GetName(): []map[string]interface{}{{
				"version": providerwrapper.GetProviderVersion(provider.GetName()),
			}},
		}},
	}

	providerDataFile, err := terraformutils.Print(providerData, map[string]struct{}{}, output, false)
	if err != nil {
		return err
	}
	PrintFile(path+"/provider."+GetFileExtension(output), providerDataFile)

	// create outputs files
	outputs := map[string]interface{}{}
	outputsByResource := map[string]map[string]interface{}{}

	for i, r := range resources {
		outputState := map[string]*states.OutputValue{}
		resourceNameRef := r.Address.Type + "_" + r.Address.Name + "_" + r.GetIDKey()
		outputsByResource[resourceNameRef] = map[string]interface{}{
			"value": "${" + r.Address.Type + "." + r.Address.Name + "." + r.GetIDKey() + "}",
		}
		outputState[r.Address.Type+"_"+r.Address.Name+"_"+r.GetIDKey()] = &states.OutputValue{
			Addr:  addrs.RootModuleInstance.OutputValue(resourceNameRef),
			Value: r.InstanceState.Value.GetAttr(r.GetIDKey()),
		}
		for _, v := range provider.GetResourceConnections() {
			for k, ids := range v {
				if (serviceName != "" && k == serviceName) || (serviceName == "" && k == r.ServiceName()) {
					if r.InstanceState.Value.Type().HasAttribute(ids[1]) {
						key := ids[1]
						if ids[1] == "self_link" || ids[1] == "id" {
							key = r.GetIDKey()
						}
						linkKey := r.Address.Type + "_" + r.Address.Name + "_" + key
						outputsByResource[linkKey] = map[string]interface{}{
							"value": "${" + r.Address.Type + "." + r.Address.Name + "." + key + "}",
						}
						outputState[linkKey] = &states.OutputValue{
							Addr:  addrs.RootModuleInstance.OutputValue(linkKey),
							Value: r.InstanceState.Value.GetAttr(ids[1]),
						}
					}
				}
			}
		}
		resources[i].Outputs = outputState
	}
	if len(outputsByResource) > 0 {
		outputs["output"] = outputsByResource
		outputsFile, err := terraformutils.Print(outputs, map[string]struct{}{}, output, false)
		if err != nil {
			return err
		}
		PrintFile(path+"/outputs."+GetFileExtension(output), outputsFile)
	}

	// group by resource by type
	typeOfServices := map[string][]terraformutils.Resource{}
	for _, r := range resources {
		typeOfServices[r.Address.Type] = append(typeOfServices[r.Address.Type], r)
	}
	if isCompact {
		err := printFile(resources, "resources", path, output)
		if err != nil {
			return err
		}
	} else {
		for k, v := range typeOfServices {
			fileName := strings.ReplaceAll(k, strings.Split(k, "_")[0]+"_", "")
			err := printFile(v, fileName, path, output)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func printFile(v []terraformutils.Resource, fileName, path, output string) error {
	for _, res := range v {
		if res.DataFiles == nil {
			continue
		}
		for fileName, content := range res.DataFiles {
			if err := os.MkdirAll(path+"/data/", os.ModePerm); err != nil {
				return err
			}
			err := ioutil.WriteFile(path+"/data/"+fileName, content, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}

	tfFile, err := terraformutils.HclPrintResource(v, map[string]interface{}{}, output)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path+"/"+fileName+"."+GetFileExtension(output), tfFile, os.ModePerm)
	if err != nil {
		return err
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

func GetFileExtension(outputFormat string) string {
	if outputFormat == "json" {
		return "tf.json"
	}
	return "tf"
}
