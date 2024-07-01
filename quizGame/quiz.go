package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Quiz Game")

	//open specific file in read-only mode
	file, err := os.Open("problems.csv")

	if err != nil {
		fmt.Println("Error:", err)
	}

	//close file
	defer file.Close()

	reader := csv.NewReader(file)
	inputReader := bufio.NewReader(os.Stdin)

	questions, err := reader.ReadAll()
	// opt, _ := getInput(inputReader)

	if err != nil {
		fmt.Println("Error in reading questions")
	}

	var score = 0
	for _, question := range questions {
		opt, _ := getInput(question[0]+" = ", inputReader)
		if opt == question[1] {
			score++
		}
	}
	fmt.Printf("you answered %d/%d questions", score, len(questions))
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}
