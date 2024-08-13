package processor

import (
	"encoding/csv"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestWriterHandlerFixture(t *testing.T) {
	gunit.Run(new(WriterHandlerFixture), t)
}

type WriterHandlerFixture struct {
	*gunit.Fixture

	handler *WriterHandler
	input   chan *Envelope
	buffer  *WriterSpyBuffer
	writer  *csv.Writer
}

func (whf *WriterHandlerFixture) Setup() {
	whf.buffer = NewWriterSpyBuffer("")
	whf.input = make(chan *Envelope, 10)
	whf.handler = NewWriterHandler(whf.input, whf.buffer)
}

func (whf *WriterHandlerFixture) TestHeaderWritten() {
	whf.handler.Handle()

	whf.So(whf.buffer.String(), should.Equal, "Status,DeliveryLine1,City,State,ZIPCode\n")
}

func (whf *WriterHandlerFixture) TestOuputClosed() {
	whf.handler.Handle()

	whf.So(whf.buffer.closed, should.Equal, 1)
}
