package openstack

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/fwaas/firewalls"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/fwaas/routerinsertion"

	"github.com/hashicorp/terraform/helper/resource"
)

// Firewall is an OpenStack firewall.
type Firewall struct {
	firewalls.Firewall
	routerinsertion.FirewallExt
}

// FirewallCreateOpts represents the attributes used when creating a new firewall.
type FirewallCreateOpts struct {
	firewalls.CreateOpts
	ValueSpecs map[string]string `json:"value_specs,omitempty"`
}

// ToFirewallCreateMap casts a CreateOptsExt struct to a map.
// It overrides firewalls.ToFirewallCreateMap to add the ValueSpecs field.
func (opts FirewallCreateOpts) ToFirewallCreateMap() (map[string]interface{}, error) {
	return BuildRequest(opts, "firewall")
}

//FirewallUpdateOpts
type FirewallUpdateOpts struct {
	firewalls.UpdateOptsBuilder
}

func (opts FirewallUpdateOpts) ToFirewallUpdateMap() (map[string]interface{}, error) {
	return BuildRequest(opts, "firewall")
}

func fwFirewallV1RefreshFunc(networkingClient *gophercloud.ServiceClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		var fw Firewall

		err := firewalls.Get(networkingClient, id).ExtractInto(&fw)
		if err != nil {
			return nil, "", err
		}

		return fw, fw.Status, nil
	}
}

func fwFirewallV1DeleteFunc(networkingClient *gophercloud.ServiceClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		fw, err := firewalls.Get(networkingClient, id).Extract()

		if err != nil {
			if _, ok := err.(gophercloud.ErrDefault404); ok {
				return "", "DELETED", nil
			}
			return nil, "", fmt.Errorf("Unexpected error: %s", err)
		}

		return fw, "DELETING", nil
	}
}
