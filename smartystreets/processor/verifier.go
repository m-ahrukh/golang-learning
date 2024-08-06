package processor

import (
	"encoding/json"
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

	var request *http.Request = smartyVerifier.buildRequest(input)

	response, _ := smartyVerifier.client.Do(request)

	var output []Candidate
	/* _ := */ json.NewDecoder(response.Body).Decode(&output)

	return AddressOutput{
		DeliveryLine1: output[0].DeliveryLine1,
		LastLine:      output[0].LastLine,
	}
}

func (smartyVerifier *SmartyVerifier) buildRequest(input AddressInput) *http.Request {
	query := make(url.Values)
	query.Set("street", input.Street1)
	query.Set("city", input.City)
	query.Set("state", input.State)
	query.Set("zipcode", input.ZIPCode)
	request, _ := http.NewRequest("GET", "/street-address?"+query.Encode(), nil)
	return request
}

type Candidate struct {
	DeliveryLine1 string `json:"delivery_line_1"`
	LastLine      string `json:"last_line"`
}
