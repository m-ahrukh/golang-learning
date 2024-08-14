package processor

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestReaderHandlerFixture(t *testing.T) {
	gunit.Run(new(ReaderHandlerFixture), t)
}

type ReaderHandlerFixture struct {
	*gunit.Fixture

	buffer *ReadWriteSpyBuffer
	output chan *Envelope
	reader *ReaderHandler
}

func (rhf *ReaderHandlerFixture) Setup() {
	rhf.buffer = NewReadWriteSpyBuffer("")
	rhf.output = make(chan *Envelope, 10)
	rhf.reader = NewReaderHandler(rhf.buffer, rhf.output)

	const header = "Street1,City,State,ZIPCode"
	rhf.writeLine(header)
}

func (rhf *ReaderHandlerFixture) TestCSVRecordSentInEnvelope() {
	rhf.writeLine("A,B,C,D")

	rhf.reader.Handle()

	rhf.So(<-rhf.output, should.Resemble, &Envelope{
		Sequence: initialSequenceValue,
		Input: AddressInput{
			Street1: "A",
			City:    "B",
			State:   "C",
			ZIPCode: "D",
		},
	})
}

func (rhf *ReaderHandlerFixture) TestAllCSVRecordsWrittenToOutput() {
	rhf.writeLine("A1,B1,C1,D1")
	rhf.writeLine("A2,B2,C2,D2")

	rhf.reader.Handle()

	rhf.So(<-rhf.output, should.Resemble, &Envelope{
		Sequence: initialSequenceValue,
		Input: AddressInput{
			Street1: "A1",
			City:    "B1",
			State:   "C1",
			ZIPCode: "D1",
		},
	})

	rhf.So(<-rhf.output, should.Resemble, &Envelope{
		Sequence: initialSequenceValue + 1,
		Input: AddressInput{
			Street1: "A2",
			City:    "B2",
			State:   "C2",
			ZIPCode: "D2",
		},
	})

	rhf.So(<-rhf.output, should.Resemble, &Envelope{Sequence: eofSequenceValue})
	rhf.So(<-rhf.output, should.BeNil)
	rhf.So(rhf.buffer.closed, should.Equal, 1)
}

func (rhf *ReaderHandlerFixture) writeLine(line string) {
	rhf.buffer.WriteString(line + "\n")
}
