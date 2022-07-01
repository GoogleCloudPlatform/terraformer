package helpers

import (
	"context"
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

func GetAllDatacenters(client ionoscloud.APIClient) ([]ionoscloud.Datacenter, error) {
	datacenters, _, err := client.DataCentersApi.DatacentersGet(context.TODO()).Depth(10).Execute()
	if err != nil {
		return nil, err
	}
	return *datacenters.Items, nil
}
