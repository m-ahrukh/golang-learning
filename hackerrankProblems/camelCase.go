package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func camelCaseSolution() {
	fmt.Println("-----CamelCase Problem-----")

	fmt.Print("Enter your string: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("error:", err)
	}

	count := 1
	for _, char := range input {
		if unicode.IsUpper(char) {
			count++
		}
	}

	fmt.Println("Number of words are:", count)
}
