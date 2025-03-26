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

package aws

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
)

var identityStoreAllowEmptyValues = []string{"tags."}

type IdentityStoreGenerator struct {
	AWSService
}

func (g *IdentityStoreGenerator) GetIdentityStoreId() (*string, error) {
	config, e := g.generateConfig()
	if e != nil {
		return nil, e
	}
	svc := ssoadmin.NewFromConfig(config)
	instances, err := svc.ListInstances(context.TODO(), &ssoadmin.ListInstancesInput{})
	if err != nil {
		return nil, err
	}
	if len(instances.Instances) == 0 {
		return nil, nil
	}
	identityStoreId := StringValue(instances.Instances[0].IdentityStoreId)
	return &identityStoreId, nil

}

func (g *IdentityStoreGenerator) InitGroupResources(identityStoreId string) error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := identitystore.NewFromConfig(config)
	p := identitystore.NewListGroupsPaginator(svc, &identitystore.ListGroupsInput{
		IdentityStoreId: aws.String(identityStoreId),
	})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, group := range page.Groups {
			groupId := StringValue(group.GroupId)
			displayName := StringValue(group.DisplayName)
			g.Resources = append(g.Resources, terraformutils.NewResource(
				identityStoreId+"/"+groupId,
				displayName,
				"aws_identitystore_group",
				"aws",
				map[string]string{
					"identity_store_id": identityStoreId,
					"description":       StringValue(group.Description),
				},
				identityStoreAllowEmptyValues,
				map[string]interface{}{},
			))
			err = g.InitGroupMembershipResources(identityStoreId, groupId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (g *IdentityStoreGenerator) InitGroupMembershipResources(identityStoreId string, groupId string) error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := identitystore.NewFromConfig(config)
	p := identitystore.NewListGroupMembershipsPaginator(svc, &identitystore.ListGroupMembershipsInput{
		GroupId:         aws.String(groupId),
		IdentityStoreId: aws.String(identityStoreId),
	})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, user := range page.GroupMemberships {
			var memberId string
			switch v := user.MemberId.(type) {
			case *types.MemberIdMemberUserId:
				memberId = v.Value // Value is string
			case *types.UnknownUnionMember:
				memberId = v.Tag
			default:
				memberId = ""
			}
			membershipId := StringValue(user.MembershipId)
			g.Resources = append(g.Resources, terraformutils.NewResource(
				identityStoreId+"/"+membershipId,
				"m-"+groupId+"-"+memberId,
				"aws_identitystore_group_membership",
				"aws",
				map[string]string{
					"identity_store_id": identityStoreId,
					"group_id":          groupId,
					"member_id":         memberId,
				},
				identityStoreAllowEmptyValues,
				map[string]interface{}{},
			))
		}
	}
	return nil
}

func (g *IdentityStoreGenerator) InitUserResources(identityStoreId string) error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := identitystore.NewFromConfig(config)
	p := identitystore.NewListUsersPaginator(svc, &identitystore.ListUsersInput{
		IdentityStoreId: aws.String(identityStoreId),
	})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, user := range page.Users {
			userId := StringValue(user.UserId)
			displayName := StringValue(user.DisplayName)
			//			name := StringValue(user.Name)
			userName := StringValue(user.UserName)
			g.Resources = append(g.Resources, terraformutils.NewResource(
				identityStoreId+"/"+userId,
				userName,
				"aws_identitystore_user",
				"aws",
				map[string]string{
					"identity_store_id": identityStoreId,
					"display_name":      displayName,
					"use_name":          userName,
				},
				identityStoreAllowEmptyValues,
				map[string]interface{}{},
			))
		}
	}
	return nil
}

func (g *IdentityStoreGenerator) InitResources() error {
	identityStoreId, e := g.GetIdentityStoreId()
	if e != nil {
		return e
	}
	if identityStoreId == nil {
		return nil
	}

	e = g.InitUserResources(*identityStoreId)
	if e != nil {
		return e
	}

	e = g.InitGroupResources(*identityStoreId)
	if e != nil {
		return e
	}

	return nil
}
