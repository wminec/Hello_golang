package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(solution("string", 3))
}

func solution(my_string string, k int) string {
	var res bytes.Buffer
	for i := 0; i < k; i++ {
		res.WriteString(my_string)
	}
	return res.String()
}
