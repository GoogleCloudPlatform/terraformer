package aws

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
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

	svc := appsync.New(config)

	// TODO:
	// * Service does not provides a `NewXXXXXXXXXXXPaginator` like other services' "NewListClustersPaginator" right now
	apis, err := svc.ListGraphqlApisRequest(&appsync.ListGraphqlApisInput{}).Send(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}

	var resources []terraform_utils.Resource
	for _, api := range apis.GraphqlApis {
		var id = *api.ApiId
		var name = *api.Name
		resources = append(resources, terraform_utils.NewSimpleResource(
			id,
			name,
			"aws_appsync_graphql_api",
			"aws",
			[]string{}))
	}

	g.Resources = resources
	return nil
}
