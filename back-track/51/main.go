package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 4
	fmt.Println(solveNQueens(n))
}

//https://leetcode-cn.com/problems/n-queens/
func solveNQueens(n int) [][]string {
	var res [][]string
	// 记录每一列摆放皇后的位置
	col := make([]bool, n)
	// 记录对角线1和2上摆放皇后的位置
	dial1, dial2 := make([]bool, 2*n-1), make([]bool, 2*n-1)
	var putQueens func(int, []int)
	// 找到第i行摆放皇后的位置
	putQueens = func(i int, row []int) {
		if i == n {
			res = append(res, formatRow(n, row))
			return
		}
		// 遍历i行的每一列j，找到摆放皇后的合适位置
		for j := 0; j < n; j++ {
			if !col[j] && !dial1[i+j] && !dial2[i-j+n-1] {
				col[j] = true
				dial1[i+j] = true
				dial2[i-j+n-1] = true
				tmp := make([]int, len(row))
				copy(tmp, row)
				tmp = append(tmp, j)
				// 在第i+1行放置皇后
				putQueens(i+1, tmp)
				// 回溯
				col[j] = false
				dial1[i+j] = false
				dial2[i-j+n-1] = false
			}
		}
	}
	putQueens(0, []int{})
	return res
}

func formatRow(n int, row []int) []string {
	var res []string
	//fmt.Println(row)
	for i := 0; i < n; i++ {
		var str []string
		for j := 0; j < n; j++ {
			//fmt.Printf("j: %d, row[i]: %d\n", j, row[i])
			if j == row[i] {
				str = append(str, "Q")
			} else {
				str = append(str, ".")
			}
		}
		res = append(res, strings.Join(str, ""))
	}
	return res
}
