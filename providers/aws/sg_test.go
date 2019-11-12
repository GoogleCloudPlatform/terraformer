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
	"github.com/aws/aws-sdk-go/service/ec2"
	"reflect"
	"testing"
)

func TestCycleReference(t *testing.T) {
	securityGroups := []*ec2.SecurityGroup{

	}

	spanningTree := findSpanningTree(securityGroups)


	if !reflect.DeepEqual(spanningTree, map[string]interface{}{

	}) {
		t.Errorf("failed to calculate spanning tree %v", spanningTree)
	}
}