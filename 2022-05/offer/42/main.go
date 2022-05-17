package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

// https://leetcode.cn/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
// 剑指 Offer 42. 连续子数组的最大和
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			// 前一个元素 nums[i-1] 为正数
			// 将前一个元素 nums[i-1] 加入到当前元素中
			nums[i] += nums[i-1]
		}
		// 若当前元素>max，则当前连续子数组的最大和
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
