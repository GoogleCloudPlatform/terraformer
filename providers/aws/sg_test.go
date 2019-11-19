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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"reflect"
	"testing"
)

func TestEmptySgs(t *testing.T) {
	var securityGroups []*ec2.SecurityGroup

	rulesToMoveOut := findRulesToMoveOut(securityGroups)

	if rulesToMoveOut != nil {
		t.Errorf("failed to calculate rules to move out %v", rulesToMoveOut)
	}
}
func TestCycleReference(t *testing.T) {
	sgA := ec2.SecurityGroup{
		GroupId: aws.String("aaaa"),
		IpPermissions: []*ec2.IpPermission{
			{
				UserIdGroupPairs: []*ec2.UserIdGroupPair{
					{
						GroupId: aws.String("bbbb"),
					},
				},
			},
			{},
		},
	}
	securityGroups := []*ec2.SecurityGroup{
		&sgA,
		{
			GroupId: aws.String("bbbb"),
			IpPermissions: []*ec2.IpPermission{
				{
					UserIdGroupPairs: []*ec2.UserIdGroupPair{
						{
							GroupId: aws.String("aaaa"),
						},
					},
				},
			},
		},
	}

	rulesToMoveOut := findRulesToMoveOut(securityGroups)

	if !reflect.DeepEqual(rulesToMoveOut[0].sourceSG, &sgA) {
		t.Errorf("failed to calculate rules to move out %v", rulesToMoveOut)
	}
}
