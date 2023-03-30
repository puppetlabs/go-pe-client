package rbac

const (
	rolesPath = "/rbac-api/v1/roles"
)

// GetRoles fetches information about all user roles.
func (c *Client) GetRoles(token string) ([]Role, error) {
	var roles []Role

	response, err := c.resty.R().
		SetHeader("X-Authentication", token).
		SetResult(&roles).
		Get(rolesPath)
	if err != nil {
		return nil, FormatError(response, err.Error())
	}

	return roles, nil
}

// CreateRole creates a role, and attaches to it the specified permissions and
// the specified users and groups. Authentication is required.
//
// If the role was created successfully then the path of the new role is
// returned, otherwise an error is returned.
func (c *Client) CreateRole(role *Role, token string) (string, error) {
	r, err := c.resty.R().
		SetBody(role).
		SetHeader("X-Authentication", token).
		Post(rolesPath)
	if err != nil {
		// This API uses a redirect with location header to indicate success.
		// Because redirects are disabled in the RBAC client an
		// error will be thrown when the redirect cannot be followed.
		if !r.IsError() && r.RawResponse.Header.Get("Location") != "" {
			// Ignore the error.
		} else {
			return "", FormatError(r, err.Error())
		}
	}

	// If the HTTP status code is >400 or there is no location header in the
	// response then the request was not successful.
	if r.IsError() || r.RawResponse.Header.Get("Location") == "" {
		return "", FormatError(r)
	}

	return r.RawResponse.Header.Get("Location"), nil
}

// Role represents an RBAC role
type Role struct {
	ID          uint         `json:"id,omitempty"`
	Permissions []Permission `json:"permissions"`
	UserIDs     []string     `json:"user_ids"`
	GroupIDs    []string     `json:"group_ids"`
	DisplayName string       `json:"display_name"`
	Description string       `json:"description"`
}

// Permission represents an RBAC permission
type Permission struct {
	ObjectType string `json:"object_type"`
	Action     string `json:"action"`
	Instance   string `json:"instance"`
}
