package main

import "fmt"

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle))
}

//https://leetcode-cn.com/problems/triangle/
// 尝试记忆化搜索求解，超时
// 时间复杂度：O(n^2)
// 空间复杂度：O(n^2), N为三角形的行数
func minimumTotal(triangle [][]int) int {
	var dfs func(int, int) int
	memo := make([][]int, len(triangle))
	for i := 0; i < len(memo); i++ {
		for j := 0; j < len(memo); j++ {
			memo[i] = append(memo[i], -1)
		}
	}
	dfs = func(i int, j int) int {
		if i == len(triangle) {
			return 0
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		return min(dfs(i+1, j), dfs(i+1, j+1)) + triangle[i][j]
	}
	return dfs(0, 0)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
