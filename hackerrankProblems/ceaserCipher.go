package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func ceaserCipherSolution() {
	fmt.Println("------Ceaser Cipher Problem------")

	fmt.Print("Enter your string: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("error:", err)
	}

	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")
	if len(parts) < 2 || len(parts) > 2 {
		fmt.Println("Error: Invlid number of arguments")
		return
	}
	str := parts[0]
	tempKey := parts[1]
	key, err := strconv.ParseInt(tempKey, 10, 0)
	if err != nil {
		key = 3
	}

	var cipheredText []rune
	for _, char := range str {
		if unicode.IsLetter(char) {
			cipheredLeter := rune(key)
			if unicode.IsUpper(char) {
				cipheredText = append(cipheredText, 'A'+(char-'A'+cipheredLeter)%26)
			} else if unicode.IsLower(char) {
				cipheredText = append(cipheredText, 'a'+(char-'a'+cipheredLeter)%26)
			}
		} else {
			cipheredText = append(cipheredText, char)
		}
	}

	fmt.Println("Original text:", str)
	fmt.Println("Key is:", key)
	fmt.Println("Ciphered text:", string(cipheredText))

}
