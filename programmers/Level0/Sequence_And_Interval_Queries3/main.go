package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4}
	queries := [][]int{{0, 3}, {1, 2}, {1, 4}}
	fmt.Println(solution(arr, queries))
}

func solution(arr []int, queries [][]int) []int {
	for _, query := range queries {
		tmp := arr[query[0]]
		arr[query[0]] = arr[query[1]]
		arr[query[1]] = tmp
	}
	return arr
}
