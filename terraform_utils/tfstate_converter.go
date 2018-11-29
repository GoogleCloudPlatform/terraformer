package terraform_utils

import (
	"encoding/json"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/hashicorp/terraform/flatmap"
	"github.com/hashicorp/terraform/terraform"
)

type TfstateConverter struct{}

func (c TfstateConverter) Convert(pathToTfstate string, metadata map[string]ResourceMetaData) ([]TerraformResource, error) {
	resources := []TerraformResource{}
	data, err := ioutil.ReadFile(pathToTfstate)
	if err != nil {
		return resources, err
	}
	tfState := terraform.State{}
	err = json.Unmarshal(data, &tfState)
	if err != nil {
		return resources, err
	}
	for _, module := range tfState.Modules {
		for key, resource := range module.Resources {
			rawItem := map[string]interface{}{}
			allAttributes := []string{}
			for key := range resource.Primary.Attributes {
				allAttributes = append(allAttributes, key)
			}
			sort.Strings(allAttributes)
			for _, key := range allAttributes {
				if strings.HasSuffix(key, ".#") && resource.Primary.Attributes[key] == "0" {
					delete(resource.Primary.Attributes, key)
					continue
				}
			}

			for _, key := range allAttributes {
				blockName := strings.Split(key, ".")[0]

				if _, exist := rawItem[blockName]; exist {
					continue
				}

				rawItem[blockName] = flatmap.Expand(resource.Primary.Attributes, blockName)
			}
			item := map[string]interface{}{}
			for key, v := range rawItem {
				switch v.(type) {
				case []interface{}:
					item[key] = v
				default:
					if _, exist := metadata[resource.Primary.ID].IgnoreKeys[key]; exist {
						continue
					}
					if v == nil {
						allowEmptyValue := false
						for pattern := range metadata[resource.Primary.ID].AllowEmptyValue {
							if strings.Contains(key, pattern) {
								allowEmptyValue = true
							}
						}
						if !allowEmptyValue {
							continue
						}
					}
					item[key] = v
				}
			}

			for key, value := range metadata[resource.Primary.ID].AdditionalFields {
				item[key] = value
			}
			resources = append(resources, TerraformResource{
				ResourceType: strings.Split(key, ".")[0],
				ResourceName: strings.Split(key, ".")[1],
				Item:         item,
				ID:           resource.Primary.ID,
				Provider:     metadata[resource.Primary.ID].Provider,
			})
		}
	}
	return resources, nil
}
