// Automatically generated, do not edit

package commercetools

import "time"

// APIClient is a standalone struct
type APIClient struct {
	Secret     string     `json:"secret,omitempty"`
	Scope      string     `json:"scope"`
	Name       string     `json:"name"`
	LastUsedAt Date       `json:"lastUsedAt,omitempty"`
	ID         string     `json:"id"`
	DeleteAt   *time.Time `json:"deleteAt,omitempty"`
	CreatedAt  *time.Time `json:"createdAt,omitempty"`
}

// APIClientDraft is a standalone struct
type APIClientDraft struct {
	Scope                   string `json:"scope"`
	Name                    string `json:"name"`
	DeleteDaysAfterCreation int    `json:"deleteDaysAfterCreation,omitempty"`
}

// APIClientPagedQueryResponse is a standalone struct
type APIClientPagedQueryResponse struct {
	Total   int         `json:"total,omitempty"`
	Results []APIClient `json:"results"`
	Offset  int         `json:"offset"`
	Count   int         `json:"count"`
}
