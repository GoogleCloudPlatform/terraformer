package ionoscloud

import (
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

type DatacenterGenerator struct {
	Service
}

func (g DatacenterGenerator) createResources(datacentersList []ionoscloud.Datacenter) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, datacenter := range datacentersList {
		if datacenter.Properties == nil || datacenter.Properties.Name == nil {
			log.Printf("[WARNING] 'nil' values in the response for Datacenter with ID %v, skipping this resource.\n", *datacenter.Id)
			continue
		}
		resources = append(resources, terraformutils.NewResource(
			*datacenter.Id,
			*datacenter.Properties.Name+"-"+*datacenter.Id,
			"ionoscloud_datacenter",
			helpers.Ionos,
			map[string]string{},
			[]string{},
			map[string]interface{}{}))
	}
	return resources
}

func (g *DatacenterGenerator) InitResources() error {
	client := g.generateClient()
	cloudAPIClient := client.CloudAPIClient
	output, err := helpers.GetAllDatacenters(*cloudAPIClient)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
