package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}

//https://leetcode-cn.com/problems/house-robber/
// 动态规划
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	memo := make([]int, n)
	memo[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		for j := i; j < n; j++ {
			var extra int
			if j+2 < n {
				extra = memo[j+2]
			}
			memo[i] = max(memo[i], nums[j]+extra)
		}
	}
	return memo[0]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
