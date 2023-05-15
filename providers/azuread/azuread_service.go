// Copyright 2019 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package azuread

import (
	"context"
	"fmt"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/manicminer/hamilton/auth"
	"github.com/manicminer/hamilton/environments"
	"github.com/manicminer/hamilton/msgraph"
)

type AzureADService struct { //nolint
	terraformutils.Service
}

type ServiceGenerator interface {
	terraformutils.ServiceGenerator
	GetResourceConnections() map[string][]string
}

func (az *AzureADService) getAuthorizer() (auth.Authorizer, error) {
	environment := environments.Global
	ctx := context.Background()
	tenantID := az.Args["tenant_id"].(string)
	clientID := az.Args["client_id"].(string)
	clientSecret := az.Args["client_secret"].(string)

	config := &auth.Config{
		Environment:            environment,
		TenantID:               tenantID,
		ClientID:               clientID,
		ClientSecret:           clientSecret,
		EnableClientSecretAuth: true,
	}
	authorizer, err := config.NewAuthorizer(ctx, config.Environment.MsGraph)
	if err != nil {
		fmt.Println(err.Error())
		log.Println(err.Error())
		return nil, err
	}
	return authorizer, nil
}

func (az *AzureADService) getUserClient() (*msgraph.UsersClient, error) {
	authorizer, err := az.getAuthorizer()
	if err != nil {
		fmt.Println(err.Error())
		log.Println(err.Error())
		return nil, err
	}

	tenantID := az.Args["tenant_id"].(string)
	client := msgraph.NewUsersClient(tenantID)
	client.BaseClient.Authorizer = authorizer

	return client, nil
}

func (az *AzureADService) getApplicationsClient() (*msgraph.ApplicationsClient, error) {
	authorizer, err := az.getAuthorizer()
	if err != nil {
		fmt.Println(err.Error())
		log.Println(err.Error())
		return nil, err
	}

	tenantID := az.Args["tenant_id"].(string)
	client := msgraph.NewApplicationsClient(tenantID)
	client.BaseClient.Authorizer = authorizer

	return client, nil
}

func (az *AzureADService) getGroupsClient() (*msgraph.GroupsClient, error) {
	authorizer, err := az.getAuthorizer()
	if err != nil {
		fmt.Println(err.Error())
		log.Println(err.Error())
		return nil, err
	}

	tenantID := az.Args["tenant_id"].(string)
	client := msgraph.NewGroupsClient(tenantID)
	client.BaseClient.Authorizer = authorizer

	return client, nil
}

func (az *AzureADService) getServicePrincipalsClient() (*msgraph.ServicePrincipalsClient, error) {
	authorizer, err := az.getAuthorizer()
	if err != nil {
		fmt.Println(err.Error())
		log.Println(err.Error())
		return nil, err
	}

	tenantID := az.Args["tenant_id"].(string)
	client := msgraph.NewServicePrincipalsClient(tenantID)
	client.BaseClient.Authorizer = authorizer

	return client, nil
}

func (az *AzureADService) getAppRoleAssignmentsClient() (*msgraph.AppRoleAssignedToClient, error) {
	authorizer, err := az.getAuthorizer()
	if err != nil {
		fmt.Println(err.Error())
		log.Println(err.Error())
		return nil, err
	}

	tenantID := az.Args["tenant_id"].(string)
	client := msgraph.NewAppRoleAssignedToClient(tenantID)
	client.BaseClient.Authorizer = authorizer

	return client, nil
}

func (az *AzureADService) GetResourceConnections() map[string][]string {
	return nil
}

func (az *AzureADService) appendSimpleResource(id string, resourceName string, resourceType string) {
	newResource := terraformutils.NewResource(id, resourceName, resourceType, az.ProviderName, map[string]string{
		"id": id,
	}, []string{}, map[string]interface{}{})
	az.Resources = append(az.Resources, newResource)
}
