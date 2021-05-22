package main

import "fmt"

// 利用hashmap求解，但是未使用到数组有序的特性

//https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/
func twoSum(nums []int, target int) []int {
	if len(nums) <= 0 {
		return nil
	}
	m := make(map[int]bool)
	for _, num := range nums {
		if _, ok := m[target-num]; ok {
			return []int{num, target - num}
		}
		m[num] = true
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
	nums1 := []int{10, 26, 30, 31, 47, 60}
	fmt.Println(twoSum(nums1, 40))
}
