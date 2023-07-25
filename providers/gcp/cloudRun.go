package gcp

import (
	"context"
	"fmt"
	"os"

	run "cloud.google.com/go/run/apiv2"
	"cloud.google.com/go/run/apiv2/runpb"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var cloudRunAllowEmptyValues = []string{""}
var cloudRunAdditionalFields = map[string]interface{}{}

type CloudRunGenerator struct {
	GCPService
}

func (g *CloudRunGenerator) InitResources() error {
	ctx := context.Background()
	filename := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	project := g.GetArgs()["project"].(string)
	location := g.GetArgs()["region"].(compute.Region).Name
	// endpoint := fmt.Sprintf("%s-aiplatform.googleapis.com:443", location)

	serviceClient, err := run.NewServicesClient(ctx, option.WithCredentialsFile(filename))
	if err != nil {
		return err
	}
	defer serviceClient.Close()

	it := serviceClient.ListServices(ctx, &runpb.ListServicesRequest{Parent: fmt.Sprintf("projects/%s/locations/%s", project, location)})
	if err := g.createServices(it); err != nil {
		return err
	}

	return nil
}

func (g *CloudRunGenerator) createServices(it *run.ServiceIterator) error {
	for {
		svc, err := it.Next()
		if err != nil {
			if err == iterator.Done {
				return nil
			}
			return err
		}

		project := g.GetArgs()["project"].(string)
		location := g.GetArgs()["region"].(compute.Region).Name

		resource := terraformutils.NewResource(
			svc.GetName(),
			svc.GetName(),
			"google_cloud_run_v2_service",
			g.GetProviderName(),
			map[string]string{
				"name":    svc.GetName(),
				"project": project,
				"region":  location,
			},
			cloudRunAllowEmptyValues,
			cloudRunAdditionalFields,
		)

		// if enc, err := json.MarshalIndent(resource, "", "    "); err == nil {
		// 	fmt.Printf("SVC: %s\n", enc)
		// }

		g.Resources = append(g.Resources, resource)

	}
}
