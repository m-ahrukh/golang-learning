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
}

func (rhf *ReaderHandlerFixture) Setup() {

}

func (rhf *ReaderHandlerFixture) TestCSVRecordSentInEnvelope() {
	buffer := NewReadWriteSpyBuffer("Street1,City,State,ZIPCode\n" + "A,B,C,D\n")
	output := make(chan *Envelope, 10)
	reader := NewReaderHandler(buffer, output)

	reader.Handle()

	rhf.So(<-output, should.Resemble, &Envelope{Input: AddressInput{
		Street1: "A",
		City:    "B",
		State:   "C",
		ZIPCode: "D",
	}})
}
