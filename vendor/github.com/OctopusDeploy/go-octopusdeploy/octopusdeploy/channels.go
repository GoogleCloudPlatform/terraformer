package octopusdeploy

import (
	"fmt"

	"github.com/dghubble/sling"
	"gopkg.in/go-playground/validator.v9"
)

type ChannelService struct {
	sling *sling.Sling
}

func NewChannelService(sling *sling.Sling) *ChannelService {
	return &ChannelService{
		sling: sling,
	}
}

type Channels struct {
	Items []Channel `json:"Items"`
	PagedResults
}

type Channel struct {
	Description string        `json:"Description"`
	ID          string        `json:"Id,omitempty"`
	IsDefault   bool          `json:"IsDefault"`
	LifecycleID string        `json:"LifecycleId"`
	Name        string        `json:"Name"`
	ProjectID   string        `json:"ProjectId"`
	Rules       []ChannelRule `json:"Rules,omitempty"`
	TenantTags  []string      `json:"TenantedDeploymentMode,omitempty"`
}

type ChannelRule struct {
	// name of Package step(s) this rule applies to
	Actions []string `json:"Actions,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// Pre-release tag
	Tag string `json:"Tag,omitempty"`

	//Use the NuGet or Maven versioning syntax (depending on the feed type)
	//to specify the range of versions to include
	VersionRange string `json:"VersionRange,omitempty"`
}

func (d *Channels) Validate() error {
	validate := validator.New()

	err := validate.Struct(d)

	if err != nil {
		return err
	}

	return nil
}

func (s *ChannelService) Get(channelID string) (*Channel, error) {
	path := fmt.Sprintf("channels/%s", channelID)
	resp, err := apiGet(s.sling, new(Channel), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Channel), nil
}

func (s *ChannelService) GetAll() (*[]Channel, error) {
	var ch []Channel

	path := "channel"

	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(Channels), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*Channels)

		for _, item := range r.Items {
			ch = append(ch, item)
		}

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &ch, nil
}

// Add adds a new channel
func (s *ChannelService) Add(channel *Channel) (*Channel, error) {
	err := ValidateChannelValues(channel)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, channel, new(Channel), "channels")

	if err != nil {
		return nil, err
	}

	return resp.(*Channel), nil
}

// ValidateChannelValues checks the values of a Channel object to see if they are suitable for
// sending to Octopus Deploy. Used when adding or updating channels.
func ValidateChannelValues(Channel *Channel) error {
	return ValidateMultipleProperties([]error{
		ValidateRequiredPropertyValue("Name", Channel.Name),
		ValidateRequiredPropertyValue("ProjectID", Channel.ProjectID),
	})
}

func NewChannel(name, description, projectID string) *Channel {
	return &Channel{
		Name:        name,
		ProjectID:   projectID,
		Description: description,
	}
}

// Delete deletes an existing channel in Octopus Deploy
func (s *ChannelService) Delete(channelid string) error {
	path := fmt.Sprintf("channels/%s", channelid)
	err := apiDelete(s.sling, path)

	if err != nil {
		return err
	}

	return nil
}

// Update updates an existing channel in Octopus Deploy
func (s *ChannelService) Update(channel *Channel) (*Channel, error) {
	err := ValidateChannelValues(channel)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("channels/%s", channel.ID)
	resp, err := apiUpdate(s.sling, channel, new(Channel), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Channel), nil
}
