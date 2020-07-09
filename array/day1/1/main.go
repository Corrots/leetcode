package main

import "fmt"

/**
#1
两数之和
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
*/
func main() {
	nums := []int{2, 7, 11, 15}
	nums2 := []int{3, 2, 4}
	nums3 := []int{3, 3}
	fmt.Println(twoSum(nums, 9))
	fmt.Println(twoSum(nums2, 6))
	fmt.Println(twoSum(nums3, 6))

}

/*
1. 暴力解法
时间复杂度：O(n^2)

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
*/

/**
2. 使用哈希表，通过一次遍历解决
*/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		leave := target - v
		if i, ok := m[leave]; ok {
			return []int{k, i}
		}
		m[v] = k
	}
	return nil
}
