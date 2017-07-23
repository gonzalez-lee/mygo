package main

import (
	"fmt"
	"github.com/gonzalez_lee/mygo/stringutil"
	"math/rand"
)

func main() {
	a, b := 1, "a"
	if a == 1 {
		fmt.Println(a, b)
	}
	fmt.Println("Hello world.")
	num1, num2 := 1, 2
	fmt.Printf("Before swaping, num1=%d, num2=%d\n", num1, num2)
	num1, num2 = swap(num1, num2)
	fmt.Printf("After swaping, num1=%d, num2=%d\n", num1, num2)
	fmt.Printf("My favorite numner is %d.\n", rand.Intn(10))
	fmt.Println(concatenate("a", "b"))
	fmt.Println(stringutil.Reverse("Hello world."))
}

func swap(a, b int) (int, int) {
	tmp := a
	a = b
	b = tmp
	return a, b
}

func concatenate(a, b string) (concat string) {
	concat = a + b
	return
}
