package main

import (
	"fmt"

	"github.com/gonzalez-lee/mygo/bean"
	"github.com/gonzalez-lee/mygo/chatbot"
	"github.com/gonzalez-lee/mygo/control"
	"github.com/gonzalez-lee/mygo/stringutil"
)

func testInter(ai chatbot.AI) {
	ai.Hello("world")
	fmt.Println(ai.Reply("Hello"))
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func sum1() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

// TestGoFunc is used to test go features.
func TestGoFunc() {
	andy := fmt.Sprintf("%s, %d", "Andy", 35)
	fmt.Println(andy)
	sum1()
	var person chatbot.AI = bean.Person{Name: "Andy", Gender: "Male", Age: 30}
	testInter(person)

	myAi := chatbot.MyAI("123")
	testInter(myAi)

	control.DemoIf(1)
	control.DemoSwitch("go")
	control.DemoTypeSwitch("go")
	control.DemoTypeSwitch(1)
	control.DemoFor()

	str := "日本语"
	runes := []rune(str)
	for i, c := range runes {
		fmt.Printf("byte at index %d is %c\n", i, c)
	}

	fmt.Printf("%s", stringutil.SubStr("日本语", 1))

	字符串 := "hello2"
	fmt.Println(字符串)
}
