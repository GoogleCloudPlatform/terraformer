package openstack

import (
	"fmt"
	"log"
	"time"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/fwaas/policies"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFWPolicyV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceFWPolicyV1Create,
		Read:   resourceFWPolicyV1Read,
		Update: resourceFWPolicyV1Update,
		Delete: resourceFWPolicyV1Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
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
				Optional: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"audited": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"shared": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"tenant_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},

			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"value_specs": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceFWPolicyV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	audited := d.Get("audited").(bool)
	opts := PolicyCreateOpts{
		policies.CreateOpts{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Audited:     &audited,
			TenantID:    d.Get("tenant_id").(string),
			Rules:       expandToStringSlice(d.Get("rules").([]interface{})),
		},
		MapValueSpecs(d),
	}

	if r, ok := d.GetOk("shared"); ok {
		shared := r.(bool)
		opts.Shared = &shared
	}

	log.Printf("[DEBUG] openstack_fw_policy_v1 create options: %#v", opts)

	policy, err := policies.Create(networkingClient, opts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating openstack_fw_policy_v1: %s", err)
	}

	log.Printf("[DEBUG] openstack_fw_policy_v1 %s created: %#v", policy.ID, policy)

	d.SetId(policy.ID)

	return resourceFWPolicyV1Read(d, meta)
}

func resourceFWPolicyV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	policy, err := policies.Get(networkingClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "Error retrieving openstack_fw_policy_v1")
	}

	log.Printf("[DEBUG] Retrieved openstack_fw_policy_v1 %s: %#v", d.Id(), policy)

	d.Set("name", policy.Name)
	d.Set("description", policy.Description)
	d.Set("shared", policy.Shared)
	d.Set("audited", policy.Audited)
	d.Set("tenant_id", policy.TenantID)
	d.Set("rules", policy.Rules)
	d.Set("region", GetRegion(d, config))

	return nil
}

func resourceFWPolicyV1Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	opts := policies.UpdateOpts{}

	if d.HasChange("name") {
		name := d.Get("name").(string)
		opts.Name = &name
	}

	if d.HasChange("description") {
		description := d.Get("description").(string)
		opts.Description = &description
	}

	if d.HasChange("rules") {
		opts.Rules = expandToStringSlice(d.Get("rules").([]interface{}))
	}

	log.Printf("[DEBUG] openstack_fw_policy_v1 %s update options: %#v", d.Id(), opts)

	err = policies.Update(networkingClient, d.Id(), opts).Err
	if err != nil {
		return fmt.Errorf("Error updating openstack_fw_policy_v1 %s: %s", d.Id(), err)
	}

	return resourceFWPolicyV1Read(d, meta)
}

func resourceFWPolicyV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	_, err = policies.Get(networkingClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "Error retrieving openstack_fw_policy_v1")
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"ACTIVE"},
		Target:     []string{"DELETED"},
		Refresh:    fwPolicyV1DeleteFunc(networkingClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      0,
		MinTimeout: 2 * time.Second,
	}

	if _, err = stateConf.WaitForState(); err != nil {
		return fmt.Errorf("Error waiting for openstack_fw_policy_v1 %s to be deleted: %s", d.Id(), err)
	}

	return nil
}
