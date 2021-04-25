package aws

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/service/swf"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
)

type SWFGenerator struct {
	AWSService
}

func (g *SWFGenerator) InitResources() error {
	regStatuses := []types.RegistrationStatus{types.RegistrationStatusRegistered, types.RegistrationStatusDeprecated}
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := swf.NewFromConfig(config)
	for _, status := range regStatuses {
		p := swf.NewListDomainsPaginator(svc, &swf.ListDomainsInput{RegistrationStatus: status})
		for p.HasMorePages() {
			page, err := p.NextPage(context.TODO())
			if err != nil {
				return err
			}
			for _, domain := range page.DomainInfos {
				g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
					*domain.Name,
					*domain.Name,
					"aws_swf_domain",
					"aws",
					[]string{},
				))
			}
		}
	}
	return nil
}
