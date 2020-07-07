package main

import "fmt"

/**
#498
对角线遍历
给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素，对角线遍历如下图所示。
*/
func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(findDiagonalOrder(matrix))
}

func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	row := len(matrix)
	col := len(matrix[0])
	i := 0
	var res []int
	for i < row+col-1 {
		x1 := 0
		if i < row {
			x1 = i
		} else {
			x1 = row - 1
		}
		y1 := i - x1

		for x1 >= 0 && y1 < col {
			res = append(res, matrix[x1][y1])
			x1--
			y1++
		}
		i++
		//
		y2 := 0
		if i < col {
			y2 = i
		} else {
			y2 = col - 1
		}
		x2 := i - y2
		for y2 >= 0 && x2 < row {
			res = append(res, matrix[x2][y2])
			x2++
			y2--
		}
		i++
	}
	return res
}
