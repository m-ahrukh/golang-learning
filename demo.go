package main

import "fmt"

// fmt is used for IO operations

func main() { //main function is the entry point of the go code
	// fmt.Println("Hello World")

	//Variables
	//var variableName dataType = value
	//variableName := value

	//if we declare any variable and dosn't use it, it will give error.

	//Strings
	var firstName string = "Mahrukh"
	var secondName = "Ameen" // it can automatically detect that it is a string
	var thirdName string     // default value of string will be null
	fmt.Println(firstName, secondName, thirdName)
	myName := "Mahrukh" // shorthand of var myName string = "Mahrukh"
	fmt.Println(myName)

	//Integers
	var age1 int = 24
	var age2 = 27
	var age3 int //default value of int is 0
	age4 := 3
	fmt.Println(age1, age2, age3, age4)

	// bits and memory
	var num1 int8 = 25 //8 bit number ranges from -127 to 128
	// var num2 int8 = 215 //gives error that number is not in the scope
	var num2 int16 = 215
	var num3 uint = 25  //unsigned int
	var num4 uint8 = 25 //unsigned 8 bit int ranging from 0 to 255
	// var num4 uint = -25 //gves error as number is signed number
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)

	//floats
	var score1 float32 = 25.98
	var score2 float64 = 7837978970920909.7
	score3 := 9090.99 //by efault it take float64 datatype
	fmt.Println(score1, score2, score3)
}
