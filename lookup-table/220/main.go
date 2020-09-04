package main

import (
	"fmt"
	"math"
)

func main() {
	nums, k1, t1 := []int{1, 2, 3, 1}, 3, 0
	fmt.Println(containsNearbyAlmostDuplicate(nums, k1, t1))
}

//https://leetcode-cn.com/problems/contains-duplicate-iii/
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j <= i+k && j < n; j++ {
			if int(math.Abs(float64(nums[i]-nums[j]))) <= t {
				return true
			}
		}
	}
	return false
}
