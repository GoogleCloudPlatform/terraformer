package aws

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/batch"
)

var BatchAllowEmptyValues = []string{"tags."}

var BatchAdditionalFields = map[string]interface{}{}

type BatchGenerator struct {
	AWSService
}

func (g *BatchGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	batchClient := batch.NewFromConfig(config)

	if err := g.loadComputeEnvironments(batchClient); err != nil {
		return err
	}
	if err := g.loadJobDefinitions(batchClient); err != nil {
		return err
	}
	if err := g.loadJobQueues(batchClient); err != nil {
		return err
	}

	return nil
}

func (g *BatchGenerator) loadComputeEnvironments(batchClient *batch.Client) error {
	p := batch.NewDescribeComputeEnvironmentsPaginator(batchClient, &batch.DescribeComputeEnvironmentsInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, computeEnvironment := range page.ComputeEnvironments {
			computeEnvironmentName := StringValue(computeEnvironment.ComputeEnvironmentName)
			g.Resources = append(g.Resources, terraformutils.NewResource(
				computeEnvironmentName,
				computeEnvironmentName,
				"aws_batch_compute_environment",
				"aws",
				map[string]string{
					"compute_environment_name": computeEnvironmentName,
				},
				BatchAllowEmptyValues,
				BatchAdditionalFields,
			))
		}
	}
	return nil
}

func (g *BatchGenerator) loadJobDefinitions(batchClient *batch.Client) error {
	p := batch.NewDescribeJobDefinitionsPaginator(batchClient, &batch.DescribeJobDefinitionsInput{
		Status: aws.String("ACTIVE"),
	})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, jobDefinition := range page.JobDefinitions {
			jobDefinitionName := StringValue(jobDefinition.JobDefinitionName) + ":" + fmt.Sprint(jobDefinition.Revision)
			g.Resources = append(g.Resources, terraformutils.NewResource(
				jobDefinitionName,
				jobDefinitionName,
				"aws_batch_job_definition",
				"aws",
				map[string]string{
					"arn": StringValue(jobDefinition.JobDefinitionArn),
				},
				BatchAllowEmptyValues,
				BatchAdditionalFields,
			))
		}
	}
	return nil
}

func (g *BatchGenerator) loadJobQueues(batchClient *batch.Client) error {
	p := batch.NewDescribeJobQueuesPaginator(batchClient, &batch.DescribeJobQueuesInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, jobQueue := range page.JobQueues {
			jobQueueName := StringValue(jobQueue.JobQueueName)
			g.Resources = append(g.Resources, terraformutils.NewResource(
				jobQueueName,
				jobQueueName,
				"aws_batch_job_queue",
				"aws",
				map[string]string{},
				BatchAllowEmptyValues,
				BatchAdditionalFields,
			))
		}
	}
	return nil
}
