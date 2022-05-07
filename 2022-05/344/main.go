package main

import "fmt"

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Printf("%s", s)
}

// https://leetcode-cn.com/problems/reverse-string/
// 344. 反转字符串
func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}
	l, r := 0, len(s)-1
	for l <= r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
