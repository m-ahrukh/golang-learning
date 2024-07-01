package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Quiz Game")

	filename := flag.String("file", "problems", "CSV file containing quiz questions")
	timeLimit := flag.Int("limit", 30, "Time Limit for each question")
	flag.Parse()

	*filename += ".csv"
	//open specific file in read-only mode
	// file, err := os.Open("problems.csv")
	file, err := os.Open(*filename)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//close file
	defer file.Close()

	reader := csv.NewReader(file)
	inputReader := bufio.NewReader(os.Stdin)

	questions, err := reader.ReadAll()
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	if err != nil {
		fmt.Println("Error in reading questions")
	}

	var score = 0
loop:
	for _, question := range questions {

		answerCh := make(chan string)

		go func() {
			opt, _ := getInput(question[0]+" = ", inputReader)
			answerCh <- opt
		}()

		select {
		case <-timer.C:
			// fmt.Println("\nYour time been been finished")
			// fmt.Println("----------Score Card----------")
			// fmt.Printf("|you answered %d/%d questions |\n", score, len(questions))
			// fmt.Printf("|Total questions: %d         |\n", len(questions))
			// fmt.Printf("|Correct answers: %d          |\n", score)
			// fmt.Printf("|Incorrect answers: %d       |\n", len(questions)-score)
			// fmt.Println("------------------------------")
			// return
			fmt.Println()
			break loop //goto statement
		case opt := <-answerCh:
			// opt, _ := getInput(question[0]+" = ", inputReader)
			if opt == question[1] {
				score++
			}
		}

	}
	fmt.Println("\nYour time been been finished")
	fmt.Println("----------Score Card----------")
	fmt.Printf("|you answered %d/%d questions |\n", score, len(questions))
	fmt.Printf("|Total questions: %d         |\n", len(questions))
	fmt.Printf("|Correct answers: %d          |\n", score)
	fmt.Printf("|Incorrect answers: %d       |\n", len(questions)-score)
	fmt.Println("------------------------------")
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}
