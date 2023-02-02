package main

import (
	"fmt"
)

func main() {
	// 提取 "the quick brown fox jumps over the lazy dog" 中每个字母的首次出现可以得到替换表
	fmt.Println(decodeMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv"))
	fmt.Println(decodeMessage("eljuxhpwnyrdgtqkviszcfmabo", "zwx hnfx lqantp mnoeius ycgk vcnjrdb"))
}

// https://leetcode.cn/problems/decode-the-message/
func decodeMessage(key string, message string) string {
	cur := byte('a')
	rules := map[rune]byte{}

	for _, c := range key {
		if c != ' ' && rules[c] == 0 {
			rules[c] = cur
			cur++
		}
	}

	m := []byte(message)
	for i, c := range message {
		if c != ' ' {
			m[i] = rules[c]
		}
	}

	return string(m)
}
