package processor

import (
	"testing"

	"github.com/smarty/gunit"
)

func TestHandlerFixture(t *testing.T) {
	gunit.Run(new(HandlerFixture), t)
}

type HandlerFixture struct {
	*gunit.Fixture

	in      chan interface{}
	out     chan interface{}
	handler *VerifyHander
}

func (handlerFixture *HandlerFixture) Setup() {
	handlerFixture.in = make(chan interface{}, 10)
	handlerFixture.out = make(chan interface{}, 10)
	handlerFixture.handler = NewVerifyHandler(handlerFixture.in, handlerFixture.out)

}
func (handlerFixture *HandlerFixture) TestVeriferRecievesInput() {

	handlerFixture.in <- 1
	close(handlerFixture.in)

	handlerFixture.handler.Listen()

	// recieved := <-handlerFixture.out
	// if recieved != 1 {
	// 	handlerFixture.Errorf("\nGot %v\nwant: 1", recieved)
	// }

	handlerFixture.AssertEqual(1, <-handlerFixture.out)
}
