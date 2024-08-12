package processor

import (
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

func (acf *AuthenticationClientFixture) TestHostnameAndSchema() {
	request := httptest.NewRequest("GET", "/path", nil)

	acf.client.Do(request)

	acf.So(acf.inner.request.Host, should.Equal, "different-company.com")
	acf.So(acf.inner.request.URL.Scheme, should.Equal, "http")
}
