package processor

import (
	"net/http"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestWireupFixture(t *testing.T) {
	gunit.Run(new(WireupFixture), t)
}

type WireupFixture struct {
	*gunit.Fixture

	reader   *ReadWriteSpyBuffer
	writer   *ReadWriteSpyBuffer
	client   *FakeHTTPClient
	pipeline *Pipeline
}

func (this *WireupFixture) Setup() {
	this.reader = NewReadWriteSpyBuffer("")
	this.writer = NewReadWriteSpyBuffer("")
	this.client = &FakeHTTPClient{}

	this.pipeline = Configure(this.reader, this.writer, this.client, 2)
}

func (this *WireupFixture) LongTestPipeline() {
	this.client.Configure(integrationJSONOutput, http.StatusOK, nil)
	this.reader.WriteString("Street1,City,State,ZIPCode")
	this.reader.WriteString("A,B,C,D")
	this.reader.WriteString("A,B,C,D")

	err := this.pipeline.Process()

	this.So(this.writer.String(), should.Equal,
		"Status,DeliveryLine1,LastLine,City,State,ZIPCode\n"+
			"Deliverable,AA,BB,CC,DD,EE\n"+
			"Deliverable,AA,BB,CC,DD,EE\n")

	this.So(err, should.BeNil)
}

const integrationJSONOutput = `
[
	{
		"delivery_line_1": "AA",
		"last_line": "BB",
		"components": {
			"city_name": "CC",
			"state_abbreviation": "DD",
			"zipcode": "EE"
		},
		{
			"analysis": {
				"dpv_match_code": "Y",
				"dpv_vacant": "N",
				"active": "Y"
			}
		}
	}
]`
