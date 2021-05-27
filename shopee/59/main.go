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
	var row, col, dirIndex int
	for i := 1; i <= n*n; i++ {
		matrix[row][col] = i
		dir := dirs[dirIndex]
		r, c := row+dir.x, col+dir.y
		// 若下一步的位置超出矩阵边界，或者是之前访问过的位置，则顺时针旋转，进入下一个方向
		if r < 0 || r >= n || c < 0 || c >= n || matrix[r][c] > 0 {
			dirIndex = (dirIndex + 1) % 4
			dir = dirs[dirIndex]
		}
		row += dir.x
		col += dir.y
	}
	return matrix
}

func main() {
	fmt.Println(generateMatrix(3))
}
