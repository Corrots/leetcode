package main

import "fmt"

/**
#283
移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
*/
func main() {
	nums1 := []int{0, 1, 0, 3, 12}
	moveZeroes(nums1)
	fmt.Println(nums1)
	// output: [1 3 12 0 0]
}

/**
//
func moveZeroes(nums []int) {
	p := 0
	for _, v := range nums {
		if v != 0 {
			nums[p] = v
			p++
		}
	}
	//
	for i := p; i < len(nums); i++ {
		nums[i] = 0
	}
}
*/

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
