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
			LastLine:      "D",
			State:         "E",
			ZIPCode:       "F",
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

func (whf *WriterHandlerFixture) TestAllEnvelopeWritten() {
	whf.input <- &Envelope{
		Output: AddressOutput{
			Status:        "A1",
			DeliveryLine1: "B1",
			City:          "C1",
			LastLine:      "D1",
			State:         "E1",
			ZIPCode:       "F1",
		},
	}
	whf.input <- &Envelope{
		Output: AddressOutput{
			Status:        "A2",
			DeliveryLine1: "B2",
			City:          "C2",
			LastLine:      "D2",
			State:         "E2",
			ZIPCode:       "F2",
		},
	}
	close(whf.input)

	whf.handler.Handle()
	outputFile := strings.TrimSpace(whf.buffer.String())
	lines := strings.Split(outputFile, "\n")
	if whf.So(lines, should.HaveLength, 3) {
		whf.So(lines[1], should.Equal, "A1,B1,C1,D1,E1,F1")
		whf.So(lines[2], should.Equal, "A2,B2,C2,D2,E2,F2")
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
