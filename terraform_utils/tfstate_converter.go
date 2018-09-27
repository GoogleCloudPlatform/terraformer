package terraform_utils

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type TfstateResource struct {
	Type      string        `json:"type"`
	DependsOn []interface{} `json:"depends_on"`
	Primary   struct {
		ID         string                 `json:"id"`
		Attributes map[string]interface{} `json:"attributes"`
		Meta       struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}
type Tfstate struct {
	Version          int    `json:"version"`
	TerraformVersion string `json:"terraform_version"`
	Serial           int    `json:"serial"`
	Lineage          string `json:"lineage"`
	Modules          []struct {
		Path    []string `json:"path"`
		Outputs struct {
		} `json:"outputs"`
		Resources map[string]TfstateResource `json:"resources"`
	} `json:"modules"`
}

func TfstateToTfConverter(pathToTfstate, provider string, ignoreKeys map[string]bool) ([]TerraformResource, error) {
	resources := []TerraformResource{}
	data, err := ioutil.ReadFile(pathToTfstate)
	if err != nil {
		return resources, err
	}
	tfState := Tfstate{}
	err = json.Unmarshal(data, &tfState)
	if err != nil {
		return resources, err
	}
	for _, module := range tfState.Modules {
		for key, resource := range module.Resources {
			item := map[string]interface{}{}
			arrayElements := map[string]map[string]map[string]interface{}{}
			hashElements := map[string]map[string]string{}
			for key := range resource.Primary.Attributes {
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
			for key, value := range resource.Primary.Attributes {
				if _, exist := ignoreKeys[key]; exist {
					continue
				}
				if !strings.Contains(key, ".") {
					item[key] = value
				} else {
					keys := strings.Split(key, ".")
					blockName := keys[0]
					if keys[len(keys)-1] == "#" || keys[len(keys)-1] == "%" {
						continue
					}
					if _, exist := arrayElements[blockName]; exist { //array Element
						if _, exist := arrayElements[blockName][keys[1]]; !exist {
							arrayElements[blockName][keys[1]] = map[string]interface{}{}
						}
						if len(keys) == 3 {
							arrayElements[blockName][keys[1]][keys[2]] = value.(string)
						} else if len(keys) == 4 {
							if _, exist := arrayElements[blockName][keys[1]][keys[2]]; !exist {
								arrayElements[blockName][keys[1]][keys[2]] = []string{}
							}
							arrayElements[blockName][keys[1]][keys[2]] = append(arrayElements[blockName][keys[1]][keys[2]].([]string), value.(string))
						}
					}
					if _, exist := hashElements[blockName]; exist { // hash Element
						item[blockName].(map[string]string)[keys[1]] = value.(string)
					}
				}
			}
			for key, elem := range arrayElements {
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
						}
					}
					item[key] = append(item[key].([]map[string]interface{}), element)
				}
			}
			resources = append(resources, TerraformResource{
				ResourceType: strings.Split(key, ".")[0],
				ResourceName: strings.Split(key, ".")[1],
				Item:         item,
				ID:           resource.Primary.ID,
				Provider:     provider,
			})
		}
	}
	return resources, nil
}
