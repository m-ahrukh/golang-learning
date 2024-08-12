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

func (authenticationClient *AuthenticationClient) Do(request *http.Request) (*http.Response, error) {
	request.URL.Scheme = authenticationClient.scheme
	request.Host = authenticationClient.hostname
	request.URL.Host = authenticationClient.hostname
	authenticationClient.inner.Do(request)
	return &http.Response{StatusCode: http.StatusTeapot}, nil
}
