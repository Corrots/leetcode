package main

import "fmt"

//删除排序数组中的重复项
func main() {
	nums1 := []int{1, 1, 2}
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums1))
	fmt.Println(removeDuplicates(nums2))
}

// 双指针解法
func removeDuplicates(nums []int) int {
	p, q := 0, 1
	for q < len(nums) {
		if nums[p] != nums[q] {
			nums[p+1] = nums[q]
			p++
		}
		q++
	}
	return p + 1
}
