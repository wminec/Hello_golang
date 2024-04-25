package main

import (
	"fmt"
)

func main() {

	n := 1
	slicer := []int{1, 5, 2}
	num_list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(solution(n, slicer, num_list))
	fmt.Println(other_solution(n, slicer, num_list))
}

func solution(n int, slicer []int, num_list []int) []int {
	res := []int{}
	num_len := len(num_list)
	switch n {
	case 1:
		// From 0 to b
		for i := 0; i <= slicer[1]; i++ {
			res = append(res, num_list[i])
		}
	case 2:
		// From a to last
		for i := slicer[0]; i < num_len; i++ {
			res = append(res, num_list[i])
		}
	case 3:
		// From a to b
		for i := slicer[0]; i <= slicer[1]; i++ {
			res = append(res, num_list[i])
		}
	case 4:
		// From a to b with interval c
		for i := slicer[0]; i <= slicer[1]; i += slicer[2] {
			res = append(res, num_list[i])
		}
	default:
		// Code to be executed when n doesn't match any case
		fmt.Println("No match found!")
	}
	return res
}

func other_solution(n int, slicer []int, num_list []int) []int {
	res := []int{}
	switch n {
	case 1:
		// From 0 to b
		res = append(res, num_list[:slicer[1]+1]...)
	case 2:
		// From a to last
		res = append(res, num_list[slicer[0]:]...)
	case 3:
		// From a to b
		res = append(res, num_list[slicer[0]:slicer[1]+1]...)
	case 4:
		// From a to b with interval c
		for i := slicer[0]; i <= slicer[1]; i += slicer[2] {
			res = append(res, num_list[i])
		}
	}

	return res
}
