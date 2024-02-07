package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type S3KeyGenerator struct {
	Service
}

func (g *S3KeyGenerator) InitResources() error {
	client := g.generateClient()
	cloudAPIClient := client.CloudAPIClient
	resourceType := "ionoscloud_s3_key"

	usersResponse, _, err := cloudAPIClient.UserManagementApi.UmUsersGet(context.TODO()).Execute()
	if err != nil {
		return err
	}
	if usersResponse.Items == nil {
		log.Printf("[WARNING] expected a response containing users but received 'nil' instead")
		return nil
	}
	for _, user := range *usersResponse.Items {
		s3KeysResponse, _, err := cloudAPIClient.UserS3KeysApi.UmUsersS3keysGet(context.TODO(), *user.Id).Depth(1).Execute()
		if err != nil {
			return err
		}
		if s3KeysResponse.Items == nil {
			log.Printf("[WARNING] expected a response containing S3 keys but received 'nil' instead, skipping search for user with ID: %v.\n",
				*user.Id)
			continue
		}
		for _, s3Key := range *s3KeysResponse.Items {
			g.Resources = append(g.Resources, terraformutils.NewResource(
				*s3Key.Id,
				*s3Key.Id,
				resourceType,
				helpers.Ionos,
				map[string]string{helpers.UserID: *user.Id},
				[]string{},
				map[string]interface{}{}))
		}
	}
	return nil
}
