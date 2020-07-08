package main

import (
	"fmt"
	"strings"
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
	return sorting(s) == sorting(t)
}

func sorting(s string) string {
	s1 := strings.Split(s, "")
	for i := 1; i < len(s1); i++ {
		e := s1[i]
		j := 0
		for j = i; j > 0 && s1[j-1] > e; j-- {
			s1[j] = s1[j-1]
		}
		s1[j] = e
	}
	return strings.Join(s1, "")
}
