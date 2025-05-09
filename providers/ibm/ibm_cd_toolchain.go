package ibm

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
	"sync"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/catalog"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev2/controllerv2"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM/continuous-delivery-go-sdk/v2/cdtektonpipelinev2"
	"github.com/IBM/continuous-delivery-go-sdk/v2/cdtoolchainv2"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
)

type ToolchainGenerator struct {
	IBMService
}

var workerIDMutex sync.RWMutex // Used in PostConvertHook
var repoMutex sync.RWMutex     // Used in PostConvertHook
var toolMutex sync.RWMutex     // Used in PostConvertHook

func (g ToolchainGenerator) loadToolchain(tcID string, tcName string) terraformutils.Resource {
	resource := terraformutils.NewSimpleResource(
		tcID,
		tcName,
		"ibm_cd_toolchain",
		"ibm",
		[]string{},
	)
	return resource
}

func (g ToolchainGenerator) loadTool(resourceType string, tID string, tName string, tcIDref string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		tID,
		tName,
		resourceType,
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"toolchain_id": tcIDref,
		})
	return resource
}

// Adds S2S authorization required by some integrations
func (g ToolchainGenerator) loadAuthPolicies(policyID string, tcIDref string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		policyID,
		normalizeResourceName("iam_authorization_policy", true),
		"ibm_iam_authorization_policy",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"source_resource_instance_id": tcIDref,
		})

	// Conflict parameters
	resource.IgnoreKeys = append(resource.IgnoreKeys,
		"^subject_attributes$",
		"^resource_attributes$",
		"^source_service_account$",
		"^transaction_id$",
	)
	return resource
}

func (g ToolchainGenerator) loadPL(plID string, plName string, plIDref string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		plID,
		plName,
		"ibm_cd_tekton_pipeline",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"pipeline_id": plIDref,
		})
	return resource
}

func (g ToolchainGenerator) loadPLProp(resourceType string, pID string, pName string, plIDref string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		pID,
		pName,
		resourceType,
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"pipeline_id": plIDref,
		})
	return resource
}

func (g ToolchainGenerator) loadPLDef(resourceType string, pID string, pName string, plIDref string, tcID string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		pID,
		pName,
		resourceType,
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"pipeline_id":         plIDref,
			"toolchain_id_actual": tcID, // removed on PostConvertHook
		})
	return resource
}

func (g ToolchainGenerator) loadPLTrigProp(resourceType string, pID string, pName string, plIDref string, trigIDref string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		pID,
		pName,
		resourceType,
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"pipeline_id": plIDref,
			"trigger_id":  trigIDref,
		})
	return resource
}

