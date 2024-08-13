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
	return &WriterHandler{
		input:  input,
		closer: output,
		writer: csv.NewWriter(output),
	}
}

func (wh *WriterHandler) Handle() {
	wh.writer.Write([]string{"Status", "DeliveryLine1", "City", "State", "ZIPCode"})

	for envelope := range wh.input {
		output := envelope.Output
		wh.writeAddressOutput(output)
	}

	// envelope := <-wh.input
	// if envelope != nil {
	// 	output := envelope.Output
	// 	wh.writeAddressOutput(output)
	// }

	wh.writer.Flush()
	wh.closer.Close()
}

func (wh *WriterHandler) writeAddressOutput(output AddressOutput) {
	wh.writeValues(
		output.Status,
		output.DeliveryLine1,
		output.City,
		output.LastLine,
		output.State,
		output.ZIPCode,
	)
}
func (wh WriterHandler) writeValues(values ...string) {
	wh.writer.Write(values)
}
