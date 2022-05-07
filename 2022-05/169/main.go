package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}

// https://leetcode-cn.com/problems/majority-element/
// 先将数组排序，而后数组中位数即为众数
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
