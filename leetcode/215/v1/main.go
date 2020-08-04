package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums1 := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(nums1, 2))
	nums2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(nums2, 4))
}

// #215
// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, l, r, index int) int {
	p := randomPartition(nums, l, r)
	if p == index {
		return nums[p]
	} else if p < index {
		return quickSelect(nums, p+1, r, index)
	}
	return quickSelect(nums, l, p-1, index)
}

// 随机选择标的点进行分区
func randomPartition(nums []int, l, r int) int {
	i := rand.Int()%(r-l+1) + l
	nums[i], nums[l] = nums[l], nums[i]
	return partition(nums, l, r)
}

func partition(nums []int, l, r int) int {
	v := nums[l]
	j := l
	for i := l; i <= r; i++ {
		if nums[i] < v {
			nums[i], nums[j+1] = nums[j+1], nums[i]
			j++
		}
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}
