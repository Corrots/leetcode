package main

import "fmt"

func main() {
	s := "abcdefg"
	fmt.Println(reverseLeftWords(s, 2))
}

// https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
// 剑指 Offer 58 - II. 左旋转字符串
func reverseLeftWords(s string, n int) string {
	bytes := []byte(s)
	var reverse func([]byte, int, int)
	reverse = func(b []byte, p int, q int) {
		for p < q {
			b[p], b[q] = b[q], b[p]
			p++
			q--
		}
	}

	// 翻转三部曲
	// 1.翻转整体
	// 翻转结果：[g f e d c b a]
	reverse(bytes, 0, len(s)-1)
	// 翻转左半边
	// 翻转结果：[c d e f g b a]
	reverse(bytes, 0, len(s)-n-1)
	// 翻转右半边
	// 翻转结果：[c d e f g a b]
	reverse(bytes, len(s)-n, len(s)-1)
	return string(bytes)
}
