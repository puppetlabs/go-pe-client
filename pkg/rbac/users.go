package rbac

import (
	"fmt"
	"time"
)

const (
	requestUsersURI       = "/users"         // #nosec - this is the uri to g et RBAC tokens
	requestCurrentUserURI = "/users/current" // #nosec - this is the uri to authenticate RBAC tokens
	requestUserURI        = "/users/"        // #nosec - this is the uri to revoke individual RBAC tokens
)

// User describes the user keys.
type User struct {
	ID               string    `json:"id"`
	Login            string    `json:"login"`
	Email            string    `json:"email,omitempty"`
	DisplayName      string    `json:"display_name,omitempty"`
	RoleIDs          []int     `json:"role_ids,omitempty"`
	IsGroup          bool      `json:"is_group,omitempty"`
	IsRemote         bool      `json:"is_remote,omitempty"`
	IsUser           bool      `json:"is_user,omitempty"`
	IsSuperUser      bool      `json:"is_superuser,omitempty"`
	IsRevoked        bool      `json:"is_revoked,omitempty"`
	LastLogin        time.Time `json:"last_login,omitempty"`
	InheritedRoleIDs []int     `json:"inherited_role_ids,omitempty"`
	GroupIDs         []string  `json:"group_ids,omitempty"`
}

// GetUsers returns all the users in the system.
func (c *Client) GetUsers(token string) ([]User, error) {
	users := []User{}
	r, err := c.resty.R().
		SetHeader("X-Authentication", token).
		SetResult(&users).
		Get(requestUsersURI)
	if err != nil {
		return nil, FormatError(r, err.Error())
	}
	if r.IsError() {
		if r.Error() != nil {
			return nil, FormatError(r)
		}
		return nil, FormatError(r)
	}
	return users, nil
}

// GetCurrentUser will return the current user details.
func (c *Client) GetCurrentUser(token string) (*User, error) {
	user := User{}
	r, err := c.resty.R().
		SetHeader("X-Authentication", token).
		SetResult(&user).
		Get(requestCurrentUserURI)
	if err != nil {
		return nil, FormatError(r, err.Error())
	}
	if r.IsError() {
		if r.Error() != nil {
			return nil, FormatError(r)
		}
		return nil, FormatError(r)
	}
	return &user, nil
}

// GetSpecificUser will return a specific user details
func (c *Client) GetSpecificUser(token string, sid string) (*User, error) {
	user := User{}
	r, err := c.resty.R().
		SetHeader("X-Authentication", token).
		SetResult(&user).
		Get(fmt.Sprintf("%s%s", requestUserURI, sid))
	if err != nil {
		return nil, FormatError(r, err.Error())
	}
	if r.IsError() {
		if r.Error() != nil {
			return nil, FormatError(r)
		}
		return nil, FormatError(r)
	}
	return &user, nil
}
