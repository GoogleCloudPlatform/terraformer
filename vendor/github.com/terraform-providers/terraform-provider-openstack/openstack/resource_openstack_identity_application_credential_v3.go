package openstack

import (
	"fmt"
	"log"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/applicationcredentials"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceIdentityApplicationCredentialV3() *schema.Resource {
	return &schema.Resource{
		Create: resourceIdentityApplicationCredentialV3Create,
		Read:   resourceIdentityApplicationCredentialV3Read,
		Delete: resourceIdentityApplicationCredentialV3Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"unrestricted": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},

			"secret": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
				ForceNew:  true,
			},

			"project_id": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},

			"roles": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"expires_at": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.ValidateRFC3339TimeString,
			},
		},
	}
}

func resourceIdentityApplicationCredentialV3Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	identityClient, err := config.identityV3Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack identity client: %s", err)
	}

	token := tokens.Get(identityClient, config.OsClient.TokenID)
	if token.Err != nil {
		return token.Err
	}

	user, err := token.ExtractUser()
	if err != nil {
		return err
	}

	createOpts := applicationcredentials.CreateOpts{
		Name:         d.Get("name").(string),
		Description:  d.Get("description").(string),
		Unrestricted: d.Get("unrestricted").(bool),
		Roles:        expandIdentityApplicationCredentialRolesV3(d.Get("roles").(*schema.Set).List()),
		ExpiresAt:    d.Get("expires_at").(string),
	}

	log.Printf("[DEBUG] openstack_identity_application_credential_v3 create options: %#v", createOpts)

	createOpts.Secret = d.Get("secret").(string)

	applicationCredential, err := applicationcredentials.Create(identityClient, user.ID, createOpts).Extract()
	if err != nil {
		if v, ok := err.(gophercloud.ErrDefault404); ok {
			return fmt.Errorf("Error creating openstack_identity_application_credential_v3: %s", v.ErrUnexpectedResponseCode.Body)
		}
		return fmt.Errorf("Error creating openstack_identity_application_credential_v3: %s", err)
	}

	d.SetId(applicationCredential.ID)

	// Secret is returned only once
	d.Set("secret", applicationCredential.Secret)

	return resourceIdentityApplicationCredentialV3Read(d, meta)
}

func resourceIdentityApplicationCredentialV3Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	identityClient, err := config.identityV3Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack identity client: %s", err)
	}

	token := tokens.Get(identityClient, config.OsClient.TokenID)
	if token.Err != nil {
		return token.Err
	}

	user, err := token.ExtractUser()
	if err != nil {
		return err
	}

	applicationCredential, err := applicationcredentials.Get(identityClient, user.ID, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "Error retrieving openstack_identity_application_credential_v3")
	}

	log.Printf("[DEBUG] Retrieved openstack_identity_application_credential_v3 %s: %#v", d.Id(), applicationCredential)

	d.Set("name", applicationCredential.Name)
	d.Set("description", applicationCredential.Description)
	d.Set("unrestricted", applicationCredential.Unrestricted)
	d.Set("roles", flattenIdentityApplicationCredentialRolesV3(applicationCredential.Roles))
	d.Set("project_id", applicationCredential.ProjectID)
	d.Set("region", GetRegion(d, config))

	if applicationCredential.ExpiresAt == (time.Time{}) {
		d.Set("expires_at", "")
	} else {
		d.Set("expires_at", applicationCredential.ExpiresAt.UTC().Format(time.RFC3339))
	}

	return nil
}

func resourceIdentityApplicationCredentialV3Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	identityClient, err := config.identityV3Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack identity client: %s", err)
	}

	token := tokens.Get(identityClient, config.OsClient.TokenID)
	if token.Err != nil {
		return token.Err
	}

	user, err := token.ExtractUser()
	if err != nil {
		return err
	}

	err = applicationcredentials.Delete(identityClient, user.ID, d.Id()).ExtractErr()
	if err != nil {
		return CheckDeleted(d, err, "Error deleting openstack_identity_application_credential_v3")
	}

	return nil
}
