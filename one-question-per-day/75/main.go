package main

import (
	"fmt"
)

func main() {
	fmt.Println(totalNQueens(4))
	fmt.Println(totalNQueens(8))
}

func totalNQueens(n int) int {
	var res int
	// 记录每一列是否有摆放皇后
	col := make([]bool, n)
	// 记录正对角线和斜对角线上是否有摆放皇后
	dial1, dial2 := make([]bool, 2*n-1), make([]bool, 2*n-1)
	var backTrack func(int)
	// 找到第i行皇后应摆放的合适位置
	backTrack = func(i int) {
		if i == n {
			res++
			return
		}
		// 遍历第i行的每一列，尝试摆放皇后
		for j := 0; j < n; j++ {
			if !col[j] && !dial1[i+j] && !dial2[i-j+n-1] {
				col[j] = true
				dial1[i+j] = true
				dial2[i-j+n-1] = true
				backTrack(i + 1)
				col[j] = false
				dial1[i+j] = false
				dial2[i-j+n-1] = false
			}
		}
	}
	backTrack(0)
	return res
}
