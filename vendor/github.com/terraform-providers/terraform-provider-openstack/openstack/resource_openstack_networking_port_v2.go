package openstack

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/attributestags"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/dns"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/extradhcpopts"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/portsbinding"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/portsecurity"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/structure"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceNetworkingPortV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkingPortV2Create,
		Read:   resourceNetworkingPortV2Read,
		Update: resourceNetworkingPortV2Update,
		Delete: resourceNetworkingPortV2Delete,
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
				ForceNew: false,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"network_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"admin_state_up": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
				Computed: true,
			},

			"mac_address": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},

			"tenant_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},

			"device_owner": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},

			"security_group_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: false,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			"no_security_groups": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
			},

			"device_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},

			"fixed_ip": {
				Type:          schema.TypeList,
				Optional:      true,
				ForceNew:      false,
				ConflictsWith: []string{"no_fixed_ip"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"ip_address": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"no_fixed_ip": {
				Type:          schema.TypeBool,
				Optional:      true,
				ForceNew:      false,
				ConflictsWith: []string{"fixed_ip"},
			},

			"allowed_address_pairs": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: false,
				Set:      resourceNetworkingPortV2AllowedAddressPairsHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address": {
							Type:     schema.TypeString,
							Required: true,
						},
						"mac_address": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"extra_dhcp_option": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
						"ip_version": {
							Type:     schema.TypeInt,
							Default:  4,
							Optional: true,
						},
					},
				},
			},

			"value_specs": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
			},

			"all_fixed_ips": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"all_security_group_ids": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"all_tags": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"port_security_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"binding": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"profile": {
							Type:             schema.TypeString,
							Optional:         true,
							ValidateFunc:     validation.ValidateJsonString,
							DiffSuppressFunc: suppressDiffPortBindingProfileV2,
							StateFunc: func(v interface{}) string {
								json, _ := structure.NormalizeJsonString(v)
								return json
							},
						},
						"vif_details": {
							Type:     schema.TypeMap,
							Computed: true,
						},
						"vif_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vnic_type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "normal",
							ValidateFunc: validation.StringInSlice([]string{
								"direct", "direct-physical", "macvtap", "normal", "baremetal", "virtio-forwarder",
							}, true),
						},
					},
				},
			},

			"dns_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"dns_assignment": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},
		},
	}
}

