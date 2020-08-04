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
	var p, q, r int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			p++
		} else if nums[i] == 1 {
			q++
		} else {
			r++
		}
	}
	for i := 0; i < p; i++ {
		nums[i] = 0
	}
	for i := p; i < p+q; i++ {
		nums[i] = 1
	}
	for i := p + q; i < len(nums); i++ {
		nums[i] = 2
	}
}
