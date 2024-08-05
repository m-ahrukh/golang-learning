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

	input       chan interface{}
	output      chan interface{}
	application *FakeVerifier
	handler     *VerifyHander
}

func (handlerFixture *HandlerFixture) Setup() {
	handlerFixture.input = make(chan interface{}, 10)
	handlerFixture.output = make(chan interface{}, 10)
	handlerFixture.application = NewFakeVerifier()
	handlerFixture.handler = NewVerifyHandler(handlerFixture.input, handlerFixture.output, handlerFixture.application)

}
func (handlerFixture *HandlerFixture) TestVeriferRecievesInput() {

	handlerFixture.input <- 42
	close(handlerFixture.input)

	handlerFixture.handler.Handle()

	handlerFixture.AssertEqual(42, <-handlerFixture.output)
	handlerFixture.AssertEqual(42, handlerFixture.application.input)
}

// ////////////////////////////////////////////////////////
type FakeVerifier struct {
	input interface{}
}

func NewFakeVerifier() *FakeVerifier {
	return &FakeVerifier{}
}

func (fakeVarifier *FakeVerifier) Verify(value interface{}) {
	fakeVarifier.input = value
}
