package ionoscloud

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

type LanGenerator struct {
	IonosCloudService
}

func (g LanGenerator) createResources(lansList []ionoscloud.Lan) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, lan := range lansList {
		resources = append(resources, terraformutils.NewSimpleResource(
			*lan.Id,
			*lan.Properties.Name+"-"+*lan.Id,
			"ionoscloud_lan",
			"ionoscloud",
			[]string{}))
	}
	return resources
}

func (g *LanGenerator) InitResources() error {
	var lansOuput []ionoscloud.Lan
	client := g.generateClient()
	cloudApiClient := client.CloudApiClient
	datacenters, err := helpers.GetAllDatacenters(*cloudApiClient)
	if err != nil {
		return err
	}
	if datacenters != nil {
		for _, datacenter := range datacenters {
			lans, _, err := cloudApiClient.LANsApi.DatacentersLansGet(context.TODO(), *datacenter.Id).Depth(10).Execute()
			if err != nil {
				return err
			}
			lansToAdd := *lans.Items
			lansOuput = append(lansOuput, lansToAdd...)
		}
	}
	g.Resources = g.createResources(lansOuput)
	return nil
}
