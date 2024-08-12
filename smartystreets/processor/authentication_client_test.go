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
	acf.client = &AuthenticationClient{}
}

func (acf *AuthenticationClientFixture) TestHostnameAndSchema() {
	acf.client = NewAuthenticationClient(acf.inner, "us-street.api.smartystreets.com", "HOSTNAME")
	request := httptest.NewRequest("GET", "/path", nil)

	acf.client.Do(request)

	acf.So(acf.inner.request.Host, should.Equal, "us-street.api.smartystreets.com")
	acf.So(acf.inner.request.URL.Scheme, should.Equal, "https")
}
