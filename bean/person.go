package bean

import (
	"fmt"
)

type Person struct {
	Name string
	Gender string 
	Age int
}

func (person Person) Hello(word string) {
	fmt.Println("Hello : I'm " + person.Name + ", word : " + word)
}

func (person Person) Reply(greeting string) string {
	return "Reply : I'm " + person.Name
}