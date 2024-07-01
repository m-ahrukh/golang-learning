package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Quiz Game")

	filename := flag.String("file", "problems.csv", "CSV file containing quiz questions")
	flag.Parse()

	//open specific file in read-only mode
	// file, err := os.Open("problems.csv")
	file, err := os.Open(*filename + ".csv")

	if err != nil {
		fmt.Println("Error:", err)
	}

	//close file
	defer file.Close()

	reader := csv.NewReader(file)
	inputReader := bufio.NewReader(os.Stdin)

	questions, err := reader.ReadAll()

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
	fmt.Printf("you answered %d/%d questions\n", score, len(questions))
	fmt.Printf("Total questions: %d\n", len(questions))
	fmt.Printf("Correct answers: %d\n", score)
	fmt.Printf("Incorrect answers: %d\n", len(questions)-score)
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}
