package processor

type VerifyHander struct {
	in  chan interface{}
	out chan interface{}
}

func NewVerifyHandler(in, out chan interface{}) *VerifyHander {
	return &VerifyHander{
		in:  in,
		out: out,
	}
}

func (verifier *VerifyHander) Listen() {
	verifier.out <- 9

}
