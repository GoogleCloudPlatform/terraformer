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
	svc := xray.NewFromConfig(config)

	p := xray.NewGetSamplingRulesPaginator(svc, &xray.GetSamplingRulesInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, samplingRule := range page.SamplingRuleRecords {
			// NOTE: Builtin rule with unmodifiable name and 10000 prirority (lowest)
			if *samplingRule.SamplingRule.RuleName != "Default" {
				g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
					*samplingRule.SamplingRule.RuleName,
					*samplingRule.SamplingRule.RuleName,
					"aws_xray_sampling_rule",
					"aws",
					xrayAllowEmptyValues))
			}
		}
	}

	return nil
}
