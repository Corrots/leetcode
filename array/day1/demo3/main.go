package main

import (
	"fmt"
	"math"
)

/*
搜索插入位置
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
你可以假设数组中无重复元素。
*/
func main() {
	var nums = []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 5))
	fmt.Println(searchInsert(nums, 2))
	fmt.Println(searchInsert(nums, 7))
	fmt.Println(searchInsert(nums, 0))
	// 2 1 4 0
}

func searchInsert(nums []int, target int) int {
	count := len(nums)
	center := int(math.Ceil(float64(count / 2)))

	var index int
	if target >= nums[center] {
		// 搜索后半部分
		for i := center; i < count; i++ {
			if target <= nums[i] {
				index = i
				break
			}
			index = i + 1
		}
	} else {
		// 搜索前半部分
		for i := 0; i < center; i++ {
			if target <= nums[i] {
				index = i
				break
			}
			index = i + 1
		}
	}
	return index
}
