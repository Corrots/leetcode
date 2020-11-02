package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
	nums3 := []int{4, 9, 5}
	nums4 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(nums3, nums4))
}

// 解法：排序+双指针
//https://leetcode-cn.com/problems/intersection-of-two-arrays/
func intersection(nums1, nums2 []int) []int {
	var res []int
	if len(nums1) == 0 || len(nums2) == 0 {
		return res
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	var p, q int
	for p < len(nums1) && q < len(nums2) {
		x, y := nums1[p], nums2[q]
		if x == y {
			if res == nil || x > res[len(res)-1] {
				//if n == 0 || res[n-1] != x {
				res = append(res, x)
			}
			p++
			q++
		} else if x > y {
			q++
		} else {
			p++
		}
	}
	return res
}
