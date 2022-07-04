package ionoscloud

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

type DatacenterGenerator struct {
	IonosCloudService
}

func (g DatacenterGenerator) createResources(datacentersList []ionoscloud.Datacenter) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, datacenter := range datacentersList {
		resources = append(resources, terraformutils.NewResource(
			*datacenter.Id,
			*datacenter.Properties.Name+"-"+*datacenter.Id,
			"ionoscloud_datacenter",
			"ionoscloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{}))
	}
	return resources
}

func (g *DatacenterGenerator) InitResources() error {
	client := g.generateClient()
	cloudApiClient := client.CloudApiClient
	output, _, err := cloudApiClient.DataCentersApi.DatacentersGet(context.TODO()).Depth(5).Execute()
	if err != nil {
		return err
	}
	g.Resources = g.createResources(*output.Items)
	return nil
}
