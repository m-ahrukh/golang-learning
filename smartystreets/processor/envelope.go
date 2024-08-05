package processor

type Envelope struct {
	Input  AddressInput
	Output AddressOutput
}

type AddressInput struct {
	Street1 string
}

type AddressOutput struct {
	DeliveryLine1 string
}