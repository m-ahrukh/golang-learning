package processor

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestPipelineFixture(t *testing.T) {
	gunit.Run(new(PipelineFixture), t)
}

type PipelineFixture struct {
	*gunit.Fixture

	reader   *ReadWriteSpyBuffer
	writer   *ReadWriteSpyBuffer
	client   *IntegrationHTTPClient
	pipeline *Pipeline
}

func (pf *PipelineFixture) Setup() {
	pf.reader = NewReadWriteSpyBuffer("")
	pf.writer = NewReadWriteSpyBuffer("")
	pf.client = &IntegrationHTTPClient{}
	pf.pipeline = NewPipeline(io.NopCloser(pf.reader), pf.writer, pf.client, 2)
}

func (pf *PipelineFixture) LongTestPipeline() {
	fmt.Fprintln(pf.reader, ("Street1,City,State,ZIPCode"))
	fmt.Fprintln(pf.reader, ("A,B,C,D"))
	fmt.Fprintln(pf.reader, ("A,B,C,D"))

	err := pf.pipeline.Process()

	pf.So(pf.writer.String(), should.Equal,
		"Status,DeliveryLine1,LastLine,City,State,ZIPCode\n"+
			"Deliverable,AA,BB,CC,DD,EE\n"+
			"Deliverable,AA,BB,CC,DD,EE\n")

	pf.So(err, should.BeNil)
}

type IntegrationHTTPClient struct{}

func (pf *IntegrationHTTPClient) Do(request *http.Request) (*http.Response, error) {
	return &http.Response{
		Body:       NewReadWriteSpyBuffer(integrationJSONOutput),
		StatusCode: http.StatusOK,
	}, nil
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
		"analysis": {
			"dpv_match_code": "Y",
			"dpv_vacant": "N",
			"active": "Y"
		}
	}
]`
