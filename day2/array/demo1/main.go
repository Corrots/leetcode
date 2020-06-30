package main

import "fmt"

/**
给定一个整数数组，判断是否存在重复元素。
如果任意一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。
*/
func main() {
	nums1 := []int{1, 2, 3, 1}
	nums2 := []int{1, 2, 3, 4}
	nums3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

	fmt.Println(containsDuplicate(nums1))
	fmt.Println(containsDuplicate(nums2))
	fmt.Println(containsDuplicate(nums3))
}

// 朴素线性查找 【超时】
//func containsDuplicate(nums []int) bool {
//	for i := 0; i < len(nums); i++ {
//		for j := 0; j < i; j++ {
//			if nums[i] == nums[j] {
//				return true
//			}
//		}
//	}
//	return false
//}

// 利用golang的map
func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = 0
	}
	return false
}
