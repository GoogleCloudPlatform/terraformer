package aws

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/service/xray"
)

var xrayAllowEmptyValues = []string{"tags."}

type XrayGenerator struct {
	AWSService
}

func (g *XrayGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := xray.New(config)

	p := xray.NewGetSamplingRulesPaginator(svc.GetSamplingRulesRequest(&xray.GetSamplingRulesInput{}))
	for p.Next(context.Background()) {
		for _, samplingRule := range p.CurrentPage().SamplingRuleRecords {
			// NOTE: Builtin rule with unmodifiable name and 10000 prirority (lowest)
			if *samplingRule.SamplingRule.RuleName != "Default" {
				g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
					*samplingRule.SamplingRule.RuleName,
					*samplingRule.SamplingRule.RuleName,
					"aws_xray_sampling_rule",
					"aws",
					xrayAllowEmptyValues))

				if err := p.Err(); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
