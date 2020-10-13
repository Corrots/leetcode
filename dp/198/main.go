package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}

//https://leetcode-cn.com/problems/house-robber/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	memo := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		memo[i] = -1
	}
	var tryRob func(int) int
	tryRob = func(index int) int {
		if index >= len(nums) {
			return 0
		}
		if memo[index] != -1 {
			return memo[index]
		}
		res := -1
		for i := index; i < len(nums); i++ {
			res = max(res, nums[i]+tryRob(i+2))
		}
		memo[index] = res
		return res
	}
	return tryRob(0)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
