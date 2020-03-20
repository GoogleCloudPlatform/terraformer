package octopusdeploy

import (
	"fmt"

	"github.com/dghubble/sling"
	"gopkg.in/go-playground/validator.v9"
)

type EnvironmentService struct {
	sling *sling.Sling
}

func NewEnvironmentService(sling *sling.Sling) *EnvironmentService {
	return &EnvironmentService{
		sling: sling,
	}
}

type Environments struct {
	Items []Environment `json:"Items"`
	PagedResults
}

type Environment struct {
	ID                         string `json:"Id"`
	Name                       string `json:"Name"`
	Description                string `json:"Description"`
	SortOrder                  int    `json:"SortOrder"`
	UseGuidedFailure           bool   `json:"UseGuidedFailure"`
	AllowDynamicInfrastructure bool   `json:"AllowDynamicInfrastructure"`
}

func (t *Environment) Validate() error {
	validate := validator.New()

	err := validate.Struct(t)

	if err != nil {
		return err
	}

	return nil
}

func NewEnvironment(name, description string, useguidedfailure bool) *Environment {
	return &Environment{
		Name:             name,
		Description:      description,
		UseGuidedFailure: useguidedfailure,
	}
}

func (s *EnvironmentService) Get(environmentid string) (*Environment, error) {
	path := fmt.Sprintf("environments/%s", environmentid)
	resp, err := apiGet(s.sling, new(Environment), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Environment), nil
}

func (s *EnvironmentService) GetAll() (*[]Environment, error) {
	var p []Environment

	path := "environments"

	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(Environments), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*Environments)

		for _, item := range r.Items {
			p = append(p, item)
		}

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

func (s *EnvironmentService) GetByName(environmentName string) (*Environment, error) {
	var foundEnvironment Environment
	environments, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, project := range *environments {
		if project.Name == environmentName {
			return &project, nil
		}
	}

	return &foundEnvironment, fmt.Errorf("no environment found with environment name %s", environmentName)
}

func (s *EnvironmentService) Add(environment *Environment) (*Environment, error) {
	resp, err := apiAdd(s.sling, environment, new(Environment), "environments")

	if err != nil {
		return nil, err
	}

	return resp.(*Environment), nil
}

func (s *EnvironmentService) Delete(environmentid string) error {
	path := fmt.Sprintf("environments/%s", environmentid)
	err := apiDelete(s.sling, path)

	if err != nil {
		return err
	}

	return nil
}

func (s *EnvironmentService) Update(environment *Environment) (*Environment, error) {
	path := fmt.Sprintf("environments/%s", environment.ID)
	resp, err := apiUpdate(s.sling, environment, new(Environment), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Environment), nil
}
