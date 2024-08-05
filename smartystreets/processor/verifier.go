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

func (smartyVerifier *SmartyVerifier) Verify(input AddressInput) AddressOutput {

	query := make(url.Values)
	query.Set("street", input.Street1)
	query.Set("city", input.City)
	query.Set("state", input.State)
	query.Set("zipcode", input.ZIPCode)

	request, _ := http.NewRequest("GET", "/street-address?"+query.Encode(), nil)

	smartyVerifier.client.Do(request)

	return AddressOutput{}
}
