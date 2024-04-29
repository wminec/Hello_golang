package main

import (
	"fmt"
)

func main() {
	num_list := []int{12, 4, 15, 1, 14}
	fmt.Println(solution(num_list))
}

func divide(num int) int {
	count := 0
	for num != 1 {
		if num%2 == 0 {
			num /= 2
			count++
		} else {
			num = (num - 1) / 2
			count++
		}
	}
	return count
}

func solution(num_list []int) int {
	res := 0
	for _, num := range num_list {
		res += divide(num)
	}

	return res
}
