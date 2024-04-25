package main

import "fmt"

func main() {
	//arr := []int{1, 2, 1, 2, 1, 10, 2, 1}
	//arr := []int{1, 1, 1}
	//arr := []int{1, 2, 1}
	arr := []int{1, 2, 1, 4, 5, 2, 9}

	fmt.Println(solution(arr))
}

func solution(arr []int) []int {
	// start and end index of the subarray
	s, e := -1, -1

	for index, i := range arr {
		if i == 2 {
			s = index
			break
		}
	}

	if s == -1 {
		return []int{-1}
	}

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 2 {
			e = i
			break
		}
	}

	return arr[s : e+1]
}
