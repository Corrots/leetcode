package main

import "fmt"

/**
#88
合并两个有序数组
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
说明：
初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
*/
func main() {
	//nums1 := []int{1, 2, 3, 0, 0, 0}
	//nums2 := []int{2, 5, 6}
	//merge(nums1, 3, nums2, 3)
	nums3 := []int{4, 0, 0, 0, 0, 0}
	nums4 := []int{1, 2, 3, 5, 6}
	merge(nums3, 1, nums4, 5)
	fmt.Println(nums3)
}

// 归并排序法的归并思路
func merge(nums1 []int, m int, nums2 []int, n int) {
	tmp := make([]int, m)
	copy(tmp, nums1)
	i, j := 0, 0
	for k := 0; k < m+n; k++ {
		if i >= m {
			nums1[k] = nums2[j]
			j++
		} else if j >= n {
			nums1[k] = tmp[i]
			i++
		} else if tmp[i] < nums2[j] {
			nums1[k] = tmp[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
	}
}
