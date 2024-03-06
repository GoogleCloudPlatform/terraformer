package ibm

import (
    "context"
    "github.com/GoogleCloudPlatform/terraformer/terraformutils"
    "github.com/IBM/continuous-delivery-go-sdk/cdtoolchainv2"
    "github.com/IBM/go-sdk-core/v5/core"
    "os"
    "log"
    "fmt"
)

type CdToolchainGenerator struct {
    IBMService
}

func (g CdToolchainGenerator) createCdToolchainResources(id, name string) terraformutils.Resource {
        resources := terraformutils.NewSimpleResource(
            id,
            name,
            "ibm_cd_toolchain",
            "ibm",
            []string{},
        )
        return resources
}

func (g *CdToolchainGenerator) InitResources() error {
    ctx := context.Background()
    auth := &core.IamAuthenticator{
        ApiKey: os.Getenv("IC_API_KEY"),
    }
    region := g.Args["region"].(string)
    apiKey := os.Getenv("IC_API_KEY")
    if apiKey == "" {
            log.Fatal("No API key set")
    }

    service, err := cdtoolchainv2.NewCdToolchainV2(&cdtoolchainv2.CdToolchainV2Options{
        Authenticator: auth,
    })
    if err != nil {
        return err
    }

    listToolchainsOptions := &cdtoolchainv2.ListToolchainsOptions{}
    //if start != "" {
    //	listToolchainsOptions.Start = &start
    //}
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

    for _, toolchain := range toolchainCollection.Toolchains {
        // Check for nil pointers and use default values if necessary
        
        id := ""
        if toolchain.ID != nil {
            id = *toolchain.ID
        }

        name := ""
        if toolchain.Name != nil {
            name = *toolchain.Name
        }

//        resource := terraformutils.NewSimpleResource(
//            id,
//            name,
//            "ibm_cd_toolchain",
//            "ibm",
//            []string{},
//        )
        // Append to the generator's internal resource list; adjust this line according to how resources are managed in your generator
        g.Resources = append(g.Resources, g.createCdToolchainResources(id, name))
	//g.AddResource(resource)
    }

    return nil
}

//func (g *CdToolchainGenerator) GetProviderData(arg ...string) map[string]interface{} {
//    return map[string]interface{}{
//        "provider":           "ibmcdtoolchain",
//        "resource_namespace": "ibm",
//    }
//}

