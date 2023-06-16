package rbac

const (
	requestAuthTokenURI  = "/rbac-api/v1/auth/token"              // #nosec - this is the uri to g et RBAC tokens
	tokenAuthenticateURI = "/rbac-api/v2/auth/token/authenticate" // #nosec - this is the uri to authenticate RBAC tokens
)

// GetRBACToken returns an auth token given user/password information
func (c *Client) GetRBACToken(authRequest *RequestKeys) (*Token, error) {
	payload := Token{}
	r, err := c.resty.R().
		SetResult(&payload).
		SetBody(authRequest).
		Post(requestAuthTokenURI)
	if err != nil {
		return nil, FormatError(r, err.Error())
	}
	if r.IsError() {
		if r.Error() != nil {
			return nil, FormatError(r)
		}
		return nil, FormatError(r)
	}
	return &payload, nil
}

// AuthenticateRBACToken returns a response with the token details or errors otherwise.
func (c *Client) AuthenticateRBACToken(token string) (*AuthenticateResponse, error) {
	authenticateRequest := &AuthenticateRequest{Token: token}

	payload := AuthenticateResponse{}
	r, err := c.resty.R().
		SetResult(&payload).
		SetBody(authenticateRequest).
		Post(tokenAuthenticateURI)
	if err != nil {
		return nil, FormatError(r, err.Error())
	}
	if r.IsError() {
		if r.Error() != nil {
			return nil, FormatError(r)
		}
		return nil, FormatError(r)
	}
	return &payload, nil
}

// Token is the returned auth token
type Token struct {
	Token string `json:"token"`
}

// RequestKeys describes the keys used by the token endpoint
type RequestKeys struct {
	Login       string `json:"login"`
	Password    string `json:"password"`
	Lifetime    string `json:"lifetime,omitempty"`
	Description string `json:"description,omitempty"`
	Client      string `json:"client,omitempty"`
	Label       string `json:"label,omitempty"`
}

// AuthenticateRequest will hold the request needed for an authenticate.
type AuthenticateRequest struct {
	Token              string `json:"token"`
	UpdateLastActivity bool   `json:"update_last_activity?"`
}

// AuthenticateResponse will hold the response from an authenticate.
type AuthenticateResponse struct {
	Description string `json:"description"`
	Login       string `json:"login"`
	RoleIDs     []int  `json:"role_ids"`
	UserID      string `json:"user_id"`
	DisplayName string `json:"display_name"`
}
