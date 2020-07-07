package main

import "fmt"

/**
#73
编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。

*/
func main() {
	matrix1 := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	matrix2 := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix1)
	setZeroes(matrix2)
	fmt.Println(matrix1)
	fmt.Println(matrix2)
}

func setZeroes(matrix [][]int) {
	row, col := make(map[int]int), make(map[int]int)
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = 1
				col[j] = 1
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if row[i] == 1 || col[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}
