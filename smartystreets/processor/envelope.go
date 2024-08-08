package processor

type Envelope struct {
	Input  AddressInput
	Output AddressOutput
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
