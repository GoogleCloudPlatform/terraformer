package terraform_utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"

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
			item := map[string]interface{}{}
			arrayElements := map[string]map[string]map[string]interface{}{}
			hashElements := map[string]map[string]string{}
			allAttributes := []string{}
			for key := range resource.Primary.Attributes {
				allAttributes = append(allAttributes, key)
			}
			sort.Strings(allAttributes)
			for _, key := range allAttributes {
				keys := strings.Split(key, ".")
				if len(keys) == 2 {
					if keys[1] == "#" {
						arrayElements[keys[0]] = map[string]map[string]interface{}{}
					} else if keys[1] == "%" {
						hashElements[keys[0]] = map[string]string{}
						item[keys[0]] = map[string]string{}
					}
				}
			}
			for _, key := range allAttributes {
				value := resource.Primary.Attributes[key]
				if _, exist := metadata[resource.Primary.ID].IgnoreKeys[key]; exist {
					continue
				}
				if value == "" {
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

				if !strings.Contains(key, ".") {
					item[key] = resource.Primary.Attributes[key]
				} else {
					keys := strings.Split(key, ".")
					blockName := keys[0]
					if _, exist := metadata[resource.Primary.ID].IgnoreKeys[blockName]; exist {
						continue
					}
					if keys[len(keys)-1] == "#" || keys[len(keys)-1] == "%" {
						continue
					}
					if _, exist := arrayElements[blockName]; exist { // array Element
						if len(strings.Split(key, ".")) == 2 {
							item[blockName] = collectArray(allAttributes, blockName, resource.Primary.Attributes, 1)
						} else {
							if _, exist := arrayElements[blockName][keys[1]]; !exist {
								arrayElements[blockName][keys[1]] = map[string]interface{}{}
							}
							if len(keys) == 3 {
								arrayElements[blockName][keys[1]][keys[2]] = value
							} else if len(keys) == 4 {
								if _, exist := arrayElements[blockName][keys[1]][keys[2]]; !exist {
									arrayElements[blockName][keys[1]][keys[2]] = []string{}
								}
								arrayElements[blockName][keys[1]][keys[2]] = append(arrayElements[blockName][keys[1]][keys[2]].([]string), value)
							} else if len(keys) == 5 {
								if _, exist := arrayElements[blockName][keys[1]][keys[2]]; !exist {
									arrayElements[blockName][keys[1]][keys[2]] = map[string]interface{}{}
								}
								arrayElements[blockName][keys[1]][keys[2]].(map[string]interface{})[keys[4]] = value
							} else {
								log.Println(blockName)
							}
						}
					}
					if _, exist := hashElements[blockName]; exist { // hash Element
						item[blockName].(map[string]string)[keys[1]] = value
					}
				}
			}
			for key, elem := range arrayElements {
				if len(elem) == 0 {
					continue
				}
				item[key] = []map[string]interface{}{}
				for _, v := range elem {
					element := map[string]interface{}{}
					for k, value := range v {
						if _, ok := value.(string); ok {
							element[k] = value.(string)
						} else if _, ok := value.([]string); ok {
							element[k] = []string{}
							for _, arrayElem := range value.([]string) {
								element[k] = append(element[k].([]string), arrayElem)
							}
						} else if _, ok := value.(map[string]interface{}); ok {
							element[k] = map[string]interface{}{}
							for kk, vv := range value.(map[string]interface{}) {
								element[k].(map[string]interface{})[kk] = vv.(string)
							}
						} else {
							log.Println(value)
						}
					}
					item[key] = append(item[key].([]map[string]interface{}), element)
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

func collectArray(keys []string, keyPrefix string, attributes map[string]string, part int) []string {
	values := []string{}
	for _, key := range keys {
		if strings.HasSuffix(key, "#") {
			continue
		}
		if strings.HasPrefix(key, keyPrefix)  {
			_, err := strconv.Atoi(strings.Split(key, ".")[part])
			if err != nil {
				collectArray(keys, keyPrefix, attributes, part+1)
			} else if len(strings.Split(key, ".")) == part+1 {
				values = append(values, attributes[key])
			}
		}
	}
	return values
}
