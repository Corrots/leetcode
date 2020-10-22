package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

// #75 颜色分类
// https://leetcode-cn.com/problems/sort-colors/
func sortColors(nums []int) {
	var counts [3]int
	for i := 0; i < len(nums); i++ {
		counts[nums[i]]++
	}

	index := 0
	for i := 0; i < counts[0]; i++ {
		nums[index] = 0
		index++
	}
	for i := 0; i < counts[1]; i++ {
		nums[index] = 1
		index++
	}
	for i := 0; i < counts[2]; i++ {
		nums[index] = 2
		index++
	}
}
