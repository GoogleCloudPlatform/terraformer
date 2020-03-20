package octopusdeploy

import (
	"fmt"

	"github.com/dghubble/sling"
	"gopkg.in/go-playground/validator.v9"
)

type FeedService struct {
	sling *sling.Sling
}

func NewFeedService(sling *sling.Sling) *FeedService {
	return &FeedService{
		sling: sling,
	}
}

type Feeds struct {
	Items []Feed `json:"Items"`
	PagedResults
}

type Feed struct {
	ID                          string         `json:"Id"`
	Name                        string         `json:"Name"`
	FeedType                    string         `json:"FeedType"`
	DownloadAttempts            int            `json:"DownloadAttempts"`
	DownloadRetryBackoffSeconds int            `json:"DownloadRetryBackoffSeconds"`
	FeedUri                     string         `json:"FeedUri"`
	EnhancedMode                bool           `json:"EnhancedMode"`
	Username                    string         `json:"Username"`
	Password                    SensitiveValue `json:"Password"`
}

func (t *Feed) Validate() error {
	validate := validator.New()

	err := validate.Struct(t)

	if err != nil {
		return err
	}

	return nil
}

func NewFeed(name, feedType string, feedUri string) *Feed {
	return &Feed{
		Name:     name,
		FeedType: feedType,
		FeedUri:  feedUri,
	}
}

func (s *FeedService) Get(feedId string) (*Feed, error) {
	path := fmt.Sprintf("feeds/%s", feedId)
	resp, err := apiGet(s.sling, new(Feed), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Feed), nil
}

func (s *FeedService) GetAll() (*[]Feed, error) {
	var p []Feed

	path := "feeds"

	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(Feeds), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*Feeds)

		for _, item := range r.Items {
			p = append(p, item)
		}

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

func (s *FeedService) GetByName(feedName string) (*Feed, error) {
	var foundFeed Feed
	feeds, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, feed := range *feeds {
		if feed.Name == feedName {
			return &feed, nil
		}
	}

	return &foundFeed, fmt.Errorf("no feed found with feed name %s", feedName)
}

func (s *FeedService) Add(feed *Feed) (*Feed, error) {
	resp, err := apiAdd(s.sling, feed, new(Feed), "feeds")

	if err != nil {
		return nil, err
	}

	return resp.(*Feed), nil
}

func (s *FeedService) Delete(feedId string) error {
	path := fmt.Sprintf("feeds/%s", feedId)
	err := apiDelete(s.sling, path)

	if err != nil {
		return err
	}

	return nil
}

func (s *FeedService) Update(feed *Feed) (*Feed, error) {
	path := fmt.Sprintf("feeds/%s", feed.ID)
	resp, err := apiUpdate(s.sling, feed, new(Feed), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Feed), nil
}
