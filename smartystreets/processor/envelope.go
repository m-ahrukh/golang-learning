package processor

const (
	initialSequenceValue = 0
	eofSequenceValue     = -1
)

var endOfFile = &Envelope{Sequence: eofSequenceValue}

type Envelope struct {
	Input    AddressInput
	Output   AddressOutput
	Sequence int
}

type AddressInput struct {
	Street1 string
	City    string
	State   string
	ZIPCode string
}

type AddressOutput struct {
	Status        string
	DeliveryLine1 string
	LastLine      string
	City          string
	State         string
	ZIPCode       string
}
