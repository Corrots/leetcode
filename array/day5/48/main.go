package main

import "fmt"

/**
#01.07
#48
旋转矩阵
给你一幅由 N × N 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。
不占用额外内存空间能否做到？
*/
func main() {
	matrix1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix2 := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(matrix1)
	rotate(matrix2)
	fmt.Println(matrix1)
	fmt.Println(matrix2)
}

/**
1 2 3
4 5 6
7 8 9
// 以4 5 6为中线翻转
7 8 9
4 5 6
1 2 3
// 以对角线753为中线翻转

*/

func rotate(matrix [][]int) {
	n := len(matrix)
	// 以水平线456为中线翻转
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[n-i-1][j] = matrix[n-i-1][j], matrix[i][j]
		}
	}
	// 以对角线753为中线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
