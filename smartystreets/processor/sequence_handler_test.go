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
	close(shf.input)

	shf.handler.Handle()

	shf.So(<-shf.output, should.Equal, envelope)
}

func (shf *SequenceHandlerFixture) TestEnvelopesReceivedOutOfOrder_BufferedUntilContiguousBlock() {
	shf.input <- &Envelope{Sequence: 4}
	shf.input <- &Envelope{Sequence: 2}
	shf.input <- &Envelope{Sequence: 0}
	shf.input <- &Envelope{Sequence: 3}
	shf.input <- &Envelope{Sequence: 1}
	close(shf.input)

	shf.handler.Handle()

	close(shf.output)

	shf.So(shf.sequenceOrder(), should.Resemble, []int{0, 1, 2, 3, 4})
	shf.So(shf.handler.buffer, should.BeEmpty)
}

func (shf *SequenceHandlerFixture) sequenceOrder() (order []int) {
	for envelope := range shf.output {
		order = append(order, envelope.Sequence)
	}
	return order
}
