package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

// https://leetcode.cn/problems/he-wei-sde-liang-ge-shu-zi-lcof/
// 剑指 Offer 57. 和为s的两个数字
func twoSum(nums []int, target int) []int {
	n := len(nums)
	if n <= 1 {
		return nil
	}
	p, q := 0, n-1
	for p < q {
		if nums[p]+nums[q] == target {
			return []int{nums[p], nums[q]}
		} else if nums[p]+nums[q] > target {
			q--
		} else {
			p++
		}
	}
	return nil
}
