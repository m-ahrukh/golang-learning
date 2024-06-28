package main

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill
func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{
			"Ice Cream":            50.00,
			"Lays":                 60.00,
			"Knorr Noodles Masala": 40.00,
		},
		tip: 10.0,
	}
	return b
}
