package rbac

const (
	requestAuthTokenURI = "/rbac-api/v1/auth/token" // #nosec - this is the uri to get RBAC tokens
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
