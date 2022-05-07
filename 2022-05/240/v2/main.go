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
	// 使用Z字形查找求解
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row, col := 0, len(matrix[0])-1
	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}

	return false
}
