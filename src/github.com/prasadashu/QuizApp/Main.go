package main

import (
	"fmt"
	"time"
)

func main() {
	questions := []string{"1+2", "2+3", "3+4", "4+5"}
	answers := []string{"3", "5", "7", "9"}
	count := 0

	var userAnswer string

	quizChannel := make(chan int)

	go func() {
		time.Sleep(time.Second * 30)
		close(quizChannel)
	}()

	for i := 0; i < len(questions); i++ {
		select {
		case <-quizChannel:
			fmt.Printf("Total score is %d out of %d", count, len(questions))
			return
		default:
			fmt.Print(questions[i] + " = ")
			fmt.Scan(&userAnswer)
			if userAnswer == answers[i] {
				count++
			}
		}
	}

	fmt.Printf("Total score is %d out of %d", count, len(questions))
}
