package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Question struct {
	question string
	answer   string
}

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal("Error while reading the file", err)
		return
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	problems, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Error reading problems")
		return
	}

	var questions []Question
	for _, question := range problems {
		q := Question{}
		q.question = question[0]
		q.answer = question[1]
		questions = append(questions, q)
	}

	score, err := askQuestions(questions, 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nTest Complete!")
	fmt.Printf("\nScore: %d/%d", score, len(problems))

}

func askQuestions(questions []Question, limit int) (int, error) {
	score := 0
	timer := time.NewTimer(time.Duration(limit) * time.Second)
	done := make(chan string)

	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			text, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			text = strings.Replace(text, "\n", "", -1)
			done <- text
		}
	}()

	for _, question := range questions {
		answer, err := askQuestion(question, timer.C, done)
		if err != nil && answer == -1 {
			return score, nil
		}
		score += answer
	}

	return score, nil
}

func askQuestion(problem Question, timer <-chan time.Time, done <-chan string) (int, error) {
	fmt.Printf("%s = ", problem.question)

	for {
		select {
		case <-timer:
			return -1, fmt.Errorf("Time limit exceeded")
		case answer := <-done:
			score := 0
			if strings.Compare(problem.answer, answer) == 0 {
				score = 1
			} else {
				return 0, fmt.Errorf("Wrong Answer")
			}

			return score, nil
		}
	}
}
