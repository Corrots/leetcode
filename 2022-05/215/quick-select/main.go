package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(nums, 2))
	nums1 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(nums1, 4))
}

// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
// 215. 数组中的第K个最大元素
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 || len(nums) < k {
		return -1
	}
	k = len(nums) - k
	l, r := 0, len(nums)-1
	for l <= r {
		p := partition(nums, l, r)
		if p == k {
			return nums[p]
		} else if p < k {
			l = p + 1
		} else {
			r = p - 1
		}
	}
	return -1
}

func partition(nums []int, l, r int) int {
	//random := rand.Int()%(r-l+1) + l
	//nums[l], nums[random] = nums[random], nums[l]
	v := nums[l]
	j := l
	for i := l + 1; i <= r; i++ {
		if nums[i] < v {
			nums[i], nums[j+1] = nums[j+1], nums[i]
			j++
		}
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}
