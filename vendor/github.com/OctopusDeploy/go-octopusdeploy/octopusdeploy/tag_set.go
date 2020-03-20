package octopusdeploy

import (
	"fmt"

	"github.com/dghubble/sling"
	"gopkg.in/go-playground/validator.v9"
)

type TagSetService struct {
	sling *sling.Sling
}

func NewTagSetService(sling *sling.Sling) *TagSetService {
	return &TagSetService{
		sling: sling,
	}
}

type TagSets struct {
	Items []TagSet `json:"Items"`
	PagedResults
}

type TagSet struct {
	ID   string `json:"Id"`
	Name string `json:"Name"`
	Tags []Tag  `json:"Tags,omitempty"`
}

type Tag struct {
	ID               string `json:"Id"`
	Name             string `json:"Name"`
	Color            string `json:"Color"`
	CanonicalTagName string `json:"CanonicalTagName"`
	Description      string `json:"Description"`
	SortOrder        int    `json:"SortOrder"`
}

func (t *TagSet) Validate() error {
	validate := validator.New()

	err := validate.Struct(t)

	if err != nil {
		return err
	}

	return nil
}

func NewTagSet(name string) *TagSet {
	return &TagSet{
		Name: name,
	}
}

func (s *TagSetService) Get(tagSetId string) (*TagSet, error) {
	path := fmt.Sprintf("tagSets/%s", tagSetId)
	resp, err := apiGet(s.sling, new(TagSet), path)

	if err != nil {
		return nil, err
	}

	return resp.(*TagSet), nil
}

func (s *TagSetService) GetAll() (*[]TagSet, error) {
	var p []TagSet

	path := "tagSets"

	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(TagSets), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*TagSets)

		for _, item := range r.Items {
			p = append(p, item)
		}

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

func (s *TagSetService) GetByName(tagSetName string) (*TagSet, error) {
	var foundTagSet TagSet
	tagSets, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, tagSet := range *tagSets {
		if tagSet.Name == tagSetName {
			return &tagSet, nil
		}
	}

	return &foundTagSet, fmt.Errorf("no tagSet found with tagSet name %s", tagSetName)
}

func (s *TagSetService) Add(tagSet *TagSet) (*TagSet, error) {
	resp, err := apiAdd(s.sling, tagSet, new(TagSet), "tagSets")

	if err != nil {
		return nil, err
	}

	return resp.(*TagSet), nil
}

func (s *TagSetService) Delete(tagSetId string) error {
	path := fmt.Sprintf("tagSets/%s", tagSetId)
	err := apiDelete(s.sling, path)

	if err != nil {
		return err
	}

	return nil
}

func (s *TagSetService) Update(tagSet *TagSet) (*TagSet, error) {
	path := fmt.Sprintf("tagSets/%s", tagSet.ID)
	resp, err := apiUpdate(s.sling, tagSet, new(TagSet), path)

	if err != nil {
		return nil, err
	}

	return resp.(*TagSet), nil
}
