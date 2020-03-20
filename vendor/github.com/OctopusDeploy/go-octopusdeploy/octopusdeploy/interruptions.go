package octopusdeploy

import (
	"fmt"
	"time"

	"github.com/dghubble/sling"
)

type InterruptionsService struct {
	sling *sling.Sling
}

func NewInterruptionService(sling *sling.Sling) *InterruptionsService {
	return &InterruptionsService{
		sling: sling,
	}
}

type Interruption struct {
	ID        string    `json:"Id"`
	Title     string    `json:"Title"`
	Created   time.Time `json:"Created"`
	IsPending bool      `json:"IsPending"`
	Form      struct {
		Values struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"Values"`
		Elements []struct {
			Name    string `json:"Name"`
			Control struct {
			} `json:"Control"`
			IsValueRequired bool `json:"IsValueRequired"`
		} `json:"Elements"`
	} `json:"Form"`
	RelatedDocumentIds          []string          `json:"RelatedDocumentIds"`
	ResponsibleTeamIds          []string          `json:"ResponsibleTeamIds"`
	ResponsibleUserID           string            `json:"ResponsibleUserId"`
	CanTakeResponsibility       bool              `json:"CanTakeResponsibility"`
	HasResponsibility           bool              `json:"HasResponsibility"`
	TaskID                      string            `json:"TaskId"`
	CorrelationID               string            `json:"CorrelationId"`
	IsLinkedToOtherInterruption bool              `json:"IsLinkedToOtherInterruption"`
	LastModifiedOn              time.Time         `json:"LastModifiedOn"`
	LastModifiedBy              string            `json:"LastModifiedBy"`
	Links                       InterruptionLinks `json:"Links"`
}

type InterruptionLinks struct {
	Self        string `json:"Self"`
	Submit      string `json:"Submit"`
	Responsible string `json:"Responsible"`
}

type Interruptions struct {
	Items []Interruption `json:"Items"`
	PagedResults
}

type InterruptionSubmitRequest struct {
	Instructions string `json:"Instructions"`
	Notes        string `json:"Notes"`
	Result       string `json:"Result"`
}

const ManualInterverventionApprove = "Proceed"
const ManualInterventionDecline = "Abort"

// Get returns the interruption matching the id
func (s *InterruptionsService) Get(id string) (*Interruption, error) {
	path := fmt.Sprintf("%v/%v", "interruptions", id)

	resp, err := apiGet(s.sling, new(Interruption), path)

	if err != nil {
		return nil, err
	}

	r := resp.(*Interruption)
	return r, nil
}

// GetAll returns all interruptions in Octopus Deploy
func (s *InterruptionsService) GetAll() ([]Interruption, error) {
	var interruptions []Interruption
	path := "interruptions"

	loadNextPage := true
	for loadNextPage {
		resp, err := apiGet(s.sling, new(Interruptions), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*Interruptions)

		for _, item := range r.Items {
			interruptions = append(interruptions, item)
		}

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}
	return interruptions, nil
}

// Submit Submits a dictionary of form values for the interruption. Only the user with responsibility for this interruption can submit this form.
func (s *InterruptionsService) Submit(i *Interruption, r *InterruptionSubmitRequest) (*Interruption, error) {
	path := i.Links.Submit

	resp, err := apiPost(s.sling, r, new(Interruption), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Interruption), nil
}

// GetResponsability Gets the user that is currently responsible for this interruption.
func (s *InterruptionsService) GetResponsability(i *Interruption) (*User, error) {
	path := i.Links.Responsible

	resp, err := apiGet(s.sling, new(User), path)
	if err != nil {
		return nil, err
	}
	return resp.(*User), nil
}

// TakeResponsability Allows the current user to take responsibility for this interruption. Only users in one of the responsible teams on this interruption can take responsibility for it.
func (s *InterruptionsService) TakeResponsability(i *Interruption) (*User, error) {
	path := i.Links.Responsible

	resp, err := apiUpdate(s.sling, nil, new(User), path)
	if err != nil {
		return nil, err
	}
	return resp.(*User), nil
}
