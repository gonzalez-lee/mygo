package stringutil

import (
	"strings"
)

func SubStr(str string, startIndex int) string {
	runes := []rune(str)
	// for _, r := range runes {
	// 	fmt.Printf("%c\n", r)
	// }
	return string(runes[startIndex:])
}

func Contains(str string, substr string) bool {
	return strings.Contains(str, substr)
}

