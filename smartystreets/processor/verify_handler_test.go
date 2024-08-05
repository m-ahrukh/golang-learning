package processor

import (
	"testing"

	"github.com/smarty/gunit"
)

func TestHandlerFixture(t *testing.T) {
	gunit.Run(new(HandlerFixture), t)
}

type HandlerFixture struct {
	*gunit.Fixture

	input       chan *Envelope
	output      chan *Envelope
	application *FakeVerifier
	handler     *VerifyHander
}

func (handlerFixture *HandlerFixture) Setup() {
	handlerFixture.input = make(chan *Envelope, 10)
	handlerFixture.output = make(chan *Envelope, 10)
	handlerFixture.application = NewFakeVerifier()
	handlerFixture.handler = NewVerifyHandler(handlerFixture.input, handlerFixture.output, handlerFixture.application)

}
func (handlerFixture *HandlerFixture) TestVeriferRecievesInput() {

	envelope := &Envelope{
		Input: AddressInput{
			Street1: "42",
		},
	}
	handlerFixture.input <- envelope
	close(handlerFixture.input)

	handlerFixture.handler.Handle()

	handlerFixture.AssertEqual(envelope, <-handlerFixture.output)
	handlerFixture.AssertEqual(envelope.Input, handlerFixture.application.input)
}

// ////////////////////////////////////////////////////////
type FakeVerifier struct {
	input AddressInput
}

func NewFakeVerifier() *FakeVerifier {
	return &FakeVerifier{}
}

func (fakeVarifier *FakeVerifier) Verify(value AddressInput) {
	fakeVarifier.input = value
}
