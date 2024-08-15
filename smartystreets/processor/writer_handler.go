package processor

import (
	"encoding/csv"
	"io"
)

type WriterHandler struct {
	input  chan *Envelope
	closer io.Closer
	writer *csv.Writer
}

func NewWriterHandler(input chan *Envelope, output io.WriteCloser) *WriterHandler {
	this := &WriterHandler{
		input:  input,
		closer: output,
		writer: csv.NewWriter(output),
	}

	this.writeValues("Status", "DeliveryLine1", "LastLine", "City", "State", "ZIPCode")

	return this
}

func (wh *WriterHandler) Handle() {

	for envelope := range wh.input {
		// if !envelope.EOF {
		output := envelope.Output
		wh.writeAddressOutput(output)
		// }
	}

	wh.writer.Flush()
	wh.closer.Close()
}

func (wh *WriterHandler) writeAddressOutput(output AddressOutput) {
	wh.writeValues(
		output.Status,
		output.DeliveryLine1,
		output.LastLine,
		output.City,
		output.State,
		output.ZIPCode,
	)
}
func (wh WriterHandler) writeValues(values ...string) {
	wh.writer.Write(values)
}
