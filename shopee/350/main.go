package main

import "sort"

/**
同#349，且不需判断结果集中是否存在改相同元素
*/

// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) <= 0 || len(nums2) <= 0 {
		return nil
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	var res []int
	var p, q int
	for p < len(nums1) && q < len(nums2) {
		n1, n2 := nums1[p], nums2[q]
		if n1 == n2 {
			res = append(res, n1)
			p++
			q++
		} else if n1 < n2 {
			p++
		} else {
			q++
		}
	}
	return res
}

func main() {

}
