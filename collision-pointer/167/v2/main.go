package main

import "fmt"

/**
#167
两数之和 II - 输入有序数组
给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。
说明:
返回的下标值（index1 和 index2）不是从零开始的。
你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
*/
func main() {
	nums1 := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums1, 9))
	fmt.Println(twoSum(nums1, 22))
	fmt.Println(twoSum(nums1, 17))
}

// 对撞指针 O(n)
func twoSum(nums []int, target int) []int {
	p, q := 0, len(nums)-1
	for p < q {
		if nums[p]+nums[q] == target {
			return []int{p + 1, q + 1}
		} else if nums[p]+nums[q] > target {
			q--
		} else {
			p++
		}
	}
	return nil
}
