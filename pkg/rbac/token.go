package rbac

import "fmt"

const (
	requestAuthTokenURI  = "/rbac-api/v1/auth/token"              // #nosec - this is the uri to g et RBAC tokens
	tokenAuthenticateURI = "/rbac-api/v2/auth/token/authenticate" // #nosec - this is the uri to authenticate RBAC tokens
	tokenRevokeURI       = "/rbac-api/v2/tokens/"                 // #nosec - this is the uri to revoke individual RBAC tokens
	tokenGenerateURI     = "/rbac-api/v1/tokens"                  // #nosec - this is the uri to generate a token.
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

func (c *Client) RevokeRBACToken(token string) error {
	payload := AuthenticateResponse{}

	r, err := c.resty.R().
		SetResult(&payload).
		Delete(fmt.Sprintf("%s%s", tokenRevokeURI, token))
	if err != nil {
		return FormatError(r, err.Error())
	}
	if r.IsError() {
		if r.Error() != nil {
			return FormatError(r)
		}
		return FormatError(r)
	}
	return nil
}

// GenerateRBACToken returns an RBAC token or errors otherwise
func (c *Client) GenerateRBACToken(token string, request TokenRequest) (string, error) {
	var payload Token

	r, err := c.resty.R().
		SetHeader("X-Authentication", token).
		SetResult(&payload).
		SetBody(request).
		Post(tokenGenerateURI)
	if err != nil {
		return "", FormatError(r, err.Error())
	}
	if r.IsError() {
		if r.Error() != nil {
			return "", FormatError(r)
		}
		return "", FormatError(r)
	}
	return payload.Token, nil
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

// TokenRequest will hold the details needed by the /tokens endpoint.
type TokenRequest struct {
	Lifetime    string `json:"lifetime,omitempty"`
	Description string `json:"description,omitempty"`
	Client      string `json:"client,omitempty"`
}
