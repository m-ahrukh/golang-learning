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
	counter := 0
	var buffer []*Envelope
	for envelope := range handler.input {
		if envelope.Sequence == counter {
			handler.output <- envelope
			counter++
			if len(buffer) > 0 {
				handler.output <- buffer[0]
				counter++
			}
		} else {
			buffer = append(buffer, envelope)
		}
	}
	// input := <-handler.input
	// handler.output <- input
	// handler.output <- <- handler.input
}
