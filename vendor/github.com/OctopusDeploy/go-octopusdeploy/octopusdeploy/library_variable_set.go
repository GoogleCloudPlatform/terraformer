package octopusdeploy

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"net/url"

	"github.com/dghubble/sling"
)

type LibraryVariableSetService struct {
	sling *sling.Sling
}

func NewLibraryVariableSetService(sling *sling.Sling) *LibraryVariableSetService {
	return &LibraryVariableSetService{
		sling: sling,
	}
}

type LibraryVariableSets struct {
	Items []LibraryVariableSet `json:"Items"`
	PagedResults
}

type LibraryVariableSet struct {
	ID            string                    `json:"Id,omitempty"`
	Name          string                    `json:"Name" validate:"required"`
	Description   string                    `json:"Description,omitempty"`
	VariableSetId string                    `json:"VariableSetId,omitempty"`
	ContentType   VariableSetContentType    `json:"ContentType" validate:"required"`
	Templates     []ActionTemplateParameter `json:"Templates,omitempty"`
}

type VariableSetContentType string

const (
	VariableSetContentType_Variables    = VariableSetContentType("Variables")
	VariableSetContentType_ScriptModule = VariableSetContentType("ScriptModule")
)

func NewLibraryVariableSet(name string) *LibraryVariableSet {
	return &LibraryVariableSet{
		Name:        name,
		ContentType: VariableSetContentType_Variables,
	}
}

// ValidateLibraryVariableSetValues checks the values of a LibraryVariableSet object to see if they are suitable for
// sending to Octopus Deploy. Used when adding or updating libraryVariableSets.
func ValidateLibraryVariableSetValues(LibraryVariableSet *LibraryVariableSet) error {
	validate := validator.New()
	err := validate.Struct(LibraryVariableSet)
	return err
}

// Get returns a single LibraryVariableSet by its Id in Octopus Deploy
func (s *LibraryVariableSetService) Get(libraryVariableSetID string) (*LibraryVariableSet, error) {
	path := fmt.Sprintf("libraryVariableSets/%s", libraryVariableSetID)
	resp, err := apiGet(s.sling, new(LibraryVariableSet), path)

	if err != nil {
		return nil, err
	}

	return resp.(*LibraryVariableSet), nil
}

// GetAll returns all libraryVariableSets in Octopus Deploy
func (s *LibraryVariableSetService) GetAll() (*[]LibraryVariableSet, error) {
	return s.get("")
}

func (s *LibraryVariableSetService) get(query string) (*[]LibraryVariableSet, error) {
	var p []LibraryVariableSet

	path := "libraryvariablesets?take=2147483647"
	if query != "" {
		path = fmt.Sprintf("%s&%s", path, query)
	}

	loadNextPage := true

	for loadNextPage { // Older Octopus Servers do not accept the take parameter, so the only choice is to page through them
		resp, err := apiGet(s.sling, new(LibraryVariableSets), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*LibraryVariableSets)

		for _, item := range r.Items {
			p = append(p, item)
		}

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

// GetByName gets an existing Library Variable Set by its name in Octopus Deploy
func (s *LibraryVariableSetService) GetByName(name string) (*LibraryVariableSet, error) {
	var foundLibraryVariableSet LibraryVariableSet
	libraryVariableSets, err := s.get(fmt.Sprintf("partialName=%s", url.PathEscape(name)))

	if err != nil {
		return nil, err
	}

	for _, libraryVariableSet := range *libraryVariableSets {
		if libraryVariableSet.Name == name {
			return &libraryVariableSet, nil
		}
	}

	return &foundLibraryVariableSet, fmt.Errorf("no Library Variable Set found with name %s", name)
}

// Add adds an new libraryVariableSet in Octopus Deploy
func (s *LibraryVariableSetService) Add(libraryVariableSet *LibraryVariableSet) (*LibraryVariableSet, error) {
	err := ValidateLibraryVariableSetValues(libraryVariableSet)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, libraryVariableSet, new(LibraryVariableSet), "libraryVariableSets")

	if err != nil {
		return nil, err
	}

	return resp.(*LibraryVariableSet), nil
}

// Delete deletes an existing libraryVariableSet in Octopus Deploy
func (s *LibraryVariableSetService) Delete(libraryVariableSetid string) error {
	path := fmt.Sprintf("libraryVariableSets/%s", libraryVariableSetid)
	err := apiDelete(s.sling, path)

	if err != nil {
		return err
	}

	return nil
}

// Update updates an existing libraryVariableSet in Octopus Deploy
func (s *LibraryVariableSetService) Update(libraryVariableSet *LibraryVariableSet) (*LibraryVariableSet, error) {
	err := ValidateLibraryVariableSetValues(libraryVariableSet)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("libraryVariableSets/%s", libraryVariableSet.ID)
	resp, err := apiUpdate(s.sling, libraryVariableSet, new(LibraryVariableSet), path)

	if err != nil {
		return nil, err
	}

	return resp.(*LibraryVariableSet), nil
}
