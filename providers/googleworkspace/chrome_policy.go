package googleworkspace

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"google.golang.org/api/chromepolicy/v1"
)

type ChromePolicyGenerator struct {
	GoogleWorkspaceService
}

func (g *ChromePolicyGenerator) InitResources() error {
	client, err := g.ChromePolicyClient()
	if err != nil {
		return err
	}

	var policySchemaList []*chromepolicy.GoogleChromePolicyV1PolicySchema
	policySchemaListResponse, err := client.Customers.PolicySchemas.List("customers/" + g.orgID).Do()
	if err != nil {
		return err
	}
	policySchemaList = append(policySchemaList, policySchemaListResponse.PolicySchemas...)
	for {
		if policySchemaListResponse.NextPageToken == "" {
			break
		}
		policySchemaListResponse, err = client.Customers.PolicySchemas.List("customers/" + g.orgID).PageToken(policySchemaListResponse.NextPageToken).Do()
		if err != nil {
			return err
		}
		policySchemaList = append(policySchemaList, policySchemaListResponse.PolicySchemas...)

	}

	policySchemaLeafs := map[string]*struct{}{}
	found := &struct{}{}
	for _, policySchema := range policySchemaList {
		splitSchema := strings.Split(policySchema.SchemaName, ".")
		splitSchema = splitSchema[:len(splitSchema)-1]
		schemaLeaf := strings.Join(splitSchema, ".")
		if _, alreadyFound := policySchemaLeafs[schemaLeaf]; !alreadyFound {
			policySchemaLeafs[schemaLeaf] = found
		}
	}

	orgUnitList, err := g.getAllOrgUnits()
	if err != nil {
		return err
	}

	for _, orgUnit := range orgUnitList {
		log.Println("Loading " + orgUnit.OrgUnitPath)
		orgUnitPolicies := []*chromepolicy.GoogleChromePolicyV1ResolvedPolicy{}
		for policySchemaLeaf, _ := range policySchemaLeafs {
			policyTargetOrgUnit := &chromepolicy.GoogleChromePolicyV1PolicyTargetKey{
				TargetResource: "orgunits/" + strings.Split(orgUnit.OrgUnitId, ":")[1],
			}
			err := retryTimeDuration(context.Background(), time.Minute, func() error {
				return client.Customers.Policies.Resolve("customers/"+g.orgID, &chromepolicy.GoogleChromePolicyV1ResolveRequest{
					PolicySchemaFilter: policySchemaLeaf + ".*",
					PolicyTargetKey:    policyTargetOrgUnit,
				}).Pages(context.Background(), func(chromePolicyResponse *chromepolicy.GoogleChromePolicyV1ResolveResponse) error {
					orgUnitPolicies = append(orgUnitPolicies, chromePolicyResponse.ResolvedPolicies...)
					return nil
				})
			})
			if err != nil {
				return err
			}
		}
		resource := g.createOrgUnitResource(strings.Split(orgUnit.OrgUnitId, ":")[1], orgUnit.Name, orgUnitPolicies)
		if resource != nil {
			g.Resources = append(g.Resources, *resource)
		}
	}
	return nil
}

type chromePolicySchema struct {
	SchemaName   string                 `json:"schema_name,omitempty"`
	SchemaValues map[string]interface{} `json:"schema_values,omitempty"`
}

func (g ChromePolicyGenerator) createOrgUnitResource(orgUnitID string, orgUnitName string, chromePolicies []*chromepolicy.GoogleChromePolicyV1ResolvedPolicy) *terraformutils.Resource {
	if len(chromePolicies) == 0 {
		return nil
	}

	var policySchemas []chromePolicySchema
	for _, chromePolicy := range chromePolicies {

		// Skipping inherited policies
		if chromePolicy.SourceKey.TargetResource != chromePolicy.TargetKey.TargetResource {
			continue
		}

		chromePolicyDefinition := chromePolicySchema{
			SchemaName: chromePolicy.Value.PolicySchema,
		}

		valueBytes, err := chromePolicy.Value.Value.MarshalJSON()
		if err != nil {
			log.Fatal("Failed to marshal Chrome Policy Value definition", err)
			return nil
		}

		err = json.Unmarshal(valueBytes, &chromePolicyDefinition.SchemaValues)
		if err != nil {
			log.Fatal("Failed to Unmarshal Chrome Policy Value definition", err)
			return nil
		}

		policySchemas = append(policySchemas, chromePolicyDefinition)
	}

	if len(policySchemas) == 0 {
		return nil
	}

	resourceName := g.EnsureStringRandomness("chrome_policy_" + orgUnitName)
	resource := terraformutils.NewResource(
		orgUnitID,
		resourceName,
		"googleworkspace_chrome_policy",
		"googleworkspace",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"org_unit_id": orgUnitID,
			"policies":    policySchemas,
		},
	)

	return &resource
}