// Goroutine helper to handle different tool types
func (g *ToolchainGenerator) HandleTool(t cdtoolchainv2.ToolModel, toolType string, tID string, tName string, tcID string, tcIDref string, waitGroup *sync.WaitGroup) error {
	defer waitGroup.Done()

	apiKey := os.Getenv("IC_API_KEY")

	// typical case. handle exceptional cases seperately
	// maps tool_type_id to the terraform resource type
	supportedTools := map[string]string{
		"appconfig":           "ibm_cd_toolchain_tool_appconfig",
		"artifactory":         "ibm_cd_toolchain_tool_artifactory",
		"bitbucketgit":        "ibm_cd_toolchain_tool_bitbucketgit",
		"private_worker":      "ibm_cd_toolchain_tool_privateworker",
		"draservicebroker":    "ibm_cd_toolchain_tool_devopsinsights",
		"eventnotifications":  "ibm_cd_toolchain_tool_eventnotifications",
		"hostedgit":           "ibm_cd_toolchain_tool_hostedgit",
		"githubconsolidated":  "ibm_cd_toolchain_tool_githubconsolidated",
		"gitlab":              "ibm_cd_toolchain_tool_gitlab",
		"hashicorpvault":      "ibm_cd_toolchain_tool_hashicorpvault",
		"jenkins":             "ibm_cd_toolchain_tool_jenkins",
		"jira":                "ibm_cd_toolchain_tool_jira",
		"keyprotect":          "ibm_cd_toolchain_tool_keyprotect",
		"nexus":               "ibm_cd_toolchain_tool_nexus",
		"customtool":          "ibm_cd_toolchain_tool_custom",
		"saucelabs":           "ibm_cd_toolchain_tool_saucelabs",
		"secretsmanager":      "ibm_cd_toolchain_tool_secretsmanager",
		"security_compliance": "ibm_cd_toolchain_tool_securitycompliance",
		"slack":               "ibm_cd_toolchain_tool_slack",
		"sonarqube":           "ibm_cd_toolchain_tool_sonarqube",
	}

	if resourceType, ok := supportedTools[toolType]; ok {
		resourceMutex.Lock()
		g.Resources = append(g.Resources, g.loadTool(resourceType, tID, tName, tcIDref))
		resourceMutex.Unlock()
	} else {
		switch toolType {
		case "pipeline":
			// Classic pipelines cannot be created using Terraform
			if t.Parameters["type"] != "tekton" {
				resourceMutex.Lock()
				g.Resources = append(g.Resources, g.loadTool("ibm_cd_toolchain_tool_pipeline", tID, tName+"--classic", tcIDref))
				resourceMutex.Unlock()
				fmt.Println("......! Only Tekton pipelines are supported in Terraform", toolType)
				return nil
			}

			resourceMutex.Lock()
			g.Resources = append(g.Resources, g.loadTool("ibm_cd_toolchain_tool_pipeline", tID, tName+"--tekton", tcIDref))
			resourceMutex.Unlock()

			plID := *(t.ID)
			plName := tName

			plIDref := fmt.Sprintf("${ibm_cd_toolchain_tool_pipeline.tfer--%s--tekton.tool_id}", tName)

			resourceMutex.Lock()
			g.Resources = append(g.Resources, g.loadPL(plID, plName, plIDref))
			resourceMutex.Unlock()

			// Get pipeline
			cdTektonPipelineServiceOptions := &cdtektonpipelinev2.CdTektonPipelineV2Options{
				Authenticator: &core.IamAuthenticator{
					ApiKey: apiKey,
				},
			}

			cdTektonPipelineService, err := cdtektonpipelinev2.NewCdTektonPipelineV2UsingExternalConfig(cdTektonPipelineServiceOptions)
			if err != nil {
				log.Print("......! Error getting pipeline information: ", err)
			}

			getTektonPipelineOptions := cdTektonPipelineService.NewGetTektonPipelineOptions(plID)

			tektonPipeline, _, err := cdTektonPipelineService.GetTektonPipeline(getTektonPipelineOptions)
			if err != nil {
				log.Print("......! Error getting pipeline information: ", err)
			}

			// Definitions
			for _, def := range tektonPipeline.Definitions {
				defID := fmt.Sprintf("%s/%s", plID, *(def.ID))
				defName := normalizeResourceName("definition", true)

				resourceMutex.Lock()
				g.Resources = append(g.Resources, g.loadPLDef("ibm_cd_tekton_pipeline_definition", defID, defName, plIDref, plID))
				resourceMutex.Unlock()
			}

			// Properties
			for _, prop := range tektonPipeline.Properties {
				pID := fmt.Sprintf("%s/%s", plID, *(prop.Name))
				pName := normalizeResourceName(*(prop.Name), true)

				resourceMutex.Lock()
				g.Resources = append(g.Resources, g.loadPLProp("ibm_cd_tekton_pipeline_property", pID, pName, plIDref))
				resourceMutex.Unlock()
			}

			// Triggers
			for _, trig := range tektonPipeline.Triggers {
				trigger := trig.(*cdtektonpipelinev2.Trigger)

				trigID := fmt.Sprintf("%s/%s", plID, *(trigger.ID))
				trigName := normalizeResourceName(*(trigger.Name), true)

				resourceMutex.Lock()
				g.Resources = append(g.Resources, g.loadPLProp("ibm_cd_tekton_pipeline_trigger", trigID, trigName, plIDref))
				resourceMutex.Unlock()

				// Trigger Properties
				for _, trigp := range trigger.Properties {
					trigpID := fmt.Sprintf("%s/%s", trigID, *(trigp.Name))
					trigpName := normalizeResourceName(*(trigp.Name), true)

					trigIDref := fmt.Sprintf("${ibm_cd_tekton_pipeline_trigger.tfer--%s.trigger_id}", trigName)

					resourceMutex.Lock()
					g.Resources = append(g.Resources, g.loadPLTrigProp("ibm_cd_tekton_pipeline_trigger_property", trigpID, trigpName, plIDref, trigIDref))
					resourceMutex.Unlock()
				}
			}
		case "pagerduty":
			// If this integration is misconfigured, it lacks the necessary fields to work in Terraform
			if *(t.State) == "configured" {
				resourceMutex.Lock()
				g.Resources = append(g.Resources, g.loadTool("ibm_cd_toolchain_tool_pagerduty", tID, tName, tcIDref))
				resourceMutex.Unlock()
			}
		default:
			fmt.Println("......! Unknown tool type", toolType)
		}
	}
	return nil
}

