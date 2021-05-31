package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	l := sort.SearchInts(nums, target)
	if l == len(nums) || nums[l] != target {
		return []int{-1, -1}
	}
	r := sort.SearchInts(nums, target+1) - 1
	return []int{l, r}
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(nums, 8))
	fmt.Println(searchRange(nums, 6))
}
