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

	request := smartyVerifier.buildRequest(input)
	response, _ := smartyVerifier.client.Do(request)

	candidates := smartyVerifier.decodeResponse(response)

	return smartyVerifier.preparingAddressOutput(candidates)
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

func (smartyVerifier *SmartyVerifier) decodeResponse(response *http.Response) (output []Candidate) {
	if response != nil {
		defer response.Body.Close()
		json.NewDecoder(response.Body).Decode(&output)
	}
	return output
}

func (smartyVerifier *SmartyVerifier) preparingAddressOutput(candidates []Candidate) AddressOutput {
	if len(candidates) == 0 {
		return AddressOutput{Status: "Invalid API Response"}
	}

	candidate := candidates[0]
	return AddressOutput{
		DeliveryLine1: candidate.DeliveryLine1,
		LastLine:      candidate.LastLine,
		City:          candidate.Components.City,
		State:         candidate.Components.State,
		ZIPCode:       candidate.Components.ZIPCode,
	}
}

type Candidate struct {
	DeliveryLine1 string `json:"delivery_line_1"`
	LastLine      string `json:"last_line"`
	Components    struct {
		City    string `json:"city_name"`
		State   string `json:"state_abbreviation"`
		ZIPCode string `json:"zipcode"`
	} `json:"components"`
}