func resourceNetworkingPortV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	securityGroups := expandToStringSlice(d.Get("security_group_ids").(*schema.Set).List())
	noSecurityGroups := d.Get("no_security_groups").(bool)

	// Check and make sure an invalid security group configuration wasn't given.
	if noSecurityGroups && len(securityGroups) > 0 {
		return fmt.Errorf("Cannot have both no_security_groups and security_group_ids set for openstack_networking_port_v2")
	}

	allowedAddressPairs := d.Get("allowed_address_pairs").(*schema.Set)
	createOpts := PortCreateOpts{
		ports.CreateOpts{
			Name:                d.Get("name").(string),
			Description:         d.Get("description").(string),
			NetworkID:           d.Get("network_id").(string),
			MACAddress:          d.Get("mac_address").(string),
			TenantID:            d.Get("tenant_id").(string),
			DeviceOwner:         d.Get("device_owner").(string),
			DeviceID:            d.Get("device_id").(string),
			FixedIPs:            expandNetworkingPortFixedIPV2(d),
			AllowedAddressPairs: expandNetworkingPortAllowedAddressPairsV2(allowedAddressPairs),
		},
		MapValueSpecs(d),
	}

	if v, ok := d.GetOkExists("admin_state_up"); ok {
		asu := v.(bool)
		createOpts.AdminStateUp = &asu
	}

	if noSecurityGroups {
		securityGroups = []string{}
		createOpts.SecurityGroups = &securityGroups
	}

	// Only set SecurityGroups if one was specified.
	// Otherwise this would mimic the no_security_groups action.
	if len(securityGroups) > 0 {
		createOpts.SecurityGroups = &securityGroups
	}

	// Declare a finalCreateOpts interface to hold either the
	// base create options or the extended DHCP options.
	var finalCreateOpts ports.CreateOptsBuilder
	finalCreateOpts = createOpts

	dhcpOpts := d.Get("extra_dhcp_option").(*schema.Set)
	if dhcpOpts.Len() > 0 {
		finalCreateOpts = extradhcpopts.CreateOptsExt{
			CreateOptsBuilder: createOpts,
			ExtraDHCPOpts:     expandNetworkingPortDHCPOptsV2Create(dhcpOpts),
		}
	}

	// Add the port security attribute if specified.
	if v, ok := d.GetOkExists("port_security_enabled"); ok {
		portSecurityEnabled := v.(bool)
		finalCreateOpts = portsecurity.PortCreateOptsExt{
			CreateOptsBuilder:   finalCreateOpts,
			PortSecurityEnabled: &portSecurityEnabled,
		}
	}

	// Add the port binding parameters if specified.
	if v, ok := d.GetOkExists("binding"); ok {
		for _, raw := range v.([]interface{}) {
			binding := raw.(map[string]interface{})
			profile := map[string]interface{}{}

			// Convert raw string into the map
			rawProfile := binding["profile"].(string)
			if len(rawProfile) > 0 {
				err := json.Unmarshal([]byte(rawProfile), &profile)
				if err != nil {
					return fmt.Errorf("Failed to unmarshal the JSON: %s", err)
				}
			}

			finalCreateOpts = portsbinding.CreateOptsExt{
				CreateOptsBuilder: finalCreateOpts,
				HostID:            binding["host_id"].(string),
				Profile:           profile,
				VNICType:          binding["vnic_type"].(string),
			}
		}
	}

	if dnsName := d.Get("dns_name").(string); dnsName != "" {
		finalCreateOpts = dns.PortCreateOptsExt{
			CreateOptsBuilder: finalCreateOpts,
			DNSName:           dnsName,
		}
	}

	log.Printf("[DEBUG] openstack_networking_port_v2 create options: %#v", finalCreateOpts)

	// Create a Neutron port and set extra options if they're specified.
	var port portExtended

	err = ports.Create(networkingClient, finalCreateOpts).ExtractInto(&port)
	if err != nil {
		return fmt.Errorf("Error creating openstack_networking_port_v2: %s", err)
	}

	log.Printf("[DEBUG] Waiting for openstack_networking_port_v2 %s to become available.", port.ID)

	stateConf := &resource.StateChangeConf{
		Target:     []string{"ACTIVE", "DOWN"},
		Refresh:    resourceNetworkingPortV2StateRefreshFunc(networkingClient, port.ID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for openstack_networking_port_v2 %s to become available: %s", port.ID, err)
	}

	d.SetId(port.ID)

	tags := networkV2AttributesTags(d)
	if len(tags) > 0 {
		tagOpts := attributestags.ReplaceAllOpts{Tags: tags}
		tags, err := attributestags.ReplaceAll(networkingClient, "ports", port.ID, tagOpts).Extract()
		if err != nil {
			return fmt.Errorf("Error setting tags on openstack_networking_port_v2 %s: %s", port.ID, err)
		}
		log.Printf("[DEBUG] Set tags %s on openstack_networking_port_v2 %s", tags, port.ID)
	}

	log.Printf("[DEBUG] Created openstack_networking_port_v2 %s: %#v", port.ID, port)
	return resourceNetworkingPortV2Read(d, meta)
}

