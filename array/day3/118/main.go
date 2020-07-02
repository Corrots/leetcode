package main

import "fmt"

/**
#118
杨辉三角
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行
在杨辉三角中，每个数是它左上方和右上方的数的和。
*/
func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	var res [][]int
	if numRows == 0 {
		return res
	}
	res = append(res, []int{1})
	for i := 1; i < numRows; i++ {
		// 在每个row的头部加入元素0
		m := []int{0}
		m = append(m, res[i-1]...)
		for j := 0; j < len(m)-1; j++ {
			m[j] = m[j] + m[j+1]
		}
		res = append(res, m)
	}
	return res
}
