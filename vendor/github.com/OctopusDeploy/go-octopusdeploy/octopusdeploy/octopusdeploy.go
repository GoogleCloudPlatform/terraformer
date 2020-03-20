package octopusdeploy

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dghubble/sling"
)

// Client is an OctopusDeploy for making OctpusDeploy API requests.
type Client struct {
	sling *sling.Sling
	// Octopus Deploy API Services
	Account            *AccountService
	Certificate        *CertificateService
	DeploymentProcess  *DeploymentProcessService
	ProjectGroup       *ProjectGroupService
	Project            *ProjectService
	ProjectTrigger     *ProjectTriggerService
	Environment        *EnvironmentService
	Feed               *FeedService
	Variable           *VariableService
	MachinePolicy      *MachinePolicyService
	Machine            *MachineService
	Lifecycle          *LifecycleService
	LibraryVariableSet *LibraryVariableSetService
	Interruption       *InterruptionsService
	TagSet             *TagSetService
	Space              *SpaceService
	Channel            *ChannelService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client, octopusURL, octopusAPIKey string) *Client {
	baseURLWithAPI := strings.TrimRight(octopusURL, "/")
	baseURLWithAPI = fmt.Sprintf("%s/api/", baseURLWithAPI)
	base := sling.New().Client(httpClient).Base(baseURLWithAPI).Set("X-Octopus-ApiKey", octopusAPIKey)
	return &Client{
		sling:              base,
		Account:            NewAccountService(base.New()),
		Certificate:        NewCertificateService(base.New()),
		DeploymentProcess:  NewDeploymentProcessService(base.New()),
		ProjectGroup:       NewProjectGroupService(base.New()),
		Project:            NewProjectService(base.New()),
		ProjectTrigger:     NewProjectTriggerService(base.New()),
		Environment:        NewEnvironmentService(base.New()),
		Feed:               NewFeedService(base.New()),
		Variable:           NewVariableService(base.New()),
		MachinePolicy:      NewMachinePolicyService(base.New()),
		Machine:            NewMachineService(base.New()),
		Lifecycle:          NewLifecycleService(base.New()),
		LibraryVariableSet: NewLibraryVariableSetService(base.New()),
		Interruption:       NewInterruptionService(base.New()),
		TagSet:             NewTagSetService(base.New()),
		Space:              NewSpaceService(base.New()),
		Channel:            NewChannelService(base.New()),
	}
}

func ForSpace(httpClient *http.Client, octopusURL, octopusAPIKey string, space *Space) *Client {
	baseURLWithAPI := strings.TrimRight(octopusURL, "/")
	baseURLWithAPI = fmt.Sprintf("%s/api/%s/", baseURLWithAPI, space.ID)
	base := sling.New().Client(httpClient).Base(baseURLWithAPI).Set("X-Octopus-ApiKey", octopusAPIKey)
	return &Client{
		sling:              base,
		Account:            NewAccountService(base.New()),
		Certificate:        NewCertificateService(base.New()),
		DeploymentProcess:  NewDeploymentProcessService(base.New()),
		ProjectGroup:       NewProjectGroupService(base.New()),
		Project:            NewProjectService(base.New()),
		ProjectTrigger:     NewProjectTriggerService(base.New()),
		Environment:        NewEnvironmentService(base.New()),
		Feed:               NewFeedService(base.New()),
		Variable:           NewVariableService(base.New()),
		MachinePolicy:      NewMachinePolicyService(base.New()),
		Machine:            NewMachineService(base.New()),
		Lifecycle:          NewLifecycleService(base.New()),
		LibraryVariableSet: NewLibraryVariableSetService(base.New()),
		TagSet:             NewTagSetService(base.New()),
		Channel:            NewChannelService(base.New()),
	}
}

type APIError struct {
	ErrorMessage  string   `json:"ErrorMessage"`
	Errors        []string `json:"Errors"`
	FullException string   `json:"FullException"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("Octopus Deploy Error Response: %v %+v %v", e.ErrorMessage, e.Errors, e.FullException)
}

// APIErrorChecker is a generic error handler for the OctopusDeploy API.
func APIErrorChecker(urlPath string, resp *http.Response, wantedResponseCode int, slingError error, octopusDeployError *APIError) error {
	if octopusDeployError.Errors != nil {
		return fmt.Errorf("octopus deploy api returned an error on endpoint %s - %s", urlPath, octopusDeployError.Errors)
	}

	if slingError != nil {
		return fmt.Errorf("cannot get endpoint %s from server. failure from http client %v", urlPath, slingError)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return ErrItemNotFound
	}

	if resp.StatusCode != wantedResponseCode {
		return fmt.Errorf("cannot get item from endpoint %s. response from server %s", urlPath, resp.Status)
	}

	return nil
}

// LoadNextPage checks if the next page should be loaded from the API. Returns the new path and a bool if the next page should be checked.
func LoadNextPage(pagedResults PagedResults) (string, bool) {
	if pagedResults.Links.PageNext != "" {
		return pagedResults.Links.PageNext, true
	}

	return "", false
}

// Generic OctopusDeploy API Get Function.
func apiGet(sling *sling.Sling, inputStruct interface{}, path string) (interface{}, error) {
	octopusDeployError := new(APIError)

	resp, err := sling.New().Get(path).Receive(inputStruct, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return inputStruct, nil
}

// Generic OctopusDeploy API Add Function. Expects a 201 response.
func apiAdd(sling *sling.Sling, inputStruct, returnStruct interface{}, path string) (interface{}, error) {
	octopusDeployError := new(APIError)

	resp, err := sling.New().Post(path).BodyJSON(inputStruct).Receive(returnStruct, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusCreated, err, octopusDeployError)

	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return returnStruct, nil
}

// apiPost post to octopus and expect a 200 response code.
func apiPost(sling *sling.Sling, inputStruct, returnStruct interface{}, path string) (interface{}, error) {
	octopusDeployError := new(APIError)

	resp, err := sling.New().Post(path).BodyJSON(inputStruct).Receive(returnStruct, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return returnStruct, nil
}

// Generic OctopusDeploy API Update Function.
func apiUpdate(sling *sling.Sling, inputStruct, returnStruct interface{}, path string) (interface{}, error) {
	octopusDeployError := new(APIError)

	resp, err := sling.New().Put(path).BodyJSON(inputStruct).Receive(returnStruct, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return returnStruct, nil
}

// Generic OctopusDeploy API Delete Function.
func apiDelete(sling *sling.Sling, path string) error {
	octopusDeployError := new(APIError)

	resp, err := sling.New().Delete(path).Receive(nil, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

	if apiErrorCheck != nil {
		return apiErrorCheck
	}

	return nil
}

// ErrItemNotFound is an OctopusDeploy error returned an item cannot be found.
var ErrItemNotFound = errors.New("cannot find the item")
