package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type CertificateGenerator struct {
	Service
}

func (g *CertificateGenerator) InitResources() error {
	client := g.generateClient()
	cloudAPIClient := client.CertificateManagerAPIClient
	resourceType := "ionoscloud_certificate"

	response, _, err := cloudAPIClient.CertificatesApi.CertificatesGet(context.TODO()).Execute()
	if err != nil {
		return err
	}
	if response.Items == nil {
		log.Printf("[WARNING] 'nil' expected a response containing certificates but received 'nil' instead.")
		return nil
	}
	certificates := *response.Items
	for _, certificate := range certificates {
		if certificate.Properties == nil || certificate.Properties.Name == nil {
			log.Printf("[WARNING] 'nil' values in the response for the certificate with ID %v, skipping this resource.\n", *certificate.Id)
			continue
		}
		g.Resources = append(g.Resources, terraformutils.NewResource(
			*certificate.Id,
			*certificate.Properties.Name+"-"+*certificate.Id,
			resourceType,
			helpers.Ionos,
			map[string]string{},
			[]string{},
			map[string]interface{}{}))
	}
	return nil
}
