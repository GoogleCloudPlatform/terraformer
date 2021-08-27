// Copyright 2020 The Terraformer Authors.
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

package gitlab

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/xanzy/go-gitlab"
)

type GroupGenerator struct {
	GitLabService
}

// Generate TerraformResources from gitlab API,
func (g *GroupGenerator) InitResources() error {
	ctx := context.Background()
	client, err := g.createClient()
	if err != nil {
		return err
	}

	group := g.Args["group"].(string)
	g.Resources = append(g.Resources, createGroups(ctx, client, group)...)

	return nil
}

func createGroups(ctx context.Context, client *gitlab.Client, groupID string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	group, _, err := client.Groups.GetGroup(groupID, gitlab.WithContext(ctx))
	if err != nil {
		log.Println(err)
		return nil
	}

	resource := terraformutils.NewSimpleResource(
		strconv.FormatInt(int64(group.ID), 10),
		getGroupResourceName(group),
		"gitlab_group",
		"gitlab",
		[]string{},
	)

	// NOTE: mirror fields from API doesn't match with the ones from terraform provider
	resource.IgnoreKeys = []string{"mirror_trigger_builds", "only_mirror_protected_branches", "mirror", "mirror_overwrites_diverged_branches"}

	resource.SlowQueryRequired = true
	resources = append(resources, resource)
	resources = append(resources, createGroupVariables(ctx, client, group)...)
	resources = append(resources, createGroupMembership(ctx, client, group)...)

	return resources
}
func createGroupVariables(ctx context.Context, client *gitlab.Client, group *gitlab.Group) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	opt := &gitlab.ListGroupVariablesOptions{}

	for {
		groupVariables, resp, err := client.GroupVariables.ListVariables(group.ID, opt, gitlab.WithContext(ctx))
		if err != nil {
			log.Println(err)
			return nil
		}

		for _, groupVariable := range groupVariables {

			resource := terraformutils.NewSimpleResource(
				fmt.Sprintf("%d:%s:%s", group.ID, groupVariable.Key, groupVariable.EnvironmentScope),
				fmt.Sprintf("%s___%s___%s", getGroupResourceName(group), groupVariable.Key, groupVariable.EnvironmentScope),
				"gitlab_group_variable",
				"gitlab",
				[]string{},
			)
			resource.SlowQueryRequired = true
			resources = append(resources, resource)
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return resources
}

func createGroupMembership(ctx context.Context, client *gitlab.Client, group *gitlab.Group) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	opt := &gitlab.ListGroupMembersOptions{}

	for {
		groupMembers, resp, err := client.Groups.ListGroupMembers(group.ID, opt, gitlab.WithContext(ctx))
		if err != nil {
			log.Println(err)
			return nil
		}

		for _, groupMember := range groupMembers {

			resource := terraformutils.NewSimpleResource(
				fmt.Sprintf("%d:%d", group.ID, groupMember.ID),
				fmt.Sprintf("%s___%s", getGroupResourceName(group), groupMember.Username),
				"gitlab_group_membership",
				"gitlab",
				[]string{},
			)
			resource.SlowQueryRequired = true
			resources = append(resources, resource)
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return resources
}

func getGroupResourceName(group *gitlab.Group) string {
	return fmt.Sprintf("%d___%s", group.ID, strings.ReplaceAll(group.FullPath, "/", "__"))
}
