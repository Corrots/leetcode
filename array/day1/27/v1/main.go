package main

import "fmt"

/*
#27
移除元素
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
*/
func main() {
	nums := []int{3, 2, 2, 3}
	nums1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, 3))
	fmt.Println(removeElement(nums1, 2))
}

func removeElement(nums []int, val int) int {
	n := len(nums)
	// 循环不变量：nums[0:k)都是不=val的，nums[k:n)均=val
	k := 0
	for i := 0; i < n; i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
