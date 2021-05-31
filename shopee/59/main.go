package main

import "fmt"

type point struct {
	x, y int
}

var dirs = []point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

//https://leetcode-cn.com/problems/spiral-matrix-ii/
func generateMatrix(n int) [][]int {
	// 初始化矩阵，所有点初始值为0
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	var row, col, index int
	for i := 1; i <= n*n; i++ {
		matrix[row][col] = i
		dir := dirs[index]
		r, c := row+dir.x, col+dir.y
		if r < 0 || r >= n || c < 0 || c >= n || matrix[r][c] > 0 {
			index = (index + 1) % 4
			dir = dirs[index]
		}
		row += dir.x
		col += dir.y
	}
	return matrix
}

func main() {
	fmt.Println(generateMatrix(3))
}
