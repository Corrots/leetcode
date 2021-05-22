package main

import "fmt"

// 利用数组有序的特性
// 对撞指针

//https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/
func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	p, q := 0, len(nums)-1
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

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
	nums1 := []int{10, 26, 30, 31, 47, 60}
	fmt.Println(twoSum(nums1, 40))
}
