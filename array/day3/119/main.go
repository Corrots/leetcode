package main

import "fmt"

/**
#119
杨辉三角 II
给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
在杨辉三角中，每个数是它左上方和右上方的数的和。
*/
func main() {
	fmt.Println(getRow(33))
}

func getRow(rowIndex int) []int {
	var res [][]int
	res = append(res, []int{1})
	for i := 1; i <= rowIndex; i++ {
		m := []int{0}
		m = append(m, res[i-1]...)
		for j := 0; j < len(m)-1; j++ {
			m[j] = m[j] + m[j+1]
		}
		res = append(res, m)
	}
	return res[rowIndex]
}
