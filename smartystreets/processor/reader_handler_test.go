package processor

import (
	"strconv"
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

func (rhf *ReaderHandlerFixture) TestAllCSVRecordsSentToOutput() {
	rhf.writeLine("A1,B1,C1,D1")
	rhf.writeLine("A2,B2,C2,D2")

	rhf.reader.Handle()

	rhf.assertRecordSent()
	rhf.assertCleanup()
}

func (rhf *ReaderHandlerFixture) assertRecordSent() {
	rhf.So(<-rhf.output, should.Resemble, buildEnvelope(initialSequenceValue))
	rhf.So(<-rhf.output, should.Resemble, buildEnvelope(initialSequenceValue+1))
}

func (rhf *ReaderHandlerFixture) assertCleanup() {
	rhf.So(<-rhf.output, should.Equal, endOfFile)
	rhf.So(<-rhf.output, should.BeNil)
	rhf.So(rhf.buffer.closed, should.Equal, 1)
}

func (rhf *ReaderHandlerFixture) writeLine(line string) {
	rhf.buffer.WriteString(line + "\n")
}

func buildEnvelope(index int) *Envelope {
	suffix := strconv.Itoa(index + 1)
	return &Envelope{
		Sequence: index,
		Input: AddressInput{
			Street1: "A" + suffix,
			City:    "B" + suffix,
			State:   "C" + suffix,
			ZIPCode: "D" + suffix,
		},
	}
}

func (rhf *ReaderHandlerFixture) TestMalformedInputReturnsError() {
	malformedRecord := "A1" //too short
	rhf.writeLine(malformedRecord)

	err := rhf.reader.Handle()

	rhf.So(err, should.NotBeNil)
}
