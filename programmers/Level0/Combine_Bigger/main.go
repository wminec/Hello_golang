package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(solution(9, 91))
}

func solution(a int, b int) int {
	n1, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	n2, _ := strconv.Atoi(strconv.Itoa(b) + strconv.Itoa(a))

	if n1-n2 < 0 {
		return n2
	} else {
		return n1
	}
}
