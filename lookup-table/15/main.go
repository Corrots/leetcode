package main

import (
	"fmt"
	"sort"
)

func main() {
	//nums := []int{-1, 0, 1, 2, -1, -4}
	nums1 := []int{0, 0, 0, 0}
	//fmt.Println(threeSum(nums))
	fmt.Println(threeSum(nums1))
}

// https://leetcode-cn.com/problems/3sum/
func threeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	//fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
