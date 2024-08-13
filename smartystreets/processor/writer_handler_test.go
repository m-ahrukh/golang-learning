package processor

import (
	"bytes"
	"encoding/csv"
	"strings"
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
	close(whf.input)
	whf.handler.Handle()

	whf.So(whf.buffer.String(), should.Equal, "Status,DeliveryLine1,City,State,ZIPCode\n")
}

func (whf *WriterHandlerFixture) TestOuputClosed() {
	close(whf.input)
	whf.handler.Handle()

	whf.So(whf.buffer.closed, should.Equal, 1)
}

func (whf *WriterHandlerFixture) TestEnvelopeWritten() {
	whf.input <- &Envelope{
		Output: AddressOutput{
			Status:        "A",
			DeliveryLine1: "B",
			City:          "C",
			State:         "D",
			ZIPCode:       "E",
			LastLine:      "F",
		},
	}
	close(whf.input)

	whf.handler.Handle()
	outputFile := strings.TrimSpace(whf.buffer.String())
	lines := strings.Split(outputFile, "\n")
	if whf.So(lines, should.HaveLength, 2) {
		whf.So(lines[1], should.Equal, "A,B,C,D,E,F")
	}
}

// /////////////////////////////////////////////////////////
type WriterSpyBuffer struct {
	*bytes.Buffer
	closed int
}

func NewWriterSpyBuffer(value string) *WriterSpyBuffer {
	return &WriterSpyBuffer{
		Buffer: bytes.NewBufferString(value),
	}
}

func (spyBuffer *WriterSpyBuffer) Close() error {
	spyBuffer.closed++
	return nil
}
