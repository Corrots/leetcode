package main

import "fmt"

/**
#724
给定一个整数类型的数组 nums，请编写一个能够返回数组 “中心索引” 的方法。
我们是这样定义数组 中心索引 的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。
如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。
*/
func main() {
	var nums = []int{1, 7, 3, 6, 5, 6}
	var nums1 = []int{1, 2, 3}
	var nums2 = []int{-1, -1, -1, -1, -1, -1}
	var nums3 = []int{-1, -1, -1, 0, 1, 1}
	var nums4 = []int{-1, -1, 0, 1, 1, 0}
	var nums5 = []int{-1, -1, 0, 1, 1, -1}
	fmt.Println(pivotIndex(nums))
	fmt.Println(pivotIndex(nums1))
	fmt.Println(pivotIndex(nums2))
	fmt.Println(pivotIndex(nums3))
	fmt.Println(pivotIndex(nums4))
	fmt.Println(pivotIndex(nums5))
}

func pivotIndex(nums []int) int {
	var total int
	for _, v := range nums {
		total += v
	}
	var left int
	for k, v := range nums {
		if left*2+v == total {
			return k
		}
		left += v
	}
	return -1
}
