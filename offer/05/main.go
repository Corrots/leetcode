package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "We are happy."
	fmt.Println(replaceSpace(str))
}

// https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
func replaceSpace(s string) string {
	var res strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			res.WriteString("%20")
		} else {
			res.WriteString(string(s[i]))
		}
	}
	return res.String()
}
