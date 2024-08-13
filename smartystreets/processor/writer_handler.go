package processor

import (
	"bytes"
	"encoding/csv"
	"io"
)

type WriterHandler struct {
	input  chan *Envelope
	closer io.Closer
	writer *csv.Writer
}

func NewWriterHandler(input chan *Envelope, output io.WriteCloser) *WriterHandler {
	return &WriterHandler{
		input:  input,
		closer: output,
		writer: csv.NewWriter(output),
	}
}

func (wh *WriterHandler) Handle() {
	wh.writer.Write([]string{"Status", "DeliveryLine1", "City", "State", "ZIPCode"})
	wh.writer.Flush()
	wh.closer.Close()
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
