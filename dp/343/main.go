package main

import (
	"fmt"
	"math"
)

func main() {
	n := 10
	fmt.Println(integerBreak(n))
}

//https://leetcode-cn.com/problems/integer-break/
func integerBreak(n int) int {
	if n < 2 {
		return 0
	}
	memo := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		memo[i] = math.MinInt64
	}
	var dfs func(int) int
	dfs = func(n int) int {
		if n == 1 {
			return 1
		}
		if memo[n] != math.MinInt64 {
			return memo[n]
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
