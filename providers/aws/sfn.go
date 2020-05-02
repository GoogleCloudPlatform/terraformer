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
	svc := sfn.New(config)

	p := sfn.NewListStateMachinesPaginator(svc.ListStateMachinesRequest(&sfn.ListStateMachinesInput{}))
	for p.Next(context.Background()) {
		for _, stateMachine := range p.CurrentPage().StateMachines {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				*stateMachine.StateMachineArn,
				*stateMachine.Name,
				"aws_sfn_state_machine",
				"aws",
				sfnAllowEmptyValues,
			))

			if err := p.Err(); err != nil {
				return err
			}
		}
	}

	pActivity := sfn.NewListActivitiesPaginator(svc.ListActivitiesRequest(&sfn.ListActivitiesInput{}))
	for pActivity.Next(context.Background()) {
		for _, stateMachine := range pActivity.CurrentPage().Activities {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				*stateMachine.ActivityArn,
				*stateMachine.Name,
				"aws_sfn_activity",
				"aws",
				sfnAllowEmptyValues,
			))

			if err := pActivity.Err(); err != nil {
				return err
			}
		}
	}

	return nil
}
