package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}

// 根据状态的定义，决定状态转移
//https://leetcode-cn.com/problems/house-robber/
// 记忆化搜索
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	memo := make([]int, len(nums))
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(index int) int {
		// 递归终止条件
		if index >= len(nums) {
			return 0
		}
		if memo[index] != -1 {
			return memo[index]
		}
		res := math.MinInt64
		for i := index; i < len(nums); i++ {
			res = max(res, nums[i]+dfs(i+2))
		}
		memo[index] = res
		return res
	}
	return dfs(0)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
