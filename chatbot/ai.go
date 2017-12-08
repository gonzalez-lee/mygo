package chatbot

import (
	"fmt"
)

type AI interface {
	Hello(word string)
	Reply(msg string) string
}

type MyAI string

func (ai MyAI) Hello(word string) {
	fmt.Println("Hello " + word)
}

func (ai MyAI) Reply(msg string) (string) {
	return msg + ", I'm my ai."
}