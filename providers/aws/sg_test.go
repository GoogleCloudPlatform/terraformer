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
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func TestEmptySgs(t *testing.T) {
	var securityGroups []types.SecurityGroup

	rulesToMoveOut := findSgsToMoveOut(securityGroups)

	if !reflect.DeepEqual(rulesToMoveOut, []string{}) {
		t.Errorf("failed to calculate rules to move out %v", rulesToMoveOut)
	}
}

func Test1CycleReference(t *testing.T) {
	sgA := types.SecurityGroup{
		GroupId: aws.String("aaaa"),
		IpPermissions: []types.IpPermission{
			{
				UserIdGroupPairs: []types.UserIdGroupPair{
					{
						GroupId: aws.String("aaaa"),
					},
				},
			},
			{},
		},
	}
	securityGroups := []types.SecurityGroup{
		sgA,
	}

	rulesToMoveOut := findSgsToMoveOut(securityGroups)

	if !reflect.DeepEqual(rulesToMoveOut, []string{}) {
		t.Errorf("failed to calculate rules to move out %v", rulesToMoveOut)
	}
}

func Test2CycleReference(t *testing.T) {
	sgA := types.SecurityGroup{
		GroupId: aws.String("aaaa"),
		IpPermissions: []types.IpPermission{
			{
				UserIdGroupPairs: []types.UserIdGroupPair{
					{
						GroupId: aws.String("bbbb"),
					},
				},
			},
		},
	}
	securityGroups := []types.SecurityGroup{
		{
			GroupId: aws.String("bbbb"),
			IpPermissions: []types.IpPermission{
				{
					UserIdGroupPairs: []types.UserIdGroupPair{
						{
							GroupId: aws.String("aaaa"),
						},
					},
				},
				{},
			},
		},
		sgA,
	}

	rulesToMoveOut := findSgsToMoveOut(securityGroups)

	if !reflect.DeepEqual(rulesToMoveOut[0], *sgA.GroupId) {
		t.Errorf("failed to calculate rules to move out %v", rulesToMoveOut)
	}
}

func TestNoCycleReference(t *testing.T) {
	sgA := types.SecurityGroup{
		GroupId: aws.String("aaaa"),
		IpPermissions: []types.IpPermission{
			{
				UserIdGroupPairs: []types.UserIdGroupPair{
					{
						GroupId: aws.String("bbbb"),
					},
				},
			},
		},
	}
	securityGroups := []types.SecurityGroup{
		{
			GroupId: aws.String("bbbb"),
			IpPermissions: []types.IpPermission{
				{},
				{},
			},
		},
		sgA,
	}

	rulesToMoveOut := findSgsToMoveOut(securityGroups)

	if len(rulesToMoveOut) != 0 {
		t.Errorf("failed to calculate rules to move out %v", rulesToMoveOut)
	}
}

func Test3Cycle1CycleReference(t *testing.T) {
	sgA := types.SecurityGroup{
		GroupId: aws.String("aaaa"),
		IpPermissions: []types.IpPermission{
			{
				UserIdGroupPairs: []types.UserIdGroupPair{
					{
						GroupId: aws.String("aaaa"),
					},
				},
			},
			{
				UserIdGroupPairs: []types.UserIdGroupPair{
					{
						GroupId: aws.String("bbbb"),
					},
				},
			},
		},
	}
	securityGroups := []types.SecurityGroup{
		sgA,
		{
			GroupId: aws.String("bbbb"),
			IpPermissions: []types.IpPermission{
				{
					UserIdGroupPairs: []types.UserIdGroupPair{
						{
							GroupId: aws.String("cccc"),
						},
					},
				},
				{},
			},
		},
		{
			GroupId: aws.String("cccc"),
			IpPermissions: []types.IpPermission{
				{
					UserIdGroupPairs: []types.UserIdGroupPair{
						{
							GroupId: aws.String("aaaa"),
						},
					},
				},
				{},
			},
		},
		{
			GroupId: aws.String("dddd"),
			IpPermissions: []types.IpPermission{
				{
					UserIdGroupPairs: []types.UserIdGroupPair{
						{
							GroupId: aws.String("aaaa"),
						},
					},
				},
				{},
			},
		},
	}

	rulesToMoveOut := findSgsToMoveOut(securityGroups)

	if !reflect.DeepEqual(rulesToMoveOut[0], *sgA.GroupId) {
		t.Errorf("failed to calculate rules to move out %v", rulesToMoveOut)
	}
}
