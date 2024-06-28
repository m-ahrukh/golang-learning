package main

import "fmt"

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

// Here we are going to create a function named format which is asscoaited with the bill struct which will be a reciever function
// format the bill
func (b *bill) format() string { //to make it associated with bill struct we need to specify the bill object
	//here (b bill) is recieved in the function
	fs := "Bill breakdown: \n" //fs is formatted string
	var total float64 = 0

	//list items
	for k, value := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", value)
		total += value
	}

	//add tip
	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)

	//total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return fs
}

//update tip
func (b *bill) updateTip(tip float64) { //here we are passing the refernece of the bill
	(*b).tip = tip //here we dereference the bill so that it can update the value and not making the copy of the object
	//b.tip = tip // this cal also update the value but it is not a good practice
}

//add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

//pointers aur automatically dereferenced
