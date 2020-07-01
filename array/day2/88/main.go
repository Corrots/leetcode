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

var nums = []int{4, 0, 0, 4, 5, 6}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 && n > 0 {
		if nums2[n-1] > nums1[m-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}
	// 若m为0，则表示nums1中所有元素都比nums2中元素大
	if m == 0 {
		// 此处i < n，不能直接遍历nums2
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
	}
}
