package main

import (
	"fmt"
	"math"
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

func containsNearbyDuplicate(nums []int, k int) bool {
	record := make(map[int]int)
	for i, v := range nums {
		if index, ok := record[v]; ok {
			if int(math.Abs(float64(i-index))) <= k {
				return true
			}
		}
		record[v] = i
	}
	return false
}
