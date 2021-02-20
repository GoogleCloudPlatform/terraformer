package aws

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
)

type AppSyncGenerator struct {
	AWSService
}

func (g *AppSyncGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}

	svc := appsync.NewFromConfig(config)

	var nextToken *string
	for {
		apis, err := svc.ListGraphqlApis(context.TODO(), &appsync.ListGraphqlApisInput{
			NextToken: nextToken,
		})
		if err != nil {
			return err
		}

		for _, api := range apis.GraphqlApis {
			var id = *api.ApiId
			var name = *api.Name
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				id,
				name,
				"aws_appsync_graphql_api",
				"aws",
				[]string{}))
		}
		nextToken = apis.NextToken
		if nextToken == nil {
			break
		}
	}

	return nil
}
