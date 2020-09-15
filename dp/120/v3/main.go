package main

import "fmt"

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle))
}

//https://leetcode-cn.com/problems/triangle/
// DP
//定义二维 dp 数组，将解法二中「自顶向下的递归」改为「自底向上的递推」
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		for j := 0; j < n+1; j++ {
			dp[i] = append(dp[i], 0)
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
		}
	}
	return dp[0][0]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
