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
// DP
func integerBreak(n int) int {
	if n == 1 {
		return 0
	}
	memo := make([]int, n+1)
	memo[1] = 1
	for i := 2; i < n+1; i++ {
		for j := 1; j < i; j++ {
			// 将i分割成 j + (i-j)
			memo[i] = max(memo[i], j*(i-j), j*memo[i-j])
		}
	}
	return memo[n]
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
