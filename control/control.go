package control

import (
	"fmt"
)

func DemoIf(condition int) {
	if condition == 1 {
		fmt.Println("condition is 1.")
	} else if condition == 2 {
		fmt.Println("branch 2")
	} else {
		fmt.Println("branch else.")
	}
}

func DemoSwitch(lang string) {
	switch lang {
	default:
		fmt.Println("No favourite language.")
	case "java":
		fmt.Println("most favourite language is java.")
		fallthrough
	case "go":
		fmt.Println("most favourite language is go.")		
	}
}

func DemoTypeSwitch(lang interface{}) {
	switch lang.(type) {
	default:
		fmt.Println("No favourite language.")
	case string:
		fmt.Printf("the argument is %s.\n", lang)
	case int, int16, int32, int64:
		fmt.Printf("the argument is %d.\n", lang)		
	}
}

func DemoFor() {
	str := "hello"
	str1 := `\xac`
	fmt.Println(str1)
	for i, _ := range str {
		fmt.Printf("character at index %d is %b\n", i, str[i])
	}
}