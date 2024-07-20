package main

import "fmt"

/*
//simple main function prints "Hello, World"
func main() {
	fmt.Println("Hello, World")
}
*/

func Hello() string {
	return "Hello, World"
}

func main() {
	fmt.Println(Hello())
}
