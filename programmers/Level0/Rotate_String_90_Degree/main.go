package main

import (
	"fmt"
)

func main() {
	var s1 string
	fmt.Scan(&s1)
	for _, char := range s1 {
		fmt.Printf("%c\n", char)
	}
}
