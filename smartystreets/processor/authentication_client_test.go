package processor

import (
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
	acf.client = NewAuthenticationClient(acf.inner, "http", "different-company.com")
}

func (acf *AuthenticationClientFixture) TestHostnameAndSchemaAddedBeforeRequestIsSent() {
	request := httptest.NewRequest("GET", "/path", nil)

	acf.client.Do(request)

	acf.So(acf.inner.request.Host, should.Equal, "different-company.com")
	acf.So(acf.inner.request.URL.Scheme, should.Equal, "http")
	acf.So(acf.inner.request.URL.Host, should.Equal, "different-company.com")
}

func (acf *AuthenticationClientFixture) TestResponseFromInnerClientReturned() {
	acf.inner.response = &http.Response{
		StatusCode: http.StatusTeapot + 1,
	}

	request := httptest.NewRequest("GET", "/path", nil)
	response, _ := acf.client.Do(request)

	if acf.So(response, should.NotBeNil) {
		acf.So(response.StatusCode, should.Equal, http.StatusTeapot+1)
	}

}
