package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	datadog_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/datadog"
)

var (
	commandTerraformInit       = "terraform init"
	commandTerraformPlan       = "terraform plan -detailed-exitcode"
	commandTerraformDestroy    = "terraform destroy -auto-approve"
	commandTerraformApply      = "terraform apply -auto-approve"
	commandTerraformOutput     = "terraform output"
	commandTerraformV13Upgrade = "terraform 0.13upgrade -yes ."
	datadogResourcesPath       = "tests/datadog/resources/"
)

type DatadogConfig struct {
	apiKey string
	appKey string
	apiURL string
}

type TerraformConfig struct {
	target string
}

type Config struct {
	Datadog      DatadogConfig
	Terraform    TerraformConfig
	logCMDOutput bool
	rootPath     string
	tfVersion    string
}

func getConfig() (*Config, error) {
	logCMDOutput := false
	if envVar := os.Getenv("LOG_CMD_OUTPUT"); envVar != "" {
		logCMDOutputEnv, err := strconv.ParseBool(envVar)
		if err != nil {
			return nil, err
		}
		logCMDOutput = logCMDOutputEnv
	}
	rootPath, _ := os.Getwd()

	return &Config{
		Datadog: DatadogConfig{
			apiKey: os.Getenv("DD_TEST_CLIENT_API_KEY"),
			appKey: os.Getenv("DD_TEST_CLIENT_APP_KEY"),
			apiURL: os.Getenv("DATADOG_HOST"),
		},
		Terraform: TerraformConfig{
			target: os.Getenv("DATADOG_TERRAFORM_TARGET"),
		},
		logCMDOutput: logCMDOutput,
		rootPath:     rootPath,
		tfVersion:    os.Getenv("DATADOG_TF_VERSION"),
	}, nil
}

func getAllServices(provider *datadog_terraforming.DatadogProvider) []string {
	var services []string
	for service := range provider.GetSupportedService() {
		if service == "timeboard" {
			continue
		}
		if service == "screenboard" {
			continue
		}
		services = append(services, service)
	}
	return services
}

func initializeDatadogProvider(cfg *Config) error {
	// Initialize the provider
	log.Print("Initializing the Datadog provider")
	if err := cmdRun(cfg, []string{commandTerraformInit}); err != nil {
		return err
	}
	log.Print("Successfully initialized  the Datadog provider")
	return nil
}

func createDatadogResource(cfg *Config) (*map[string][]string, error) {
	// Create terraform -target flags if targets are passed
	var terraformTargets []string
	if v := cfg.Terraform.target; v != "" {
		vArr := strings.Split(v, ":")
		for _, terraformTarget := range vArr {
			terraformTargetFlag := fmt.Sprintf("-target=%s", terraformTarget)
			terraformTargets = append(terraformTargets, terraformTargetFlag)
		}
	}

	// Create resources
	log.Print("Creating resources")
	if err := cmdRun(cfg, []string{commandTerraformApply, strings.Join(terraformTargets, " ")}); err != nil {
		return nil, err
	}

	// Get output of created resources and parse the data into a map
	output, err := terraformOutput()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resources := parseTerraformOutput(string(output))
	log.Printf("Created resources: \n%s", string(output))

	return resources, nil
}

func terraformOutput() ([]byte, error) {
	output, err := exec.Command("sh", "-c", commandTerraformOutput).Output()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return output, nil
}

func terraformPlan(cfg *Config) error {
	log.Print("Running terraform plan against resources")
	err := cmdRun(cfg, []string{commandTerraformPlan})
	if err != nil {
		return err
	}

	log.Print("terraform plan did not generate diffs")
	return nil
}

func destroyDatadogResources(cfg *Config) error {
	// Destroy created resources in the /tests/datadog/resources directory
	err := os.Chdir(cfg.rootPath + "/" + datadogResourcesPath)
	if err != nil {
		return err
	}
	log.Print("Destroying resources")
	if err := cmdRun(cfg, []string{commandTerraformDestroy}); err != nil {
		return err
	}
	_ = os.Chdir(cfg.rootPath)

	return nil
}

func parseTerraformOutput(output string) *map[string][]string {
	outputArr := strings.Split(output, "\n")
	resources := map[string][]string{}
	for _, resourceOutput := range outputArr {
		if len(resourceOutput) > 0 {
			resourceArr := strings.Split(resourceOutput, " = ")
			resourceID := resourceArr[len(resourceArr)-1]
			// Get resource name
			re := regexp.MustCompile("_(.*?)(--|_)")
			match := re.FindStringSubmatch(resourceArr[0])
			resourceName := match[1]

			resources[resourceName] = append(resources[resourceName], resourceID)
		}
	}
	return &resources
}

func cmdRun(cfg *Config, args []string) error {
	terraformAPIKeyEnvVariable := fmt.Sprintf("DATADOG_API_KEY=%s", cfg.Datadog.apiKey)
	terraformAppKeyEnvVariable := fmt.Sprintf("DATADOG_APP_KEY=%s", cfg.Datadog.appKey)

	cmd := exec.Command("sh", "-c", strings.Join(args, " "))
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, terraformAPIKeyEnvVariable, terraformAppKeyEnvVariable)
	if cfg.logCMDOutput {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err := cmd.Run()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
