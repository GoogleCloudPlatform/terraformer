package gcp

import (
	"fmt"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"golang.org/x/net/context"
	"google.golang.org/api/cloudresourcemanager/v1"
)

type OrganizationPolicyGenerator struct {
	GCPService
}

// Run on routersList and create for each TerraformResource
func (g OrganizationPolicyGenerator) createResources(ctx context.Context, policiesList *cloudresourcemanager.OrganizationsListOrgPoliciesCall) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	if err := policiesList.Pages(ctx, func(page *cloudresourcemanager.ListOrgPoliciesResponse) error {
		for _, obj := range page.Policies {
			resources = append(resources, terraform_utils.NewResource(
				fmt.Sprintf("%s/%s", g.GetArgs()["organization"].(string), obj.Constraint),
				fmt.Sprintf("%s-%s", g.GetArgs()["organization"].(string), obj.Constraint),
				"google_organization_policy",
				"google",
				map[string]string{
					"org_id": g.GetArgs()["organization"].(string),
				},
				[]string{},
				map[string]interface{}{},
			))
		}
		return nil
	}); err != nil {
		log.Println(err)
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each routers create 1 TerraformResource
// Need routers name as ID for terraform resource
func (g *OrganizationPolicyGenerator) InitResources() error {
	ctx := context.Background()
	service, err := cloudresourcemanager.NewService(ctx)
	if err != nil {
		return err
	}
	resource := "organizations/" + g.GetArgs()["organization"].(string)
	policiesList := service.Organizations.ListOrgPolicies(resource, &cloudresourcemanager.ListOrgPoliciesRequest{
		PageSize: 100,
	})
	g.Resources = g.createResources(ctx, policiesList)
	return nil
}
