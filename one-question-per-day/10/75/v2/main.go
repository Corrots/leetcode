package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

// #75 颜色分类
// https://leetcode-cn.com/problems/sort-colors/
// 使用三路快排的思想解题
func sortColors(nums []int) {
	n := len(nums)
	// 循环不变量: nums[0:zero]都是0, nums[two:n-1]都是2
	zero, two := -1, n
	i := 0
	// 循环终止条件, i<two
	for i < two {
		if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			nums[two-1], nums[i] = nums[i], nums[two-1]
			two--
		} else if nums[i] == 0 {
			nums[zero+1], nums[i] = nums[i], nums[zero+1]
			zero++
			i++
		}
	}
}
