package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	nums3 := []int{4, 9, 5}
	nums4 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(nums1, nums2))
	fmt.Println(intersection(nums3, nums4))
}

// https://leetcode-cn.com/problems/intersection-of-two-arrays/
func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	set := make(map[int]bool)
	for _, v := range nums1 {
		set[v] = true
	}
	var res []int
	for _, v := range nums2 {
		if set[v] {
			res = append(res, v)
			set[v] = false
		}
	}
	return res
}
