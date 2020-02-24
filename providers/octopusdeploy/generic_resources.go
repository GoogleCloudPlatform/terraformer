package octopusdeploy

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

type GenericGenerator struct {
	OctopusDeployService
	APIService string
}

// InitResources initialize the process to generate the Terraform resources from the
// Octopus Deploy API.
func (g *GenericGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*octopusdeploy.Client) error{
		g.createResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *GenericGenerator) createResources(client *octopusdeploy.Client) error {
	switch strings.ToLower(g.APIService) {
	case "accounts":
		resources, err := client.Account.GetAll()
		if err != nil {
			return err
		}

		for _, ressource := range *resources {
			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				fmt.Sprintf("%s", ressource.ID),
				fmt.Sprintf("%s", ressource.Name),
				"octopusdeploy_account",
				g.ProviderName,
				[]string{},
			))
		}
	case "certificates":
		resources, err := client.Certificate.GetAll()
		if err != nil {
			return err
		}

		for _, ressource := range *resources {
			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				fmt.Sprintf("%s", ressource.ID),
				fmt.Sprintf("%s", ressource.Name),
				"octopusdeploy_certificate",
				g.ProviderName,
				[]string{},
			))
		}
	// case "channels":
	// TODO: Somehow there is an issue with the channels:
	//    2020/02/24 16:35:55 octopusdeploy importing... channels
	//    2020/02/24 16:35:55 cannot find the item

	// 	resources, err := client.Channel.GetAll()
	// 	if err != nil {
	// 		return err
	// 	}

	// 	for _, ressource := range *resources {
	// 		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
	// 			fmt.Sprintf("%s", ressource.ID),
	// 			fmt.Sprintf("%s", ressource.Name),
	// 			"octopusdeploy_channel",
	// 			g.ProviderName,
	// 			[]string{},
	// 		))
	// 	}
	case "environments":
		resources, err := client.Environment.GetAll()
		if err != nil {
			return err
		}

		for _, ressource := range *resources {
			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				fmt.Sprintf("%s", ressource.ID),
				fmt.Sprintf("%s", ressource.Name),
				"octopusdeploy_environment",
				g.ProviderName,
				[]string{},
			))
		}
	case "feeds":
		resources, err := client.Feed.GetAll()
		if err != nil {
			return err
		}

		for _, ressource := range *resources {
			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				fmt.Sprintf("%s", ressource.ID),
				fmt.Sprintf("%s", ressource.Name),
				"octopusdeploy_feed",
				g.ProviderName,
				[]string{},
			))
		}
	case "libraryvariablesets":
		resources, err := client.LibraryVariableSet.GetAll()
		if err != nil {
			return err
		}

		for _, ressource := range *resources {
			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				fmt.Sprintf("%s", ressource.ID),
				fmt.Sprintf("%s", ressource.Name),
				"octopusdeploy_library_variable_set",
				g.ProviderName,
				[]string{},
			))
		}
	case "lifecycles":
		resources, err := client.Lifecycle.GetAll()
		if err != nil {
			return err
		}

		for _, ressource := range *resources {
			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				fmt.Sprintf("%s", ressource.ID),
				fmt.Sprintf("%s", ressource.Name),
				"octopusdeploy_lifecycle",
				g.ProviderName,
				[]string{},
			))
		}
	case "projects":
		resources, err := client.Project.GetAll()
		if err != nil {
			return err
		}

		for _, ressource := range *resources {
			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				fmt.Sprintf("%s", ressource.ID),
				fmt.Sprintf("%s", ressource.Name),
				"octopusdeploy_project",
				g.ProviderName,
				[]string{},
			))
		}
	case "projectgroups":
		resources, err := client.ProjectGroup.GetAll()
		if err != nil {
			return err
		}

		for _, ressource := range *resources {
			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				fmt.Sprintf("%s", ressource.ID),
				fmt.Sprintf("%s", ressource.Name),
				"octopusdeploy_project_group",
				g.ProviderName,
				[]string{},
			))
		}
	case "projecttriggers":
		resources, err := client.ProjectTrigger.GetAll()
		if err != nil {
			return err
		}

		for _, ressource := range *resources {
			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				fmt.Sprintf("%s", ressource.ID),
				fmt.Sprintf("%s", ressource.Name),
				"octopusdeploy_project_deployment_target_trigger",
				g.ProviderName,
				[]string{},
			))
		}
	case "tagsets":
		resources, err := client.TagSet.GetAll()
		if err != nil {
			return err
		}

		for _, ressource := range *resources {
			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				fmt.Sprintf("%s", ressource.ID),
				fmt.Sprintf("%s", ressource.Name),
				"octopusdeploy_tag_set",
				g.ProviderName,
				[]string{},
			))
		}
		// case "variables":
		// TODO: This cannot generate a `variables.tf` file as there is already one.

		// projects, err := client.Project.GetAll()
		// if err != nil {
		// 	return err
		// }

		// for _, project := range *projects {
		// 	// Variable.GetAll() returns all the variables for a specific project ID.
		// 	resources, err := client.Variable.GetAll(project.ID)
		// 	if err != nil {
		// 		return err
		// 	}

		// 	for _, ressource := range resources.Variables {
		// 		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
		// 			fmt.Sprintf("%s", ressource.ID),
		// 			fmt.Sprintf("%s", ressource.Name),
		// 			"octopusdeploy_variable",
		// 			g.ProviderName,
		// 			[]string{},
		// 		))
		// 	}
		// }
	}

	return nil
}
