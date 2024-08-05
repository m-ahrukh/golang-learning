package processor

import "testing"

type Verifier struct {
	in  chan interface{}
	out chan interface{}
}

func NewVerifier(in, out chan interface{}) *Verifier {
	return &Verifier{
		in:  in,
		out: out,
	}
}

func (verifier *Verifier) Listen() {
	verifier.out <- 1

}

func TestVeriferRecievesInput(t *testing.T) {
	in := make(chan interface{}, 10)
	out := make(chan interface{}, 10)
	verifier := NewVerifier(in, out)

	in <- 1

	close(in)

	verifier.Listen()

	recieved := <-out
	if recieved != 1 {
		t.Errorf("\nGot %v\nwant: 1", recieved)
	}
}
