package main

import (
	"fmt"
)

func main() {
	x_arr := []bool{true, false, true, false}

	fmt.Println(solution(x_arr[0], x_arr[1], x_arr[2], x_arr[3]))
}

func solution(x1 bool, x2 bool, x3 bool, x4 bool) bool {
	if (x1 || x2) && (x3 || x4) {
		return true
	} else {
		return false
	}
}
