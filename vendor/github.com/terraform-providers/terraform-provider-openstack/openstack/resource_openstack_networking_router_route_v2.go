package openstack

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/routers"
)

func resourceNetworkingRouterRouteV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkingRouterRouteV2Create,
		Read:   resourceNetworkingRouterRouteV2Read,
		Delete: resourceNetworkingRouterRouteV2Delete,
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

			"router_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"destination_cidr": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"next_hop": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceNetworkingRouterRouteV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	routerID := d.Get("router_id").(string)
	osMutexKV.Lock(routerID)
	defer osMutexKV.Unlock(routerID)

	r, err := routers.Get(networkingClient, routerID).Extract()
	if err != nil {
		return CheckDeleted(d, err, "Error getting openstack_networking_router_v2")
	}

	log.Printf("[DEBUG] Retrieved openstack_networking_router_v2 %s: %#v", routerID, r)

	routes := r.Routes
	dstCIDR := d.Get("destination_cidr").(string)
	nextHop := d.Get("next_hop").(string)
	exists := false

	for _, route := range routes {
		if route.DestinationCIDR == dstCIDR && route.NextHop == nextHop {
			exists = true
			break
		}
	}

	if exists {
		log.Printf("[DEBUG] openstack_networking_router_v2 %s already has route to %s via %s", routerID, dstCIDR, nextHop)
		return resourceNetworkingRouterRouteV2Read(d, meta)
	}

	routes = append(routes, routers.Route{
		DestinationCIDR: dstCIDR,
		NextHop:         nextHop,
	})
	updateOpts := routers.UpdateOpts{
		Routes: routes,
	}
	log.Printf("[DEBUG] openstack_networking_router_v2 %s update options: %#v", routerID, updateOpts)
	_, err = routers.Update(networkingClient, routerID, updateOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error updating openstack_networking_router_v2: %s", err)
	}

	d.SetId(resourceNetworkingRouterRouteV2BuildID(routerID, dstCIDR, nextHop))

	return resourceNetworkingRouterRouteV2Read(d, meta)
}

func resourceNetworkingRouterRouteV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	idFromResource, dstCIDR, nextHop, err := resourceNetworkingRouterRouteV2ParseID(d.Id())
	if err != nil {
		return fmt.Errorf("Error reading openstack_networking_router_route_v2 ID %s: %s", d.Id(), err)
	}

	routerID := d.Get("router_id").(string)
	if routerID == "" {
		routerID = idFromResource
	}
	d.Set("router_id", routerID)

	r, err := routers.Get(networkingClient, routerID).Extract()
	if err != nil {
		return CheckDeleted(d, err, "Error getting openstack_networking_router_v2")
	}

	log.Printf("[DEBUG] Retrieved openstack_networking_router_v2 %s: %#v", routerID, r)

	for _, route := range r.Routes {
		if route.DestinationCIDR == dstCIDR && route.NextHop == nextHop {
			d.Set("destination_cidr", dstCIDR)
			d.Set("next_hop", nextHop)
			break
		}
	}

	d.Set("region", GetRegion(d, config))

	return nil
}

func resourceNetworkingRouterRouteV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkingV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	routerID := d.Get("router_id").(string)
	osMutexKV.Lock(routerID)
	defer osMutexKV.Unlock(routerID)

	r, err := routers.Get(networkingClient, routerID).Extract()
	if err != nil {
		return CheckDeleted(d, err, "Error getting openstack_networking_router_v2")
	}

	log.Printf("[DEBUG] Retrieved openstack_networking_router_v2 %s: %#v", routerID, r)

	dstCIDR := d.Get("destination_cidr").(string)
	nextHop := d.Get("next_hop").(string)

	oldRoutes := r.Routes
	newRoute := []routers.Route{}

	for _, route := range oldRoutes {
		if route.DestinationCIDR != dstCIDR || route.NextHop != nextHop {
			newRoute = append(newRoute, route)
		}
	}

	if len(oldRoutes) == len(newRoute) {
		return fmt.Errorf("Can't find route to %s via %s on openstack_networking_router_v2 %s", dstCIDR, nextHop, routerID)
	}

	log.Printf("[DEBUG] Deleting openstack_networking_router_v2 %s route to %s via %s", routerID, dstCIDR, nextHop)
	updateOpts := routers.UpdateOpts{
		Routes: newRoute,
	}
	_, err = routers.Update(networkingClient, routerID, updateOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error updating openstack_networking_router_v2: %s", err)
	}

	return nil
}
