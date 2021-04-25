// Copyright 2018 The Terraformer Authors.
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

package gcp

import (
	"context"
	"log"
	"regexp"

	admin "cloud.google.com/go/iam/admin/apiv1"
	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/iterator"
	adminpb "google.golang.org/genproto/googleapis/iam/admin/v1"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var IamAllowEmptyValues = []string{"tags."}

var IamAdditionalFields = map[string]interface{}{}

type IamGenerator struct {
	GCPService
}

func (g IamGenerator) createServiceAccountResources(serviceAccountsIterator *admin.ServiceAccountIterator) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	re := regexp.MustCompile(`^[a-z]`)
	for {
		serviceAccount, err := serviceAccountsIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println("error with service account:", err)
			continue
		}
		if !re.MatchString(serviceAccount.Email) {
			log.Printf("skipping %s: service account email must start with [a-z]\n", serviceAccount.Name)
			continue
		}
		resources = append(resources, terraformutils.NewSimpleResource(
			serviceAccount.Name,
			serviceAccount.UniqueId,
			"google_service_account",
			g.ProviderName,
			IamAllowEmptyValues,
		))
	}
	return resources
}

func (g *IamGenerator) createIamCustomRoleResources(rolesResponse *adminpb.ListRolesResponse, project string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, role := range rolesResponse.Roles {
		if role.Deleted {
			// Note: no need to log that the resource has been deleted
			continue
		}
		resources = append(resources, terraformutils.NewResource(
			role.Name,
			role.Name,
			"google_project_iam_custom_role",
			g.ProviderName,
			map[string]string{
				"role_id": role.Name,
				"project": project,
			},
			IamAllowEmptyValues,
			map[string]interface{}{
				"stage": role.Stage.String(),
			},
		))
	}

	return resources
}

func (g *IamGenerator) createIamMemberResources(policy *cloudresourcemanager.Policy, project string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, b := range policy.Bindings {
		for _, m := range b.Members {
			resources = append(resources, terraformutils.NewResource(
				b.Role+m,
				b.Role+m,
				"google_project_iam_member",
				g.ProviderName,
				map[string]string{
					"role":    b.Role,
					"project": project,
					"member":  m,
				},
				IamAllowEmptyValues,
				IamAdditionalFields,
			))
		}
	}

	return resources
}

func (g *IamGenerator) InitResources() error {
	ctx := context.Background()

	projectID := g.GetArgs()["project"].(string)
	client, err := admin.NewIamClient(ctx)
	if err != nil {
		return err
	}
	serviceAccountsIterator := client.ListServiceAccounts(ctx, &adminpb.ListServiceAccountsRequest{Name: "projects/" + projectID})
	rolesResponse, err := client.ListRoles(ctx, &adminpb.ListRolesRequest{Parent: "projects/" + projectID})
	if err != nil {
		return err
	}

	cm, err := cloudresourcemanager.NewService(context.Background())
	if err != nil {
		return err
	}
	rb := &cloudresourcemanager.GetIamPolicyRequest{}
	policyResponse, err := cm.Projects.GetIamPolicy(projectID, rb).Context(context.Background()).Do()
	if err != nil {
		return err
	}

	g.Resources = g.createServiceAccountResources(serviceAccountsIterator)
	g.Resources = append(g.Resources, g.createIamCustomRoleResources(rolesResponse, projectID)...)
	g.Resources = append(g.Resources, g.createIamMemberResources(policyResponse, projectID)...)
	return nil
}
