package main

import (
	"github.com/gonzalez_lee/mygo/stringutil"
	"fmt"
	"github.com/gonzalez_lee/mygo/chatbot"
	"github.com/gonzalez_lee/mygo/bean"
	"github.com/gonzalez_lee/mygo/control"
)

func testInter(ai chatbot.AI) {
	ai.Hello("world")
	fmt.Println(ai.Reply("Hello"))
}

func main() {
	var person chatbot.AI = bean.Person{Name:"Andy", Gender:"Male", Age:30}
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

	字符串 := "hello"
	fmt.Println(字符串)
	fmt.Println(字符串)
}
