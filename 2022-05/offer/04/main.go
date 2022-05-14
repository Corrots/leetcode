package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	fmt.Println(findNumberIn2DArray(matrix, 5))
	fmt.Println(findNumberIn2DArray(matrix, 20))

	fmt.Println(findNumberIn2DArray([][]int{
		{1, 1},
	}, 2))
}

// https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
// 剑指 Offer 04. 二维数组中的查找
// Z字形查找
// 查找的初始位置为 matrix[0][ len(matrix[0])-1]
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	rows, cols := len(matrix), len(matrix[0])
	p, q := 0, cols-1
	for p < rows && q >= 0 {
		if matrix[p][q] == target {
			return true
		} else if matrix[p][q] > target {
			q--
		} else {
			p++
		}
	}
	return false
}
