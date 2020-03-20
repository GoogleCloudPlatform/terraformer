package octopusdeploy

import (
	"fmt"

	"github.com/dghubble/sling"
)

type ProjectService struct {
	sling *sling.Sling
}

func NewProjectService(sling *sling.Sling) *ProjectService {
	return &ProjectService{
		sling: sling,
	}
}

type Projects struct {
	Items []Project `json:"Items"`
	PagedResults
}

type Project struct {
	AutoCreateRelease               bool                        `json:"AutoCreateRelease"`
	AutoDeployReleaseOverrides      []AutoDeployReleaseOverride `json:"AutoDeployReleaseOverrides"`
	DefaultGuidedFailureMode        string                      `json:"DefaultGuidedFailureMode,omitempty"`
	DefaultToSkipIfAlreadyInstalled bool                        `json:"DefaultToSkipIfAlreadyInstalled"`
	DeploymentProcessID             string                      `json:"DeploymentProcessId"`
	Description                     string                      `json:"Description"`
	DiscreteChannelRelease          bool                        `json:"DiscreteChannelRelease"`
	ID                              string                      `json:"Id,omitempty"`
	IncludedLibraryVariableSetIds   []string                    `json:"IncludedLibraryVariableSetIds"`
	IsDisabled                      bool                        `json:"IsDisabled"`
	LifecycleID                     string                      `json:"LifecycleId"`
	Name                            string                      `json:"Name"`
	ProjectConnectivityPolicy       ProjectConnectivityPolicy   `json:"ProjectConnectivityPolicy"`
	ProjectGroupID                  string                      `json:"ProjectGroupId"`
	ReleaseCreationStrategy         ReleaseCreationStrategy     `json:"ReleaseCreationStrategy"`
	Slug                            string                      `json:"Slug"`
	Templates                       []ActionTemplateParameter   `json:"Templates,omitempty"`
	TenantedDeploymentMode          TenantedDeploymentMode      `json:"TenantedDeploymentMode,omitempty"`
	VariableSetID                   string                      `json:"VariableSetId"`
	VersioningStrategy              VersioningStrategy          `json:"VersioningStrategy"`
}

func NewProject(name, lifeCycleID, projectGroupID string) *Project {
	return &Project{
		Name:                     name,
		DefaultGuidedFailureMode: "EnvironmentDefault",
		LifecycleID:              lifeCycleID,
		ProjectGroupID:           projectGroupID,
		VersioningStrategy: VersioningStrategy{
			Template: "#{Octopus.Version.LastMajor}.#{Octopus.Version.LastMinor}.#{Octopus.Version.NextPatch}",
		},
		ProjectConnectivityPolicy: ProjectConnectivityPolicy{
			AllowDeploymentsToNoTargets: false,
			SkipMachineBehavior:         "None",
		},
	}
}

// ValidateProjectValues checks the values of a Project object to see if they are suitable for
// sending to Octopus Deploy. Used when adding or updating projects.
func ValidateProjectValues(Project *Project) error {
	return ValidateMultipleProperties([]error{
		ValidatePropertyValues("SkipMachineBehavior", Project.ProjectConnectivityPolicy.SkipMachineBehavior, ValidProjectConnectivityPolicySkipMachineBehaviors),
		ValidatePropertyValues("DefaultGuidedFailureMode", Project.DefaultGuidedFailureMode, ValidProjectDefaultGuidedFailureModes),
		ValidateRequiredPropertyValue("LifecycleID", Project.LifecycleID),
		ValidateRequiredPropertyValue("Name", Project.Name),
		ValidateRequiredPropertyValue("ProjectGroupID", Project.ProjectGroupID),
	})
}

// Get returns a single project by its projectid in Octopus Deploy
func (s *ProjectService) Get(projectid string) (*Project, error) {
	path := fmt.Sprintf("projects/%s", projectid)
	resp, err := apiGet(s.sling, new(Project), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Project), nil
}

// GetAll returns all projects in Octopus Deploy
func (s *ProjectService) GetAll() (*[]Project, error) {
	var p []Project

	path := "projects"

	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(Projects), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*Projects)

		for _, item := range r.Items {
			p = append(p, item)
		}

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

// GetByName gets an existing project by its project name in Octopus Deploy
func (s *ProjectService) GetByName(projectName string) (*Project, error) {
	var foundProject Project
	projects, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, project := range *projects {
		if project.Name == projectName {
			return &project, nil
		}
	}

	return &foundProject, fmt.Errorf("no project found with project name %s", projectName)
}

// Add adds an new project in Octopus Deploy
func (s *ProjectService) Add(project *Project) (*Project, error) {
	err := ValidateProjectValues(project)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, project, new(Project), "projects")

	if err != nil {
		return nil, err
	}

	return resp.(*Project), nil
}

// Delete deletes an existing project in Octopus Deploy
func (s *ProjectService) Delete(projectid string) error {
	path := fmt.Sprintf("projects/%s", projectid)
	err := apiDelete(s.sling, path)

	if err != nil {
		return err
	}

	return nil
}

// Update updates an existing project in Octopus Deploy
func (s *ProjectService) Update(project *Project) (*Project, error) {
	err := ValidateProjectValues(project)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("projects/%s", project.ID)
	resp, err := apiUpdate(s.sling, project, new(Project), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Project), nil
}
