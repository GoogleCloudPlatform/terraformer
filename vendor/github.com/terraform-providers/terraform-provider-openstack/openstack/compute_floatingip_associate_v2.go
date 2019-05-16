package openstack

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/floatingips"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	nfloatingips "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/floatingips"
)

func parseComputeFloatingIPAssociateId(id string) (string, string, string, error) {
	idParts := strings.Split(id, "/")
	if len(idParts) < 3 {
		return "", "", "", fmt.Errorf("Unable to determine floating ip association ID")
	}

	floatingIP := idParts[0]
	instanceId := idParts[1]
	fixedIP := idParts[2]

	return floatingIP, instanceId, fixedIP, nil
}

func computeFloatingIPAssociateV2NetworkExists(networkClient *gophercloud.ServiceClient, floatingIP string) (bool, error) {
	listOpts := nfloatingips.ListOpts{
		FloatingIP: floatingIP,
	}
	allPages, err := nfloatingips.List(networkClient, listOpts).AllPages()
	if err != nil {
		return false, err
	}

	allFips, err := nfloatingips.ExtractFloatingIPs(allPages)
	if err != nil {
		return false, err
	}

	if len(allFips) > 1 {
		return false, fmt.Errorf("more than one floating IP with %s address found", floatingIP)
	}

	if len(allFips) == 0 {
		return false, nil
	}

	return true, nil
}

func computeFloatingIPAssociateV2ComputeExists(computeClient *gophercloud.ServiceClient, floatingIP string) (bool, error) {
	// If the Network API isn't available, fall back to the deprecated Compute API.
	allPages, err := floatingips.List(computeClient).AllPages()
	if err != nil {
		return false, err
	}

	allFips, err := floatingips.ExtractFloatingIPs(allPages)
	if err != nil {
		return false, err
	}

	for _, f := range allFips {
		if f.IP == floatingIP {
			return true, nil
		}
	}

	return false, nil
}

func computeFloatingIPAssociateV2CheckAssociation(
	computeClient *gophercloud.ServiceClient, instanceId, floatingIP string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		instance, err := servers.Get(computeClient, instanceId).Extract()
		if err != nil {
			return instance, "", err
		}

		var associated bool
		for _, networkAddresses := range instance.Addresses {
			for _, element := range networkAddresses.([]interface{}) {
				address := element.(map[string]interface{})
				if address["OS-EXT-IPS:type"] == "floating" && address["addr"] == floatingIP {
					associated = true
				}
			}
		}

		if associated {
			return instance, "ASSOCIATED", nil
		}

		return instance, "NOT_ASSOCIATED", nil
	}
}
