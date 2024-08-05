package processor

import (
	"net/http"
	"testing"

	"github.com/smarty/gunit"
)

func TestVerifierFixture(t *testing.T) {
	gunit.Run(new(VerifierFixture), t)
}

type VerifierFixture struct {
	*gunit.Fixture

	client   *FakeHTTPClient
	verifier *SmartyVerifier
}

func (verifierFixture *VerifierFixture) Setup() {
	verifierFixture.client = &FakeHTTPClient{}
	verifierFixture.verifier = NewSmartyVerifier(verifierFixture.client)
}

func NewSmartyVerifier(client HTTPClient) *SmartyVerifier {
	return &SmartyVerifier{
		client: client,
	}
}

func (verifierFixture *VerifierFixture) TestRequestComposedProperly() {
	input := AddressInput{
		Street1: "Street1",
		City:    "City",
		State:   "State",
		ZIPCode: "ZIPCode",
	}

	verifierFixture.verifier.Verify(input)
	verifierFixture.AssertEqual("GET", verifierFixture.client.request.Method)
	verifierFixture.AssertEqual("/street-address", verifierFixture.client.request.URL.Path)
	verifierFixture.AssertQueryStringValue("street", "Street1")
	verifierFixture.AssertQueryStringValue("city", "City")
	verifierFixture.AssertQueryStringValue("state", "State")
	verifierFixture.AssertQueryStringValue("zipcode", "ZIPCode")
	// this.AssertEqual("/street-address?street=Street1&city=City", this.client.request.URL.String())

}

func (verifierFixture *VerifierFixture) AssertQueryStringValue(key, expected string) {
	query := verifierFixture.client.request.URL.Query()

	verifierFixture.AssertEqual(expected, query.Get(key))
}

func (verifierFixture *VerifierFixture) rawQuery() string {
	return verifierFixture.client.request.URL.RawQuery
}

// ///////////////////////////////////////////////////////
type FakeHTTPClient struct {
	request *http.Request
}

func (this *FakeHTTPClient) Do(request *http.Request) (*http.Response, error) {
	this.request = request
	return nil, nil
}
