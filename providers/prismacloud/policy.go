package prismacloud

import (
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/policy"
	"log"
)

// PolicyGenerator ...
type PolicyGenerator struct {
	PrismaCloudService
}

func (g *PolicyGenerator) createResources(policies []policy.Policy) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, p := range policies {
		log.Printf("Policy info : %#v", p)
		resources = append(resources, g.createResource(p.PolicyId))
	}
	return resources
}

func (g *PolicyGenerator) createResource(policyID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		policyID,
		fmt.Sprintf("policy_%s", policyID),
		"prismacloud_policy",
		"prismacloud",
		[]string{},
	)
}

// InitResources Generate Terraform Resources from PrismaCloud API,
func (g *PolicyGenerator) InitResources() error {
	configFilePath := g.Args["config_file_path"].(string)
	con := &pc.Client{}
	err := con.Initialize(configFilePath)
	if err != nil {
		return err
	}
	policies, err := policy.List(con, nil)
	if err != nil {
		return err
	}
	resources := g.createResources(policies)
	log.Printf("Resources list : %#v", resources)
	g.Resources = append(g.Resources, resources...)
	return nil
}
