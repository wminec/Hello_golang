package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 string
	var a int
	fmt.Scan(&s1, &a)
	fmt.Println(strings.Repeat(s1, a))
}
