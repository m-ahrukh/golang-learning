package processor

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestSequenceHandler(t *testing.T) {
	gunit.Run(new(SequenceHandlerFixture), t)
}

type SequenceHandlerFixture struct {
	*gunit.Fixture

	input   chan *Envelope
	output  chan *Envelope
	handler *SequenceHandler
}

func (shf *SequenceHandlerFixture) Setup() {
	shf.input = make(chan *Envelope, 10)
	shf.output = make(chan *Envelope, 10)
	shf.handler = NewSequenceHandler(shf.input, shf.output)
}

func (shf *SequenceHandlerFixture) TestExpectedEnvelopeSentToOutput() {
	envelope := &Envelope{Sequence: 0}
	shf.input <- envelope
	shf.handler.Handle()

	shf.So(<-shf.output, should.Equal, envelope)
}
