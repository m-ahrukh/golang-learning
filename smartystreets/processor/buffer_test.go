package processor

import (
	"encoding/csv"
	"io"
	"strings"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestSpyBuffer(t *testing.T) {
	gunit.Run(new(SpyBuffer), t)
}

type SpyBuffer struct {
	*gunit.Fixture
}

func (this *SpyBuffer) Setup() {

}

func (this *SpyBuffer) Test() {
	buffer := NewReadWriteSpyBuffer("")
	buffer.WriteString("Hello, World!")

	this.So(buffer.String(), should.Equal, "Hello, World!")
	raw, err := io.ReadAll(buffer)
	this.So(string(raw), should.Equal, "Hello, World!")
	this.So(err, should.BeNil)

	reader := csv.NewReader(strings.NewReader("Hello, World!"))
	record, err2 := reader.Read()
	this.So(record, should.Resemble, []string{"Hello", " World!"})
	this.So(err2, should.BeNil)
}
