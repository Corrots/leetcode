package main

import (
	"fmt"
	"math"
)

func main() {
	s1 := []string{"bella", "label", "roller"}
	fmt.Println(commonChars(s1))
	s2 := []string{"cool", "lock", "cook"}
	fmt.Println(commonChars(s2))
}

//https://leetcode-cn.com/problems/find-common-characters/
func commonChars(A []string) []string {
	if len(A) == 0 {
		return nil
	}
	var minFreq [26]int
	for i := 0; i < 26; i++ {
		minFreq[i] = math.MaxInt64
	}
	for i := 0; i < len(A); i++ {
		// 统计每个字符串中字符出现的次数
		var freq [26]int
		for _, c := range A[i] {
			freq[c-'a']++
		}
		for i, f := range freq {
			if f < minFreq[i] {
				minFreq[i] = f
			}
		}
	}
	var res []string
	for i := byte(0); i < 26; i++ {
		for j := 0; j < minFreq[i]; j++ {
			res = append(res, string(i+'a'))
		}
	}
	return res
}
