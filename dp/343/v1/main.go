package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(integerBreak(2))
	n := 10
	fmt.Println(integerBreak(n))
}

//https://leetcode-cn.com/problems/integer-break/
// 记忆化搜索解法
func integerBreak(n int) int {
	if n == 1 {
		return 1
	}
	// 初始化记忆化数组memo
	memo := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		memo[i] = math.MinInt64
	}
	var dfs func(int) int
	// 将n分割成至少2部分，求分割后数字的最大乘积
	dfs = func(n int) int {
		if n == 1 {
			return 1
		}
		if memo[n] != math.MinInt64 {
			return memo[n]
		}
		res := math.MinInt64
		for i := 1; i < n; i++ {
			// n-i不分割，直接i*(n-i)
			res = max(res, i*(n-i), i*dfs(n-i))
		}
		memo[n] = res
		return res
	}
	return dfs(n)
}

func max(args ...int) int {
	res := math.MinInt64
	for i := 0; i < len(args); i++ {
		if args[i] > res {
			res = args[i]
		}
	}
	return res
}
