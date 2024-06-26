package main

import (
	"bufio"
	"fmt" // fmt is used for IO operations
	"math"
	"os"
	"sort"
	"strconv"
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
	// passByValue()
	// pointers()
	// structs()
	// recieverFunctions()
	// recieverFunctionsWithPointers()
	// userInputs()
	// switchStatement()
	// parsingFloats()
	// savingFlies()
	interfaces()

	//capacity (kitny size ki array bni hui) and size (elements in array) in slice
	// no exceptions in Go. errors hongy

}

func print(a interface{}) {
	fmt.Println(a)
}

//capital kreingy to export ho jayga package se bahir bhi visible ho jayga
//small kreingy to package scope mein rhyga

//struct will implement all the available interfaces implicitly

type Maker interface {
	make(a int)
	makeWith(x int)
}

type coffeeMaker struct {
}

func (cm coffeeMaker) make(a int) {
	fmt.Printf("I'm making for %d\n", a)
}

func (cm coffeeMaker) makeWith(a int) {
	fmt.Printf("I'm making with %d\n", a)
}

func process(m Maker, a int) {
	m.make(a)
	m.makeWith(10)
}

//empty interface -> any interface

// shape interface
type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

// square methods
func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return s.length * 4
}

// circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("circumference of %T is: %0.2f\n", s, s.circumf())
}

func textSlices() {
	a := make([]int, 0, 20) //capacity hm define krty hain to avoid memory wastage

	fmt.Println(len(a), cap(a))
	a = append(a, 8)
	fmt.Println(len(a), cap(a))
	a = append(a, 4)
	fmt.Println(len(a), cap(a))
	a = append(a, 6)
	fmt.Println(len(a), cap(a))
	a = append(a, 5)
	fmt.Println(len(a), cap(a))
	a = append(a, 1)
	a = append(a, 0)
	a = append(a, 3)
	a = append(a, 12)
	a = append(a, 18)
	fmt.Println(len(a), cap(a))
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

func pointers() {
	name := "Mahrukh Ameen"
	fmt.Println(name)
	fmt.Println("memory address of name is:", &name) //&name returns the memory address
	memory := &name
	updateNameUsingPointer(memory) //stores the memory address of name variabl
	fmt.Println("memory address:", memory)
	fmt.Println("value at memory address:", *memory) //returns the value that will be present at that memory address
	//it will update the original value of name variable
	fmt.Println(name)
}

func structs() {
	//structs are the custom data types having multiple datatype variables
	//it is basically a bluprint to describe type of data
	myBill := newBill("Mahrukh's Bill")
	fmt.Println("Name of Bill:", myBill.name)
	fmt.Println("Items purchased")
	for k, value := range myBill.items {
		fmt.Println(k, "-", value)
	}
	fmt.Println("Tip:", myBill.tip)
}

func recieverFunctions() {
	myBill := newBill("Mahrukh's Bill")
	fmt.Println("Name of Bill:", myBill.name)
	fmt.Println(myBill.format())
}

func recieverFunctionsWithPointers() {
	myBill := newBill("Mahrukh's Bill")
	fmt.Println("Name of Bill:", myBill.name)
	myBill.updateTip(15)
	myBill.addItem("Macaronni", 100.00)
	fmt.Println(myBill.format())
}

func userInputs() {
	myBill := createBill()
	fmt.Println(myBill)
	promptOptions(myBill)
}

func switchStatement() {
	myBill := createBill()

	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose Option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item Price: ", reader)
		fmt.Println(name, price)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		fmt.Println(tip)
	case "s":
		fmt.Println("You chose s")
	default:
		fmt.Println("That was not a valid option")
		promptOptions(myBill)
	}
}

func parsingFloats() {
	myBill := createBill()

	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose Option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item Price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions(myBill)
		}
		myBill.addItem(name, p)
		myBill.format()
		fmt.Println("Item added - ", name, price)
		promptOptions(myBill)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions(myBill)
		}
		myBill.updateTip(t)
		myBill.format()
		fmt.Println("Tip added: ", tip)
		promptOptions(myBill)
	case "s":
		fmt.Println("You chose to save the bill")
	default:
		fmt.Println("That was not a valid option")
		promptOptions(myBill)
	}
}

func savingFlies() {
	myBill := createBill()

	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose Option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item Price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Price must be a number")
			opt = promptOptions(myBill)
			fmt.Println("-->", opt)
		}
		myBill.addItem(name, p)
		// myBill.format()
		fmt.Println("Item added - ", name, price)
		opt = promptOptions(myBill)
		fmt.Println("-->", opt)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Price must be a number")
			opt = promptOptions(myBill)
			fmt.Println("-->", opt)
		}
		myBill.updateTip(t)
		myBill.format()
		fmt.Println("Tip added: ", tip)
		opt = promptOptions(myBill)
		fmt.Println("-->", opt)
	case "s":
		myBill.save()
		fmt.Println("You csaved thi file - ", myBill.name)
	default:
		fmt.Println("That was not a valid option")
		opt = promptOptions(myBill)
		fmt.Println("-->", opt)
	}
}

func sayGreetings(name string) {
	fmt.Println("hello", name)
}

func interfaces() {
	cm := coffeeMaker{}
	process(cm, 5)

	shape := []shape{
		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}

	for _, v := range shape {
		printShapeInfo(v)
		fmt.Println("----")
	}
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

func updateNameUsingPointer(name *string) {
	*name = "Mahrukh"
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b bill) string {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose Option (a - add item, s - save bill, t - add tip): ", reader)
	// fmt.Println(opt)
	return opt
}
