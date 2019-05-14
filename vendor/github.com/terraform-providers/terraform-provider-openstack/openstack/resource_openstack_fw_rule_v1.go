package openstack

import (
	"fmt"
	"log"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/fwaas/policies"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/fwaas/rules"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFWRuleV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceFWRuleV1Create,
		Read:   resourceFWRuleV1Read,
		Update: resourceFWRuleV1Update,
		Delete: resourceFWRuleV1Delete,
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
				Optional: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},

			"action": {
				Type:     schema.TypeString,
				Required: true,
			},

			"ip_version": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},

			"source_ip_address": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"destination_ip_address": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"source_port": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"destination_port": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"tenant_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"value_specs": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceFWRuleV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	enabled := d.Get("enabled").(bool)
	ruleConfiguration := RuleCreateOpts{
		rules.CreateOpts{
			Name:                 d.Get("name").(string),
			Description:          d.Get("description").(string),
			Protocol:             expandFWRuleV1Protocol(d.Get("protocol").(string)),
			Action:               d.Get("action").(string),
			IPVersion:            expandFWRuleV1IPVersion(d.Get("ip_version").(int)),
			SourceIPAddress:      d.Get("source_ip_address").(string),
			DestinationIPAddress: d.Get("destination_ip_address").(string),
			SourcePort:           d.Get("source_port").(string),
			DestinationPort:      d.Get("destination_port").(string),
			Enabled:              &enabled,
			TenantID:             d.Get("tenant_id").(string),
		},
		MapValueSpecs(d),
	}

	log.Printf("[DEBUG] openstack_fw_rule_v1 create options: %#v", ruleConfiguration)

	rule, err := rules.Create(networkingClient, ruleConfiguration).Extract()
	if err != nil {
		return fmt.Errorf("Error creating openstack_fw_rule_v1: %s", err)
	}

	log.Printf("[DEBUG] Created openstack_fw_rule_v1 %s: %#v", rule.ID, rule)

	d.SetId(rule.ID)

	return resourceFWRuleV1Read(d, meta)
}

func resourceFWRuleV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	rule, err := rules.Get(networkingClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "Error retrieving openstack_fw_rule_v1")
	}

	log.Printf("[DEBUG] Retrieved openstack_fw_rule_v1 %s: %#v", d.Id(), rule)

	d.Set("action", rule.Action)
	d.Set("name", rule.Name)
	d.Set("description", rule.Description)
	d.Set("ip_version", rule.IPVersion)
	d.Set("source_ip_address", rule.SourceIPAddress)
	d.Set("destination_ip_address", rule.DestinationIPAddress)
	d.Set("source_port", rule.SourcePort)
	d.Set("destination_port", rule.DestinationPort)
	d.Set("enabled", rule.Enabled)

	if rule.Protocol == "" {
		d.Set("protocol", "any")
	} else {
		d.Set("protocol", rule.Protocol)
	}

	d.Set("region", GetRegion(d, config))

	return nil
}

func resourceFWRuleV1Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	var updateOpts rules.UpdateOpts
	if d.HasChange("name") {
		name := d.Get("name").(string)
		updateOpts.Name = &name
	}

	if d.HasChange("description") {
		description := d.Get("description").(string)
		updateOpts.Description = &description
	}

	if d.HasChange("protocol") {
		protocol := d.Get("protocol").(string)
		updateOpts.Protocol = &protocol
	}

	if d.HasChange("action") {
		action := d.Get("action").(string)
		updateOpts.Action = &action
	}

	if d.HasChange("ip_version") {
		ipVersion := expandFWRuleV1IPVersion(d.Get("ip_version").(int))
		updateOpts.IPVersion = &ipVersion
	}

	if d.HasChange("source_ip_address") {
		sourceIPAddress := d.Get("source_ip_address").(string)
		updateOpts.SourceIPAddress = &sourceIPAddress

		// Also include the ip_version.
		ipVersion := expandFWRuleV1IPVersion(d.Get("ip_version").(int))
		updateOpts.IPVersion = &ipVersion
	}

	if d.HasChange("source_port") {
		sourcePort := d.Get("source_port").(string)
		if sourcePort == "" {
			sourcePort = "0"
		}
		updateOpts.SourcePort = &sourcePort

		// Also include the protocol.
		protocol := d.Get("protocol").(string)
		updateOpts.Protocol = &protocol
	}

	if d.HasChange("destination_ip_address") {
		destinationIPAddress := d.Get("destination_ip_address").(string)
		updateOpts.DestinationIPAddress = &destinationIPAddress

		// Also include the ip_version.
		ipVersion := expandFWRuleV1IPVersion(d.Get("ip_version").(int))
		updateOpts.IPVersion = &ipVersion
	}

	if d.HasChange("destination_port") {
		destinationPort := d.Get("destination_port").(string)
		if destinationPort == "" {
			destinationPort = "0"
		}

		updateOpts.DestinationPort = &destinationPort

		// Also include the protocol.
		protocol := d.Get("protocol").(string)
		updateOpts.Protocol = &protocol
	}

	if d.HasChange("enabled") {
		enabled := d.Get("enabled").(bool)
		updateOpts.Enabled = &enabled
	}

	log.Printf("[DEBUG] openstack_fw_rule_v1 %s update options: %#v", d.Id(), updateOpts)
	err = rules.Update(networkingClient, d.Id(), updateOpts).Err
	if err != nil {
		return fmt.Errorf("Error updating openstack_fw_rule_v1 %s: %s", d.Id(), err)
	}

	return resourceFWRuleV1Read(d, meta)
}

func resourceFWRuleV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	rule, err := rules.Get(networkingClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "Error retrieving openstack_fw_rule_v1")
	}

	if rule.PolicyID != "" {
		_, err := policies.RemoveRule(networkingClient, rule.PolicyID, rule.ID).Extract()
		if err != nil {
			return fmt.Errorf("Error removing openstack_fw_rule_v1 %s from policy %s: %s", d.Id(), rule.PolicyID, err)
		}
	}

	err = rules.Delete(networkingClient, d.Id()).ExtractErr()
	if err != nil {
		return fmt.Errorf("Error deleting openstack_fw_rule_v1 %s: %s", d.Id(), err)
	}

	return nil
}
