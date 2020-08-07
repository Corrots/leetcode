package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	nums3 := []int{4, 9, 5}
	nums4 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(nums1, nums2))
	fmt.Println(intersect(nums3, nums4))
}

// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
func intersect(nums1 []int, nums2 []int) []int {
	record := make(map[int]int)
	for k := range nums1 {
		record[nums1[k]]++
	}
	var res []int
	for _, v := range nums2 {
		if record[v] > 0 {
			res = append(res, v)
			record[v]--
		}
	}
	return res
}
