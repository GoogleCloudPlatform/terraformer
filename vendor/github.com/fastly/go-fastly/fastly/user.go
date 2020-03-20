package fastly

import (
	"fmt"
	"sort"
	"time"
)

// User represents a user of the Fastly API and web interface.
type User struct {
	ID                     string     `mapstructure:"id"`
	Login                  string     `mapstructure:"login"`
	Name                   string     `mapstructure:"name"`
	Role                   string     `mapstructure:"role"`
	CustomerID             string     `mapstructure:"customer_id"`
	EmailHash              string     `mapstructure:"email_hash"`
	LimitServices          bool       `mapstructure:"limit_services"`
	Locked                 bool       `mapstructure:"locked"`
	RequireNewPassword     bool       `mapstructure:"require_new_password"`
	TwoFactorAuthEnabled   bool       `mapstructure:"two_factor_auth_enabled"`
	TwoFactorSetupRequired bool       `mapstructure:"two_factor_setup_required"`
	CreatedAt              *time.Time `mapstructure:"created_at"`
	UpdatedAt              *time.Time `mapstructure:"updated_at"`
	DeletedAt              *time.Time `mapstructure:"deleted_at"`
}

// usersByLogin is a sortable list of users.
type usersByName []*User

// Len, Swap, and Less implement the sortable interface.
func (s usersByName) Len() int      { return len(s) }
func (s usersByName) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s usersByName) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

// ListCustomerUsersInput is used as input to the ListCustomerUsers function.
type ListCustomerUsersInput struct {
	CustomerID string
}

// ListCustomerUsers returns the full list of users belonging to a specific
// customer.
func (c *Client) ListCustomerUsers(i *ListCustomerUsersInput) ([]*User, error) {
	if i.CustomerID == "" {
		return nil, ErrMissingCustomerID
	}

	path := fmt.Sprintf("/customer/%s/users", i.CustomerID)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var u []*User
	if err := decodeJSON(&u, resp.Body); err != nil {
		return nil, err
	}
	sort.Stable(usersByName(u))
	return u, nil
}

// GetCurrentUser retrieves the user information for the authenticated user.
func (c *Client) GetCurrentUser() (*User, error) {
	resp, err := c.Get("/current_user", nil)
	if err != nil {
		return nil, err
	}

	var u *User
	if err := decodeJSON(&u, resp.Body); err != nil {
		return nil, err
	}

	return u, nil
}

// GetUserInput is used as input to the GetUser function.
type GetUserInput struct {
	ID string
}

// GetUser retrieves the user information for the user with the given
// id. If no user exists for the given id, the API returns a 404 response.
func (c *Client) GetUser(i *GetUserInput) (*User, error) {
	if i.ID == "" {
		return nil, ErrMissingID
	}

	path := fmt.Sprintf("/user/%s", i.ID)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var u *User
	if err := decodeJSON(&u, resp.Body); err != nil {
		return nil, err
	}

	return u, nil
}

// CreateUserInput is used as input to the CreateUser function.
type CreateUserInput struct {
	Login string `form:"login"`
	Name  string `form:"name"`

	Role string `form:"role,omitempty"`
}

// CreateUser creates a new API token with the given information.
func (c *Client) CreateUser(i *CreateUserInput) (*User, error) {
	if i.Login == "" {
		return nil, ErrMissingLogin
	}

	if i.Name == "" {
		return nil, ErrMissingName
	}

	resp, err := c.PostForm("/user", i, nil)
	if err != nil {
		return nil, err
	}

	var u *User
	if err := decodeJSON(&u, resp.Body); err != nil {
		return nil, err
	}
	return u, nil
}

// UpdateUserInput is used as input to the UpdateUser function.
type UpdateUserInput struct {
	ID string `form:"-"`

	Name string `form:"name,omitempty"`
	Role string `form:"role,omitempty"`
}

// UpdateUser updates the user with the given input.
func (c *Client) UpdateUser(i *UpdateUserInput) (*User, error) {
	if i.ID == "" {
		return nil, ErrMissingID
	}

	path := fmt.Sprintf("/user/%s", i.ID)
	resp, err := c.PutForm(path, i, nil)
	if err != nil {
		return nil, err
	}

	var u *User
	if err := decodeJSON(&u, resp.Body); err != nil {
		return nil, err
	}
	return u, nil
}

// DeleteUserInput is used as input to the DeleteUser function.
type DeleteUserInput struct {
	ID string
}

// DeleteUser revokes a specific token by its ID.
func (c *Client) DeleteUser(i *DeleteUserInput) error {
	if i.ID == "" {
		return ErrMissingID
	}

	path := fmt.Sprintf("/user/%s", i.ID)
	resp, err := c.Delete(path, nil)
	if err != nil {
		return err
	}

	var r *statusResp
	if err := decodeJSON(&r, resp.Body); err != nil {
		return err
	}
	if !r.Ok() {
		return fmt.Errorf("Not Ok")
	}
	return nil
}