func resourceNetworkingPortV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	var port portExtended
	err = ports.Get(networkingClient, d.Id()).ExtractInto(&port)
	if err != nil {
		return CheckDeleted(d, err, "Error getting openstack_networking_port_v2")
	}

	log.Printf("[DEBUG] Retrieved openstack_networking_port_v2 %s: %#v", d.Id(), port)

	d.Set("name", port.Name)
	d.Set("description", port.Description)
	d.Set("admin_state_up", port.AdminStateUp)
	d.Set("network_id", port.NetworkID)
	d.Set("mac_address", port.MACAddress)
	d.Set("tenant_id", port.TenantID)
	d.Set("device_owner", port.DeviceOwner)
	d.Set("device_id", port.DeviceID)

	networkV2ReadAttributesTags(d, port.Tags)

	// Set a slice of all returned Fixed IPs.
	// This will be in the order returned by the API,
	// which is usually alpha-numeric.
	d.Set("all_fixed_ips", expandNetworkingPortFixedIPToStringSlice(port.FixedIPs))

	// Set all security groups.
	// This can be different from what the user specified since
	// the port can have the "default" group automatically applied.
	d.Set("all_security_group_ids", port.SecurityGroups)

	d.Set("allowed_address_pairs", flattenNetworkingPortAllowedAddressPairsV2(port.MACAddress, port.AllowedAddressPairs))
	d.Set("extra_dhcp_option", flattenNetworkingPortDHCPOptsV2(port.ExtraDHCPOptsExt))
	d.Set("port_security_enabled", port.PortSecurityEnabled)
	d.Set("binding", flattenNetworkingPortBindingV2(port))
	d.Set("dns_name", port.DNSName)
	d.Set("dns_assignment", port.DNSAssignment)

	d.Set("region", GetRegion(d, config))

	return nil
}

