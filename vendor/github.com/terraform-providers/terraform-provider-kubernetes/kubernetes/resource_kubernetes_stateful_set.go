package kubernetes

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	api "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	pkgApi "k8s.io/apimachinery/pkg/types"
	kubernetes "k8s.io/client-go/kubernetes"
)

func resourceKubernetesStatefulSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceKubernetesStatefulSetCreate,
		Read:   resourceKubernetesStatefulSetRead,
		Update: resourceKubernetesStatefulSetUpdate,
		Delete: resourceKubernetesStatefulSetDelete,
		Exists: resourceKubernetesStatefulSetExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"metadata": namespacedMetadataSchema("stateful set", true),
			"spec": {
				Type:        schema.TypeList,
				Description: "Spec defines the desired identities of pods in this set.",
				Required:    true,
				MaxItems:    1,
				MinItems:    1,
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: statefulSetSpecFields(false),
				},
			},
		},
	}
}

func resourceKubernetesStatefulSetCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)
	metadata := expandMetadata(d.Get("metadata").([]interface{}))
	spec, err := expandStatefulSetSpec(d.Get("spec").([]interface{}))
	if err != nil {
		return err
	}
	statefulSet := api.StatefulSet{
		ObjectMeta: metadata,
		Spec:       spec,
	}
	log.Printf("[INFO] Creating new StatefulSet: %#v", statefulSet)

	out, err := conn.AppsV1().StatefulSets(metadata.Namespace).Create(&statefulSet)

	if err != nil {
		return err
	}
	log.Printf("[INFO] Submitted new StatefulSet: %#v", out)

	id := buildId(out.ObjectMeta)
	d.SetId(id)

	log.Printf("[INFO] StatefulSet %s created", id)

	return resourceKubernetesStatefulSetRead(d, meta)
}

func resourceKubernetesStatefulSetExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	conn := meta.(*kubernetes.Clientset)

	namespace, name, err := idParts(d.Id())
	if err != nil {
		return false, err
	}

	log.Printf("[INFO] Checking StatefulSet %s", name)
	_, err = conn.AppsV1().StatefulSets(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return false, nil
		}
		log.Printf("[DEBUG] Received error: %#v", err)
	}
	return true, err
}

func resourceKubernetesStatefulSetRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	id := d.Id()
	namespace, name, err := idParts(id)
	if err != nil {
		return fmt.Errorf("Error parsing resource ID: %#v", err)
	}
	log.Printf("[INFO] Reading stateful set %s", id)
	statefulSet, err := conn.AppsV1().StatefulSets(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		switch {
		case errors.IsNotFound(err):
			log.Printf("[DEBUG] Stateful Set %q was not found in Namespace %q - removing from state!", namespace, name)
			d.SetId("")
			return nil
		default:
			log.Printf("[DEBUG] Error reading stateful set: %#v", err)
			return err
		}
	}
	log.Printf("[INFO] Received stateful set: %#v", statefulSet)
	if d.Set("metadata", flattenMetadata(statefulSet.ObjectMeta)) != nil {
		return fmt.Errorf("Error setting `metadata`: %+v", err)
	}
	sss, err := flattenStatefulSetSpec(statefulSet.Spec)
	if err != nil {
		return fmt.Errorf("Error flattening `spec`: %+v", err)
	}
	err = d.Set("spec", sss)
	if err != nil {
		return fmt.Errorf("Error setting `spec`: %+v", err)
	}
	return nil
}

func resourceKubernetesStatefulSetUpdate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)
	namespace, name, err := idParts(d.Id())
	if err != nil {
		return fmt.Errorf("Error parsing resource ID: %#v", err)
	}
	ops := patchMetadata("metadata.0.", "/metadata/", d)

	if d.HasChange("spec") {
		log.Println("[TRACE] StatefulSet.Spec has changes")
		specPatch, err := patchStatefulSetSpec(d)
		if err != nil {
			return err
		}
		ops = append(ops, specPatch...)
	}

	data, err := ops.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Failed to marshal update operations for StatefulSet: %s", err)
	}
	log.Printf("[INFO] Updating StatefulSet %q: %v", name, string(data))
	out, err := conn.AppsV1().StatefulSets(namespace).Patch(name, pkgApi.JSONPatchType, data)
	if err != nil {
		return fmt.Errorf("Failed to update StatefulSet: %s", err)
	}
	log.Printf("[INFO] Submitted updated StatefulSet: %#v", out)

	return resourceKubernetesStatefulSetRead(d, meta)
}

func resourceKubernetesStatefulSetDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	namespace, name, err := idParts(d.Id())
	if err != nil {
		return fmt.Errorf("Error parsing resource ID: %#v", err)
	}
	log.Printf("[INFO] Deleting StatefulSet: %#v", name)
	err = conn.AppsV1().StatefulSets(namespace).Delete(name, nil)
	if err != nil {
		return err
	}
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		out, err := conn.AppsV1().StatefulSets(namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			switch {
			case errors.IsNotFound(err):
				return nil
			default:
				return resource.NonRetryableError(err)
			}
		}

		log.Printf("[DEBUG] Current state of StatefulSet: %#v", out.Status.Conditions)
		e := fmt.Errorf("StatefulSet %s still exists %#v", name, out.Status.Conditions)
		return resource.RetryableError(e)
	})
	if err != nil {
		return err
	}

	log.Printf("[INFO] StatefulSet %s deleted", name)

	return nil
}
