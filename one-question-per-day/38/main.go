package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(countAndSay(2))
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(5))
}

// 	外观数列
//	1
// 	11
//	21
//	1211
//	111221
//	312211
//	13112221
//	1113213211
//https://leetcode-cn.com/problems/count-and-say/
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}
	str := countAndSay(n - 1)
	length := len(str)
	var res strings.Builder
	count := 1
	for i := 1; i < length; i++ {
		if str[i] != str[i-1] {
			res.WriteString(strconv.Itoa(count) + string(str[i-1]))
			count = 1
		} else {
			count++
		}
		if i+1 == length {
			res.WriteString(strconv.Itoa(count) + string(str[i]))
			break
		}
	}
	return res.String()
}
