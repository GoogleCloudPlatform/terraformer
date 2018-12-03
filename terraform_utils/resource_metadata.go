package terraform_utils

type ResourceMetaData struct {
	Provider         string
	IgnoreKeys       map[string]bool
	AllowEmptyValue  map[string]bool
	AdditionalFields map[string]string
}

func NewResourcesMetaData(resources []TerraformResource, ignoreKeys, allowEmptyValue map[string]bool, AdditionalFields map[string]string) map[string]ResourceMetaData {
	data := map[string]ResourceMetaData{}
	for _, resource := range resources {
		data[resource.ID] = ResourceMetaData{
			Provider:         resource.Provider,
			IgnoreKeys:       ignoreKeys,
			AllowEmptyValue:  allowEmptyValue,
			AdditionalFields: AdditionalFields,
		}
	}
	return data
}
