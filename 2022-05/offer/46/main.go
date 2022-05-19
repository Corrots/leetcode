package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(translateNum(12258))
}

// https://leetcode.cn/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
// 剑指 Offer 46. 把数字翻译成字符串
//
func translateNum(num int) int {
	s := strconv.Itoa(num)
	//fmt.Println(len(s))
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		num := (s[i-2]-'0')*10 + (s[i-1] - '0')
		if num >= 10 && num <= 25 {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n]
}
