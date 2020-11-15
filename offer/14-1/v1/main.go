package main

import "fmt"

func main() {
	fmt.Println(cuttingRope(10))
}

//https://leetcode-cn.com/problems/jian-sheng-zi-lcof/
func cuttingRope(n int) int {
	return cut(n)
}

// 将n进行分割(至少两部分)，获得的最大乘积
func cut(n int) int {
	if n == 1 {
		return 1
	}
	res := -1
	for i := 1; i < n; i++ {
		res = max(res, i*(n-i), i*cut(n-i))
	}
	return res
}

func max(args ...int) int {
	var res int
	for _, v := range args {
		if v > res {
			res = v
		}
	}
	return res
}
