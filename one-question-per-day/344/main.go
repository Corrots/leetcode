package main

import "fmt"

func main() {
	s1 := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s1)
	fmt.Printf("%s\n", s1)
	s2 := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	reverseString(s2)
	fmt.Printf("%s\n", s2)
}

//https://leetcode-cn.com/problems/reverse-string/
func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}
	p, q := 0, len(s)-1
	for p < q {
		// 小优化：只有当s[p] != s[q]时才需执行交换操作
		if s[p] != s[q] {
			s[p], s[q] = s[q], s[p]
		}
		p++
		q--
	}
}
