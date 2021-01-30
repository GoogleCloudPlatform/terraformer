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
)

type IAMGenerator struct {
	IBMService
}

func (g IAMGenerator) loadUserPolicies(policyID string, user string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		fmt.Sprintf("%s/%s", user, policyID),
		policyID,
		"ibm_iam_user_policy",
		"ibm",
		[]string{})
	return resources
}

func (g IAMGenerator) loadAccessGroups(grpID string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		grpID,
		grpID,
		"ibm_iam_access_group",
		"ibm",
		[]string{})
	return resources
}

func (g IAMGenerator) loadAccessGroupMembers(grpID string, dependsOn []string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
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
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
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
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
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
			return fmt.Errorf("Error retrieving access group policy: %s", err)
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

	return nil
}
