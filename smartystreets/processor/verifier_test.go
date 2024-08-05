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

func (this *VerifierFixture) Setup() {
	this.client = &FakeHTTPClient{}
	this.verifier = NewSmartyVerifier(this.client)
}

func NewSmartyVerifier(client HTTPClient) *SmartyVerifier {
	return &SmartyVerifier{
		client: client,
	}
}

func (this *VerifierFixture) Test() {

}

// ///////////////////////////////////////////////////////
type FakeHTTPClient struct {
}

func (this *FakeHTTPClient) Do(*http.Request) (*http.Response, error) {
	return nil, nil
}
