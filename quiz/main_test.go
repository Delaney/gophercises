package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAnswerQuestion(t *testing.T) {
	timer := time.NewTimer(time.Duration(3) * time.Second).C
	done := make(chan string)

	var question Question
	question.question = "2+3"
	question.answer = "5"

	allDone := make(chan bool)
	var answer int
	var err error

	go func() {
		answer, err = askQuestion(question, timer, done)
		allDone <- true
	}()
	done <- "5"

	<-allDone
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, answer, 1)
}
