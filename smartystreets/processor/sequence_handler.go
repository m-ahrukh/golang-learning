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
	// var buffer []*Envelope
	var sequence = make(map[int]*Envelope)

	for envelope := range handler.input {
		sequence[envelope.Sequence] = envelope

		for {
			next, found := sequence[counter]
			if !found {
				break
			}
			handler.output <- next
			counter++
		}

		// if envelope.Sequence == counter {
		// 	handler.output <- envelope
		// 	counter++
		// 	if len(buffer) > 0 {
		// 		handler.output <- buffer[0]
		// 		counter++
		// 	}
		// } else {
		// 	buffer = append(buffer, envelope)
		// }
	}
	// input := <-handler.input
	// handler.output <- input
	// handler.output <- <- handler.input
}
