package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {
	// Open CSV file
	file, err := os.Open("quiz.csv")

	// Check if there is any error opening CSV file
	if err != nil {
		fmt.Println("Error while opening CSV file: ", err)
	}

	// Read contents of CSV file
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

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

	for i := 0; i < len(records); i++ {
		select {
		case <-quizChannel:
			fmt.Printf("Total score is %d out of %d", count, len(records))
			return
		default:
			fmt.Print(records[i][0] + " = ")
			fmt.Scan(&userAnswer)
			if userAnswer == records[i][1] {
				count++
			}
		}
	}

	fmt.Printf("Total score is %d out of %d", count, len(records))
}
