package main

import "fmt"

func main() {
	s1 := "abcabcbb"
	s2 := "bbbbb"
	s3 := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s1))
	fmt.Println(lengthOfLongestSubstring(s2))
	fmt.Println(lengthOfLongestSubstring(s3))
}

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	var freq [256]int
	res := 0
	l, r := 0, -1
	for l < len(s) {
		if r+1 < len(s) && freq[s[r+1]] == 0 {
			r++
			freq[s[r]]++
		} else {
			freq[s[l]]--
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
