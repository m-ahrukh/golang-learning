package processor

import "net/http"

type AuthenticationClient struct {
	inner    HTTPClient
	scheme   string
	hostname string
}

func NewAuthenticationClient(inner HTTPClient, scheme, hostname string) *AuthenticationClient {
	return &AuthenticationClient{
		inner:    inner,
		scheme:   scheme,
		hostname: hostname,
	}
}

func (authenticationClient *AuthenticationClient) Do(*http.Request) (*http.Response, error) {
	panic("implement")
}
