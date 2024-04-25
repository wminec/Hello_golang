package main

import (
	"fmt"
	"strings"
)

func main() {
	order := []string{"cafelatte", "americanoice", "hotcafelatte", "anything"}
	fmt.Println(solution(order))
}

func solution(order []string) int {
	sum := 0

	for _, item := range order {
		if strings.Contains(item, "latte") {
			sum += 5000
		} else if strings.Contains(item, "americano") || item == "anything" {
			sum += 4500
		}
	}

	return sum
}
