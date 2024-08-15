package processor

type SequenceHandler struct {
	input   chan *Envelope
	output  chan *Envelope
	counter int
	buffer  map[int]*Envelope
}

func NewSequenceHandler(input, output chan *Envelope) *SequenceHandler {
	return &SequenceHandler{
		input:   input,
		output:  output,
		counter: initialSequenceValue,
		buffer:  make(map[int]*Envelope),
	}
}

func (handler *SequenceHandler) Handle() {
	for envelope := range handler.input {
		handler.processEnvelope(envelope)
	}
	// close(handler.input)
	close(handler.output)
}

func (handler *SequenceHandler) processEnvelope(envelope *Envelope) {
	handler.buffer[envelope.Sequence] = envelope
	handler.sendBufferedEnvelopesInOrder()
}

func (handler *SequenceHandler) sendBufferedEnvelopesInOrder() {
	for {
		next, found := handler.buffer[handler.counter]
		if !found {
			break
		}
		handler.sendNextEnvelope(next)
	}
}

func (handler *SequenceHandler) sendNextEnvelope(envelope *Envelope) {
	if envelope.EOF {
		close(handler.input)
	} else {
		handler.output <- envelope
	}
	delete(handler.buffer, handler.counter)
	handler.counter++
}
