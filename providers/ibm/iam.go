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

package ibm

import (
	"fmt"
	"os"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv1"
	"github.com/IBM-Cloud/bluemix-go/api/iamuum/iamuumv2"
	"github.com/IBM-Cloud/bluemix-go/api/usermanagement/usermanagementv2"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
)

type IAMGenerator struct {
	IBMService
}

func (g IAMGenerator) loadUserPolicies(policyID string, user string) terraformutils.Resource {
	resources := terraformutils.NewSimpleResource(
		fmt.Sprintf("%s/%s", user, policyID),
		policyID,
		"ibm_iam_user_policy",
		"ibm",
		[]string{})
	return resources
}

func (g IAMGenerator) loadAccessGroups(grpID string) terraformutils.Resource {
	resources := terraformutils.NewSimpleResource(
		grpID,
		grpID,
		"ibm_iam_access_group",
		"ibm",
		[]string{})
	return resources
}

func (g IAMGenerator) loadServiceIDs(serviceID string) terraformutils.Resource {
	resources := terraformutils.NewSimpleResource(
		serviceID,
		serviceID,
		"ibm_iam_service_id",
		"ibm",
		[]string{})
	return resources
}

func (g IAMGenerator) loadAuthPolicies(policyID string) terraformutils.Resource {
	resources := terraformutils.NewSimpleResource(
		policyID,
		policyID,
		"ibm_iam_authorization_policy",
		"ibm",
		[]string{})
	return resources
}

func (g IAMGenerator) loadCustomRoles(roleID string) terraformutils.Resource {
	resources := terraformutils.NewSimpleResource(
		roleID,
		roleID,
		"ibm_iam_custom_role",
		"ibm",
		[]string{})
	return resources
}

