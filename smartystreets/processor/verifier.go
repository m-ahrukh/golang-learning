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

	output, err := smartyVerifier.decodeResponse(response)
	if err != nil {
		return AddressOutput{Status: "Invalid API Response JSON"}
	}

	return smartyVerifier.translateCandidate(output[0])
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

func (smartyVerifier *SmartyVerifier) decodeResponse(response *http.Response) ([]Candidate, error) {
	var output []Candidate
	err := json.NewDecoder(response.Body).Decode(&output)
	// if err != nil {
	// 	fmt.Println("-----------------err:", err)
	// }
	return output, err
}

func (smartyVerifier *SmartyVerifier) translateCandidate(candidate Candidate) AddressOutput {
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
