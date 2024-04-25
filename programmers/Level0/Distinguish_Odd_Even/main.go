package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	if a%2 == 0 {
		fmt.Printf("%d is even", a)
	} else {
		fmt.Printf("%d is odd", a)
	}
}
