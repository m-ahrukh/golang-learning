package processor

import (
	"bytes"
	"errors"
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
		City:    "City",
		State:   "State",
		ZIPCode: "ZIPCode",
	}

	verifierFixture.client.Configure("[{}]", http.StatusOK, nil)
	verifierFixture.verifier.Verify(input)

	verifierFixture.So(verifierFixture.client.request.Method, should.Equal, "GET")
	verifierFixture.So(verifierFixture.client.request.URL.Path, should.Equal, "/street-address")

	verifierFixture.AssertEqual("/street-address", verifierFixture.client.request.URL.Path)
	verifierFixture.AssertQueryStringValue("street", input.Street1)
	verifierFixture.AssertQueryStringValue("city", input.City)
	verifierFixture.AssertQueryStringValue("state", input.State)
	verifierFixture.AssertQueryStringValue("zipcode", input.ZIPCode)

}

func (verifierFixture *VerifierFixture) AssertQueryStringValue(key, expected string) {
	query := verifierFixture.client.request.URL.Query()

	verifierFixture.So(query.Get(key), should.Equal, expected)
}

func (verifierFixture *VerifierFixture) TestResponseParsed() {

	verifierFixture.client.Configure(rawJSONOutput, http.StatusOK, nil)
	result := verifierFixture.verifier.Verify(AddressInput{})

	verifierFixture.So(result.DeliveryLine1, should.Equal, "1 Santa Claus Ln")
	verifierFixture.So(result.LastLine, should.Equal, "North Pole AK 99705-9901")
	verifierFixture.So(result.City, should.Equal, "North Pole")
	verifierFixture.So(result.State, should.Equal, "AK")
	verifierFixture.So(result.ZIPCode, should.Equal, "99705")
}

const rawJSONOutput = `
[
	{
		"delivery_line_1": "1 Santa Claus Ln",
		"last_line": "North Pole AK 99705-9901",
		"components": {
			"city_name": "North Pole",
			"state_abbreviation": "AK",
			"zipcode": "99705"
		}
	}
]`

func (verifierFixture *VerifierFixture) TestMalformedJSONHandled() {
	const malformedRawJSONOutput = `alert('Hello, World' DROP TABLE Users)`
	verifierFixture.client.Configure(malformedRawJSONOutput, http.StatusOK, nil)
	result := verifierFixture.verifier.Verify(AddressInput{})
	verifierFixture.So(result.Status, should.Equal, "Invalid API Response")
}

func (verifierFixture *VerifierFixture) TestHTTPErrorHandler() {
	verifierFixture.client.Configure("", 0, errors.New("GOPHERS!"))

	result := verifierFixture.verifier.Verify(AddressInput{})
	verifierFixture.So(result.Status, should.Equal, "Invalid API Response")
}

func (verifierFixture *VerifierFixture) TestHTTPResponseBodyClosed() {
	verifierFixture.client.Configure(rawJSONOutput, http.StatusOK, nil)
	verifierFixture.verifier.Verify(AddressInput{})

	verifierFixture.So(verifierFixture.client.responseBody.closed, should.Equal, 1)

}

// ///////////////////////////////////////////////////////
type FakeHTTPClient struct {
	request      *http.Request
	response     *http.Response
	responseBody *SpyBuffer
	err          error
}

func (fakeHTTPClient *FakeHTTPClient) Configure(responseText string, statusCode int, err error) {
	if err == nil {
		fakeHTTPClient.responseBody = NewSpyBuffer(responseText)
		fakeHTTPClient.response = &http.Response{
			Body:       fakeHTTPClient.responseBody,
			StatusCode: statusCode,
		}
	}
	fakeHTTPClient.err = err
}

func (fakeHTTPClient *FakeHTTPClient) Do(request *http.Request) (*http.Response, error) {
	fakeHTTPClient.request = request
	return fakeHTTPClient.response, fakeHTTPClient.err
}

// ///////////////////////////////////////////////////////
type SpyBuffer struct {
	*bytes.Buffer
	closed int
}

func NewSpyBuffer(value string) *SpyBuffer {
	return &SpyBuffer{
		Buffer: bytes.NewBufferString(value),
	}
}

func (spyBuffer *SpyBuffer) Close() error {
	spyBuffer.closed++
	spyBuffer.Buffer.Reset()
	return nil
}
