package main

import "fmt"

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(nums, 2))
}

//https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
func findKthLargest(nums []int, k int) int {
	k = len(nums) - k
	l, r := 0, len(nums)-1
	for l <= r {
		p := partition(nums, l, r)
		if p == k {
			return nums[p]
		} else if k < p {
			r = p - 1
		} else {
			l = p + 1
		}
	}
	return -1
}

func partition(nums []int, l, r int) int {
	v := nums[l]
	// <v和>v元素的分界点，初始值为l
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
