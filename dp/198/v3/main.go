package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}

//https://leetcode-cn.com/problems/house-robber/
// 动态规划，优化状态转移
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	memo := make([]int, n)
	memo[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		// 选择1：放弃当前房子，从下一个房子开始
		// 选择2：抢劫当前房子，从i+2后的房子开始
		extra := 0
		if i+2 < n {
			extra = memo[i+2]
		}
		memo[i] = max(memo[i+1], nums[i]+extra)
	}
	return memo[0]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
