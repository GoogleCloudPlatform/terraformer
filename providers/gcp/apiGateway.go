package gcp

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"

	apigateway "cloud.google.com/go/apigateway/apiv1"
	pb "cloud.google.com/go/apigateway/apiv1/apigatewaypb"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

var apiGatewaysAllowEmptyValues = []string{""}

var apiGatewaysAdditionalFields = map[string]interface{}{}

type ApiGatewayGenerator struct {
	GCPService
}

func (g *ApiGatewayGenerator) InitResources() error {
	ctx := context.Background()
	filename := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	client, err := apigateway.NewClient(ctx, option.WithCredentialsFile(filename))
	if err != nil {
		return err
	}
	defer client.Close()

	g.Resources = []terraformutils.Resource{}

	apisIterator := client.ListApis(ctx, &pb.ListApisRequest{Parent: fmt.Sprintf("projects/%s/locations/global", g.GetArgs()["project"].(string))})
	if err := g.createApis(client, apisIterator); err != nil {
		log.Println(err)
	}

	return nil
}

func (g *ApiGatewayGenerator) createApis(client *apigateway.Client, it *apigateway.ApiIterator) error {
	for {
		api, err := it.Next()
		if err != nil {
			return err
		}

		project := g.GetArgs()["project"].(string)
		location := g.GetArgs()["region"].(compute.Region).Name

		g.Resources = append(g.Resources, terraformutils.NewResource(
			api.GetName(),
			api.GetDisplayName(),
			"google_api_gateway_api",
			g.GetProviderName(),
			map[string]string{
				"name":    api.GetName(),
				"project": project,
				"region":  location,
			},
			apiGatewaysAllowEmptyValues,
			map[string]interface{}{
				"api_id": api.GetDisplayName(),
			},
		))

		configsIterator := client.ListApiConfigs(context.Background(), &pb.ListApiConfigsRequest{Parent: fmt.Sprintf("projects/%s/locations/global/apis/%s", project, api.GetDisplayName())})
		if err := g.createConfigs(client, configsIterator); err != nil {
			log.Println(err)
		}

		gatewayIterator := client.ListGateways(context.Background(), &pb.ListGatewaysRequest{Parent: fmt.Sprintf("projects/%s/locations/%s", project, location)})
		if err := g.createGateways(gatewayIterator); err != nil {
			log.Println(err)
		}
	}
}

func (g *ApiGatewayGenerator) createConfigs(client *apigateway.Client, it *apigateway.ApiConfigIterator) error {
	for {
		obj, err := it.Next()
		if err != nil {
			return err
		}

		project := g.GetArgs()["project"].(string)
		location := g.GetArgs()["region"].(compute.Region).Name

		g.Resources = append(g.Resources, terraformutils.NewResource(
			obj.GetName(),
			obj.GetDisplayName(),
			"google_api_gateway_api_config",
			g.GetProviderName(),
			map[string]string{
				"name":    obj.GetName(),
				"project": project,
				"region":  location,
			},
			apiGatewaysAllowEmptyValues,
			apiGatewaysAdditionalFields,
		))
	}
}

func (g *ApiGatewayGenerator) createGateways(it *apigateway.GatewayIterator) error {
	for {
		obj, err := it.Next()
		if err != nil {
			return err
		}

		project := g.GetArgs()["project"].(string)
		location := g.GetArgs()["region"].(compute.Region).Name

		projectRe := regexp.MustCompile(`^projects\/([^\/]+)\/`)

		g.Resources = append(g.Resources, terraformutils.NewResource(
			obj.GetName(),
			obj.GetDisplayName(),
			"google_api_gateway_gateway",
			g.GetProviderName(),
			map[string]string{
				"name":    obj.GetName(),
				"project": project,
				"region":  location,
			},
			apiGatewaysAllowEmptyValues,
			map[string]interface{}{
				"display_name": obj.GetDisplayName(),
				"api_config":   projectRe.ReplaceAllString(obj.GetApiConfig(), "projects/"+project+"/"),
			},
		))
	}
}
