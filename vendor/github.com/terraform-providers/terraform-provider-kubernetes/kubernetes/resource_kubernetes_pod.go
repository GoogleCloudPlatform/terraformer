package kubernetes

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	api "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	pkgApi "k8s.io/apimachinery/pkg/types"
	kubernetes "k8s.io/client-go/kubernetes"
)

func resourceKubernetesPod() *schema.Resource {
	return &schema.Resource{
		Create: resourceKubernetesPodCreate,
		Read:   resourceKubernetesPodRead,
		Update: resourceKubernetesPodUpdate,
		Delete: resourceKubernetesPodDelete,
		Exists: resourceKubernetesPodExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"metadata": namespacedMetadataSchema("pod", true),
			"spec": {
				Type:        schema.TypeList,
				Description: "Specification of the desired behavior of the pod.",
				Required:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: podSpecFields(false),
				},
			},
		},
	}
}
func resourceKubernetesPodCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	metadata := expandMetadata(d.Get("metadata").([]interface{}))
	spec, err := expandPodSpec(d.Get("spec").([]interface{}))
	if err != nil {
		return err
	}

	spec.AutomountServiceAccountToken = ptrToBool(false)

	pod := api.Pod{
		ObjectMeta: metadata,
		Spec:       spec,
	}

	log.Printf("[INFO] Creating new pod: %#v", pod)
	out, err := conn.CoreV1().Pods(metadata.Namespace).Create(&pod)

	if err != nil {
		return err
	}
	log.Printf("[INFO] Submitted new pod: %#v", out)

	d.SetId(buildId(out.ObjectMeta))

	stateConf := &resource.StateChangeConf{
		Target:  []string{"Running"},
		Pending: []string{"Pending"},
		Timeout: d.Timeout(schema.TimeoutCreate),
		Refresh: func() (interface{}, string, error) {
			out, err := conn.CoreV1().Pods(metadata.Namespace).Get(metadata.Name, metav1.GetOptions{})
			if err != nil {
				log.Printf("[ERROR] Received error: %#v", err)
				return out, "Error", err
			}

			statusPhase := fmt.Sprintf("%v", out.Status.Phase)
			log.Printf("[DEBUG] Pods %s status received: %#v", out.Name, statusPhase)
			return out, statusPhase, nil
		},
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		lastWarnings, wErr := getLastWarningsForObject(conn, out.ObjectMeta, "Pod", 3)
		if wErr != nil {
			return wErr
		}
		return fmt.Errorf("%s%s", err, stringifyEvents(lastWarnings))
	}
	log.Printf("[INFO] Pod %s created", out.Name)

	return resourceKubernetesPodRead(d, meta)
}

func resourceKubernetesPodUpdate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}

	ops := patchMetadata("metadata.0.", "/metadata/", d)
	if d.HasChange("spec") {
		specOps, err := patchPodSpec("/spec", "spec.0.", d)
		if err != nil {
			return err
		}
		ops = append(ops, specOps...)
	}
	data, err := ops.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Failed to marshal update operations: %s", err)
	}

	log.Printf("[INFO] Updating pod %s: %s", d.Id(), ops)

	out, err := conn.CoreV1().Pods(namespace).Patch(name, pkgApi.JSONPatchType, data)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Submitted updated pod: %#v", out)

	d.SetId(buildId(out.ObjectMeta))
	return resourceKubernetesPodRead(d, meta)
}

func resourceKubernetesPodRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}

	log.Printf("[INFO] Reading pod %s", name)
	pod, err := conn.CoreV1().Pods(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		log.Printf("[DEBUG] Received error: %#v", err)
		return err
	}
	log.Printf("[INFO] Received pod: %#v", pod)

	err = d.Set("metadata", flattenMetadata(pod.ObjectMeta))
	if err != nil {
		return err
	}

	podSpec, err := flattenPodSpec(pod.Spec)
	if err != nil {
		return err
	}

	err = d.Set("spec", podSpec)
	if err != nil {
		return err
	}
	return nil

}

func resourceKubernetesPodDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}

	log.Printf("[INFO] Deleting pod: %#v", name)
	err = conn.CoreV1().Pods(namespace).Delete(name, nil)
	if err != nil {
		return err
	}

	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		out, err := conn.CoreV1().Pods(namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			if statusErr, ok := err.(*errors.StatusError); ok && statusErr.ErrStatus.Code == 404 {
				return nil
			}
			return resource.NonRetryableError(err)
		}

		log.Printf("[DEBUG] Current state of pod: %#v", out.Status.Phase)
		e := fmt.Errorf("Pod %s still exists (%s)", name, out.Status.Phase)
		return resource.RetryableError(e)
	})
	if err != nil {
		return err
	}

	log.Printf("[INFO] Pod %s deleted", name)

	d.SetId("")
	return nil
}

func resourceKubernetesPodExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	conn := meta.(*kubernetes.Clientset)

	namespace, name, err := idParts(d.Id())
	if err != nil {
		return false, err
	}

	log.Printf("[INFO] Checking pod %s", name)
	_, err = conn.CoreV1().Pods(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		if statusErr, ok := err.(*errors.StatusError); ok && statusErr.ErrStatus.Code == 404 {
			return false, nil
		}
		log.Printf("[DEBUG] Received error: %#v", err)
	}
	return true, err
}
