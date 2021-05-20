package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

// 1.初始化第0个元素为最大值max
// 2.从元素1开始向后遍历，若当前元素+前一个元素>当前元素，则说明前一个元素非负，
// 3.将累加后的结果与max比较，若大于，则更新max为累加后的值
// 4.若小于，max保持原值

// https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
