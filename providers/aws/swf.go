package aws

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/service/swf"
)

type SWFGenerator struct {
	AWSService
}

func (g *SWFGenerator) InitResources() error {
	regStatuses := []swf.RegistrationStatus{swf.RegistrationStatusRegistered, swf.RegistrationStatusDeprecated}
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := swf.New(config)
	for _, status := range regStatuses {
		p := swf.NewListDomainsPaginator(svc.ListDomainsRequest(&swf.ListDomainsInput{RegistrationStatus: status}))
		for p.Next(context.Background()) {
			for _, domain := range p.CurrentPage().DomainInfos {
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
