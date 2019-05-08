package kubernetes

import (
	"errors"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Expanders

func expandStatefulSetSpec(s []interface{}) (v1.StatefulSetSpec, error) {
	obj := v1.StatefulSetSpec{}
	if len(s) == 0 || s[0] == nil {
		return obj, nil
	}
	in := s[0].(map[string]interface{})

	if v, ok := in["pod_management_policy"].(string); ok {
		obj.PodManagementPolicy = v1.PodManagementPolicyType(v)
	}

	if v, ok := in["replicas"].(int); ok && v > 0 {
		obj.Replicas = ptrToInt32(int32(v))
	}

	if v, ok := in["revision_history_limit"].(int); ok {
		obj.RevisionHistoryLimit = ptrToInt32(int32(v))
	}

	if v, ok := in["selector"].([]interface{}); ok && len(v) > 0 {
		obj.Selector = expandLabelSelector(v)
	}

	if v, ok := in["service_name"].(string); ok {
		obj.ServiceName = v
	}

	if v, ok := in["update_strategy"].([]interface{}); ok {
		us, err := expandStatefulSetSpecUpdateStrategy(v)
		if err != nil {
			return obj, err
		}
		obj.UpdateStrategy = us
	}

	template, err := expandPodTemplate(in["template"].([]interface{}))
	if err != nil {
		return obj, err
	}
	obj.Template = template

	if v, ok := in["volume_claim_template"].([]interface{}); ok {
		obj.VolumeClaimTemplates = []corev1.PersistentVolumeClaim{}
		if len(v) == 0 || v[0] == nil {
			return obj, nil
		}
		for _, pvc := range v {
			p, err := expandPersistenVolumeClaim(pvc.(map[string]interface{}))
			if err != nil {
				return obj, err
			}
			obj.VolumeClaimTemplates = append(obj.VolumeClaimTemplates, p)
		}
	}
	return obj, nil
}
func expandStatefulSetSpecUpdateStrategy(s []interface{}) (v1.StatefulSetUpdateStrategy, error) {
	ust := v1.StatefulSetUpdateStrategy{}
	if len(s) == 0 {
		return ust, nil
	}
	us, ok := s[0].(map[string]interface{})
	if !ok {
		return ust, errors.New("failed to expand 'spec.update_strategy'")
	}
	t, ok := us["type"].(string)
	if !ok {
		return ust, errors.New("failed to expand 'spec.update_strategy.type'")
	}
	ust.Type = v1.StatefulSetUpdateStrategyType(t)
	ru, ok := us["rolling_update"].([]interface{})
	if !ok {
		return ust, errors.New("failed to unroll 'spec.update_strategy.rolling_update'")
	}
	if len(ru) > 0 {
		u := v1.RollingUpdateStatefulSetStrategy{}
		r, ok := ru[0].(map[string]interface{})
		if !ok {
			return ust, errors.New("failed to expand 'spec.update_strategy.rolling_update'")
		}
		p, ok := r["partition"].(int)
		if !ok {
			return ust, errors.New("failed to expand 'spec.update_strategy.rolling_update.partition'")
		}
		u.Partition = ptrToInt32(int32(p))
		ust.RollingUpdate = &u
	}
	log.Printf("[DEBUG] Expanded StatefulSet.Spec.UpdateStrategy: %#v", ust)
	return ust, nil
}

func expandStatefulSetSelectors(s []interface{}) (metav1.LabelSelector, error) {
	obj := metav1.LabelSelector{}
	if len(s) == 0 || s[0] == nil {
		return obj, nil
	}
	in := s[0].(map[string]interface{})
	log.Printf("[DEBUG] StatefulSet Selector: %#v", in)
	if v, ok := in["match_labels"].(map[string]interface{}); ok {
		log.Printf("[DEBUG] StatefulSet Selector MatchLabels: %#v", v)
		ml := make(map[string]string)
		for k, l := range v {
			ml[k] = l.(string)
			log.Printf("[DEBUG] StatefulSet Selector MatchLabel: %#v -> %#v", k, v)
		}
		obj.MatchLabels = ml
	}
	if v, ok := in["match_expressions"].([]interface{}); ok {
		log.Printf("[DEBUG] StatefulSet Selector MatchExpressions: %#v", v)
		me, err := expandMatchExpressions(v)
		if err != nil {
			return obj, err
		}
		obj.MatchExpressions = me
	}
	return obj, nil
}

func expandMatchExpressions(in []interface{}) ([]metav1.LabelSelectorRequirement, error) {
	if len(in) == 0 {
		return []metav1.LabelSelectorRequirement{}, nil
	}
	obj := make([]metav1.LabelSelectorRequirement, len(in))
	for i, c := range in {
		p := c.(map[string]interface{})
		if v, ok := p["key"].(string); ok {
			obj[i].Key = v
		}
		if v, ok := p["operator"].(metav1.LabelSelectorOperator); ok {
			obj[i].Operator = v
		}
		if v, ok := p["values"].(*schema.Set); ok {
			obj[i].Values = schemaSetToStringArray(v)
		}
	}
	return obj, nil
}

// Flattners

func flattenStatefulSetSpec(spec v1.StatefulSetSpec) ([]interface{}, error) {
	att := make(map[string]interface{})

	if spec.PodManagementPolicy != "" {
		att["pod_management_policy"] = spec.PodManagementPolicy
	}
	if spec.Replicas != nil {
		att["replicas"] = *spec.Replicas
	}
	if spec.RevisionHistoryLimit != nil {
		att["revision_history_limit"] = *spec.RevisionHistoryLimit
	}
	if spec.Selector != nil {
		att["selector"] = flattenLabelSelector(spec.Selector)
	}
	if spec.ServiceName != "" {
		att["service_name"] = spec.ServiceName
	}
	template, err := flattenPodTemplateSpec(spec.Template)
	if err != nil {
		return []interface{}{att}, err
	}
	att["template"] = template
	att["volume_claim_template"] = flattenPersistentVolumeClaim(spec.VolumeClaimTemplates)
	att["update_strategy"] = flattenStatefulSetSpecUpdateStrategy(spec.UpdateStrategy)

	return []interface{}{att}, nil
}

func flattenPodTemplateSpec(t corev1.PodTemplateSpec) ([]interface{}, error) {
	template := make(map[string]interface{})

	template["metadata"] = flattenMetadata(t.ObjectMeta)
	spec, err := flattenPodSpec(t.Spec)
	if err != nil {
		return []interface{}{template}, err
	}
	template["spec"] = spec

	return []interface{}{template}, nil
}

func flattenPersistentVolumeClaim(in []corev1.PersistentVolumeClaim) []interface{} {
	pvcs := make([]interface{}, 0, len(in))

	for _, pvc := range in {
		p := make(map[string]interface{})
		p["metadata"] = flattenMetadata(pvc.ObjectMeta)
		p["spec"] = flattenPersistentVolumeClaimSpec(pvc.Spec)
		pvcs = append(pvcs, p)
	}
	return pvcs
}

func flattenStatefulSetSpecUpdateStrategy(s v1.StatefulSetUpdateStrategy) []interface{} {
	att := make(map[string]interface{})

	att["type"] = s.Type
	if s.RollingUpdate != nil {
		ru := make(map[string]interface{})
		if s.RollingUpdate.Partition != nil {
			ru["partition"] = *s.RollingUpdate.Partition
		}
		att["rolling_update"] = []interface{}{ru}
	}
	return []interface{}{att}
}

// Patchers

func patchStatefulSetSpec(d *schema.ResourceData) (PatchOperations, error) {
	ops := PatchOperations{}

	if d.HasChange("spec.0.replicas") {
		log.Printf("[TRACE] StatefulSet.Spec.Replicas has changes")
		if v, ok := d.Get("spec.0.replicas").(int); ok {
			ops = append(ops, &ReplaceOperation{
				Path:  "/spec/replicas",
				Value: v,
			})
		}
	}

	if d.HasChange("spec.0.template") {
		log.Printf("[TRACE] StatefulSet.Spec.Template has changes")
		t, err := patchPodTemplateSpec("spec.0.template.0.", "/spec/template/", d)
		if err != nil {
			return ops, err
		}
		ops = append(ops, t...)
	}

	if d.HasChange("spec.0.update_strategy") {
		log.Printf("[TRACE] StatefulSet.Spec.UpdateStrategy has changes")
		u, err := patchUpdateStrategy("spec.0.update_strategy.0.", "/spec/updateStrategy/", d)
		if err != nil {
			return ops, err
		}
		ops = append(ops, u...)
	}
	return ops, nil
}

func patchPodTemplateSpec(keyPrefix, pathPrefix string, d *schema.ResourceData) (PatchOperations, error) {
	ops := PatchOperations{}

	if d.HasChange(keyPrefix + "metadata") {
		log.Printf("[TRACE] StatefulSet.Spec.Template.Metadata has changes")
		m := patchMetadata(keyPrefix+"metadata.0.", pathPrefix+"metadata/", d)
		ops = append(ops, m...)
	}

	if d.HasChange(keyPrefix + "spec") {
		log.Printf("[TRACE] StatefulSet.Spec.Template.Spec has changes")
		p, err := patchPodSpec(pathPrefix+"spec.0.", keyPrefix+"spec/", d)
		if err != nil {
			return ops, err
		}
		ops = append(ops, p...)
	}

	return ops, nil
}

func patchUpdateStrategy(keyPrefix, pathPrefix string, d *schema.ResourceData) (PatchOperations, error) {
	ops := PatchOperations{}

	if d.HasChange(keyPrefix + "type") {
		log.Printf("[TRACE] StatefulSet.Spec.UpdateStrategy.Type has changes")
		oldV, newV := d.GetChange(keyPrefix + "type")
		o := oldV.(string)
		n := newV.(string)
		if len(o) != 0 && len(n) == 0 {
			return ops, fmt.Errorf("Spec.UpdateStrategy.Type cannot be empty")
		}
		if len(o) == 0 && len(n) != 0 {
			ops = append(ops, &AddOperation{
				Path:  pathPrefix + "type",
				Value: n,
			})
		} else {
			ops = append(ops, &ReplaceOperation{
				Path:  pathPrefix + "type",
				Value: n,
			})
		}
	}

	if d.HasChange(keyPrefix + "rolling_update") {
		o, n := d.GetChange(keyPrefix + "rolling_update")
		log.Printf("[TRACE] StatefulSet.Spec.UpdateStrategy.RollingUpdate has changes: %#v | %#v", o, n)

		if len(o.([]interface{})) > 0 && len(n.([]interface{})) == 0 {
			ops = append(ops, &RemoveOperation{
				Path: pathPrefix + "rollingUpdate",
			})
		}

		if len(o.([]interface{})) == 0 && len(n.([]interface{})) > 0 {
			ops = append(ops, &AddOperation{
				Path:  pathPrefix + "rollingUpdate",
				Value: struct{}{},
			})
			ops = append(ops, &AddOperation{
				Path:  pathPrefix + "rollingUpdate/partition",
				Value: d.Get(keyPrefix + "rolling_update.0.partition").(int),
			})
		}

		if len(o.([]interface{})) > 0 && len(n.([]interface{})) > 0 {
			r, err := patchUpdateStrategyRollingUpdate(keyPrefix+"rolling_update.0.", pathPrefix+"rollingUpdate/", d)
			if err != nil {
				return ops, err
			}
			ops = append(ops, r...)
		}
	}

	return ops, nil
}

func patchUpdateStrategyRollingUpdate(keyPrefix, pathPrefix string, d *schema.ResourceData) (PatchOperations, error) {
	ops := PatchOperations{}
	if d.HasChange(keyPrefix + "partition") {
		log.Printf("[TRACE] StatefulSet.Spec.UpdateStrategy.RollingUpdate.Partition has changes")
		if p, ok := d.Get(keyPrefix + "partition").(int); ok {
			ops = append(ops, &ReplaceOperation{
				Path:  pathPrefix + "partition",
				Value: p,
			})
		}
	}
	return ops, nil
}
