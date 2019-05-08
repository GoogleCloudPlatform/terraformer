package kubernetes

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func podTemplateFields(isUpdatable bool) map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"metadata": metadataSchema("stateful set", true),
		"spec": {
			Type:        schema.TypeList,
			Description: "Spec of the pods owned by the stateful set",
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: podSpecFields(false),
			},
		},
	}
	return s
}
