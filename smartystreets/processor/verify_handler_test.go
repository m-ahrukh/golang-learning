package processor

import "testing"

func TestVeriferRecievesInput(t *testing.T) {
	handler, out := setup()

	handler.Listen()

	recieved := <-out
	if recieved != 1 {
		t.Errorf("\nGot %v\nwant: 1", recieved)
	}
}

func setup() (handler *VerifyHander, out chan interface{}) {
	in := make(chan interface{}, 10)
	out = make(chan interface{}, 10)
	verifier := NewVerifyHandler(in, out)
	in <- 1
	close(in)

	return verifier, out
}
