package processor

type VerifyHander struct {
	input    chan *Envelope
	output   chan *Envelope
	verifier Verifier
}

type Verifier interface {
	Verify(AddressInput) AddressOutput
}

func NewVerifyHandler(input, output chan *Envelope, verifier Verifier) *VerifyHander {
	return &VerifyHander{
		input:    input,
		output:   output,
		verifier: verifier,
	}
}

func (verifier *VerifyHander) Handle() {
	//looping over the channel
	for envelope := range verifier.input {
		envelope.Output = verifier.verifier.Verify(envelope.Input)
		verifier.output <- envelope
	}
}