// Called within InitResources when IBM_CD_TOOLCHAIN_INCLUDE_S2S is set
func getS2SPolicies(sess *session.Session, targetTcID string) (map[string][]iampolicymanagementv1.Policy, error) {
	apiKey := os.Getenv("IC_API_KEY")

	emptyPolicies := map[string][]iampolicymanagementv1.Policy{}

	iamPolicyOptions := &iampolicymanagementv1.IamPolicyManagementV1Options{
		URL: "https://iam.cloud.ibm.com",
		Authenticator: &core.IamAuthenticator{
			ApiKey: apiKey,
		},
	}

	iamPolicyClient, err := iampolicymanagementv1.NewIamPolicyManagementV1(iamPolicyOptions)
	if err != nil {
		return emptyPolicies, err
	}

	userInfo, err := fetchUserDetails(sess, 2)
	if err != nil {
		return emptyPolicies, err
	}
	accountID := userInfo.userAccount

	listAuthPolicyOptions := iampolicymanagementv1.ListPoliciesOptions{
		AccountID: core.StringPtr(accountID),
		Type:      core.StringPtr("authorization"),
	}

	authPolicyList, _, err := iamPolicyClient.ListPolicies(&listAuthPolicyOptions)
	if err != nil {
		return emptyPolicies, fmt.Errorf("error retrieving authorization policy: %s", err)
	}
	authPolicies := authPolicyList.Policies

	s2sPolicies := map[string][]iampolicymanagementv1.Policy{} // map of toolchain id to s2s policies under it

	for _, ap := range authPolicies {
		for _, a := range ap.Subjects[0].Attributes {
			if *(a.Name) != "serviceInstance" {
				continue
			}
			if (targetTcID != "" && *(a.Value) == targetTcID) || targetTcID == "" {
				// get s2s policies for target toolchain
				if _, ok := s2sPolicies[*(a.Value)]; !ok {
					s2sPolicies[*(a.Value)] = []iampolicymanagementv1.Policy{}
				}
				s2sPolicies[*(a.Value)] = append(s2sPolicies[*(a.Value)], ap)
			}
		}
	}
	return s2sPolicies, nil
}

