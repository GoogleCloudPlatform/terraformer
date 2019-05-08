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

func resourceKubernetesServiceAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceKubernetesServiceAccountCreate,
		Read:   resourceKubernetesServiceAccountRead,
		Exists: resourceKubernetesServiceAccountExists,
		Update: resourceKubernetesServiceAccountUpdate,
		Delete: resourceKubernetesServiceAccountDelete,

		// This resource is not importable because the API doesn't offer
		// any way to differentiate between default & user-defined secret
		// after the account was created.

		Schema: map[string]*schema.Schema{
			"metadata": namespacedMetadataSchema("service account", true),
			"image_pull_secret": {
				Type:        schema.TypeSet,
				Description: "A list of references to secrets in the same namespace to use for pulling any images in pods that reference this Service Account. More info: http://kubernetes.io/docs/user-guide/secrets#manually-specifying-an-imagepullsecret",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
							Optional:    true,
						},
					},
				},
			},
			"secret": {
				Type:        schema.TypeSet,
				Description: "A list of secrets allowed to be used by pods running using this Service Account. More info: http://kubernetes.io/docs/user-guide/secrets",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
							Optional:    true,
						},
					},
				},
			},
			"automount_service_account_token": {
				Type:        schema.TypeBool,
				Description: "True to enable automatic mounting of the service account token",
				Optional:    true,
				Default:     false,
			},
			"default_secret_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceKubernetesServiceAccountCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	metadata := expandMetadata(d.Get("metadata").([]interface{}))
	svcAcc := api.ServiceAccount{
		AutomountServiceAccountToken: ptrToBool(d.Get("automount_service_account_token").(bool)),
		ObjectMeta:                   metadata,
		ImagePullSecrets:             expandLocalObjectReferenceArray(d.Get("image_pull_secret").(*schema.Set).List()),
		Secrets:                      expandServiceAccountSecrets(d.Get("secret").(*schema.Set).List(), ""),
	}
	log.Printf("[INFO] Creating new service account: %#v", svcAcc)
	out, err := conn.CoreV1().ServiceAccounts(metadata.Namespace).Create(&svcAcc)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Submitted new service account: %#v", out)
	d.SetId(buildId(out.ObjectMeta))

	// Here we get the only chance to identify and store default secret name
	// so we can avoid showing it in diff as it's not managed by Terraform
	var resp *api.ServiceAccount
	err = resource.Retry(30*time.Second, func() *resource.RetryError {
		var err error
		resp, err = conn.CoreV1().ServiceAccounts(out.Namespace).Get(out.Name, metav1.GetOptions{})
		if err != nil {
			return resource.NonRetryableError(err)
		}
		if len(resp.Secrets) > len(svcAcc.Secrets) {
			return nil
		}
		return resource.RetryableError(fmt.Errorf("Waiting for default secret of %q to appear", d.Id()))
	})

	diff := diffObjectReferences(svcAcc.Secrets, resp.Secrets)
	if len(diff) > 1 {
		return fmt.Errorf("Expected 1 generated default secret, %d found: %s", len(diff), diff)
	}

	defaultSecret := diff[0]
	d.Set("default_secret_name", defaultSecret.Name)

	return resourceKubernetesServiceAccountRead(d, meta)
}

func diffObjectReferences(origOrs []api.ObjectReference, ors []api.ObjectReference) []api.ObjectReference {
	var diff []api.ObjectReference
	uniqueRefs := make(map[string]*api.ObjectReference, 0)
	for _, or := range origOrs {
		uniqueRefs[or.Name] = &or
	}

	for _, or := range ors {
		_, found := uniqueRefs[or.Name]
		if !found {
			diff = append(diff, or)
		}
	}

	return diff
}

func resourceKubernetesServiceAccountRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}

	log.Printf("[INFO] Reading service account %s", name)
	svcAcc, err := conn.CoreV1().ServiceAccounts(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		log.Printf("[DEBUG] Received error: %#v", err)
		return err
	}
	log.Printf("[INFO] Received service account: %#v", svcAcc)
	err = d.Set("metadata", flattenMetadata(svcAcc.ObjectMeta))
	if err != nil {
		return err
	}
	d.Set("image_pull_secret", flattenLocalObjectReferenceArray(svcAcc.ImagePullSecrets))

	defaultSecretName := d.Get("default_secret_name").(string)
	log.Printf("[DEBUG] Default secret name is %q", defaultSecretName)
	secrets := flattenServiceAccountSecrets(svcAcc.Secrets, defaultSecretName)
	log.Printf("[DEBUG] Flattened secrets: %#v", secrets)
	d.Set("secret", secrets)

	return nil
}

func resourceKubernetesServiceAccountUpdate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}

	ops := patchMetadata("metadata.0.", "/metadata/", d)
	if d.HasChange("image_pull_secret") {
		v := d.Get("image_pull_secret").(*schema.Set).List()
		ops = append(ops, &ReplaceOperation{
			Path:  "/imagePullSecrets",
			Value: expandLocalObjectReferenceArray(v),
		})
	}
	if d.HasChange("secret") {
		v := d.Get("secret").(*schema.Set).List()
		defaultSecretName := d.Get("default_secret_name").(string)

		ops = append(ops, &ReplaceOperation{
			Path:  "/secrets",
			Value: expandServiceAccountSecrets(v, defaultSecretName),
		})
	}
	data, err := ops.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Failed to marshal update operations: %s", err)
	}
	log.Printf("[INFO] Updating service account %q: %v", name, string(data))
	out, err := conn.CoreV1().ServiceAccounts(namespace).Patch(name, pkgApi.JSONPatchType, data)
	if err != nil {
		return fmt.Errorf("Failed to update service account: %s", err)
	}
	log.Printf("[INFO] Submitted updated service account: %#v", out)
	d.SetId(buildId(out.ObjectMeta))

	return resourceKubernetesServiceAccountRead(d, meta)
}

func resourceKubernetesServiceAccountDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}

	log.Printf("[INFO] Deleting service account: %#v", name)
	err = conn.CoreV1().ServiceAccounts(namespace).Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	log.Printf("[INFO] Service account %s deleted", name)

	d.SetId("")
	return nil
}

func resourceKubernetesServiceAccountExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	conn := meta.(*kubernetes.Clientset)

	namespace, name, err := idParts(d.Id())
	if err != nil {
		return false, err
	}

	log.Printf("[INFO] Checking service account %s", name)
	_, err = conn.CoreV1().ServiceAccounts(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		if statusErr, ok := err.(*errors.StatusError); ok && statusErr.ErrStatus.Code == 404 {
			return false, nil
		}
		log.Printf("[DEBUG] Received error: %#v", err)
	}
	return true, err
}
