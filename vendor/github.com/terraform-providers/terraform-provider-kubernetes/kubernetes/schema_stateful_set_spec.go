package kubernetes

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func statefulSetSpecFields(isUpdatable bool) map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"pod_management_policy": {
			Type:        schema.TypeString,
			Description: "Controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down.",
			Optional:    true,
			ForceNew:    true,
			Computed:    true,
			ValidateFunc: validation.StringInSlice([]string{
				"OrderedReady",
				"Parallel",
			}, false),
		},
		"replicas": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      1,
			Description:  "The desired number of replicas of the given Template, in the sense that they are instantiations of the same Template. Value must be a positive integer.",
			ValidateFunc: validatePositiveInteger,
		},
		"revision_history_limit": {
			Type:         schema.TypeInt,
			Optional:     true,
			ForceNew:     true,
			Computed:     true,
			ValidateFunc: validatePositiveInteger,
			Description:  "The maximum number of revisions that will be maintained in the StatefulSet's revision history. The default value is 10.",
		},
		"selector": {
			Type:        schema.TypeList,
			Description: "A label query over pods that should match the replica count. It must match the pod template's labels.",
			Required:    true,
			ForceNew:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: labelSelectorFields(),
			},
		},
		"service_name": {
			Type:        schema.TypeString,
			Description: "The name of the service that governs this StatefulSet. This service must exist before the StatefulSet, and is responsible for the network identity of the set.",
			Required:    true,
			ForceNew:    true,
		},
		"template": {
			Type:        schema.TypeList,
			Description: "The object that describes the pod that will be created if insufficient replicas are detected. Each pod stamped out by the StatefulSet will fulfill this Template.",
			Required:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: podTemplateFields(isUpdatable),
			},
		},
		"update_strategy": {
			Type:        schema.TypeList,
			Description: "The strategy that the StatefulSet controller will use to perform updates.",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"type": {
						Type:        schema.TypeString,
						Description: "Indicates the type of the StatefulSet update strategy. Default is RollingUpdate",
						Optional:    true,
						Default:     "RollingUpdate",
						ValidateFunc: validation.StringInSlice([]string{
							"RollingUpdate",
							"OnDelete",
						}, false),
					},
					"rolling_update": {
						Type:        schema.TypeList,
						Description: "RollingUpdate strategy type for StatefulSet",
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"partition": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "Indicates the ordinal at which the StatefulSet should be partitioned. Default value is 0.",
									Default:     0,
								},
							},
						},
					},
				},
			},
		},
		"volume_claim_template": {
			Type:        schema.TypeList,
			Optional:    true,
			ForceNew:    true,
			Description: "A list of claims that pods are allowed to reference. Every claim in this list must have at least one matching (by name) volumeMount in one container in the template.",
			Elem: &schema.Resource{
				Schema: persistentVolumeClaimFields(),
			},
		},
	}
	return s
}
