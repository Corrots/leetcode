package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(nums))
}

// https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
func findRepeatNumber(nums []int) int {
	// 先对数组进行排序
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return nums[i]
		}
	}
	return -1
}
