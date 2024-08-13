package processor

type SequenceHandler struct {
	input   chan *Envelope
	output  chan *Envelope
	counter int
	buffer  map[int]*Envelope
}

func NewSequenceHandler(input, output chan *Envelope) *SequenceHandler {
	return &SequenceHandler{
		input:  input,
		output: output,
		buffer: make(map[int]*Envelope),
	}
}

func (handler *SequenceHandler) Handle() {

	for envelope := range handler.input {
		handler.processEnvelope(envelope)
	}
	handler.buffer = make(map[int]*Envelope)
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
		handler.output <- next
		delete(handler.buffer, handler.counter)
		handler.counter++
	}
}
