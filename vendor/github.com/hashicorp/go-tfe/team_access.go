package tfe

import (
	"context"
	"errors"
	"fmt"
	"net/url"
)

// Compile-time proof of interface implementation.
var _ TeamAccesses = (*teamAccesses)(nil)

// TeamAccesses describes all the team access related methods that the
// Terraform Enterprise API supports.
//
// TFE API docs:
// https://www.terraform.io/docs/enterprise/api/team-access.html
type TeamAccesses interface {
	// List all the team accesses for a given workspace.
	List(ctx context.Context, options TeamAccessListOptions) (*TeamAccessList, error)

	// Add team access for a workspace.
	Add(ctx context.Context, options TeamAccessAddOptions) (*TeamAccess, error)

	// Read a team access by its ID.
	Read(ctx context.Context, teamAccessID string) (*TeamAccess, error)

	// Remove team access from a workspace.
	Remove(ctx context.Context, teamAccessID string) error
}

// teamAccesses implements TeamAccesses.
type teamAccesses struct {
	client *Client
}

// AccessType represents a team access type.
type AccessType string

// List all available team access types.
const (
	AccessAdmin AccessType = "admin"
	AccessPlan  AccessType = "plan"
	AccessRead  AccessType = "read"
	AccessWrite AccessType = "write"
)

// TeamAccessList represents a list of team accesses.
type TeamAccessList struct {
	*Pagination
	Items []*TeamAccess
}

// TeamAccess represents the workspace access for a team.
type TeamAccess struct {
	ID     string     `jsonapi:"primary,team-workspaces"`
	Access AccessType `jsonapi:"attr,access"`

	// Relations
	Team      *Team      `jsonapi:"relation,team"`
	Workspace *Workspace `jsonapi:"relation,workspace"`
}

// TeamAccessListOptions represents the options for listing team accesses.
type TeamAccessListOptions struct {
	ListOptions
	WorkspaceID *string `url:"filter[workspace][id],omitempty"`
}

func (o TeamAccessListOptions) valid() error {
	if !validString(o.WorkspaceID) {
		return errors.New("workspace ID is required")
	}
	if !validStringID(o.WorkspaceID) {
		return errors.New("invalid value for workspace ID")
	}
	return nil
}

// List all the team accesses for a given workspace.
func (s *teamAccesses) List(ctx context.Context, options TeamAccessListOptions) (*TeamAccessList, error) {
	if err := options.valid(); err != nil {
		return nil, err
	}

	req, err := s.client.newRequest("GET", "team-workspaces", &options)
	if err != nil {
		return nil, err
	}

	tal := &TeamAccessList{}
	err = s.client.do(ctx, req, tal)
	if err != nil {
		return nil, err
	}

	return tal, nil
}

// TeamAccessAddOptions represents the options for adding team access.
type TeamAccessAddOptions struct {
	// For internal use only!
	ID string `jsonapi:"primary,team-workspaces"`

	// The type of access to grant.
	Access *AccessType `jsonapi:"attr,access"`

	// The team to add to the workspace
	Team *Team `jsonapi:"relation,team"`

	// The workspace to which the team is to be added.
	Workspace *Workspace `jsonapi:"relation,workspace"`
}

func (o TeamAccessAddOptions) valid() error {
	if o.Access == nil {
		return errors.New("access is required")
	}
	if o.Team == nil {
		return errors.New("team is required")
	}
	if o.Workspace == nil {
		return errors.New("workspace is required")
	}
	return nil
}

// Add team access for a workspace.
func (s *teamAccesses) Add(ctx context.Context, options TeamAccessAddOptions) (*TeamAccess, error) {
	if err := options.valid(); err != nil {
		return nil, err
	}

	// Make sure we don't send a user provided ID.
	options.ID = ""

	req, err := s.client.newRequest("POST", "team-workspaces", &options)
	if err != nil {
		return nil, err
	}

	ta := &TeamAccess{}
	err = s.client.do(ctx, req, ta)
	if err != nil {
		return nil, err
	}

	return ta, nil
}

// Read a team access by its ID.
func (s *teamAccesses) Read(ctx context.Context, teamAccessID string) (*TeamAccess, error) {
	if !validStringID(&teamAccessID) {
		return nil, errors.New("invalid value for team access ID")
	}

	u := fmt.Sprintf("team-workspaces/%s", url.QueryEscape(teamAccessID))
	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	ta := &TeamAccess{}
	err = s.client.do(ctx, req, ta)
	if err != nil {
		return nil, err
	}

	return ta, nil
}

// Remove team access from a workspace.
func (s *teamAccesses) Remove(ctx context.Context, teamAccessID string) error {
	if !validStringID(&teamAccessID) {
		return errors.New("invalid value for team access ID")
	}

	u := fmt.Sprintf("team-workspaces/%s", url.QueryEscape(teamAccessID))
	req, err := s.client.newRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	return s.client.do(ctx, req, nil)
}
