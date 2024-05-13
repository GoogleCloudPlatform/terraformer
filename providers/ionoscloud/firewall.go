package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type FirewallGenerator struct {
	Service
}

func (g *FirewallGenerator) InitResources() error {
	client := g.generateClient()
	cloudAPIClient := client.CloudAPIClient
	resourceType := "ionoscloud_firewall"

	datacenters, err := helpers.GetAllDatacenters(*cloudAPIClient)
	if err != nil {
		return err
	}
	for _, datacenter := range datacenters {
		servers, _, err := cloudAPIClient.ServersApi.DatacentersServersGet(context.TODO(), *datacenter.Id).Execute()
		if err != nil {
			return err
		}
		if servers.Items == nil {
			log.Printf(
				"[WARNING] expected a response containing servers but received 'nil' instead, skipping search for datacenter with ID: %v.\n",
				*datacenter.Id)
			continue
		}
		for _, server := range *servers.Items {
			nics, _, err := cloudAPIClient.NetworkInterfacesApi.DatacentersServersNicsGet(context.TODO(), *datacenter.Id, *server.Id).Execute()
			if err != nil {
				return err
			}
			if nics.Items == nil {
				log.Printf(
					"[WARNING] expected a response containing NICs but received 'nil' instead, skipping search for server with ID: %v, datacenter ID: %v.\n",
					*server.Id,
					*datacenter.Id)
				continue
			}
			lastNicIdx := len(*nics.Items) - 1
			for nicIdx, nic := range *nics.Items {
				firewalls, _, err := cloudAPIClient.FirewallRulesApi.DatacentersServersNicsFirewallrulesGet(context.TODO(), *datacenter.Id, *server.Id, *nic.Id).Depth(1).Execute()
				if err != nil {
					return err
				}
				if firewalls.Items == nil {
					log.Printf(
						"[WARNING] expected a response containing firewall rules but received 'nil' instead, skipping search for NIC with ID: %v, server ID: %v, datacenter ID: %v.\n",
						*nic.Id,
						*server.Id,
						*datacenter.Id)
					continue
				}
				lastFirewallIdx := len(*firewalls.Items) - 1
				for firewallIdx, firewall := range *firewalls.Items {
					// Skip the last firewall rule for the last NIC since this one will be added
					// to the server separately.
					if nicIdx == lastNicIdx && firewallIdx == lastFirewallIdx {
						continue
					}
					if firewall.Properties == nil || firewall.Properties.Name == nil {
						log.Printf(
							"[WARNING] 'nil' values in the response for the firewall rule with ID %v, NIC ID: %v, server ID: %v, datacenter ID: %v, skipping this resource.\n",
							*firewall.Id,
							*nic.Id,
							*server.Id,
							*datacenter.Id,
						)
						continue
					}
					g.Resources = append(g.Resources, terraformutils.NewResource(
						*firewall.Id,
						*firewall.Properties.Name+"-"+*firewall.Id,
						resourceType,
						helpers.Ionos,
						map[string]string{helpers.DcID: *datacenter.Id, helpers.ServerID: *server.Id, helpers.NicID: *nic.Id},
						[]string{},
						map[string]interface{}{}))
				}
			}
		}
	}
	return nil
}
