package main

import (
	"fmt"
	"sort"
	"strings"
)

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

	//Printing and Formatting

	//Print
	fmt.Print("Hello, ")
	fmt.Print("world\n")

	//println
	fmt.Println("Hello, ")
	fmt.Println("world")
	fmt.Println("My name is", myName, "and my age is", age1)

	//formatted String (printf) -> %_ = format specifier
	fmt.Printf("my age is %v and my name is %s\n", age1, myName)
	fmt.Printf("my age is %q and my name is name is %q\n", age2, myName)
	fmt.Printf("age is of type %T\n", age1)
	fmt.Printf("you scored %f points! \n", 255.5)
	fmt.Printf("you scored %0.1f points! \n", 255.5)

	//Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v\n", age1, myName)
	fmt.Println("saved string is:", str)

	//arrays and slices
	var ages [3]int = [3]int{24, 27, 49} //arrays of fixed size
	names := [4]string{"Mahrukh", "Ali", "Malaika", "Kalsoom"}
	// var ages1 = [3]int{24, 27, 49}
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	//slices (use arrays under the hood)
	var scores = []int{100, 50, 60} ///we can change the size of the array s we didn't define the size
	scores[2] = 99
	scores = append(scores, 85)
	fmt.Println(scores)

	//slice range
	range1 := names[1:3] //starting from position: to that position -1
	range2 := names[2:]
	range3 := names[:3]

	range1 = append(range1, "Ameen")
	fmt.Println(range1)
	fmt.Println(range2)
	fmt.Println(range3)

	greetings := "hello there"
	fmt.Println(strings.ReplaceAll(greetings, "hello", "hi"))
	fmt.Println(strings.ToUpper(greetings))
	fmt.Println(strings.Index(greetings, "th"))
	fmt.Println(strings.Split(greetings, " "))

	fmt.Println("original string ", greetings)

	ageArray := []int{88, 67, 22, 40, 90, 10, 45, 67}
	sort.Ints(ageArray) // sort the slice of integers
	fmt.Println(ageArray)

	index := sort.SearchInts(ageArray, 40) //first sort the array and then search the element
	fmt.Println(index)

	namesArray := []string{"babar", "mahrukh", "malaika", "ali"}
	sort.Strings(namesArray) //sort the strings slice
	fmt.Println(namesArray)

	fmt.Println(sort.SearchStrings(namesArray, "babar"))
}
