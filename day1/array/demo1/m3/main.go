package main

import "fmt"

func main() {
	var nums = []int{1, 7, 3, 6, 5, 6}
	var nums1 = []int{1, 2, 3}
	var nums2 = []int{-1, -1, -1, -1, -1, -1}
	var nums3 = []int{-1, -1, -1, 0, 1, 1}
	var nums4 = []int{-1, -1, 0, 1, 1, 0}
	var nums5 = []int{-1, -1, 0, 1, 1, -1}
	fmt.Println(pivotIndex(nums))
	fmt.Println(pivotIndex(nums1))
	fmt.Println(pivotIndex(nums2))
	fmt.Println(pivotIndex(nums3))
	fmt.Println(pivotIndex(nums4))
	fmt.Println(pivotIndex(nums5))
}

func pivotIndex(nums []int) int {
	var total int
	for i := range nums {
		total += nums[i]
	}

	lSum := 0
	center := -1
	for i, num := range nums {
		if lSum*2+num == total {
			center = i
			break
		}
		lSum += num
	}
	return center
}
