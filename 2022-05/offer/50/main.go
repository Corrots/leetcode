package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("abaccdeff"))
}

// https://leetcode.cn/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/
// 剑指 Offer 50. 第一个只出现一次的字符
func firstUniqChar(s string) byte {
	var res byte
	if len(s) == 0 {
		return 0
	}
	for i := 0; i < len(s); i++ {
		res ^= s[i]
	}
	fmt.Println(string(res))
	return res
}
