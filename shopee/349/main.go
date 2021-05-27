package main

import "sort"

/**
排序+双指针
先对两个数组进行排序，然后使用两个指针遍历两个数组。
初始时，两个指针分别指向两个数组的头部。每次比较两个指针指向的两个数组中的数字，如果两个数字不相等，则将指向较小数字的指针右移一位，
如果两个数字相等，且该数字在结果集中不存在，则将其加入结果集，同时将两个指针都右移一位。
当至少有一个指针超出数组范围时，遍历结束。
*/

// https://leetcode-cn.com/problems/intersection-of-two-arrays/
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
		if n1 == n2 && (res == nil || n1 != res[len(res)-1]) {
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
