package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(integerBreak(2))
	fmt.Println(integerBreak(10))
}

//https://leetcode-cn.com/problems/integer-break/
// 基础递归解法
func integerBreak(n int) int {
	var dfs func(int) int
	dfs = func(n int) int {
		if n == 1 {
			return 1
		}
		res := math.MinInt64
		for i := 1; i < n; i++ {
			res = max(res, i*(n-i), i*dfs(n-i))
		}
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
