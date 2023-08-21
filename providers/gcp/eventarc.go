package gcp

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"

	"google.golang.org/api/compute/v1"
	"google.golang.org/api/eventarc/v1"
)

var eventarcAllowEmptyValues = []string{""}

var eventarcAdditionalFields = map[string]interface{}{}

type EventarcGenerator struct {
	GCPService
}

func (g *EventarcGenerator) generateTriggers(ctx context.Context, triggerList *eventarc.ProjectsLocationsTriggersListCall) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	if err := triggerList.Pages(ctx, func(page *eventarc.ListTriggersResponse) error {
		for _, trigger := range page.Triggers {
			resource := terraformutils.NewResource(
				trigger.Name,
				trigger.Name,
				"google_eventarc_trigger",
				g.ProviderName,
				map[string]string{
					"name":     trigger.Name,
					"project":  g.GetArgs()["project"].(string),
					"location": g.GetArgs()["region"].(compute.Region).Name,
				},
				eventarcAllowEmptyValues,
				eventarcAdditionalFields,
			)
			resources = append(resources, resource)
		}
		return nil
	}); err != nil {
		log.Println(err)
	}
	return resources
}

func (g *EventarcGenerator) InitResources() error {
	ctx := context.Background()
	eventarcService, err := eventarc.NewService(ctx)
	if err != nil {
		return err
	}

	project := g.GetArgs()["project"].(string)
	region := g.GetArgs()["region"].(compute.Region).Name

	clusterList := eventarcService.Projects.Locations.Triggers.List("projects/" + project + "/locations/" + region)
	g.Resources = g.generateTriggers(ctx, clusterList)

	return nil
}
