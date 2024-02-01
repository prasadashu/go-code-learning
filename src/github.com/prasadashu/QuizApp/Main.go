package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problems struct {
	question string
	answer   string
}

func parseProblems(records [][]string) []problems {
	problemList := make([]problems, len(records))

	for index, line := range records {
		problemList[index] = problems{
			question: line[0],
			answer:   line[1],
		}
	}

	return problemList
}

func main() {
	// Get config variables
	quizFile := flag.String("f", "quiz.csv", "path to quiz file")
	quizTime := flag.Int("t", 30, "time duration for the quiz in seconds")
	flag.Parse()

	// Open CSV file
	file, err := os.Open(*quizFile)

	// Check if there is any error opening CSV file
	if err != nil {
		fmt.Println("Error while opening CSV file: ", err)
		os.Exit(1)
	}

	// Read contents of CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	// Check if there is any error while parsing CSV file
	if err != nil {
		fmt.Println("Error while parsing CSV file: ", err)
	}

	// Get problem list
	problemList := parseProblems(records)

	// Count of right answers
	count := 0

	// Channel for communication between Goroutine and Main function
	answerChannel := make(chan string)

	timer := time.NewTimer(time.Duration(*quizTime) * (time.Second))

	for i := 0; i < len(problemList); i++ {
		fmt.Print(problemList[i].question + " = ")

		go func() {
			// Save user's anwers
			var userAnswer string
			fmt.Scan(&userAnswer)
			answerChannel <- userAnswer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nTotal score is %d out of %d", count, len(problemList))
			return
		case answer := <-answerChannel:
			if answer == problemList[i].answer {
				count++
			}
		}
	}

	fmt.Printf("\nTotal score is %d out of %d", count, len(records))
	return
}
