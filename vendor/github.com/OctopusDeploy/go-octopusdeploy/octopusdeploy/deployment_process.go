package octopusdeploy

import (
	"fmt"

	"github.com/dghubble/sling"
	"gopkg.in/go-playground/validator.v9"
)

type DeploymentProcessService struct {
	sling *sling.Sling
}

func NewDeploymentProcessService(sling *sling.Sling) *DeploymentProcessService {
	return &DeploymentProcessService{
		sling: sling,
	}
}

type DeploymentProcesses struct {
	Items []DeploymentProcess `json:"Items"`
	PagedResults
}

type DeploymentProcess struct {
	ID             string           `json:"Id,omitempty"`
	LastModifiedBy string           `json:"LastModifiedBy,omitempty"`
	LastModifiedOn string           `json:"LastModifiedOn,omitempty"`
	LastSnapshotID string           `json:"LastSnapshotId,omitempty"`
	Links          Links            `json:"Links,omitempty"`
	ProjectID      string           `json:"ProjectId,omitempty"`
	Steps          []DeploymentStep `json:"Steps,omitempty"`
	Version        int32            `json:"Version"`
}

type DeploymentStep struct {
	ID                 string                           `json:"Id,omitempty"`
	Name               string                           `json:"Name"`
	PackageRequirement DeploymentStepPackageRequirement `json:"PackageRequirement,omitempty"`                                         // may need its own model / enum
	Properties         map[string]string                `json:"Properties"`                                                           // TODO: refactor to use the PropertyValueResource for handling sensitive values - https://blog.gopheracademy.com/advent-2016/advanced-encoding-decoding/
	Condition          DeploymentStepCondition          `json:"Condition,omitempty" validate:"oneof=Success Failure Always Variable"` // variable option adds a Property "Octopus.Action.ConditionVariableExpression"
	StartTrigger       DeploymentStepStartTrigger       `json:"StartTrigger,omitempty" validate:"oneof=StartAfterPrevious StartWithPrevious"`
	Actions            []DeploymentAction               `json:"Actions,omitempty"`
}

type DeploymentAction struct {
	ID                            string             `json:"Id,omitempty"`
	Name                          string             `json:"Name"`
	ActionType                    string             `json:"ActionType"`
	IsDisabled                    bool               `json:"IsDisabled"`
	IsRequired                    bool               `json:"IsRequired"`
	WorkerPoolId                  string             `json:"WorkerPoolId,omitempty"`
	CanBeUsedForProjectVersioning bool               `json:"CanBeUsedForProjectVersioning"`
	Environments                  []string           `json:"Environments,omitempty"`
	ExcludedEnvironments          []string           `json:"ExcludedEnvironments,omitempty"`
	Channels                      []string           `json:"Channels,omitempty"`
	TenantTags                    []string           `json:"TenantTags,omitempty"`
	Properties                    map[string]string  `json:"Properties"` // TODO: refactor to use the PropertyValueResource for handling sensitive values - https://blog.gopheracademy.com/advent-2016/advanced-encoding-decoding/
	Packages                      []PackageReference `json:"Packages,omitempty"`
}

type DeploymentStepPackageRequirement string

const (
	DeploymentStepPackageRequirement_LetOctopusDecide         = DeploymentStepPackageRequirement("LetOctopusDecide")
	DeploymentStepPackageRequirement_BeforePackageAcquisition = DeploymentStepPackageRequirement("BeforePackageAcquisition")
	DeploymentStepPackageRequirement_AfterPackageAcquisition  = DeploymentStepPackageRequirement("AfterPackageAcquisition")
)

type DeploymentStepCondition string

const (
	DeploymentStepCondition_Success  = DeploymentStepCondition("Success")
	DeploymentStepCondition_Failure  = DeploymentStepCondition("Failure")
	DeploymentStepCondition_Always   = DeploymentStepCondition("Always")
	DeploymentStepCondition_Variable = DeploymentStepCondition("Variable")
)

type DeploymentStepStartTrigger string

const (
	DeploymentStepStartTrigger_StartAfterPrevious = DeploymentStepStartTrigger("StartAfterPrevious")
	DeploymentStepStartTrigger_StartWithPrevious  = DeploymentStepStartTrigger("StartWithPrevious")
)

type PackageReference struct {
	ID                  string            `json:"Id,omitempty"`
	Name                string            `json:"Name,omitempty"`
	PackageId           string            `json:"PackageId,omitempty"`
	FeedId              string            `json:"FeedId"`
	AcquisitionLocation string            `json:"AcquisitionLocation"` // This can be an expression
	Properties          map[string]string `json:"Properties"`
}

const (
	PackageAcquisitionLocation_Server          = "Server"
	PackageAcquisitionLocation_ExecutionTarget = "ExecutionTarget"
	PackageAcquisitionLocation_NotAcquired     = "NotAcquired"
)

func (d *DeploymentProcess) Validate() error {
	validate := validator.New()

	err := validate.Struct(d)

	if err != nil {
		return err
	}

	return nil
}

func (s *DeploymentProcessService) Get(deploymentProcessID string) (*DeploymentProcess, error) {
	path := fmt.Sprintf("deploymentprocesses/%s", deploymentProcessID)
	resp, err := apiGet(s.sling, new(DeploymentProcess), path)

	if err != nil {
		return nil, err
	}

	return resp.(*DeploymentProcess), nil
}

func (s *DeploymentProcessService) GetAll() (*[]DeploymentProcess, error) {
	var dp []DeploymentProcess

	path := "deploymentprocesses"

	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(DeploymentProcesses), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*DeploymentProcesses)

		for _, item := range r.Items {
			dp = append(dp, item)
		}

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &dp, nil
}

func (s *DeploymentProcessService) Update(deploymentProcess *DeploymentProcess) (*DeploymentProcess, error) {
	path := fmt.Sprintf("deploymentprocesses/%s", deploymentProcess.ID)
	resp, err := apiUpdate(s.sling, deploymentProcess, new(DeploymentProcess), path)

	if err != nil {
		return nil, err
	}

	return resp.(*DeploymentProcess), nil
}
