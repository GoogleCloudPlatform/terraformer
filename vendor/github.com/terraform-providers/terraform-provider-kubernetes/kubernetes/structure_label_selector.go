package kubernetes

import (
	"github.com/hashicorp/terraform/helper/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Flatteners

func flattenLabelSelector(in *metav1.LabelSelector) []interface{} {
	att := make(map[string]interface{})
	if len(in.MatchLabels) > 0 {
		att["match_labels"] = in.MatchLabels
	}
	if len(in.MatchExpressions) > 0 {
		att["match_expressions"] = flattenLabelSelectorRequirement(in.MatchExpressions)
	}
	if len(att) > 0 {
		return []interface{}{att}
	}
	return []interface{}{}
}

func flattenLabelSelectorRequirement(in []metav1.LabelSelectorRequirement) []interface{} {
	att := make([]interface{}, len(in), len(in))
	for i, n := range in {
		m := make(map[string]interface{})
		m["key"] = n.Key
		m["operator"] = n.Operator
		m["values"] = newStringSet(schema.HashString, n.Values)
		att[i] = m
	}
	return att
}

// Expanders

func expandLabelSelector(l []interface{}) *metav1.LabelSelector {
	if len(l) == 0 || l[0] == nil {
		return &metav1.LabelSelector{}
	}
	in := l[0].(map[string]interface{})
	obj := &metav1.LabelSelector{}
	if v, ok := in["match_labels"].(map[string]interface{}); ok && len(v) > 0 {
		obj.MatchLabels = expandStringMap(v)
	}
	if v, ok := in["match_expressions"].([]interface{}); ok && len(v) > 0 {
		obj.MatchExpressions = expandLabelSelectorRequirement(v)
	}
	return obj
}

func expandLabelSelectorRequirement(l []interface{}) []metav1.LabelSelectorRequirement {
	if len(l) == 0 || l[0] == nil {
		return []metav1.LabelSelectorRequirement{}
	}
	obj := make([]metav1.LabelSelectorRequirement, len(l), len(l))
	for i, n := range l {
		in := n.(map[string]interface{})
		obj[i] = metav1.LabelSelectorRequirement{
			Key:      in["key"].(string),
			Operator: metav1.LabelSelectorOperator(in["operator"].(string)),
			Values:   sliceOfString(in["values"].(*schema.Set).List()),
		}
	}
	return obj
}
