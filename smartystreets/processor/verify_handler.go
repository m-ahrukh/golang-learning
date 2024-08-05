package processor

type VerifyHander struct {
	in          chan *Envelope
	out         chan *Envelope
	application Verifier
}

type Verifier interface {
	Verify(AddressInput)
}

func NewVerifyHandler(in, out chan *Envelope, application Verifier) *VerifyHander {
	return &VerifyHander{
		in:          in,
		out:         out,
		application: application,
	}
}

func (verifier *VerifyHander) Handle() {
	envelope := <-verifier.in

	verifier.application.Verify(envelope.Input)

	verifier.out <- envelope
}
