package opal

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/opalsecurity/opal-go"
)

type ResourceGenerator struct {
	OpalService
}

func (g *ResourceGenerator) InitResources() error {
	client, err := g.newClient()
	if err != nil {
		return fmt.Errorf("unable to list opal resources: %v", err)
	}

	resources, _, err := client.ResourcesApi.GetResources(context.TODO()).Execute()
	if err != nil {
		return fmt.Errorf("unable to list opal resources: %v", err)
	}

	var opalResources []*opal.Resource
	for {
		for _, resource := range resources.Results {
			resourceRef := resource
			opalResources = append(opalResources, &resourceRef)
		}

		if !resources.HasNext() || resources.Next.Get() == nil {
			break
		}

		resources, _, err = client.ResourcesApi.GetResources(context.TODO()).Cursor(*resources.Next.Get()).Execute()
		if err != nil {
			return fmt.Errorf("unable to list opal resources: %v", err)
		}
	}

	opalResourceByID := make(map[string]*opal.Resource)
	for _, resource := range opalResources {
		opalResourceByID[resource.ResourceId] = resource
	}

	seenNames := make(map[string]bool)
	for _, resource := range opalResources {
		tfname := *resource.Name
		if resource.ResourceType != nil &&
			*resource.ResourceType == opal.RESOURCETYPEENUM_AWS_SSO_PERMISSION_SET &&
			resource.ParentResourceId != nil {
			parentAccount, ok := opalResourceByID[*resource.ParentResourceId]
			if !ok {
				return fmt.Errorf("could not find account for permission set: %#v", resource)
			}
			tfname = fmt.Sprintf("%s_%s", *parentAccount.Name, *resource.Name)
		}

		if seenNames[tfname] {
			tfname = tfname + "_" + resource.ResourceId[:8]
		} else {
			seenNames[tfname] = true
		}

		tfname = normalizeResourceName(tfname)
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			resource.ResourceId,
			tfname,
			"opal_resource",
			"opal",
			[]string{},
		))
	}

	return nil
}
