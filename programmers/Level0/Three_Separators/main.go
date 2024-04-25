package main

import (
	"fmt"
)

func main() {
	myStr := "baconlettucetomato"
	//myStr := "abcd"
	//myStr := "fcababf"
	//myStr := "cabafb"
	//myStr := "cabab"

	fmt.Println(solution(myStr))
}

func solution(myStr string) []string {
	separator := []byte{'a', 'b', 'c'}
	res := []string{}
	tmp_str := ""
	myStr_len := len(myStr)

	for i := 0; i < myStr_len; i++ {
		if myStr[i] != separator[0] && myStr[i] != separator[1] && myStr[i] != separator[2] {
			tmp_str += string(myStr[i])
		} else {
			if tmp_str != "" {
				res = append(res, tmp_str)
				tmp_str = ""
			} else {
				continue
			}
		}
	}

	if tmp_str != "" {
		res = append(res, tmp_str)
	} else if len(res) == 0 {
		return []string{"EMPTY"}
	}

	return res
}
