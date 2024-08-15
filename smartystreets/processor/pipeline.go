package processor

import (
	"io"
)

type Pipeline struct {
	reader  io.ReadCloser
	writer  io.WriteCloser
	client  HTTPClient
	workers int

	verifier      Verifier
	verifyInput   chan *Envelope
	sequenceInput chan *Envelope
	writerInput   chan *Envelope
}

func NewPipeline(reader io.ReadCloser, writer io.WriteCloser, client HTTPClient, workers int) *Pipeline {
	return &Pipeline{
		reader:  reader,
		writer:  writer,
		client:  client,
		workers: workers,

		verifier:      NewSmartyVerifier(client),
		verifyInput:   make(chan *Envelope, 1024),
		sequenceInput: make(chan *Envelope, 1024),
		writerInput:   make(chan *Envelope, 1024),
	}
}

func (pipeline *Pipeline) Process() (err error) {
	pipeline.startVerifyHandlers()

	go func() {
		err = NewReaderHandler(pipeline.reader, pipeline.verifyInput).Handle()
	}()

	pipeline.startSequence()
	pipeline.awaitWriteHandler()

	return err
}

func (pipeline *Pipeline) startVerifyHandlers() {
	for i := 0; i < pipeline.workers; i++ {
		go NewVerifyHandler(pipeline.verifyInput, pipeline.sequenceInput, pipeline.verifier).Handle()
	}
}

func (pipeline *Pipeline) startSequence() {
	go NewSequenceHandler(pipeline.sequenceInput, pipeline.writerInput).Handle()
}

func (pipeline *Pipeline) awaitWriteHandler() {
	NewWriterHandler(pipeline.writerInput, pipeline.writer).Handle()
}
