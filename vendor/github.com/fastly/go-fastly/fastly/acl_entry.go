package fastly

import (
	"fmt"
	"sort"
	"time"
)

type ACLEntry struct {
	ServiceID string `mapstructure:"service_id"`
	ACLID     string `mapstructure:"acl_id"`

	ID        string     `mapstructure:"id"`
	IP        string     `mapstructure:"ip"`
	Subnet    string     `mapstructure:"subnet"`
	Negated   bool       `mapstructure:"negated"`
	Comment   string     `mapstructure:"comment"`
	CreatedAt *time.Time `mapstructure:"created_at"`
	UpdatedAt *time.Time `mapstructure:"updated_at"`
	DeletedAt *time.Time `mapstructure:"deleted_at"`
}

// entriesById is a sortable list of ACL entries.
type entriesById []*ACLEntry

// Len, Swap, and Less implements the sortable interface.
func (s entriesById) Len() int      { return len(s) }
func (s entriesById) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s entriesById) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

// ListACLEntriesInput is the input parameter to ListACLEntries function.
type ListACLEntriesInput struct {
	Service string
	ACL     string
}

// ListACLEntries return a list of entries for an ACL
func (c *Client) ListACLEntries(i *ListACLEntriesInput) ([]*ACLEntry, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.ACL == "" {
		return nil, ErrMissingACL
	}

	path := fmt.Sprintf("/service/%s/acl/%s/entries", i.Service, i.ACL)

	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var es []*ACLEntry
	if err := decodeJSON(&es, resp.Body); err != nil {
		return nil, err
	}

	sort.Stable(entriesById(es))

	return es, nil
}

// GetACLEntryInput is the input parameter to GetACLEntry function.
type GetACLEntryInput struct {
	Service string
	ACL     string
	ID      string
}

// GetACLEntry returns a single ACL entry based on its ID.
func (c *Client) GetACLEntry(i *GetACLEntryInput) (*ACLEntry, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.ACL == "" {
		return nil, ErrMissingACL
	}

	if i.ID == "" {
		return nil, ErrMissingID
	}

	path := fmt.Sprintf("/service/%s/acl/%s/entry/%s", i.Service, i.ACL, i.ID)

	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var e *ACLEntry
	if err := decodeJSON(&e, resp.Body); err != nil {
		return nil, err
	}

	return e, nil
}

// CreateACLEntryInput the input parameter to CreateACLEntry function.
type CreateACLEntryInput struct {
	// Required fields
	Service string
	ACL     string
	IP      string `form:"ip"`

	// Optional fields
	Subnet  string `form:"subnet,omitempty"`
	Negated bool   `form:"negated,omitempty"`
	Comment string `form:"comment,omitempty"`
}

// CreateACLEntry creates and returns a new ACL entry.
func (c *Client) CreateACLEntry(i *CreateACLEntryInput) (*ACLEntry, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.ACL == "" {
		return nil, ErrMissingACL
	}

	if i.IP == "" {
		return nil, ErrMissingIP
	}

	path := fmt.Sprintf("/service/%s/acl/%s/entry", i.Service, i.ACL)

	resp, err := c.PostForm(path, i, nil)
	if err != nil {
		return nil, err
	}

	var e *ACLEntry
	if err := decodeJSON(&e, resp.Body); err != nil {
		return nil, err
	}

	return e, nil
}

// DeleteACLEntryInput the input parameter to DeleteACLEntry function.
type DeleteACLEntryInput struct {
	Service string
	ACL     string
	ID      string
}

// DeleteACLEntry deletes an entry from an ACL based on its ID
func (c *Client) DeleteACLEntry(i *DeleteACLEntryInput) error {
	if i.Service == "" {
		return ErrMissingService
	}

	if i.ACL == "" {
		return ErrMissingACL
	}

	if i.ID == "" {
		return ErrMissingID
	}

	path := fmt.Sprintf("/service/%s/acl/%s/entry/%s", i.Service, i.ACL, i.ID)

	resp, err := c.Delete(path, nil)
	if err != nil {
		return err
	}

	var r *statusResp
	if err := decodeJSON(&r, resp.Body); err != nil {
		return err
	}

	if !r.Ok() {
		return fmt.Errorf("Not OK")
	}

	return nil

}

// UpdateACLEntryInput is the input parameter to UpdateACLEntry function.
type UpdateACLEntryInput struct {
	// Required fields
	Service string
	ACL     string
	ID      string

	// Optional fields
	IP      string `form:"ip,omitempty"`
	Subnet  string `form:"subnet,omitempty"`
	Negated bool   `form:"negated,omitempty"`
	Comment string `form:"comment,omitempty"`
}

// UpdateACLEntry updates an ACL entry
func (c *Client) UpdateACLEntry(i *UpdateACLEntryInput) (*ACLEntry, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.ACL == "" {
		return nil, ErrMissingACL
	}

	if i.ID == "" {
		return nil, ErrMissingID
	}

	path := fmt.Sprintf("/service/%s/acl/%s/entry/%s", i.Service, i.ACL, i.ID)

	resp, err := c.RequestForm("PATCH", path, i, nil)
	if err != nil {
		return nil, err
	}

	var e *ACLEntry
	if err := decodeJSON(&e, resp.Body); err != nil {
		return nil, err
	}

	return e, nil
}

type BatchModifyACLEntriesInput struct {
	Service string `json:"-,"`
	ACL     string `json:"-,"`

	Entries []*BatchACLEntry `json:"entries"`
}

type BatchACLEntry struct {
	Operation BatchOperation `json:"op"`
	ID        string         `json:"id,omitempty"`
	IP        string         `json:"ip,omitempty"`
	Subnet    string         `json:"subnet,omitempty"`
	Negated   bool           `json:"negated,omitempty"`
	Comment   string         `json:"comment,omitempty"`
}

func (c *Client) BatchModifyACLEntries(i *BatchModifyACLEntriesInput) error {

	if i.Service == "" {
		return ErrMissingService
	}

	if i.ACL == "" {
		return ErrMissingACL
	}

	if len(i.Entries) > BatchModifyMaximumOperations {
		return ErrBatchUpdateMaximumOperationsExceeded
	}

	path := fmt.Sprintf("/service/%s/acl/%s/entries", i.Service, i.ACL)
	resp, err := c.PatchJSON(path, i, nil)
	if err != nil {
		return err
	}

	var batchModifyResult map[string]string
	if err := decodeJSON(&batchModifyResult, resp.Body); err != nil {
		return err
	}

	return nil
}
