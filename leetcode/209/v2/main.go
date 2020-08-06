package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(7, nums))
}

// https://leetcode-cn.com/problems/minimum-size-subarray-sum/
// 滑动窗口
func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	res := math.MaxInt64
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum >= s {
				res = min(res, j-i+1)
				break
			}
		}
	}
	if res == math.MaxInt64 {
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
