package main

import (
	"regexp"
	"strings"
)

func main() {
	str1 := "aaaaa"
	str2 := "bbbbb"
	println(solution(str1, str2))
}

func solution(str1 string, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)
	var res strings.Builder

	b1, _ := regexp.MatchString("[a-z]", str1)
	b2, _ := regexp.MatchString("[a-z]", str2)

	if b1 && b2 && len1-len2 == 0 && len1 >= 1 && len1 <= 10 {
		for i := 0; i < len1; i++ {
			res.WriteByte(str1[i])
			res.WriteByte(str2[i])
		}
	}

	return res.String()
}
