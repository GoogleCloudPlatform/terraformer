package terraform_utils

type TerraformResource struct {
	ResourceType string
	ResourceName string
	Item         interface{}
	ID           string
	Provider     string
	Attributes   map[string]string
}

func NewTerraformResource(ID, resourceName, resourceType, provider string, item interface{}, attributes map[string]string) TerraformResource {
	return TerraformResource{
		ResourceType: resourceType,
		ResourceName: TfSanitize(resourceName),
		Item:         item,
		ID:           ID,
		Provider:     provider,
		Attributes:   attributes,
	}
}
