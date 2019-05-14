package openstack

import (
	"fmt"
	"log"
	"time"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/fwaas/firewalls"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/fwaas/routerinsertion"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFWFirewallV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceFWFirewallV1Create,
		Read:   resourceFWFirewallV1Read,
		Update: resourceFWFirewallV1Update,
		Delete: resourceFWFirewallV1Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
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

			"policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"admin_state_up": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"tenant_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},

			"associated_routers": {
				Type:          schema.TypeSet,
				Optional:      true,
				Elem:          &schema.Schema{Type: schema.TypeString},
				Set:           schema.HashString,
				ConflictsWith: []string{"no_routers"},
				Computed:      true,
			},

			"no_routers": {
				Type:          schema.TypeBool,
				Optional:      true,
				ConflictsWith: []string{"associated_routers"},
			},

			"value_specs": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceFWFirewallV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	var createOpts firewalls.CreateOptsBuilder

	adminStateUp := d.Get("admin_state_up").(bool)
	createOpts = FirewallCreateOpts{
		firewalls.CreateOpts{
			Name:         d.Get("name").(string),
			Description:  d.Get("description").(string),
			PolicyID:     d.Get("policy_id").(string),
			AdminStateUp: &adminStateUp,
			TenantID:     d.Get("tenant_id").(string),
		},
		MapValueSpecs(d),
	}

	associatedRoutersRaw := d.Get("associated_routers").(*schema.Set).List()
	if len(associatedRoutersRaw) > 0 {
		var routerIds []string
		for _, v := range associatedRoutersRaw {
			routerIds = append(routerIds, v.(string))
		}

		createOpts = &routerinsertion.CreateOptsExt{
			CreateOptsBuilder: createOpts,
			RouterIDs:         routerIds,
		}
	}

	if d.Get("no_routers").(bool) {
		routerIds := make([]string, 0)
		createOpts = &routerinsertion.CreateOptsExt{
			CreateOptsBuilder: createOpts,
			RouterIDs:         routerIds,
		}
	}

	log.Printf("[DEBUG] openstack_fw_firewall_v1 create options: %#v", createOpts)

	firewall, err := firewalls.Create(networkingClient, createOpts).Extract()
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] openstack_fw_firewall_v1 created: %#v", firewall)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PENDING_CREATE"},
		Target:     []string{"ACTIVE", "INACTIVE"},
		Refresh:    fwFirewallV1RefreshFunc(networkingClient, firewall.ID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      0,
		MinTimeout: 2 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for openstack_fw_firewall_v1 to become active: %s", err)
	}

	d.SetId(firewall.ID)

	return resourceFWFirewallV1Read(d, meta)
}

func resourceFWFirewallV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	var firewall Firewall
	err = firewalls.Get(networkingClient, d.Id()).ExtractInto(&firewall)
	if err != nil {
		return CheckDeleted(d, err, "Error retrieving openstack_fw_firewall_v1")
	}

	log.Printf("[DEBUG] Retrieved openstack_fw_firewall_v1 %s: %#v", d.Id(), firewall)

	d.Set("name", firewall.Name)
	d.Set("description", firewall.Description)
	d.Set("policy_id", firewall.PolicyID)
	d.Set("admin_state_up", firewall.AdminStateUp)
	d.Set("tenant_id", firewall.TenantID)
	d.Set("associated_routers", firewall.RouterIDs)
	d.Set("region", GetRegion(d, config))

	return nil
}

func resourceFWFirewallV1Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	// PolicyID is required
	opts := firewalls.UpdateOpts{
		PolicyID: d.Get("policy_id").(string),
	}

	if d.HasChange("name") {
		name := d.Get("name").(string)
		opts.Name = &name
	}

	if d.HasChange("description") {
		description := d.Get("description").(string)
		opts.Description = &description
	}

	if d.HasChange("admin_state_up") {
		adminStateUp := d.Get("admin_state_up").(bool)
		opts.AdminStateUp = &adminStateUp
	}

	var updateOpts firewalls.UpdateOptsBuilder
	var routerIds []string
	if d.HasChange("associated_routers") || d.HasChange("no_routers") {
		// 'no_routers' = true means 'associated_routers' will be empty...
		if d.Get("no_routers").(bool) {
			routerIds = make([]string, 0)
		} else {
			associatedRoutersRaw := d.Get("associated_routers").(*schema.Set).List()
			for _, v := range associatedRoutersRaw {
				routerIds = append(routerIds, v.(string))
			}
		}

		updateOpts = routerinsertion.UpdateOptsExt{
			UpdateOptsBuilder: opts,
			RouterIDs:         routerIds,
		}
	} else {
		updateOpts = opts
	}

	log.Printf("[DEBUG] openstack_fw_firewall_v1 %s update options: %#v", d.Id(), updateOpts)

	err = firewalls.Update(networkingClient, d.Id(), updateOpts).Err
	if err != nil {
		return err
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PENDING_CREATE", "PENDING_UPDATE"},
		Target:     []string{"ACTIVE", "INACTIVE"},
		Refresh:    fwFirewallV1RefreshFunc(networkingClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutUpdate),
		Delay:      0,
		MinTimeout: 2 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for openstack_fw_firewall_v1 %s to become active: %s", d.Id(), err)
	}

	return resourceFWFirewallV1Read(d, meta)
}

func resourceFWFirewallV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	_, err = firewalls.Get(networkingClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "Error retrieving openstack_fw_firewall_v1")
	}

	// Ensure the firewall was fully created/updated before being deleted.
	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PENDING_CREATE", "PENDING_UPDATE"},
		Target:     []string{"ACTIVE", "INACTIVE"},
		Refresh:    fwFirewallV1RefreshFunc(networkingClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutUpdate),
		Delay:      0,
		MinTimeout: 2 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for openstack_fw_firewall_v1 %s to become active: %s", d.Id(), err)
	}

	err = firewalls.Delete(networkingClient, d.Id()).Err
	if err != nil {
		return fmt.Errorf("Error deleting openstack_fw_firewall_v1 %s: %s", d.Id(), err)
	}

	stateConf = &resource.StateChangeConf{
		Pending:    []string{"DELETING"},
		Target:     []string{"DELETED"},
		Refresh:    fwFirewallV1DeleteFunc(networkingClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      0,
		MinTimeout: 2 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for openstack_fw_firewall_v1 %s to delete: %s", d.Id(), err)
	}

	return err
}
