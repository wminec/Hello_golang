package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(solution(2, 91))
}

func solution(a int, b int) int {
	cn, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	mn := 2 * a * b
	if cn >= mn {
		return cn
	} else {
		return mn
	}
}
