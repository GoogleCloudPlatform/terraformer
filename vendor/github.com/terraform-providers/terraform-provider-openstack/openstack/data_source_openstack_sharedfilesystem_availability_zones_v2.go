package openstack

import (
	"fmt"
	"sort"

	"github.com/gophercloud/gophercloud/openstack/sharedfilesystems/v2/availabilityzones"
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceSharedFilesystemAvailabilityZonesV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSharedFilesystemAvailabilityZonesV2Read,
		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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

func dataSourceSharedFilesystemAvailabilityZonesV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sharedfilesystemV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack sharedfilesystem client: %s", err)
	}

	allPages, err := availabilityzones.List(client).AllPages()
	if err != nil {
		return fmt.Errorf("Error retrieving openstack_sharedfilesystem_availability_zones_v2: %s", err)
	}
	zoneInfo, err := availabilityzones.ExtractAvailabilityZones(allPages)
	if err != nil {
		return fmt.Errorf("Error extracting openstack_sharedfilesystem_availability_zones_v2 from response: %s", err)
	}

	var zones []string
	for _, z := range zoneInfo {
		zones = append(zones, z.Name)
	}

	// sort.Strings sorts in place, returns nothing
	sort.Strings(zones)

	d.SetId(hashcode.Strings(zones))
	d.Set("names", zones)
	d.Set("region", GetRegion(d, config))

	return nil
}
