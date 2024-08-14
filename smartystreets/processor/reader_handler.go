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

	for {
		record, err := rh.reader.Read()
		if err == io.EOF {
			break
		} else {
			//TODO: warn user of malformed file???
		}
		rh.output <- &Envelope{
			Input: createinput(record),
		}
	}
}

func createinput(record []string) AddressInput {
	return AddressInput{
		Street1: record[0],
		City:    record[1],
		State:   record[2],
		ZIPCode: record[3],
	}
}

func (rh *ReaderHandler) skipHeader() {
	rh.reader.Read()
}
