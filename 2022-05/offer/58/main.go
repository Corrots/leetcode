package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "the sky is blue"
	fmt.Println(reverseWords(s))
	fmt.Println(reverseWords("a good   example"))
}

// https://leetcode.cn/problems/fan-zhuan-dan-ci-shun-xu-lcof/
// 剑指 Offer 58 - I. 翻转单词顺序
func reverseWords(s string) string {
	slice := strings.Split(strings.Trim(s, " "), " ")
	var res []string
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] != "" {
			res = append(res, slice[i])
		}
	}
	return strings.Join(res, " ")
}
