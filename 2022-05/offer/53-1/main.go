package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(search(nums, 8))
}

// https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
// 统计一个数字在排序数组中出现的次数。
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	// 搜索第一个=target的元素
	start := binarySearch(nums, l, r, target)
	// 搜索第一个>target的元素
	end := binarySearch(nums, start, r, target+1)
	return end - start
}

func binarySearch(nums []int, l, r, target int) int {
	for l <= r {
		mid := (r-l)/2 + l
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
