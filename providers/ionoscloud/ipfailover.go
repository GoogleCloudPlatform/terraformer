package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	uuid "github.com/gofrs/uuid/v3"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
)

type IPFailoverGenerator struct {
	Service
}

func (g *IPFailoverGenerator) InitResources() error {
	client := g.generateClient()
	cloudAPIClient := client.CloudAPIClient
	datacenters, err := helpers.GetAllDatacenters(*cloudAPIClient)
	resourceType := "ionoscloud_ipfailover"
	if err != nil {
		return err
	}
	for _, datacenter := range datacenters {
		lans, _, err := cloudAPIClient.LANsApi.DatacentersLansGet(context.TODO(), *datacenter.Id).Depth(1).Execute()
		if err != nil {
			return err
		}
		if lans.Items == nil {
			log.Printf(
				"[WARNING] expected a response containing LANs but received 'nil' instead, skipping search for datacenter with ID: %v",
				*datacenter.Id)
			continue
		}
		for _, lan := range *lans.Items {
			if lan.Properties == nil || lan.Properties.IpFailover == nil {
				log.Printf(
					"[WARNING] 'nil' values in the response for LAN with ID %v, datacenter ID: %v, skipping this resource",
					*lan.Id,
					*datacenter.Id,
				)
				continue
			}
			for _, ipFailover := range *lan.Properties.IpFailover {
				// Generate the ID of the resource using the IP
				id := uuid.NewV5(uuid.NewV5(uuid.NamespaceURL, "https://github.com/ionos-cloud/terraform-provider-ionoscloud"), *ipFailover.Ip).String()
				g.Resources = append(g.Resources, terraformutils.NewResource(
					id,
					id,
					resourceType,
					helpers.Ionos,
					map[string]string{helpers.DcID: *datacenter.Id, "lan_id": *lan.Id, "ip": *ipFailover.Ip, "nicuuid": *ipFailover.NicUuid},
					[]string{},
					map[string]interface{}{}))
			}
		}
	}
	return nil
}
