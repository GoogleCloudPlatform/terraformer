package aws

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type CognitoGenerator struct {
	AWSService
}

func (g *CognitoGenerator) loadIdentityPools(svc *cognitoidentity.Client) error {
	// TODO:
	// * Cognito does not provides a `NewIdentityPoolsPaginator` like other services' "NewListClustersPaginator"
	// ? Replace if it being avaialble, 60 being the max under constraint currently
	const maxIdentityPool = 60
	pools, err := svc.ListIdentityPoolsRequest(&cognitoidentity.ListIdentityPoolsInput{MaxResults: aws.Int64(maxIdentityPool)}).Send(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}

	for _, pool := range pools.IdentityPools {
		var id = *pool.IdentityPoolId
		var resourceName = *pool.IdentityPoolName
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			id,
			resourceName,
			"aws_cognito_identity_pool",
			"aws",
			[]string{}))
	}

	return nil
}

func (g *CognitoGenerator) loadUserPools(svc *cognitoidentityprovider.Client) error {
	// ? Replace if it being avaialble, 60 being the max under constraint currently
	const maxUserPool = 60
	req := svc.ListUserPoolsRequest(&cognitoidentityprovider.ListUserPoolsInput{MaxResults: aws.Int64(maxUserPool)})
	p := cognitoidentityprovider.NewListUserPoolsPaginator(req)

	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, pool := range page.UserPools {
			id := *pool.Id
			resourceName := *pool.Name
			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				id,
				resourceName,
				"aws_cognito_user_pool",
				"aws",
				[]string{}))
		}
	}

	if err := p.Err(); err != nil {
		log.Println(p.Err())
		return err
	}
	return nil
}

func (g *CognitoGenerator) InitResources() error {

	config, e := g.generateConfig()
	if e != nil {
		return e
	}

	svcCognitoIdentity := cognitoidentity.New(config)
	if err := g.loadIdentityPools(svcCognitoIdentity); err != nil {
		return err
	}
	svcCognitoIdentityProvider := cognitoidentityprovider.New(config)
	if err := g.loadUserPools(svcCognitoIdentityProvider); err != nil {
		return err
	}

	return nil
}
