package octopusdeploy

import (
	"fmt"

	"github.com/dghubble/sling"
)

type ProjectTriggerService struct {
	sling *sling.Sling
}

func NewProjectTriggerService(sling *sling.Sling) *ProjectTriggerService {
	return &ProjectTriggerService{
		sling: sling,
	}
}

type ProjectTriggers struct {
	Items []ProjectTrigger `json:"Items"`
	PagedResults
}

type ProjectTrigger struct {
	Action     ProjectTriggerAction `json:"Action"`
	Filter     ProjectTriggerFilter `json:"Filter"`
	ID         string               `json:"Id,omitempty"`
	IsDisabled bool                 `json:"IsDisabled,omitempty"`
	Name       string               `json:"Name"`
	ProjectID  string               `json:"ProjectId,omitempty"`
}

type ProjectTriggerFilter struct {
	DateOfMonth         string   `json:"DateOfMonth"`
	DayNumberOfMonth    string   `json:"DayNumberOfMonth"`
	DayOfWeek           string   `json:"DayOfWeek"`
	EnvironmentIds      []string `json:"EnvironmentIds,omitempty"`
	EventCategories     []string `json:"EventCategories,omitempty"`
	EventGroups         []string `json:"EventGroups,omitempty"`
	FilterType          string   `json:"FilterType"`
	MonthlyScheduleType string   `json:"MonthlyScheduleType"`
	Roles               []string `json:"Roles"`
	StartTime           string   `json:"StartTime"`
	Timezone            string   `json:"Timezone"`
}

type ProjectTriggerAction struct {
	ActionType                                 string `json:"ActionType"`
	DestinationEnvironmentID                   string `json:"DestinationEnvironmentId"`
	ShouldRedeployWhenMachineHasBeenDeployedTo bool   `json:"ShouldRedeployWhenMachineHasBeenDeployedTo"`
	SourceEnvironmentID                        string `json:"SourceEnvironmentId"`
}

func (t *ProjectTrigger) AddEventGroups(eventGroups []string) {
	for _, e := range eventGroups {
		t.Filter.EventGroups = append(t.Filter.EventGroups, e)
	}
}

func (t *ProjectTrigger) AddEventCategories(eventCategories []string) {
	for _, e := range eventCategories {
		t.Filter.EventCategories = append(t.Filter.EventCategories, e)
	}
}

func NewProjectDeploymentTargetTrigger(name, projectID string, shouldRedeploy bool, roles, eventGroups, eventCategories []string) *ProjectTrigger {
	return &ProjectTrigger{
		Action: ProjectTriggerAction{
			ActionType: "AutoDeploy",
			ShouldRedeployWhenMachineHasBeenDeployedTo: shouldRedeploy,
		},
		Filter: ProjectTriggerFilter{
			EventCategories: eventCategories,
			EventGroups:     eventGroups,
			FilterType:      "MachineFilter",
			Roles:           roles,
		},
		Name:      name,
		ProjectID: projectID,
	}
}

func (s *ProjectTriggerService) Get(projectTriggerID string) (*ProjectTrigger, error) {
	path := fmt.Sprintf("projecttriggers/%s", projectTriggerID)

	resp, err := apiGet(s.sling, new(ProjectTrigger), path)

	if err != nil {
		return nil, err
	}

	return resp.(*ProjectTrigger), nil
}

func (s *ProjectTriggerService) GetByProjectID(projectID string) (*[]ProjectTrigger, error) {
	var triggersByProject []ProjectTrigger

	triggers, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, pt := range *triggers {
		triggersByProject = append(triggersByProject, pt)
	}

	return &triggersByProject, nil
}

func (s *ProjectTriggerService) GetAll() (*[]ProjectTrigger, error) {
	var pt []ProjectTrigger

	path := "projecttriggers"

	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(ProjectTriggers), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*ProjectTriggers)

		for _, item := range r.Items {
			pt = append(pt, item)
		}

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &pt, nil
}

func (s *ProjectTriggerService) Add(projectTrigger *ProjectTrigger) (*ProjectTrigger, error) {
	resp, err := apiAdd(s.sling, projectTrigger, new(ProjectTrigger), "projecttriggers")

	if err != nil {
		return nil, err
	}

	return resp.(*ProjectTrigger), nil
}

func (s *ProjectTriggerService) Delete(projectTriggerID string) error {
	path := fmt.Sprintf("projecttriggers/%s", projectTriggerID)
	err := apiDelete(s.sling, path)

	if err != nil {
		return err
	}

	return nil
}

func (s *ProjectTriggerService) Update(projectTrigger *ProjectTrigger) (*ProjectTrigger, error) {
	path := fmt.Sprintf("projecttriggers/%s", projectTrigger.ID)
	resp, err := apiUpdate(s.sling, projectTrigger, new(ProjectTrigger), path)

	if err != nil {
		return nil, err
	}

	return resp.(*ProjectTrigger), nil
}
