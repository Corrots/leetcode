package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(7, nums))
}

// https://leetcode-cn.com/problems/minimum-size-subarray-sum/
// 滑动窗口
func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := len(nums) + 1
	sum := 0
	l, r := 0, -1
	for l < len(nums) {
		if r+1 < len(nums) && sum < s {
			r++
			sum += nums[r]
		} else {
			sum -= nums[l]
			l++
		}
		if sum >= s {
			res = min(res, r-l+1)
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