func (g IAMGenerator) loadServicePolicies(serviceID, policyID string, dependsOn []string) terraformutils.Resource {
	resources := terraformutils.NewResource(
		fmt.Sprintf("%s/%s", serviceID, policyID),
		policyID,
		"ibm_iam_service_policy",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

func (g IAMGenerator) loadAccessGroupMembers(grpID string, dependsOn []string) terraformutils.Resource {
	resources := terraformutils.NewResource(
		fmt.Sprintf("%s/%s", grpID, grpID),
		grpID,
		"ibm_iam_access_group_members",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

func (g IAMGenerator) loadAccessGroupPolicies(grpID, policyID string, dependsOn []string) terraformutils.Resource {
	resources := terraformutils.NewResource(
		fmt.Sprintf("%s/%s", grpID, policyID),
		policyID,
		"ibm_iam_access_group_policy",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

func (g IAMGenerator) loadAccessGroupDynamicPolicies(grpID, ruleID string, dependsOn []string) terraformutils.Resource {
	resources := terraformutils.NewResource(
		fmt.Sprintf("%s/%s", grpID, ruleID),
		ruleID,
		"ibm_iam_access_group_dynamic_rule",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

func (g *IAMGenerator) InitResources() error {
	bmxConfig := &bluemix.Config{
		BluemixAPIKey: os.Getenv("IC_API_KEY"),
	}
	sess, err := session.New(bmxConfig)
	if err != nil {
		return err
	}

	userManagementAPI, err := usermanagementv2.New(sess)
	if err != nil {
		return err
	}
	err = authenticateAPIKey(sess)
	if err != nil {
		return err
	}
	generation := envFallBack([]string{"Generation"}, "2")
	gen, err := strconv.Atoi(generation)
	if err != nil {
		return err
	}
	userInfo, err := fetchUserDetails(sess, gen)
	if err != nil {
		return err
	}
	accountID := userInfo.userAccount

	users, err := userManagementAPI.UserInvite().GetUsers(userInfo.userAccount)
	if err != nil {
		return err
	}
	iampap, err := iampapv1.New(sess)
	if err != nil {
		return err
	}

	for _, u := range users.Resources {
		// User policies
		policies, err := iampap.V1Policy().List(iampapv1.SearchParams{
			AccountID: accountID,
			IAMID:     u.IamID,
			Type:      iampapv1.AccessPolicyType,
		})
		if err != nil {
			return err
		}
		for _, p := range policies {
			g.Resources = append(g.Resources, g.loadUserPolicies(p.ID, u.Email))
		}
	}

	iamuumClient, err := iamuumv2.New(sess)
	if err != nil {
		return err
	}

	agrps, err := iamuumClient.AccessGroup().List(accountID)
	if err != nil {
		return err
	}
	for _, group := range agrps {
		g.Resources = append(g.Resources, g.loadAccessGroups(group.ID))
		var dependsOn []string
		dependsOn = append(dependsOn,
			"ibm_iam_access_group."+terraformutils.TfSanitize(group.ID))
		g.Resources = append(g.Resources, g.loadAccessGroupMembers(group.ID, dependsOn))

		policies, err := iampap.V1Policy().List(iampapv1.SearchParams{
			AccountID:     accountID,
			AccessGroupID: group.ID,
			Type:          iampapv1.AccessPolicyType,
		})
		if err != nil {
			return fmt.Errorf("error retrieving access group policy: %s", err)
		}
		for _, p := range policies {
			g.Resources = append(g.Resources, g.loadAccessGroupPolicies(group.ID, p.ID, dependsOn))
		}

		dynamicPolicies, err := iamuumClient.DynamicRule().List(group.ID)
		if err != nil {
			return err
		}
		for _, d := range dynamicPolicies {
			g.Resources = append(g.Resources, g.loadAccessGroupDynamicPolicies(group.ID, d.RuleID, dependsOn))
		}
	}

	// service id and service policy
	apiKey := os.Getenv("IC_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("no API key set")
	}

	iamIDurl := "https://iam.cloud.ibm.com"
	iamOptions := &iamidentityv1.IamIdentityV1Options{
		URL: envFallBack([]string{"IBMCLOUD_IAM_API_ENDPOINT"}, iamIDurl),
		Authenticator: &core.IamAuthenticator{
			ApiKey: apiKey,
		},
	}

	iamPolicyOptions := &iampolicymanagementv1.IamPolicyManagementV1Options{
		URL: envFallBack([]string{"IBMCLOUD_IAM_API_ENDPOINT"}, iamIDurl),
		Authenticator: &core.IamAuthenticator{
			ApiKey: apiKey,
		},
	}

	iamIDClient, err := iamidentityv1.NewIamIdentityV1(iamOptions)
	if err != nil {
		return err
	}

	iamPolicyClient, err := iampolicymanagementv1.NewIamPolicyManagementV1(iamPolicyOptions)
	if err != nil {
		return err
	}

	start := ""
	allrecs := []iamidentityv1.ServiceID{}
	var pg int64 = 100

	for {
		listServiceIDOptions := iamidentityv1.ListServiceIdsOptions{
			AccountID: &accountID,
			Pagesize:  &pg,
		}
		if start != "" {
			listServiceIDOptions.Pagetoken = &start
		}

		serviceIDs, resp, err := iamIDClient.ListServiceIds(&listServiceIDOptions)
		if err != nil {
			return fmt.Errorf("[ERROR] Error listing Service Ids %s %s", err, resp)
		}
		start = GetNextIAM(serviceIDs.Next)
		allrecs = append(allrecs, serviceIDs.Serviceids...)
		if start == "" {
			break
		}
	}

	// loop through all service IDs and fetch policies correspponds to each service ID
	for _, service := range allrecs {
		g.Resources = append(g.Resources, g.loadServiceIDs(*service.ID))
		var dependsOn []string
		dependsOn = append(dependsOn,
			"ibm_iam_service_id."+terraformutils.TfSanitize(*service.ID))

		listServicePolicyOptions := iampolicymanagementv1.ListPoliciesOptions{
			AccountID: core.StringPtr(accountID),
			IamID:     core.StringPtr(*service.IamID),
			Type:      core.StringPtr("access"),
		}

		policyList, _, err := iamPolicyClient.ListPolicies(&listServicePolicyOptions)
		policies := policyList.Policies

		if err != nil {
			return fmt.Errorf("error retrieving service policy: %s", err)
		}

		for _, p := range policies {
			g.Resources = append(g.Resources, g.loadServicePolicies(*service.ID, *p.ID, dependsOn))
		}
	}

	// Authorization policy
	listAuthPolicyOptions := iampolicymanagementv1.ListPoliciesOptions{
		AccountID: core.StringPtr(accountID),
		Type:      core.StringPtr("authorization"),
	}

	authPolicyList, _, err := iamPolicyClient.ListPolicies(&listAuthPolicyOptions)
	authPolicies := authPolicyList.Policies

	if err != nil {
		return fmt.Errorf("error retrieving authorization policy: %s", err)
	}

	for _, ap := range authPolicies {
		g.Resources = append(g.Resources, g.loadAuthPolicies(*ap.ID))
	}

	// Custom role
	listCustomRoleOptions := iampolicymanagementv1.ListRolesOptions{
		AccountID: core.StringPtr(accountID),
	}

	rolesList, _, err := iamPolicyClient.ListRoles(&listCustomRoleOptions)
	customRoles := rolesList.CustomRoles

	if err != nil {
		return fmt.Errorf("error retrieving custom roles: %s", err)
	}

	for _, r := range customRoles {
		g.Resources = append(g.Resources, g.loadCustomRoles(*r.ID))
	}

	return nil
}
