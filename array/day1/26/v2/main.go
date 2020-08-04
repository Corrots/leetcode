package main

import "fmt"

/**
#26
删除排序数组中的重复项
给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
*/
func main() {
	nums1 := []int{1, 1, 2}
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums1))
	fmt.Println(removeDuplicates(nums2))
}

// 考虑特殊场景，如果数组中没有重复项，只在q-p>1时才进行复制
func removeDuplicates(nums []int) int {
	p, q := 0, 1
	for q < len(nums) {
		if nums[p] != nums[q] {
			if q-p > 1 {
				nums[p+1] = nums[q]
			}
			p++
		}
		q++
	}
	return p + 1
}
