package main

import "fmt"

// 从右上角的元素开始查找
// 若>target，排除当前列
// 若<target，排除当前行

//https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	rows, cols := len(matrix), len(matrix[0])
	row, col := 0, cols-1
	for row < rows && col >= 0 {
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

func main() {
	matrix := [][]int{{1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}}
	fmt.Println(findNumberIn2DArray(matrix, 7))
}
