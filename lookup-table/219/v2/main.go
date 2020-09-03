package main

import (
	"fmt"
)

/**
#219
存在重复元素 II
给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值 至多为 k。
*/
func main() {
	nums1 := []int{1, 2, 3, 1}
	nums2 := []int{1, 0, 1, 1}
	nums3 := []int{1, 2, 3, 1, 2, 3}

	fmt.Println(containsNearbyDuplicate(nums1, 3))
	fmt.Println(containsNearbyDuplicate(nums2, 1))
	fmt.Println(containsNearbyDuplicate(nums3, 2))
	// true true false
}

// 时间: O(n)
// 空间: O(k)
func containsNearbyDuplicate(nums []int, k int) bool {
	record := make(map[int]bool)
	for i, v := range nums {
		if record[v] {
			return true
		}
		record[v] = true
		// 保持record中最多有k个元素
		if len(record) == k+1 {
			delete(record, nums[i-k])
		}
	}
	return false
}
