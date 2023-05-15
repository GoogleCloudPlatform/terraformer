package gcp

import (
	"context"

	cloudbuild "cloud.google.com/go/cloudbuild/apiv1"
	pb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const cbMaxPageSize = 50

type CloudBuildGenerator struct {
	GCPService
}

// InitResources generates TerraformResources from GCP API.
func (g *CloudBuildGenerator) InitResources() error {
	ctx := context.Background()

	c, err := cloudbuild.NewClient(ctx)
	if err != nil {
		return err
	}

	var (
		triggers      []*pb.BuildTrigger
		nextPageToken string
	)

	for {
		req := &pb.ListBuildTriggersRequest{
			ProjectId: g.GetArgs()["project"].(string),
			PageToken: nextPageToken,
			PageSize:  cbMaxPageSize,
		}

		res, err := c.ListBuildTriggers(ctx, req)
		if err != nil {
			return err
		}

		triggers = append(triggers, res.Triggers...)
		nextPageToken = res.NextPageToken

		if nextPageToken == "" {
			break
		}
	}

	g.Resources = g.createBuildTriggers(triggers)
	return nil
}

func (g *CloudBuildGenerator) createBuildTriggers(triggers []*pb.BuildTrigger) []terraformutils.Resource {
	var resources []terraformutils.Resource

	for _, trigger := range triggers {
		resources = append(resources, terraformutils.NewResource(
			trigger.GetId(),
			trigger.GetName(),
			"google_cloudbuild_trigger",
			g.ProviderName,
			map[string]string{
				"project": g.GetArgs()["project"].(string),
			},
			[]string{},
			map[string]interface{}{
				"filename": trigger.GetFilename(),
			},
		))
	}

	return resources
}
