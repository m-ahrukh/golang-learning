package processor

type VerifyHander struct {
	in          chan interface{}
	out         chan interface{}
	application Verifier
}

type Verifier interface {
	Verify(interface{})
}

func NewVerifyHandler(in, out chan interface{}, application Verifier) *VerifyHander {
	return &VerifyHander{
		in:          in,
		out:         out,
		application: application,
	}
}

func (verifier *VerifyHander) Handle() {
	recieved := <-verifier.in

	verifier.application.Verify(recieved)

	verifier.out <- recieved
}
