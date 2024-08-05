package processor

import (
	"net/http"
	"testing"

	"github.com/smarty/assertions/should"
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
		// City:    "City & City",
		City:    "City",
		State:   "State",
		ZIPCode: "ZIPCode",
	}

	verifierFixture.verifier.Verify(input)

	verifierFixture.So(verifierFixture.client.request.Method, should.Equal, "GET")
	// verifierFixture.AssertEqual("GET", verifierFixture.client.request.Method)
	verifierFixture.So(verifierFixture.client.request.URL.Path, should.Equal, "/street-address")

	verifierFixture.AssertEqual("/street-address", verifierFixture.client.request.URL.Path)
	verifierFixture.AssertQueryStringValue("street", "Street1")
	// verifierFixture.AssertQueryStringValue("city", "City & City")
	verifierFixture.AssertQueryStringValue("city", "City")
	verifierFixture.AssertQueryStringValue("state", "State")
	verifierFixture.AssertQueryStringValue("zipcode", "ZIPCode")
	// verifierFixture.Assert(strings.Contains(verifierFixture.client.request.URL.RawQuery, "%26"))

	// verifierFixture.AssertEqual("/street-address?street=Street1&city=City", this.client.request.URL.String())

}

func (verifierFixture *VerifierFixture) AssertQueryStringValue(key, expected string) {
	query := verifierFixture.client.request.URL.Query()

	verifierFixture.So(query.Get(key), should.Equal, expected)
}

func (verifierFixture *VerifierFixture) rawQuery() string {
	return verifierFixture.client.request.URL.RawQuery
}

// ///////////////////////////////////////////////////////
type FakeHTTPClient struct {
	request *http.Request
}

func (fakeHTTPClient *FakeHTTPClient) Do(request *http.Request) (*http.Response, error) {
	fakeHTTPClient.request = request
	return nil, nil
}
