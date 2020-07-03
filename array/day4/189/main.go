package main

import "fmt"

/**
#189
旋转数组
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数
*/
func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	nums2 := []int{-1, -100, 3, 99}
	rotate(nums1, 3)
	fmt.Println(nums1)
	rotate(nums2, 2)
	fmt.Println(nums2)
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	count := len(nums)
	copy(nums, append(nums[count-k:], nums[:count-k]...))
}
