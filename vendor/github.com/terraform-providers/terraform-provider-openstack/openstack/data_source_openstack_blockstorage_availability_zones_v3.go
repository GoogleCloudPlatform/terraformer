package openstack

import (
	"fmt"
	"sort"

	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/availabilityzones"
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func dataSourceBlockStorageAvailabilityZonesV3() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBlockStorageAvailabilityZonesV3Read,
		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},

			"state": {
				Type:         schema.TypeString,
				Default:      "available",
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"available", "unavailable"}, true),
			},

			"names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceBlockStorageAvailabilityZonesV3Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.blockStorageV3Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack block storage client: %s", err)
	}

	allPages, err := availabilityzones.List(client).AllPages()
	if err != nil {
		return fmt.Errorf("Error retrieving openstack_blockstorage_availability_zones_v3: %s", err)
	}
	zoneInfo, err := availabilityzones.ExtractAvailabilityZones(allPages)
	if err != nil {
		return fmt.Errorf("Error extracting openstack_blockstorage_availability_zones_v3 from response: %s", err)
	}

	stateBool := d.Get("state").(string) == "available"
	var zones []string
	for _, z := range zoneInfo {
		if z.ZoneState.Available == stateBool {
			zones = append(zones, z.ZoneName)
		}
	}

	// sort.Strings sorts in place, returns nothing
	sort.Strings(zones)

	d.SetId(hashcode.Strings(zones))
	d.Set("names", zones)
	d.Set("region", GetRegion(d, config))

	return nil
}
