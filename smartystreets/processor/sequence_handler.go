package processor

type SequenceHandler struct {
	input  chan *Envelope
	output chan *Envelope
}

func NewSequenceHandler(input, output chan *Envelope) *SequenceHandler {
	return &SequenceHandler{
		input:  input,
		output: output,
	}
}

func (handler *SequenceHandler) Handle() {
	input := <-handler.input
	handler.output <- input
	// handler.output <- <- handler.input
}
