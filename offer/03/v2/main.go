package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(nums))
}

// TODO 尝试空间复杂度O(1)的解法
//https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
func findRepeatNumber(nums []int) int {
	// 先排序后，重复的元素必然相邻
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return nums[i-1]
		}
	}
	return -1
}
