package processor

import (
	"net/http"
	"net/url"
)

// http server interface
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type SmartyVerifier struct {
	client HTTPClient
}

func (this *SmartyVerifier) Verify(input AddressInput) AddressOutput {

	query := make(url.Values)
	query.Set("street", input.Street1)

	request, _ := http.NewRequest("GET", "/street-address?"+query.Encode(), nil)
	this.client.Do(request)

	return AddressOutput{}
}
