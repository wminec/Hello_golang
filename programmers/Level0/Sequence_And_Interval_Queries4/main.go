package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 1, 2, 4, 3}
	queries := [][]int{{0, 4, 1}, {0, 3, 2}, {0, 3, 3}}
	fmt.Println(solution(arr, queries))
}

func solution(arr []int, queries [][]int) []int {
	for _, query := range queries {
		start, end := query[0], query[1]
		k := query[2]
		for i := start; i <= end; i++ {
			if i%k == 0 {
				arr[i] += 1
			}
		}
	}
	return arr
}
