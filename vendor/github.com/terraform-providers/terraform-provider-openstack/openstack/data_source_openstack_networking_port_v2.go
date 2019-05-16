package openstack

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/structure"
	"github.com/hashicorp/terraform/helper/validation"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/dns"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
)

func dataSourceNetworkingPortV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkingPortV2Read,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"port_id": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"admin_state_up": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"network_id": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"tenant_id": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"device_owner": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"mac_address": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"device_id": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"fixed_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.SingleIP(),
			},

			"status": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"security_group_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"allowed_address_pairs": {
				Type:     schema.TypeSet,
				Computed: true,
				Set:      resourceNetworkingPortV2AllowedAddressPairsHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mac_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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

			"all_tags": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"extra_dhcp_option": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_version": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},

			"binding": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"profile": {
							Type:     schema.TypeString,
							Computed: true,
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
							Computed: true,
						},
					},
				},
			},

			"dns_name": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"dns_assignment": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},
		},
	}
}

func dataSourceNetworkingPortV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	listOpts := ports.ListOpts{}
	var listOptsBuilder ports.ListOptsBuilder

	if v, ok := d.GetOk("port_id"); ok {
		listOpts.ID = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		listOpts.Name = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		listOpts.Description = v.(string)
	}

	if v, ok := d.GetOkExists("admin_state_up"); ok {
		asu := v.(bool)
		listOpts.AdminStateUp = &asu
	}

	if v, ok := d.GetOk("network_id"); ok {
		listOpts.NetworkID = v.(string)
	}

	if v, ok := d.GetOk("status"); ok {
		listOpts.Status = v.(string)
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		listOpts.TenantID = v.(string)
	}

	if v, ok := d.GetOk("project_id"); ok {
		listOpts.ProjectID = v.(string)
	}

	if v, ok := d.GetOk("device_owner"); ok {
		listOpts.DeviceOwner = v.(string)
	}

	if v, ok := d.GetOk("mac_address"); ok {
		listOpts.MACAddress = v.(string)
	}

	if v, ok := d.GetOk("device_id"); ok {
		listOpts.DeviceID = v.(string)
	}

	tags := networkV2AttributesTags(d)
	if len(tags) > 0 {
		listOpts.Tags = strings.Join(tags, ",")
	}

	listOptsBuilder = listOpts

	if v, ok := d.GetOk("dns_name"); ok {
		listOptsBuilder = dns.PortListOptsExt{
			ListOptsBuilder: listOptsBuilder,
			DNSName:         v.(string),
		}
	}

	allPages, err := ports.List(networkingClient, listOptsBuilder).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to list openstack_networking_ports_v2: %s", err)
	}

	var allPorts []portExtended

	err = ports.ExtractPortsInto(allPages, &allPorts)
	if err != nil {
		return fmt.Errorf("Unable to retrieve openstack_networking_ports_v2: %s", err)
	}

	if len(allPorts) == 0 {
		return fmt.Errorf("No openstack_networking_port_v2 found")
	}

	var portsList []portExtended

	// Filter returned Fixed IPs by a "fixed_ip".
	if v, ok := d.GetOk("fixed_ip"); ok {
		for _, p := range allPorts {
			for _, ipObject := range p.FixedIPs {
				if v.(string) == ipObject.IPAddress {
					portsList = append(portsList, p)
				}
			}
		}
		if len(portsList) == 0 {
			log.Printf("No openstack_networking_port_v2 found after the 'fixed_ip' filter")
			return fmt.Errorf("No openstack_networking_port_v2 found")
		}
	} else {
		portsList = allPorts
	}

	securityGroups := expandToStringSlice(d.Get("security_group_ids").(*schema.Set).List())
	if len(securityGroups) > 0 {
		var sgPorts []portExtended
		for _, p := range portsList {
			for _, sg := range p.SecurityGroups {
				if strSliceContains(securityGroups, sg) {
					sgPorts = append(sgPorts, p)
				}
			}
		}
		if len(sgPorts) == 0 {
			log.Printf("[DEBUG] No openstack_networking_port_v2 found after the 'security_group_ids' filter")
			return fmt.Errorf("No openstack_networking_port_v2 found")
		}
		portsList = sgPorts
	}

	if len(portsList) > 1 {
		return fmt.Errorf("More than one openstack_networking_port_v2 found (%d)", len(portsList))
	}

	port := portsList[0]

	log.Printf("[DEBUG] Retrieved openstack_networking_port_v2 %s: %+v", port.ID, port)
	d.SetId(port.ID)

	d.Set("port_id", port.ID)
	d.Set("name", port.Name)
	d.Set("description", port.Description)
	d.Set("admin_state_up", port.AdminStateUp)
	d.Set("network_id", port.NetworkID)
	d.Set("tenant_id", port.TenantID)
	d.Set("project_id", port.ProjectID)
	d.Set("device_owner", port.DeviceOwner)
	d.Set("mac_address", port.MACAddress)
	d.Set("device_id", port.DeviceID)
	d.Set("region", GetRegion(d, config))
	d.Set("all_tags", port.Tags)
	d.Set("all_security_group_ids", port.SecurityGroups)
	d.Set("all_fixed_ips", expandNetworkingPortFixedIPToStringSlice(port.FixedIPs))
	d.Set("allowed_address_pairs", flattenNetworkingPortAllowedAddressPairsV2(port.MACAddress, port.AllowedAddressPairs))
	d.Set("extra_dhcp_option", flattenNetworkingPortDHCPOptsV2(port.ExtraDHCPOptsExt))
	d.Set("binding", flattenNetworkingPortBindingV2(port))
	d.Set("dns_name", port.DNSName)
	d.Set("dns_assignment", port.DNSAssignment)

	return nil
}
