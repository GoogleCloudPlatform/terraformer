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

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	adminpb "google.golang.org/genproto/googleapis/iam/admin/v1"

	"cloud.google.com/go/iam/admin/apiv1"
	"google.golang.org/api/iterator"
)

var IamAllowEmptyValues = []string{"tags."}

var IamAdditionalFields = map[string]string{}

type IamGenerator struct {
	GCPService
}

func (IamGenerator) createResources(serviceAccountsIterator *admin.ServiceAccountIterator) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	for {
		serviceAccount, err := serviceAccountsIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println("error with service account:", err)
			continue
		}
		resources = append(resources, terraform_utils.NewResource(
			serviceAccount.Name,
			serviceAccount.UniqueId,
			"google_service_account",
			"google",
			map[string]string{},
			IamAllowEmptyValues,
			IamAdditionalFields,
		))
	}
	return resources
}

// TODO ALL
func (g *IamGenerator) InitResources() error {
	ctx := context.Background()

	projectID := g.GetArgs()["project"].(string)
	client, err := admin.NewIamClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	serviceAccountsIterator := client.ListServiceAccounts(ctx, &adminpb.ListServiceAccountsRequest{Name: "projects/" + projectID})

	g.Resources = g.createResources(serviceAccountsIterator)
	g.PopulateIgnoreKeys()
	return nil

}
