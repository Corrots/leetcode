package main

import "fmt"

func main() {
	fmt.Println(cuttingRope(10))
}

//https://leetcode-cn.com/problems/jian-sheng-zi-lcof/
func cuttingRope(n int) int {
	if n == 1 {
		return 0
	}
	memo := make([]int, n+1)
	memo[1] = 1
	for i := 2; i < n+1; i++ {
		for j := 1; j < i; j++ {
			memo[i] = max(memo[i], j*(i-j), j*memo[i-j])
		}
	}
	return memo[n]
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
