package confluent

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	v2 "github.com/confluentinc/ccloud-sdk-go-v2/iam/v2"
)

type ServiceAccountGenerator struct {
	ConfluentService
}

func (g ServiceAccountGenerator) createResources(sas []v2.IamV2ServiceAccount) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, sa := range sas {
		resources = append(resources, terraformutils.NewSimpleResource(
			sa.GetId(),
			sa.GetDisplayName(),
			"confluent_service_account",
			"confluent",
			[]string{}))
	}
	return resources
}
