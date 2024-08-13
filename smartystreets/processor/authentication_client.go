package processor

import (
	"net/http"
)

type AuthenticationClient struct {
	inner     HTTPClient
	scheme    string
	hostname  string
	authId    string
	authToken string
}

func NewAuthenticationClient(inner HTTPClient, scheme, hostname, authId, authToken string) *AuthenticationClient {
	return &AuthenticationClient{
		inner:     inner,
		scheme:    scheme,
		hostname:  hostname,
		authId:    authId,
		authToken: authToken,
	}
}

func (authenticationClient *AuthenticationClient) Do(request *http.Request) (*http.Response, error) {
	request.URL.Scheme = authenticationClient.scheme
	request.Host = authenticationClient.hostname
	request.URL.Host = authenticationClient.hostname
	query := request.URL.Query()
	query.Set("auth-id", authenticationClient.authId)
	query.Set("auth-token", authenticationClient.authToken)
	request.URL.RawQuery = query.Encode()
	return authenticationClient.inner.Do(request)
}
