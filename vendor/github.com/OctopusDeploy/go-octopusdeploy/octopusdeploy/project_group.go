package octopusdeploy

import (
	"fmt"

	"github.com/dghubble/sling"
	"gopkg.in/go-playground/validator.v9"
)

type ProjectGroupService struct {
	sling *sling.Sling
}

func NewProjectGroupService(sling *sling.Sling) *ProjectGroupService {
	return &ProjectGroupService{
		sling: sling,
	}
}

type ProjectGroups struct {
	Items []ProjectGroup `json:"Items"`
	PagedResults
}

type ProjectGroup struct {
	Description       string   `json:"Description,omitempty"`
	EnvironmentIds    []string `json:"EnvironmentIds"`
	ID                string   `json:"Id,omitempty"`
	LastModifiedBy    string   `json:"LastModifiedBy,omitempty"`
	LastModifiedOn    string   `json:"LastModifiedOn,omitempty"`
	Links             Links    `json:"Links,omitempty"`
	Name              string   `json:"Name,omitempty" validate:"required"`
	RetentionPolicyID string   `json:"RetentionPolicyId,omitempty"`
}

func (p *ProjectGroup) Validate() error {
	validate := validator.New()

	err := validate.Struct(p)

	if err != nil {
		return err
	}

	return nil
}

func NewProjectGroup(name string) *ProjectGroup {
	return &ProjectGroup{
		Name: name,
	}
}

func (s *ProjectGroupService) Get(projectGroupID string) (*ProjectGroup, error) {
	path := fmt.Sprintf("projectgroups/%s", projectGroupID)
	resp, err := apiGet(s.sling, new(ProjectGroup), path)

	if err != nil {
		return nil, err
	}

	return resp.(*ProjectGroup), nil
}

func (s *ProjectGroupService) GetAll() (*[]ProjectGroup, error) {
	var pg []ProjectGroup

	path := "projectgroups"

	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(ProjectGroups), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*ProjectGroups)

		for _, item := range r.Items {
			pg = append(pg, item)
		}

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &pg, nil
}

func (s *ProjectGroupService) Add(projectGroup *ProjectGroup) (*ProjectGroup, error) {
	resp, err := apiAdd(s.sling, projectGroup, new(ProjectGroup), "projectgroups")

	if err != nil {
		return nil, err
	}

	return resp.(*ProjectGroup), nil
}

func (s *ProjectGroupService) Delete(projectGroupID string) error {
	path := fmt.Sprintf("projectgroups/%s", projectGroupID)
	err := apiDelete(s.sling, path)

	if err != nil {
		return err
	}

	return nil
}

func (s *ProjectGroupService) Update(projectGroup *ProjectGroup) (*ProjectGroup, error) {
	path := fmt.Sprintf("projectgroups/%s", projectGroup.ID)
	resp, err := apiUpdate(s.sling, projectGroup, new(ProjectGroup), path)

	if err != nil {
		return nil, err
	}

	return resp.(*ProjectGroup), nil
}
