package main

import (
	"fmt"
)

func main() {
	nums := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	nums2 := [][]int{{1, 3, 5}}
	fmt.Println(searchMatrix(nums, 5))
	fmt.Println(searchMatrix(nums2, 5))
}

// https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
func searchMatrix(matrix [][]int, target int) bool {
	// 使用二分查找法求解
	for _, row := range matrix {
		l, r := 0, len(row)-1
		for l <= r {
			mid := (r-l)/2 + l
			if row[mid] == target {
				return true
			} else if row[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return false
}
