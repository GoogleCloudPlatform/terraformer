package openstack

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/rules"
)

func resourceNetworkingSecGroupRuleV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkingSecGroupRuleV2Create,
		Read:   resourceNetworkingSecGroupRuleV2Read,
		Delete: resourceNetworkingSecGroupRuleV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: false,
				ForceNew: true,
			},

			"direction": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"ethertype": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"port_range_min": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},

			"port_range_max": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},

			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},

			"remote_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},

			"remote_ip_prefix": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
				StateFunc: func(v interface{}) string {
					return strings.ToLower(v.(string))
				},
			},

			"security_group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"tenant_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
		},
	}
}

func resourceNetworkingSecGroupRuleV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	portRangeMin := d.Get("port_range_min").(int)
	portRangeMax := d.Get("port_range_max").(int)
	protocol := d.Get("protocol").(string)

	if protocol == "" {
		if portRangeMin != 0 || portRangeMax != 0 {
			return fmt.Errorf("A protocol must be specified when using port_range_min and port_range_max for openstack_networking_secgroup_rule_v2")
		}
	}

	opts := rules.CreateOpts{
		Description:    d.Get("description").(string),
		SecGroupID:     d.Get("security_group_id").(string),
		PortRangeMin:   d.Get("port_range_min").(int),
		PortRangeMax:   d.Get("port_range_max").(int),
		RemoteGroupID:  d.Get("remote_group_id").(string),
		RemoteIPPrefix: d.Get("remote_ip_prefix").(string),
		ProjectID:      d.Get("tenant_id").(string),
	}

	if v, ok := d.GetOk("direction"); ok {
		direction, err := resourceNetworkingSecGroupRuleV2Direction(v.(string))
		if err != nil {
			return err
		}
		opts.Direction = direction
	}

	if v, ok := d.GetOk("ethertype"); ok {
		ethertype, err := resourceNetworkingSecGroupRuleV2EtherType(v.(string))
		if err != nil {
			return err
		}
		opts.EtherType = ethertype
	}

	if v, ok := d.GetOk("protocol"); ok {
		protocol, err := resourceNetworkingSecGroupRuleV2Protocol(v.(string))
		if err != nil {
			return err
		}
		opts.Protocol = protocol
	}

	log.Printf("[DEBUG] openstack_networking_secgroup_rule_v2 create options: %#v", opts)

	sgRule, err := rules.Create(networkingClient, opts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating openstack_networking_secgroup_rule_v2: %s", err)
	}

	d.SetId(sgRule.ID)

	log.Printf("[DEBUG] Created openstack_networking_secgroup_rule_v2 %s: %#v", sgRule.ID, sgRule)
	return resourceNetworkingSecGroupRuleV2Read(d, meta)
}

func resourceNetworkingSecGroupRuleV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	sgRule, err := rules.Get(networkingClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "Error getting openstack_networking_secgroup_rule_v2")
	}

	log.Printf("[DEBUG] Retrieved openstack_networking_secgroup_rule_v2 %s: %#v", d.Id(), sgRule)

	d.Set("description", sgRule.Description)
	d.Set("direction", sgRule.Direction)
	d.Set("ethertype", sgRule.EtherType)
	d.Set("protocol", sgRule.Protocol)
	d.Set("port_range_min", sgRule.PortRangeMin)
	d.Set("port_range_max", sgRule.PortRangeMax)
	d.Set("remote_group_id", sgRule.RemoteGroupID)
	d.Set("remote_ip_prefix", sgRule.RemoteIPPrefix)
	d.Set("security_group_id", sgRule.SecGroupID)
	d.Set("tenant_id", sgRule.TenantID)
	d.Set("region", GetRegion(d, config))

	return nil
}

func resourceNetworkingSecGroupRuleV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	if err := rules.Delete(networkingClient, d.Id()).ExtractErr(); err != nil {
		return CheckDeleted(d, err, "Error deleting openstack_networking_secgroup_rule_v2")
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"ACTIVE"},
		Target:     []string{"DELETED"},
		Refresh:    resourceNetworkingSecGroupRuleV2StateRefreshFunc(networkingClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for openstack_networking_secgroup_rule_v2 %s to delete: %s", d.Id(), err)
	}

	d.SetId("")
	return nil
}