func (g *ToolchainGenerator) InitResources() error {
	region := g.Args["region"].(string)

	guidRegex := regexp.MustCompile("[0-9a-fA-F]{8}-([0-9a-fA-F]{4}-){3}[0-9a-fA-F]{12}$")

	targetTcID := os.Getenv("IBM_CD_TOOLCHAIN_TARGET")
	if targetTcID != "" && !guidRegex.MatchString(targetTcID) {
		log.Fatal("Env variable IBM_CD_TOOLCHAIN_TARGET is not a GUID")
	}

	apiKey := os.Getenv("IC_API_KEY")
	if apiKey == "" {
		log.Fatal("No API key set")
	}

	bmxConfig := &bluemix.Config{
		BluemixAPIKey: apiKey,
		Region:        region,
	}

	sess, err := session.New(bmxConfig)
	if err != nil {
		return err
	}

	err = authenticateAPIKey(sess)
	if err != nil {
		return err
	}

	catalogClient, err := catalog.New(sess)
	if err != nil {
		return err
	}

	controllerClient, err := controllerv2.New(sess)
	if err != nil {
		return err
	}

	serviceID, err := catalogClient.ResourceCatalog().FindByName("toolchain", true)
	if err != nil {
		return err
	}

	query := controllerv2.ServiceInstanceQuery{
		ServiceID: serviceID[0].ID,
	}

	tcInstances, err := controllerClient.ResourceServiceInstanceV2().ListInstances(query)
	if err != nil {
		return err
	}

	// Get s2s policies
	s2sPolicies := map[string][]iampolicymanagementv1.Policy{}

	includeS2S := os.Getenv("IBM_CD_TOOLCHAIN_INCLUDE_S2S")
	if includeS2S != "" {
		s2sPolicies, err = getS2SPolicies(sess, targetTcID)
		if err != nil {
			return err
		}
	}

	var toolWG sync.WaitGroup

	// Iterate over toolchains to get tools
	for _, tc := range tcInstances {
		// Get toolchain ids, double-checking if they are valid
		crnSplit := strings.Split(tc.ID, ":")
		if len(crnSplit) < 8 {
			fmt.Println("received invalid CRN format from Resource Controller, skipping...")
			continue
		}

		tcID := crnSplit[7]

		if !guidRegex.MatchString(tcID) {
			fmt.Println("received invalid CRN format from Resource Controller, skipping...")
			continue
		}

		if targetTcID != "" && tcID != targetTcID {
			continue
		}

		if tc.RegionID == region {
			tcName := normalizeResourceName(tc.Name, true)
			tcIDref := fmt.Sprintf("${ibm_cd_toolchain.tfer--%s.id}", tcName)

			resourceMutex.Lock()
			g.Resources = append(g.Resources, g.loadToolchain(tcID, tcName))
			resourceMutex.Unlock()

			fmt.Println("=== FOUND TOOLCHAIN", tcID, "WITH NAME", tcName)

			// Get tools
			toolchainClientOptions := &cdtoolchainv2.CdToolchainV2Options{
				Authenticator: &core.IamAuthenticator{
					ApiKey: apiKey,
				},
			}

			toolchainClient, err := cdtoolchainv2.NewCdToolchainV2UsingExternalConfig(toolchainClientOptions)
			if err != nil {
				return err
			}

			listToolsOptions := toolchainClient.NewListToolsOptions(tcID)

			listToolsOptions.SetLimit(150) // 150 is max num tools per toolchain

			tools, _, err := toolchainClient.ListTools(listToolsOptions)
			if err != nil {
				return err
			}

			if includeS2S != "" {
				// Add toolchain's s2s policies (some tools require it)
				for _, pol := range s2sPolicies[tcID] {
					resourceMutex.Lock()
					g.Resources = append(g.Resources, g.loadAuthPolicies(*(pol.ID), tcIDref))
					resourceMutex.Unlock()
				}
			}

			for _, t := range tools.Tools {
				toolType := *(t.ToolTypeID)
				tID := fmt.Sprintf("%s/%s", tcID, *(t.ID))

				// Name won't always exist in Parameters
				var tName string

				if t.Parameters["name"] != nil {
					tName = normalizeResourceName(t.Parameters["name"].(string), true)
				} else {
					tName = normalizeResourceName(toolType, true)
				}

				toolWG.Add(1)
				go g.HandleTool(t, toolType, tID, tName, tcID, tcIDref, &toolWG)
			}
		}
	}
	toolWG.Wait()
	return nil
}

// Goroutine helper to collect worker IDs for TektonPipelinePostProcess
func (g *ToolchainGenerator) updateWorkerIDs(i int, res terraformutils.Resource, workerIDs map[string]string) {
	resID := g.Resources[i].InstanceState.ID
	wkrIDSplit := strings.Split(resID, "/")
	if resID == "" || len(wkrIDSplit) != 2 {
		return
	}
	workerID := wkrIDSplit[1]
	workerIDMutex.Lock()
	workerIDs[workerID] = res.InstanceInfo.ResourceAddress().String()
	workerIDMutex.Unlock()
}

