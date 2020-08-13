package main

import (
	"fmt"
)

/**
#242
有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
*/
func main() {
	s1, t1 := "anagram", "nagaram"
	fmt.Println(isAnagram(s1, t1))
	s2, t2 := "rat", "car"
	fmt.Println(isAnagram(s2, t2))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var count [26]int
	for _, v := range s {
		count[v-'a']++
	}
	for _, v := range t {
		count[v-'a']--
	}
	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}
