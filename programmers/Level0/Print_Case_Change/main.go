package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World!"
	convertedStr := convertCase(str)
	fmt.Println(convertedStr)
}

func convertCase(str string) string {
	var convertedStr strings.Builder

	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			convertedStr.WriteRune(char - 32)
		} else if char >= 'A' && char <= 'Z' {
			convertedStr.WriteRune(char + 32)
		} else {
			convertedStr.WriteRune(char)
		}
	}

	return convertedStr.String()
}
