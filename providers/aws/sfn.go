package aws

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
)

var sfnAllowEmptyValues = []string{"tags."}

type SfnGenerator struct {
	AWSService
}

func (g *SfnGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := sfn.NewFromConfig(config)

	p := sfn.NewListStateMachinesPaginator(svc, &sfn.ListStateMachinesInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, stateMachine := range page.StateMachines {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				*stateMachine.StateMachineArn,
				*stateMachine.Name,
				"aws_sfn_state_machine",
				"aws",
				sfnAllowEmptyValues,
			))
		}
	}

	pActivity := sfn.NewListActivitiesPaginator(svc, &sfn.ListActivitiesInput{})
	for pActivity.HasMorePages() {
		pActivityNextPage, err := pActivity.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, stateMachine := range pActivityNextPage.Activities {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				*stateMachine.ActivityArn,
				*stateMachine.Name,
				"aws_sfn_activity",
				"aws",
				sfnAllowEmptyValues,
			))
		}
	}

	return nil
}
