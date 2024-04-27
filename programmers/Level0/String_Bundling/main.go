package main

import "fmt"

func main() {
	strArr := []string{"a", "bc", "d", "efg", "hi"}
	fmt.Println(solution(strArr))
}

func solution(strArr []string) int {
	res := 0
	tmpArr := make([]int, 30)

	for _, str := range strArr {
		tmpArr[len(str)-1]++
	}

	for _, i := range tmpArr {
		if i > res {
			res = i
		}
	}

	return res
}
