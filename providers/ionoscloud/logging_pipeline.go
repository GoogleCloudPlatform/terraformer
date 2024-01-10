package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type LoggingPipelineGenerator struct {
	Service
}

func (g *LoggingPipelineGenerator) InitResources() error {
	client := g.generateClient()
	loggingAPIClient := client.LoggingAPIClient
	resourceType := "ionoscloud_logging_pipeline"

	response, _, err := loggingAPIClient.PipelinesApi.PipelinesGet(context.TODO()).Execute()
	if err != nil {
		return err
	}
	if response.Items == nil {
		log.Printf("[WARNING] expected a response containing pipelines, but received 'nil' instead")
		return nil
	}
	pipelines := *response.Items
	for _, pipeline := range pipelines {
		if pipeline.Properties == nil || pipeline.Properties.Name == nil {
			log.Printf("[WARNING] 'nil' values in the response for the pipeline with ID: %v, skipping this resource", *pipeline.Id)
			continue
		}
		g.Resources = append(g.Resources, terraformutils.NewResource(
			*pipeline.Id,
			*pipeline.Properties.Name+"-"+*pipeline.Id,
			resourceType,
			helpers.Ionos,
			map[string]string{},
			[]string{},
			map[string]interface{}{}))
	}
	return nil
}
