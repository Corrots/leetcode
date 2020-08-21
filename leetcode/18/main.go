package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-3, -2, -1, 0, 0, 1, 2, 3}
	fmt.Println(fourSum(nums, 0))
	nums1 := []int{0, 0, 0, 0}
	fmt.Println(fourSum(nums1, 0))
}

// https://leetcode-cn.com/problems/4sum/
func fourSum(nums []int, target int) [][]int {
	var res [][]int
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)
	//fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, len(nums)-1
			for l < r {
				sum := nums[i] + nums[l] + nums[j] + nums[r]
				if sum < target {
					l++
				} else if sum > target {
					r--
				} else {
					res = append(res, []int{nums[i], nums[l], nums[j], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				}
			}
		}
	}
	return res
}