// Goroutine helper to collect repos for TektonDefinitionPostProcess
func (g *ToolchainGenerator) updateRepos(i int, res terraformutils.Resource, repos map[string](map[string]string)) {
	params, ok := g.Resources[i].Item["parameters"].([]interface{})
	if !ok || len(params) == 0 {
		return
	}
	paramsMap, ok := params[0].(map[string]interface{})
	if !ok {
		return
	}
	if tcID, ok := g.Resources[i].InstanceState.Attributes["toolchain_id"]; ok {
		repoMutex.Lock()
		if repos[tcID] == nil {
			repos[tcID] = make(map[string]string)
		}
		repos[tcID][paramsMap["repo_url"].(string)] = res.InstanceInfo.ResourceAddress().String()
		repoMutex.Unlock()
	}
}

func (g *ToolchainGenerator) PostConvertHook() error {
	workerIDs := map[string]string{}
	repos := map[string](map[string]string){}
	tools := map[string]string{}

	var resWG sync.WaitGroup
	for i, res := range g.Resources {
		resWG.Add(1)
		go func() {
			defer resWG.Done()

			switch res.InstanceInfo.Type {
			case "ibm_cd_toolchain_tool_privateworker":
				g.updateWorkerIDs(i, res, workerIDs)
			case "ibm_cd_toolchain_tool_bitbucketgit":
				g.updateRepos(i, res, repos)
			case "ibm_cd_toolchain_tool_hostedgit":
				g.updateRepos(i, res, repos)
			case "ibm_cd_toolchain_tool_gitlab":
				g.updateRepos(i, res, repos)
			case "ibm_cd_toolchain_tool_githubconsolidated":
				g.updateRepos(i, res, repos)
			}

			// Collect tools for TektonPropertyPostProcess
			if strings.HasPrefix(res.InstanceInfo.Type, "ibm_cd_toolchain_tool_") {
				if tID, ok := g.Resources[i].InstanceState.Attributes["tool_id"]; ok {
					toolMutex.Lock()
					tools[tID] = res.InstanceInfo.ResourceAddress().String()
					toolMutex.Unlock()
				}
			}
		}()
	}
	resWG.Wait()

	for i, res := range g.Resources {
		switch res.InstanceInfo.Type {
		case "ibm_cd_tekton_pipeline":
			g.TektonPipelinePostProcess(i, res, workerIDs)

		case "ibm_cd_tekton_pipeline_definition":
			g.TektonDefinitionPostProcess(i, res, repos)

		case "ibm_cd_tekton_pipeline_property":
			g.TektonPropertyPostProcess(i, res, tools)

		case "ibm_cd_tekton_pipeline_trigger_property":
			g.TektonPropertyPostProcess(i, res, tools)

		case "ibm_cd_toolchain_tool_jenkins":
			g.JenkinsPostProcess(i, res)

		case "ibm_cd_toolchain_tool_bitbucketgit":
			g.GitRepositoryPostProcess(i, res)

		case "ibm_cd_toolchain_tool_hostedgit":
			g.GitRepositoryPostProcess(i, res)

		case "ibm_cd_toolchain_tool_gitlab":
			g.GitRepositoryPostProcess(i, res)

		case "ibm_cd_toolchain_tool_githubconsolidated":
			g.GitRepositoryPostProcess(i, res)
		}
	}

	return nil
}

// PostConvertHook helper to add private workers refs to tekton pipelines
func (g *ToolchainGenerator) TektonPipelinePostProcess(i int, res terraformutils.Resource, workerIDs map[string]string) {
	worker, ok := g.Resources[i].Item["worker"].([]interface{})
	if !ok {
		return
	}
	workerMap, ok := worker[0].(map[string]interface{})
	if !ok {
		return
	}
	plWorkerID := workerMap["id"]
	if plWorkerID == nil || plWorkerID == "public" {
		return
	}
	if wkr, ok := workerIDs[plWorkerID.(string)]; ok {
		workerMap["id"] = fmt.Sprintf("${%s.tool_id}", wkr)
		return
	}
}

