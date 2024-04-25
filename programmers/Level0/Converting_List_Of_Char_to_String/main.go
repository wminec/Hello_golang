package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := []string{"a", "b", "c"}
	fmt.Println(solution(arr))
}

func solution(arr []string) string {
	return strings.Join(arr[:], "")
}
