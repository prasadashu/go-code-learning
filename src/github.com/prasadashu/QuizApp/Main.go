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
	// Save user's anwers
	var userAnswer string

	// Channel for communication between Goroutine and Main function
	quizChannel := make(chan int)

	go func() {
		time.Sleep(time.Second * 30)
		close(quizChannel)
	}()

	for i := 0; i < len(problemList); i++ {
		select {
		case <-quizChannel:
			fmt.Printf("Total score is %d out of %d", count, len(problemList))
			return
		default:
			fmt.Print(problemList[i].question + " = ")
			fmt.Scan(&userAnswer)
			if userAnswer == problemList[i].answer {
				count++
			}
		}
	}

	fmt.Printf("Total score is %d out of %d", count, len(records))
}
