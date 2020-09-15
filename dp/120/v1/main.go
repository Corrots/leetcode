package main

import "fmt"

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle))
}

//https://leetcode-cn.com/problems/triangle/
// 递归求解，超时
func minimumTotal(triangle [][]int) int {
	var dfs func(int, int) int
	dfs = func(i int, j int) int {
		if i == len(triangle) {
			return 0
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
