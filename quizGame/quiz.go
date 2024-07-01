package main

import (
	"encoding/csv"
	"fmt"
	"os"
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

	questions, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error in reading questions")
	}

	for _, question := range questions {
		fmt.Println(question[0][0], "+", question[0][0], "=", question[1])
	}
}
