package main

import "fmt"

func main() {
	fmt.Println(numWays(7))
	fmt.Println(numWays(3))
}

// https://leetcode.cn/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
// 剑指 Offer 10- II. 青蛙跳台阶问题
// n=0和n=1时，都有1种跳法
func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	memo := make([]int, n+1)
	memo[0] = 1
	memo[1] = 1
	for i := 2; i <= n; i++ {
		// 对结果取模
		memo[i] = (memo[i-1] + memo[i-2]) % (1e9 + 7)
	}
	return memo[n]
}
