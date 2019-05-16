package openstack

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/fwaas/policies"

	"github.com/hashicorp/terraform/helper/resource"
)

// PolicyCreateOpts represents the attributes used when creating a new firewall policy.
type PolicyCreateOpts struct {
	policies.CreateOpts
	ValueSpecs map[string]string `json:"value_specs,omitempty"`
}

// ToPolicyCreateMap casts a CreateOpts struct to a map.
// It overrides policies.ToFirewallPolicyCreateMap to add the ValueSpecs field.
func (opts PolicyCreateOpts) ToFirewallPolicyCreateMap() (map[string]interface{}, error) {
	return BuildRequest(opts, "firewall_policy")
}

func fwPolicyV1DeleteFunc(networkingClient *gophercloud.ServiceClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		err := policies.Delete(networkingClient, id).Err
		if err == nil {
			return "", "DELETED", nil
		}

		if _, ok := err.(gophercloud.ErrDefault409); ok {
			// This error usually means that the policy is attached
			// to a firewall. At this point, the firewall is probably
			// being delete. So, we retry a few times.
			return nil, "ACTIVE", nil
		}

		return nil, "ACTIVE", err
	}
}
