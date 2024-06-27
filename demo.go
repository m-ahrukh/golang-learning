package main

import (
	"fmt" // fmt is used for IO operations
	"math"
	"sort"
	"strings"
)

var score = 99.5 //if we define this in our main function and then access this in greeting file, it will give error that score is undefined. this happened because its scope is changed

func main() { //main function is the entry point of the go code
	// fmt.Println("Hello World")

	// learnVariables()
	// learnBitsAndMemory()
	// arraysAndSlies()
	// standardLibraryFunctions()
	// loops()
	// booleansAndConditionals()
	// functions()
	// functionsReturningMultipleValues()
	// packageScope()
	// maps()
	passByValue()
}

func learnVariables() {
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
}

func learnBitsAndMemory() {
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

func printingAndFormatting() {
	myName := "Mahrukh"
	age1 := 24
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
	fmt.Printf("my age is %q and my name is name is %q\n", age1, myName)
	fmt.Printf("age is of type %T\n", age1)
	fmt.Printf("you scored %f points! \n", 255.5)
	fmt.Printf("you scored %0.1f points! \n", 255.5)

	//Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v\n", age1, myName)
	fmt.Println("saved string is:", str)
}

func arraysAndSlies() {
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
}

func standardLibraryFunctions() {
	//standard library function
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

func loops() {
	//loops
	x := 0
	namesArray := []string{"babar", "mahrukh", "malaika", "ali"}
	//similar to while loop
	for x < 5 {
		fmt.Println("value of x is ", x+1)
		x++
	}

	//traditional for loop
	for i := 0; i < 5; i++ {
		fmt.Println("value of i is ", i)
	}

	for i := 0; i < len(namesArray); i++ {
		fmt.Println(namesArray[i])
	}

	for index, value := range namesArray {
		fmt.Printf("value at index %v is %v\n", index, value)
	}

	for _, value := range namesArray { //here _ defines the index (which is neither declares not used)
		fmt.Printf("value is %v\n", value)
		value = "new string" ///kind of local copy but it doesn't effect the original slice
	}
	fmt.Println(namesArray)
}

func booleansAndConditionals() {
	//booleans and conditionals
	age1 := 45
	names := [4]string{"Mahrukh", "Ali", "Malaika", "Kalsoom"}

	fmt.Println(age1 < 50)
	fmt.Println(age1 >= 50)
	fmt.Println(age1 == 45)
	fmt.Println(age1 != 50)

	if age1 < 30 {
		fmt.Println("age is less that 30")
	} else if age1 < 40 {
		fmt.Println("age is less that 40")
	} else {
		fmt.Println("age is not less that 45")
	}

	for index, value := range names {
		if index == 3 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}
		fmt.Printf("the value at pos %v is %v\n", index, value)
	}
}

func functions() {
	//functions
	//some built-in functions like len(), append()
	sayGreetings("Babar")
	sayBye("Babar")

	namesSlice := []string{"Mahrukh", "Ali", "Malaika"}
	cycleNames(namesSlice, sayGreetings)
	cycleNames(namesSlice, sayBye)

	a1 := circleArea(10.5)
	fmt.Printf("circle1 raduis is %0.3f\n", a1)

	a2 := circleArea(4)
	fmt.Printf("circle2 raduis is %0.3f\n", a2)
}

func functionsReturningMultipleValues() {
	firstName, secondName := getInitials("Mahrukh Ameen")
	fmt.Println(firstName, secondName)

	firstName1, secondName1 := getInitials("Mahrukh")
	fmt.Println(firstName1, secondName1)
}

func packageScope() {
	//for package scope, we have to run both files simultaneously like "go run file1.go file2.go"
	//if we did.nt do that, it didnt recognize the variables of the other file package scope
	sayHello("Mahrukh")
	for _, value := range points {
		fmt.Print(value, " ")
	}
	fmt.Println()
	showScroe()
}

func maps() {
	menu := map[string]float64{ //map [key]value pair
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])
	//looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	//int as key type
	phonebook := map[int]string{
		0322214567: "Person X",
		8980980808: "Person Y",
		5766775979: "Person Z",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[1])
	phonebook[8980980808] = "Person A"
	for k, v := range phonebook {
		fmt.Println(k, "-", v)
	}
}

func passByValue() {
	//makes copy of value hen passed in functions
	//variables splits in to 2 groups
	//Group A -> string, int, float, boolean, array, structs
	//Group B -> slices, maps, functions

	//Group A types -> string, int, float, boolean, array, structs
	name := "Mahrukh Ameen"
	updateName(name) //variable passed to a function, go creates copy of the variable
	//so this is a copy of the variable not the actual variable.
	//inside the function we are just updating the cpy of the variable not the actual one.
	fmt.Println(name) // prints Mahrukh Ameen

	//if we want to update the original value by passing it in a funtion,
	//we should have to return that value in the function in group A.
	name = updateNameFun(name)
	fmt.Println(name)

	//Group B Types -> slices, maps, functions
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	updateMenu(menu) //it updates the original variable menu
	//if we update the value by passing value to the funtion of group B,
	//it updates the original value
	fmt.Println(menu)
}

func sayGreetings(name string) {
	fmt.Println("hello", name)
}

func sayBye(name string) {
	fmt.Println("Goodbye", name)
}

func cycleNames(names []string, fun func(string)) { //we can call functin inside a function as a parameter
	for _, value := range names {
		fun(value)
	}
}

func circleArea(radius float64) float64 { //2nd float64 is the return type of the function
	return math.Pi * radius * radius
}

func getInitials(name string) (string, string) {
	s := strings.ToUpper(name)
	names := strings.Split(s, " ")
	var initials []string
	for _, value := range names {
		initials = append(initials, value[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func updateName(name string) {
	name = "Mahrukh Babar"
}

func updateNameFun(name string) string {
	name = "Mahrukh Babar"
	return name
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}
