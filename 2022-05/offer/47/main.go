package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(maxValue(grid))
}

// https://leetcode.cn/problems/li-wu-de-zui-da-jie-zhi-lcof/
// 剑指 Offer 47. 礼物的最大价值
// 状态转移方程： f(i, j) = max(f(i, j-1), f(i-1, j)) + grid[i, j]
func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += max(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[m-1][n-1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
