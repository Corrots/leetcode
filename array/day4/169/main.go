package main

import (
	"fmt"
)

/**
#169
多数元素
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
*/
func main() {
	nums1 := []int{3, 2, 3}
	nums2 := []int{2, 2, 1, 1, 1, 2, 2}
	nums3 := []int{1}
	fmt.Println(majorityElement(nums1))
	fmt.Println(majorityElement(nums2))
	fmt.Println(majorityElement(nums3))
}

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
		if m[v] > len(nums)/2 {
			return v
		}
	}
	return 0
}
