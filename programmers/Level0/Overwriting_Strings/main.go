package main

import (
	"fmt"
)

func main() {
	my_string := "He11oWord1"
	overwite_string := "lloWorl"
	s := 2
	fmt.Println(solution(my_string, overwite_string, s))
}

func solution(my_string string, overwite_string string, s int) string {
	res := my_string[:s] + overwite_string + my_string[(s+len(overwite_string)):]
	return res
}
