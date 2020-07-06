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
	res = append(res, []int{1})
	for i := 1; i < numRows-1; i++ {
		line := []int{0}
		line = append(line, res[i-1]...)
		for j := 0; j < len(line)-1; j++ {
			line[j] = line[j] + line[j+1]
		}
		res = append(res, line)
	}
	return res
}
