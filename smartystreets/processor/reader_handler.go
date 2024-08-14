package processor

import (
	"encoding/csv"
	"io"
)

type ReaderHandler struct {
	reader *csv.Reader
	closer io.Closer
	output chan *Envelope
}

func NewReaderHandler(reader io.ReadCloser, output chan *Envelope) *ReaderHandler {
	return &ReaderHandler{
		reader: csv.NewReader(reader),
		closer: reader,
		output: output,
	}
}

func (rh *ReaderHandler) Handle() {
	rh.skipHeader()

	record, _ := rh.reader.Read()
	envelope := &Envelope{
		Input: AddressInput{
			Street1: record[0],
			City:    record[1],
			State:   record[2],
			ZIPCode: record[3],
		},
	}

	rh.output <- envelope
}

func (rh *ReaderHandler) skipHeader() {
	rh.reader.Read()
}
