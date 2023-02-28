package rbac

const (
	rbacRoles = "/rbac-api/v1/roles"
)

// CreateRole creates a role, and attaches to it the specified permissions and
// the specified users and groups. Authentication is required.
//
// If the role was created successfully then the path of the new role is
// returned. In all other cases an error is returned, cast it to rbac.APIError
// to get the HTTP status code & message.
func (c *Client) CreateRole(roles *Role, token string) (string, error) {
	r, err := c.resty.R().
		SetBody(roles).
		SetHeader("X-Authentication", token).
		Post(rbacRoles)

	if err != nil {
		// On success the response is HTTP 303 with a redirect. The RBAC client
		// considers this an error because it's configured to not follow
		// redirects.
		if r.StatusCode() != 303 {
			return "", FormatError(r, err.Error())
		}
	}
	if r.IsError() {
		return "", FormatError(r)
	}

	return r.RawResponse.Header.Get("Location"), nil
}

// Role represents an RBAC role
type Role struct {
	Permissions []Permission `json:"permissions"`
	UserIDs     []string     `json:"user_ids"`
	GroupIDs    []string     `json:"group_ids"`
	DisplayName string       `json:"display_name"`
	Description string       `json:"description,omitempty"`
}

// Permission represents an RBAC permission
type Permission struct {
	ObjectType string `json:"object_type"`
	Action     string `json:"action"`
	Instance   string `json:"instance"`
}
