package aws

import (
	"context"
	"github.com/zclconf/go-cty/cty"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type CognitoGenerator struct {
	AWSService
}

const CognitoMaxResults = 60 // Required field for Cognito API

func (g *CognitoGenerator) loadIdentityPools(svc *cognitoidentity.Client) error {
	p := cognitoidentity.NewListIdentityPoolsPaginator(svc, &cognitoidentity.ListIdentityPoolsInput{
		MaxResults: *aws.Int32(CognitoMaxResults),
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
				resourceName,
				"aws_cognito_identity_pool",
				"aws",
				[]string{}))
		}
	}

	return nil
}

func (g *CognitoGenerator) loadUserPools(svc *cognitoidentityprovider.Client) error {
	p := cognitoidentityprovider.NewListUserPoolsPaginator(svc, &cognitoidentityprovider.ListUserPoolsInput{
		MaxResults: *aws.Int32(CognitoMaxResults),
	})

	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, pool := range page.UserPools {
			id := *pool.Id
			resourceName := *pool.Name
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				id,
				resourceName,
				"aws_cognito_user_pool",
				"aws",
				[]string{}))
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
	if err := g.loadUserPools(svcCognitoIdentityProvider); err != nil {
		return err
	}

	return nil
}

func (g *CognitoGenerator) PostConvertHook() error {
	for _, r := range g.Resources {
		if r.Address.Type != "aws_cognito_user_pool" {
			continue
		}
		instanceStateMap := r.InstanceState.Value.AsValueMap()
		if r.InstanceState.Value.HasIndex(cty.StringVal("sms_verification_message")) == cty.True {
			if r.InstanceState.Value.GetAttr("verification_message_template").AsValueSlice()[0].HasIndex(cty.StringVal("sms_message")) == cty.True {
				delete(instanceStateMap, "sms_verification_message")
			}
		}
		if r.InstanceState.Value.HasIndex(cty.StringVal("email_verification_message")) == cty.True {
			if r.InstanceState.Value.GetAttr("verification_message_template").AsValueSlice()[0].HasIndex(cty.StringVal("email_message")) == cty.True {
				delete(instanceStateMap, "email_verification_message")
			}
		}
		if r.InstanceState.Value.HasIndex(cty.StringVal("email_verification_subject")) == cty.True {
			if r.InstanceState.Value.GetAttr("verification_message_template").AsValueSlice()[0].HasIndex(cty.StringVal("email_subject")) == cty.True {
				delete(instanceStateMap, "email_verification_subject")
			}
		}
		r.InstanceState.Value = cty.ObjectVal(instanceStateMap)
	}
	return nil
}
