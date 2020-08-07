package main

import "fmt"

func main() {
	s := "cbaebabacd"
	fmt.Println(findAnagrams(s, "abc"))
	s1 := "abab"
	fmt.Println(findAnagrams(s1, "ab"))
}

// https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/
func findAnagrams(s string, p string) []int {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := range p {
		need[p[i]]++
	}
	valid := 0
	var res []int
	var left, right int
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		//
		for right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res
}
