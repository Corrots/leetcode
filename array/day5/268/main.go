package main

import "fmt"

/**
#268
缺失数字
给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数
*/
func main() {
	nums1 := []int{3, 0, 1}
	nums2 := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(nums1))
	fmt.Println(missingNumber(nums2))
}

func missingNumber(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v] = 0
	}

	for i := 0; i < len(nums)+1; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return 0
}
