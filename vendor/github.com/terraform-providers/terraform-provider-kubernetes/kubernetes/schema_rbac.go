package kubernetes

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func rbacRoleRefSchema(kind string) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"api_group": {
			Type:        schema.TypeString,
			Description: "The API group of the user. Always `rbac.authorization.k8s.io`",
			Required:    true,
			Default:     "rbac.authorization.k8s.io",
		},
		"kind": {
			Type:        schema.TypeString,
			Description: "The kind of resource.",
			Default:     kind,
			Required:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Description: "The name of the User to bind to.",
			Required:    true,
		},
	}
}

func rbacSubjectSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"api_group": {
			Type:        schema.TypeString,
			Description: "The API group of the user. Always `rbac.authorization.k8s.io`",
			Optional:    true,
			Default:     "rbac.authorization.k8s.io",
		},
		"kind": {
			Type:        schema.TypeString,
			Description: "The kind of resource.",
			Required:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Description: "The name of the resource to bind to.",
			Required:    true,
		},
		"namespace": {
			Type:        schema.TypeString,
			Description: "The Namespace of the ServiceAccount",
			Optional:    true,
			Default:     "default",
		},
	}
}
