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
	"log"
	"math/rand"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

type InstanceGroupGenerator struct {
	IBMService
	importDone  chan bool
	fatalErrors chan error
}

func (g *InstanceGroupGenerator) loadInstanceGroup(instanceGroupID, instanceGroupName string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		instanceGroupID,
		instanceGroupName,
		"ibm_is_instance_group",
		"ibm",
		[]string{})
	return resources
}

func (g *InstanceGroupGenerator) loadInstanceGroupManger(instanceGroupID, instanceGroupManagerID, managerName string) terraformutils.Resource {
	if managerName == "" {
		managerName = fmt.Sprintf("manager-%d-%d", rand.Intn(100), rand.Intn(50))
	}
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		fmt.Sprintf("%s/%s", instanceGroupID, instanceGroupManagerID),
		managerName,
		"ibm_is_instance_group_manager",
		"ibm",
		[]string{})
	return resources
}

func (g *InstanceGroupGenerator) loadInstanceGroupMangerPolicy(instanceGroupID, instanceGroupManagerID, PolicyID, policyName string) terraformutils.Resource {
	if policyName == "" {
		policyName = fmt.Sprintf("manager-%d-%d", rand.Intn(100), rand.Intn(50))
	}
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		fmt.Sprintf("%s/%s/%s", instanceGroupID, instanceGroupManagerID, PolicyID),
		policyName,
		"ibm_is_instance_group_manager_policy",
		"ibm",
		[]string{})
	return resources
}

func (g *InstanceGroupGenerator) handlePolicies(sess *vpcv1.VpcV1, instanceGroupID, instanceGroupManagerID string, policies []string) {
	for _, instanceGroupManagerPolicyID := range policies {
		getInstanceGroupManagerPolicyOptions := vpcv1.GetInstanceGroupManagerPolicyOptions{
			ID:                     &instanceGroupManagerPolicyID,
			InstanceGroupID:        &instanceGroupID,
			InstanceGroupManagerID: &instanceGroupManagerID,
		}
		data, response, err := sess.GetInstanceGroupManagerPolicy(&getInstanceGroupManagerPolicyOptions)
		if err != nil {
			g.fatalErrors <- fmt.Errorf("Error Getting InstanceGroup Manager Policy: %s\n%s", err, response)
		}
		instanceGroupManagerPolicy := data.(*vpcv1.InstanceGroupManagerPolicy)
		g.Resources = append(g.Resources, g.loadInstanceGroupMangerPolicy(instanceGroupID,
			instanceGroupManagerID,
			instanceGroupManagerPolicyID,
			*instanceGroupManagerPolicy.Name))
	}
}

func (g *InstanceGroupGenerator) handleManagers(sess *vpcv1.VpcV1, instanceGroupID string, managers []string) {
	for _, instanceGroupManagerID := range managers {
		getInstanceGroupManagerOptions := vpcv1.GetInstanceGroupManagerOptions{
			ID:              &instanceGroupManagerID,
			InstanceGroupID: &instanceGroupID,
		}
		instanceGroupManager, response, err := sess.GetInstanceGroupManager(&getInstanceGroupManagerOptions)
		if err != nil {
			g.fatalErrors <- fmt.Errorf("Error Getting InstanceGroup Manager: %s\n%s", err, response)
		}
		g.Resources = append(g.Resources, g.loadInstanceGroupManger(instanceGroupID, instanceGroupManagerID, *instanceGroupManager.Name))

		policies := make([]string, 0)

		for i := 0; i < len(instanceGroupManager.Policies); i++ {
			policies = append(policies, string(*(instanceGroupManager.Policies[i].ID)))
		}

		go g.handlePolicies(sess, instanceGroupID, instanceGroupManagerID, policies)
	}
}

func (g *InstanceGroupGenerator) handleInstanceGroups(sess *vpcv1.VpcV1) {
	// Support for pagination
	start := ""
	allrecs := []vpcv1.InstanceGroup{}
	for {
		listInstanceGroupOptions := vpcv1.ListInstanceGroupsOptions{}
		if start != "" {
			listInstanceGroupOptions.Start = &start
		}
		instanceGroupsCollection, response, err := sess.ListInstanceGroups(&listInstanceGroupOptions)
		if err != nil {
			g.fatalErrors <- fmt.Errorf("Error Fetching InstanceGroups %s\n%s", err, response)
		}
		start = GetNext(instanceGroupsCollection.Next)
		allrecs = append(allrecs, instanceGroupsCollection.InstanceGroups...)
		if start == "" {
			break
		}
	}

	for _, instanceGroup := range allrecs {
		instanceGoupID := *instanceGroup.ID
		g.Resources = append(g.Resources, g.loadInstanceGroup(instanceGoupID, *instanceGroup.Name))
		managers := make([]string, 0)
		for i := 0; i < len(instanceGroup.Managers); i++ {
			managers = append(managers, string(*(instanceGroup.Managers[i].ID)))
		}
		go g.handleManagers(sess, instanceGoupID, managers)
	}
	close(g.importDone)
}

func (g *InstanceGroupGenerator) InitResources() error {
	apiKey := os.Getenv("IBMCLOUD_API_KEY")
	if apiKey == "" {
		log.Fatal("No API key set")
	}

	// Instantiate the service with an API key based IAM authenticator
	sess, err := vpcv1.NewVpcV1(&vpcv1.VpcV1Options{
		Authenticator: &core.IamAuthenticator{
			ApiKey: apiKey,
		},
	})
	if err != nil {
		log.Fatal("Error creating VPC Service.")
		return err
	}

	g.importDone = make(chan bool)
	g.fatalErrors = make(chan error)

	go g.handleInstanceGroups(sess)

	select {
	case <-g.importDone:
		log.Println("Done Importing")
		break
	case err := <-g.fatalErrors:
		close(g.fatalErrors)
		log.Fatal("Error: ", err)
		return err
	}
	return nil
}