// PostConvertHook helper to add repo depends_on to tekton pipeline definitions
func (g *ToolchainGenerator) TektonDefinitionPostProcess(i int, res terraformutils.Resource, repos map[string](map[string]string)) {
	defSource, ok := g.Resources[i].Item["source"].([]interface{})
	if !ok || len(defSource) == 0 {
		return
	}
	defSourceMap, ok := defSource[0].(map[string]interface{})
	if !ok {
		return
	}
	defProps, ok := defSourceMap["properties"].([]interface{})
	if !ok || len(defProps) == 0 {
		return
	}
	defPropsMap, ok := defProps[0].(map[string]interface{})
	if !ok {
		return
	}
	tcID, ok := g.Resources[i].Item["toolchain_id_actual"]
	if !ok {
		return
	}
	if repo, ok := repos[tcID.(string)][defPropsMap["url"].(string)]; ok {
		g.Resources[i].Item["depends_on"] = []string{repo}
	}
	delete(g.Resources[i].Item, "toolchain_id_actual")
}

// PostConvertHook helper to add tool refs to tekton pipeline properties and additional escape appconfig substitution
func (g *ToolchainGenerator) TektonPropertyPostProcess(i int, res terraformutils.Resource, tools map[string]string) {
	target, ok := g.Resources[i].Item["value"].(string)
	if !ok {
		return
	}

	// escape appconfig values -- ${...} is interpreted as a template in terraform
	g.Resources[i].Item["value"] = strings.ReplaceAll(g.Resources[i].Item["value"].(string), "${", "$${")

	// add tool integration ref to tekton definitions
	if g.Resources[i].Item["type"] != "integration" {
		return
	}

	if tool, ok := tools[target]; ok {
		g.Resources[i].Item["value"] = fmt.Sprintf("${%s.tool_id}", tool)
		return
	}
	fmt.Println("......! Could not link pipeline property of type integration:", res.InstanceInfo.ResourceAddress().String())
}

// PostConvertHook helper to remove Jenkins webhook_url from tf files, which is supposed to be sensitive and computed
func (g *ToolchainGenerator) JenkinsPostProcess(i int, res terraformutils.Resource) {
	params, ok := g.Resources[i].Item["parameters"].([]interface{})
	if !ok || len(params) == 0 {
		return
	}
	paramsMap, ok := params[0].(map[string]interface{})
	if !ok {
		return
	}
	if _, ok := paramsMap["webhook_url"].(string); ok {
		delete(paramsMap, "webhook_url")
	}
}

// PostConvertHook helper to generate initialization block and remove computed values for git repo resources
func (g *ToolchainGenerator) GitRepositoryPostProcess(i int, res terraformutils.Resource) {
	// Handle Initialization args
	params, ok := g.Resources[i].Item["parameters"].([]interface{})
	if !ok || len(params) == 0 {
		return
	}
	paramsMap, ok := params[0].(map[string]interface{})
	if !ok {
		return
	}
	initMap := map[string]interface{}{}

	initMap["git_id"] = paramsMap["git_id"]
	initMap["type"] = paramsMap["type"] // this will always be "link"
	initMap["repo_url"] = paramsMap["repo_url"]
	initMap["private_repo"] = paramsMap["private_repo"]

	// additional parameters
	if res.InstanceInfo.Type == "ibm_cd_toolchain_tool_githubconsolidated" {
		initMap["blind_connection"] = paramsMap["blind_connection"]
		initMap["auto_init"] = paramsMap["auto_init"]
	} else if res.InstanceInfo.Type == "ibm_cd_toolchain_tool_gitlab" {
		initMap["blind_connection"] = paramsMap["blind_connection"]
	}

	// add to initialization accordingly
	g.Resources[i].Item["initialization"] = initMap

	// add missing initialization to terraform state attributes
	g.Resources[i].InstanceState.Attributes["initialization.#"] = "1"
	for key, val := range initMap {
		g.Resources[i].InstanceState.Attributes["initialization.0."+key] = val.(string)
	}

	// only include non-computed parameters
	includeParams := []string{"api_token", "auth_type", "enable_traceability", "integration_owner", "toolchain_issues_enabled"}

	for key := range paramsMap {
		if !slices.Contains(includeParams, key) {
			delete(g.Resources[i].Item["parameters"].([]interface{})[0].(map[string]interface{}), key)
			delete(g.Resources[i].InstanceState.Attributes, key)
		}
	}
}
