package aws

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

var CognitoAllowEmptyValues = []string{"tags."}

var CognitoAdditionalFields = map[string]interface{}{}

type CognitoGenerator struct {
	AWSService
}

const CognitoMaxResults = 60 // Required field for Cognito API

func (g *CognitoGenerator) loadIdentityPools(svc *cognitoidentity.Client) error {
	p := cognitoidentity.NewListIdentityPoolsPaginator(svc, &cognitoidentity.ListIdentityPoolsInput{
		MaxResults: aws.Int32(CognitoMaxResults),
	})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, pool := range page.IdentityPools {
			var id = *pool.IdentityPoolId
			var resourceName = *pool.IdentityPoolName
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				id,
				resourceName+"_"+id,
				"aws_cognito_identity_pool",
				"aws",
				[]string{}))
		}
	}

	return nil
}

func (g *CognitoGenerator) loadUserPools(svc *cognitoidentityprovider.Client) ([]string, error) {
	p := cognitoidentityprovider.NewListUserPoolsPaginator(svc, &cognitoidentityprovider.ListUserPoolsInput{
		MaxResults: aws.Int32(CognitoMaxResults),
	})

	var userPoolIds []string
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}
		for _, pool := range page.UserPools {
			id := *pool.Id
			resourceName := *pool.Name
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				id,
				resourceName+"_"+id,
				"aws_cognito_user_pool",
				"aws",
				[]string{}))

			userPoolIds = append(userPoolIds, *pool.Id)
		}
	}
	return userPoolIds, nil
}

func (g *CognitoGenerator) loadUserPoolClients(svc *cognitoidentityprovider.Client, userPoolIds []string) error {
	for _, userPoolID := range userPoolIds {
		p := cognitoidentityprovider.NewListUserPoolClientsPaginator(svc, &cognitoidentityprovider.ListUserPoolClientsInput{
			UserPoolId: aws.String(userPoolID),
			MaxResults: aws.Int32(CognitoMaxResults),
		})

		for p.HasMorePages() {
			page, err := p.NextPage(context.TODO())
			if err != nil {
				return err
			}
			for _, poolClient := range page.UserPoolClients {
				id := *poolClient.ClientId
				resourceName := *poolClient.ClientName
				g.Resources = append(g.Resources, terraformutils.NewResource(
					id,
					resourceName+"_"+id,
					"aws_cognito_user_pool_client",
					"aws",
					map[string]string{
						"user_pool_id": *poolClient.UserPoolId,
					},
					CognitoAllowEmptyValues,
					CognitoAdditionalFields))
			}
		}
	}
	return nil
}

func (g *CognitoGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}

	svcCognitoIdentity := cognitoidentity.NewFromConfig(config)
	if err := g.loadIdentityPools(svcCognitoIdentity); err != nil {
		return err
	}
	svcCognitoIdentityProvider := cognitoidentityprovider.NewFromConfig(config)

	userPoolIds, err := g.loadUserPools(svcCognitoIdentityProvider)
	if err != nil {
		return err
	}
	if err = g.loadUserPoolClients(svcCognitoIdentityProvider, userPoolIds); err != nil {
		return err
	}

	return nil
}

func (g *CognitoGenerator) PostConvertHook() error {
	for _, r := range g.Resources {
		if r.InstanceInfo.Type != "aws_cognito_user_pool" {
			continue
		}
		if _, ok := r.InstanceState.Attributes["admin_create_user_config.0.unused_account_validity_days"]; ok {
			if _, okpp := r.InstanceState.Attributes["admin_create_user_config.0.unused_account_validity_days"]; okpp {
				delete(r.Item["admin_create_user_config"].([]interface{})[0].(map[string]interface{}), "unused_account_validity_days")
			}
		}
		if _, ok := r.InstanceState.Attributes["sms_verification_message"]; ok {
			if _, oktmp := r.InstanceState.Attributes["verification_message_template.0.sms_message"]; oktmp {
				delete(r.Item, "sms_verification_message")
			}
		}
		if _, ok := r.InstanceState.Attributes["email_verification_message"]; ok {
			if _, oktmp := r.InstanceState.Attributes["verification_message_template.0.email_message"]; oktmp {
				delete(r.Item, "email_verification_message")
			}
		}
		if _, ok := r.InstanceState.Attributes["email_verification_subject"]; ok {
			if _, oktmp := r.InstanceState.Attributes["verification_message_template.0.email_subject"]; oktmp {
				delete(r.Item, "email_verification_subject")
			}
		}
	}
	return nil
}
