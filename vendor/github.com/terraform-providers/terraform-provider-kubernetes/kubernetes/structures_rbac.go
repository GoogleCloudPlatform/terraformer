package kubernetes

import (
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	api "k8s.io/api/rbac/v1"
)

func expandRBACRoleRef(in interface{}) api.RoleRef {
	if in == nil {
		return api.RoleRef{}
	}
	ref := api.RoleRef{}
	m := in.(map[string]interface{})
	if v, ok := m["api_group"]; ok {
		ref.APIGroup = v.(string)
	}
	if v, ok := m["kind"].(string); ok {
		ref.Kind = string(v)
	}
	if v, ok := m["name"]; ok {
		ref.Name = v.(string)
	}

	return ref
}

func expandRBACSubjects(in []interface{}) []api.Subject {
	if len(in) == 0 || in[0] == nil {
		return []api.Subject{}
	}
	subjects := make([]api.Subject, 0, len(in))
	for i := range in {
		subject := api.Subject{}
		m := in[i].(map[string]interface{})
		if v, ok := m["api_group"]; ok {
			subject.APIGroup = v.(string)
		}
		if v, ok := m["kind"].(string); ok {
			subject.Kind = string(v)
		}
		if v, ok := m["name"]; ok {
			subject.Name = v.(string)
		}
		if v, ok := m["namespace"]; ok {
			subject.Namespace = v.(string)
		}
		subjects = append(subjects, subject)
	}
	return subjects
}

func flattenRBACRoleRef(in api.RoleRef) interface{} {
	att := make(map[string]interface{})

	if in.APIGroup != "" {
		att["api_group"] = in.APIGroup
	}
	att["kind"] = in.Kind
	att["name"] = in.Name
	return att
}

func flattenRBACSubjects(in []api.Subject) []interface{} {
	att := make([]interface{}, 0, len(in))
	for _, n := range in {
		m := make(map[string]interface{})
		if n.APIGroup != "" {
			m["api_group"] = n.APIGroup
		}
		m["kind"] = n.Kind
		m["name"] = n.Name
		if n.Namespace != "" {
			m["namespace"] = n.Namespace
		}
		att = append(att, m)
	}
	return att
}

// Patch Ops
func patchRbacSubject(d *schema.ResourceData) PatchOperations {
	o, n := d.GetChange("subject")
	oldsubjects := expandRBACSubjects(o.([]interface{}))
	newsubjects := expandRBACSubjects(n.([]interface{}))
	ops := make([]PatchOperation, 0, len(newsubjects)+len(oldsubjects))

	common := len(newsubjects)
	if common > len(oldsubjects) {
		common = len(oldsubjects)
	}
	for i, v := range newsubjects[:common] {
		ops = append(ops, &ReplaceOperation{
			Path:  "/subjects/" + strconv.Itoa(i),
			Value: v,
		})
	}
	if len(oldsubjects) > len(newsubjects) {
		for i := len(newsubjects); i < len(oldsubjects); i++ {
			ops = append(ops, &RemoveOperation{
				Path: "/subjects/" + strconv.Itoa(len(oldsubjects)-i),
			})
		}
	}
	if len(newsubjects) > len(oldsubjects) {
		for i, v := range newsubjects[common:] {
			ops = append(ops, &AddOperation{
				Path:  "/subjects/" + strconv.Itoa(common+i),
				Value: v,
			})
		}
	}
	return ops
}