func resourceNetworkingPortV2Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	securityGroups := expandToStringSlice(d.Get("security_group_ids").(*schema.Set).List())
	noSecurityGroups := d.Get("no_security_groups").(bool)

	// Check and make sure an invalid security group configuration wasn't given.
	if noSecurityGroups && len(securityGroups) > 0 {
		return fmt.Errorf("Cannot have both no_security_groups and security_group_ids set for openstack_networking_port_v2")
	}

	var hasChange bool
	var updateOpts ports.UpdateOpts

	if d.HasChange("allowed_address_pairs") {
		hasChange = true
		allowedAddressPairs := d.Get("allowed_address_pairs").(*schema.Set)
		aap := expandNetworkingPortAllowedAddressPairsV2(allowedAddressPairs)
		updateOpts.AllowedAddressPairs = &aap
	}

	if d.HasChange("no_security_groups") {
		if noSecurityGroups {
			hasChange = true
			v := []string{}
			updateOpts.SecurityGroups = &v
		}
	}

	if d.HasChange("security_group_ids") {
		hasChange = true
		updateOpts.SecurityGroups = &securityGroups
	}

	if d.HasChange("name") {
		hasChange = true
		name := d.Get("name").(string)
		updateOpts.Name = &name
	}

	if d.HasChange("description") {
		hasChange = true
		description := d.Get("description").(string)
		updateOpts.Description = &description
	}

	if d.HasChange("admin_state_up") {
		hasChange = true
		asu := d.Get("admin_state_up").(bool)
		updateOpts.AdminStateUp = &asu
	}

	if d.HasChange("device_owner") {
		hasChange = true
		deviceOwner := d.Get("device_owner").(string)
		updateOpts.DeviceOwner = &deviceOwner
	}

	if d.HasChange("device_id") {
		hasChange = true
		deviceId := d.Get("device_id").(string)
		updateOpts.DeviceID = &deviceId
	}

	if d.HasChange("fixed_ip") || d.HasChange("no_fixed_ip") {
		fixedIPs := expandNetworkingPortFixedIPV2(d)
		if fixedIPs != nil {
			hasChange = true
			updateOpts.FixedIPs = fixedIPs
		}
	}

	var finalUpdateOpts ports.UpdateOptsBuilder
	finalUpdateOpts = updateOpts

	if d.HasChange("port_security_enabled") {
		hasChange = true
		portSecurityEnabled := d.Get("port_security_enabled").(bool)
		finalUpdateOpts = portsecurity.PortUpdateOptsExt{
			UpdateOptsBuilder:   finalUpdateOpts,
			PortSecurityEnabled: &portSecurityEnabled,
		}
	}

	// Next, perform any dhcp option changes.
	if d.HasChange("extra_dhcp_option") {
		hasChange = true

		o, n := d.GetChange("extra_dhcp_option")
		oldDHCPOpts := o.(*schema.Set)
		newDHCPOpts := n.(*schema.Set)

		deleteDHCPOpts := oldDHCPOpts.Difference(newDHCPOpts)
		addDHCPOpts := newDHCPOpts.Difference(oldDHCPOpts)

		updateExtraDHCPOpts := expandNetworkingPortDHCPOptsV2Update(deleteDHCPOpts, addDHCPOpts)
		finalUpdateOpts = extradhcpopts.UpdateOptsExt{
			UpdateOptsBuilder: finalUpdateOpts,
			ExtraDHCPOpts:     updateExtraDHCPOpts,
		}
	}

	// Next, perform port binding option changes.
	if d.HasChange("binding") {
		hasChange = true

		profile := map[string]interface{}{}

		// default options, when unsetting the port bindings
		newOpts := portsbinding.UpdateOptsExt{
			UpdateOptsBuilder: finalUpdateOpts,
			HostID:            new(string),
			Profile:           profile,
			VNICType:          "normal",
		}

		for _, raw := range d.Get("binding").([]interface{}) {
			binding := raw.(map[string]interface{})

			// Convert raw string into the map
			rawProfile := binding["profile"].(string)
			if len(rawProfile) > 0 {
				err := json.Unmarshal([]byte(rawProfile), &profile)
				if err != nil {
					return fmt.Errorf("Failed to unmarshal the JSON: %s", err)
				}
			}

			hostID := binding["host_id"].(string)
			newOpts = portsbinding.UpdateOptsExt{
				UpdateOptsBuilder: finalUpdateOpts,
				HostID:            &hostID,
				Profile:           profile,
				VNICType:          binding["vnic_type"].(string),
			}
		}

		finalUpdateOpts = newOpts
	}

	if d.HasChange("dns_name") {
		hasChange = true

		dnsName := d.Get("dns_name").(string)
		finalUpdateOpts = dns.PortUpdateOptsExt{
			UpdateOptsBuilder: finalUpdateOpts,
			DNSName:           &dnsName,
		}
	}

	// At this point, perform the update for all "standard" port changes.
	if hasChange {
		log.Printf("[DEBUG] openstack_networking_port_v2 %s update options: %#v", d.Id(), finalUpdateOpts)
		_, err = ports.Update(networkingClient, d.Id(), finalUpdateOpts).Extract()
		if err != nil {
			return fmt.Errorf("Error updating OpenStack Neutron Port: %s", err)
		}
	}

	// Next, perform any required updates to the tags.
	if d.HasChange("tags") {
		tags := networkV2UpdateAttributesTags(d)
		tagOpts := attributestags.ReplaceAllOpts{Tags: tags}
		tags, err := attributestags.ReplaceAll(networkingClient, "ports", d.Id(), tagOpts).Extract()
		if err != nil {
			return fmt.Errorf("Error setting tags on openstack_networking_port_v2 %s: %s", d.Id(), err)
		}
		log.Printf("[DEBUG] Set tags %s on openstack_networking_port_v2 %s", tags, d.Id())
	}

	return resourceNetworkingPortV2Read(d, meta)
}

func resourceNetworkingPortV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	if err := ports.Delete(networkingClient, d.Id()).ExtractErr(); err != nil {
		return CheckDeleted(d, err, "Error deleting openstack_networking_port_v2")
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"ACTIVE"},
		Target:     []string{"DELETED"},
		Refresh:    resourceNetworkingPortV2StateRefreshFunc(networkingClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for openstack_networking_port_v2 %s to delete: %s", d.Id(), err)
	}

	d.SetId("")
	return nil
}
