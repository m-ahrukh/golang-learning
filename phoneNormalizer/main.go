package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Phone Normalizer Problem")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your phone number: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
	}
	phoneNumber := strings.TrimSpace(input)

	fmt.Println("Phone Number:", phoneNumber)
}
