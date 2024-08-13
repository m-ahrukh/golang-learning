package processor

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestAuthenticationClient(t *testing.T) {
	gunit.Run(new(AuthenticationClientFixture), t)
}

type AuthenticationClientFixture struct {
	*gunit.Fixture

	inner  *FakeHTTPClient
	client *AuthenticationClient
}

func (acf *AuthenticationClientFixture) Setup() {
	acf.inner = &FakeHTTPClient{}
	acf.client = NewAuthenticationClient(acf.inner, "http", "different-company.com", "authid", "authtoken")
}

func (acf *AuthenticationClientFixture) TestProvidedInformationAddedBeforeRequestIsSent() {
	request := httptest.NewRequest("GET", "/path?existingKey=existingValue", nil)

	acf.client.Do(request)

	acf.assertRequestConnectionInformation()
	acf.assertQueryStringIncludesAuthentication()
}

func (acf *AuthenticationClientFixture) TestResponseAndErrorFromInnerClientReturned() {
	acf.inner.response = &http.Response{
		StatusCode: http.StatusTeapot,
	}
	acf.inner.err = errors.New("HTTP Error")
	request := httptest.NewRequest("GET", "/path", nil)
	response, err := acf.client.Do(request)

	acf.So(response.StatusCode, should.Equal, http.StatusTeapot)
	acf.So(err.Error(), should.Equal, "HTTP Error")
}

func (acf *AuthenticationClientFixture) assertQueryStringValue(key string, expectedString string) {
	acf.So(acf.inner.request.URL.Query().Get(key), should.Equal, expectedString)
}

func (acf *AuthenticationClientFixture) assertRequestConnectionInformation() {
	acf.So(acf.inner.request.Host, should.Equal, "different-company.com")
	acf.So(acf.inner.request.URL.Scheme, should.Equal, "http")
	acf.So(acf.inner.request.URL.Host, should.Equal, "different-company.com")
}

func (acf *AuthenticationClientFixture) assertQueryStringIncludesAuthentication() {
	acf.assertQueryStringValue("auth-id", "authid")
	acf.assertQueryStringValue("auth-token", "authtoken")
	acf.assertQueryStringValue("existingKey", "existingValue")
}
