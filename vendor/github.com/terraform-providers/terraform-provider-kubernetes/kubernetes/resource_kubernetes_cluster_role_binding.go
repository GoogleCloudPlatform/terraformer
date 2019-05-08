package kubernetes

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	api "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	pkgApi "k8s.io/apimachinery/pkg/types"
	kubernetes "k8s.io/client-go/kubernetes"
)

func resourceKubernetesClusterRoleBinding() *schema.Resource {
	return &schema.Resource{
		Create: resourceKubernetesClusterRoleBindingCreate,
		Read:   resourceKubernetesClusterRoleBindingRead,
		Exists: resourceKubernetesClusterRoleBindingExists,
		Update: resourceKubernetesClusterRoleBindingUpdate,
		Delete: resourceKubernetesClusterRoleBindingDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"metadata": metadataSchema("clusterRoleBinding", false),
			"role_ref": {
				Type:        schema.TypeMap,
				Description: "RoleRef references the Cluster Role for this binding",
				Required:    true,
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: rbacRoleRefSchema("ClusterRole"),
				},
			},
			"subject": {
				Type:        schema.TypeList,
				Description: "Subjects defines the entities to bind a ClusterRole to.",
				Required:    true,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: rbacSubjectSchema(),
				},
			},
		},
	}
}

func resourceKubernetesClusterRoleBindingCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	metadata := expandMetadata(d.Get("metadata").([]interface{}))
	binding := &api.ClusterRoleBinding{
		ObjectMeta: metadata,
		RoleRef:    expandRBACRoleRef(d.Get("role_ref").(interface{})),
		Subjects:   expandRBACSubjects(d.Get("subject").([]interface{})),
	}
	log.Printf("[INFO] Creating new ClusterRoleBinding: %#v", binding)
	binding, err := conn.Rbac().ClusterRoleBindings().Create(binding)

	if err != nil {
		return err
	}
	log.Printf("[INFO] Submitted new ClusterRoleBinding: %#v", binding)
	d.SetId(metadata.Name)

	return resourceKubernetesClusterRoleBindingRead(d, meta)
}

func resourceKubernetesClusterRoleBindingRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	name := d.Id()
	log.Printf("[INFO] Reading ClusterRoleBinding %s", name)
	binding, err := conn.Rbac().ClusterRoleBindings().Get(name, meta_v1.GetOptions{})
	if err != nil {
		log.Printf("[DEBUG] Received error: %#v", err)
		return err
	}

	log.Printf("[INFO] Received ClusterRoleBinding: %#v", binding)
	err = d.Set("metadata", flattenMetadata(binding.ObjectMeta))
	if err != nil {
		return err
	}

	flattenedRef := flattenRBACRoleRef(binding.RoleRef)
	log.Printf("[DEBUG] Flattened ClusterRoleBinding roleRef: %#v", flattenedRef)
	err = d.Set("role_ref", flattenedRef)
	if err != nil {
		return err
	}

	flattenedSubjects := flattenRBACSubjects(binding.Subjects)
	log.Printf("[DEBUG] Flattened ClusterRoleBinding subjects: %#v", flattenedSubjects)
	err = d.Set("subject", flattenedSubjects)
	if err != nil {
		return err
	}

	return nil
}

func resourceKubernetesClusterRoleBindingUpdate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	name := d.Id()

	ops := patchMetadata("metadata.0.", "/metadata/", d)
	if d.HasChange("subject") {
		diffOps := patchRbacSubject(d)
		ops = append(ops, diffOps...)
	}
	data, err := ops.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Failed to marshal update operations: %s", err)
	}
	log.Printf("[INFO] Updating ClusterRoleBinding %q: %v", name, string(data))
	out, err := conn.Rbac().ClusterRoleBindings().Patch(name, pkgApi.JSONPatchType, data)
	if err != nil {
		return fmt.Errorf("Failed to update ClusterRoleBinding: %s", err)
	}
	log.Printf("[INFO] Submitted updated ClusterRoleBinding: %#v", out)
	d.SetId(out.ObjectMeta.Name)

	return resourceKubernetesClusterRoleBindingRead(d, meta)
}

func resourceKubernetesClusterRoleBindingDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	name := d.Id()
	log.Printf("[INFO] Deleting ClusterRoleBinding: %#v", name)
	err := conn.Rbac().ClusterRoleBindings().Delete(name, &meta_v1.DeleteOptions{})
	if err != nil {
		return err
	}
	log.Printf("[INFO] ClusterRoleBinding %s deleted", name)

	d.SetId("")
	return nil
}

func resourceKubernetesClusterRoleBindingExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	conn := meta.(*kubernetes.Clientset)

	name := d.Id()
	log.Printf("[INFO] Checking ClusterRoleBinding %s", name)
	_, err := conn.Rbac().ClusterRoleBindings().Get(name, meta_v1.GetOptions{})
	if err != nil {
		if statusErr, ok := err.(*errors.StatusError); ok && statusErr.ErrStatus.Code == 404 {
			return false, nil
		}
		log.Printf("[DEBUG] Received error: %#v", err)
	}
	return true, err
}
