package ibm

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/GoogleCloudPlatform/terraformer/terraformutils"
    "github.com/IBM/continuous-delivery-go-sdk/cdtoolchainv2"
    "github.com/IBM/continuous-delivery-go-sdk/cdtektonpipelinev2"
    "github.com/IBM/go-sdk-core/v5/core"
)

type CdTektonPipelineGenerator struct {
    IBMService
}

func (g *CdTektonPipelineGenerator) createTektonPipelineResource(toolchainToolId, pipelineId, pipelineName string) terraformutils.Resource {
    // Assuming the ID format "<toolchain_id>/<pipeline_id>" to ensure uniqueness
    name := fmt.Sprintf("%s/%s", pipelineName, pipelineId)
    return terraformutils.NewSimpleResource(
        pipelineId,
        name,
        "ibm_cd_tekton_pipeline",
        "ibm",
        []string{},
    )
}

func (g *CdTektonPipelineGenerator) InitResources() error {
    ctx := context.Background()
    apiKey := os.Getenv("IC_API_KEY")
    if apiKey == "" {
        log.Fatal("No API key set")
    }
    region := g.Args["region"].(string)

    auth := &core.IamAuthenticator{ApiKey: apiKey}
    service, err := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
        Authenticator: auth,
    })
    tektonService, err := cdtektonpipelinev2.NewCdTektonPipelineV2(&cdtektonpipelinev2.CdTektonPipelineV2Options{
        Authenticator: auth,
    })
    if err != nil {
        return err
    }

    // First, list all toolchains
    listToolchainsOptions := &cdtoolchainv2.ListToolchainsOptions{}
    if rg := g.Args["resource_group"].(string); rg != "" {
            rg, err = GetResourceGroupID(apiKey, rg, region)
            if err != nil {
                    return fmt.Errorf("Error Fetching Resource Group Id %s", err)
            }
            listToolchainsOptions.ResourceGroupID = &rg
    } 

    toolchainCollection, _, err := service.ListToolchainsWithContext(ctx, listToolchainsOptions)
    if err != nil {
        return err
    }

    // Then, for each toolchain, list and process its pipelines
    for _, toolchain := range toolchainCollection.Toolchains {
        toolchainId := *toolchain.ID
        listPipelinesOptions := &cdtoolchainv2.ListToolsOptions{
            ToolchainID: &toolchainId,
        }
        toolchaintoolCollection, _, err := service.ListToolsWithContext(ctx, listPipelinesOptions)
        if err != nil {
            log.Printf("Error listing pipelines for toolchain %s: %v", toolchainId, err)
            continue // Skip to the next toolchain if we can't list pipelines for this one
        }
        for _, toolchaintool := range toolchaintoolCollection.Tools {
            fmt.Printf("Toolchaintool: %s ToolType %s\n", *toolchaintool.ID, *toolchaintool.ToolTypeID)
            if *toolchaintool.ToolTypeID == "pipeline" {

                pipelineId := *toolchaintool.ID
                // Get Tekton pipeline details
                getPipelineOptions := &cdtektonpipelinev2.GetTektonPipelineOptions{
                        ID: &pipelineId,
                }
                pipeline, _, err := tektonService.GetTektonPipelineWithContext(ctx, getPipelineOptions)
                if err != nil {
                    log.Fatalf("Failed to get Tekton pipeline details for pipeline ID %s: %v", pipelineId, err)
                }
                if pipeline.ID != nil && pipeline.Name != nil {
                    g.Resources = append(g.Resources, g.createTektonPipelineResource(toolchainId, pipelineId, *pipeline.Name))
                }
            }
	}
    }

    return nil
}

