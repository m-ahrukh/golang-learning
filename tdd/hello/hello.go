package main

import "fmt"

/*
//simple main function prints "Hello, World"
func main() {
	fmt.Println("Hello, World")
}
*/

/*
//function without any parameter
func Hello() string {
	return "Hello, World"
}
*/

/*
//function with parameter name
func Hello(name string) string {
	return "Hello, " + name
}
*/

/*
//replace string with Constant
const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name
}
*/

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Mahrukh"))
}
