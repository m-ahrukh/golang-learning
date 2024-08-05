package processor

import (
	"strings"
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
	envelope := handlerFixture.enqueueEnvelope("street")
	close(handlerFixture.input)

	handlerFixture.handler.Handle()

	handlerFixture.AssertEqual("STREET", envelope.Output.DeliveryLine1)
	handlerFixture.AssertEqual(envelope, <-handlerFixture.output)
}

func (handlerFixture *HandlerFixture) enqueueEnvelope(street1 string) *Envelope {
	envelope := &Envelope{
		Input: AddressInput{
			Street1: street1,
		},
	}
	handlerFixture.input <- envelope

	return envelope
}

func (handlerFixture *HandlerFixture) TestInputQueueDrained() {
	envelope1 := handlerFixture.enqueueEnvelope("41")
	envelope2 := handlerFixture.enqueueEnvelope("42")
	envelope3 := handlerFixture.enqueueEnvelope("43")
	//read from the channel until the channel is closed.
	//if we do not close the channel, it will cause the
	//deadlock that says all go routines are asleep
	close(handlerFixture.input)

	handlerFixture.handler.Handle()

	handlerFixture.AssertEqual(envelope1, <-handlerFixture.output)
	handlerFixture.AssertEqual(envelope2, <-handlerFixture.output)
	handlerFixture.AssertEqual(envelope3, <-handlerFixture.output)
}

// ////////////////////////////////////////////////////////
type FakeVerifier struct {
	input  AddressInput
	output AddressOutput
}

func NewFakeVerifier() *FakeVerifier {
	return &FakeVerifier{}
}

func (fakeVarifier *FakeVerifier) Verify(value AddressInput) AddressOutput {
	fakeVarifier.input = value
	return AddressOutput{DeliveryLine1: strings.ToUpper(value.Street1)}
}
