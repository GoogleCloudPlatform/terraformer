package openstack

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/trunks"
)

func dataSourceNetworkingTrunkV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkingTrunkV2Read,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"trunk_id": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"port_id": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"admin_state_up": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"project_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"sub_port": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"segmentation_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"segmentation_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
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
		},
	}
}

func dataSourceNetworkingTrunkV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	listOpts := trunks.ListOpts{}

	if v, ok := d.GetOk("name"); ok {
		listOpts.Name = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		listOpts.Description = v.(string)
	}

	if v, ok := d.GetOk("trunk_id"); ok {
		listOpts.ID = v.(string)
	}

	if v, ok := d.GetOk("port_id"); ok {
		listOpts.PortID = v.(string)
	}

	if v, ok := d.GetOkExists("admin_state_up"); ok {
		asu := v.(bool)
		listOpts.AdminStateUp = &asu
	}

	if v, ok := d.GetOk("project_id"); ok {
		listOpts.ProjectID = v.(string)
	}

	if v, ok := d.GetOk("status"); ok {
		listOpts.Status = v.(string)
	}

	tags := networkV2AttributesTags(d)
	if len(tags) > 0 {
		listOpts.Tags = strings.Join(tags, ",")
	}

	pages, err := trunks.List(networkingClient, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to retrieve trunks: %s", err)
	}

	allTrunks, err := trunks.ExtractTrunks(pages)
	if err != nil {
		return fmt.Errorf("Unable to extract trunks: %s", err)
	}

	if len(allTrunks) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allTrunks) > 1 {
		return fmt.Errorf("Your query returned more than one result." +
			" Please try a more specific search criteria")
	}

	trunk := allTrunks[0]

	log.Printf("[DEBUG] Retrieved Trunk %s: %+v", trunk.ID, trunk)
	d.SetId(trunk.ID)

	d.Set("region", GetRegion(d, config))
	d.Set("name", trunk.Name)
	d.Set("description", trunk.Description)
	d.Set("port_id", trunk.PortID)
	d.Set("admin_state_up", trunk.AdminStateUp)
	d.Set("status", trunk.Status)
	d.Set("project_id", trunk.ProjectID)
	d.Set("all_tags", trunk.Tags)

	subports := make([]map[string]interface{}, len(trunk.Subports))
	for i, trunkSubport := range trunk.Subports {
		subports[i] = make(map[string]interface{})
		subports[i]["port_id"] = trunkSubport.PortID
		subports[i]["segmentation_type"] = trunkSubport.SegmentationType
		subports[i]["segmentation_id"] = trunkSubport.SegmentationID
	}
	if err = d.Set("sub_port", subports); err != nil {
		return fmt.Errorf("Unable to set sub_port for trunk %s: %s", d.Id(), err)
	}

	return nil
}
