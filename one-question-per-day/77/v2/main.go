package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
}

//https://leetcode-cn.com/problems/combinations/
func combine(n int, k int) [][]int {
	res = [][]int{}
	if n <= 0 || k <= 0 || k > n {
		return res
	}
	getCombinations(n, k, 1, []int{})
	return res
}

var res [][]int

func getCombinations(n, k, start int, c []int) {
	if len(c) == k {
		tmp := make([]int, len(c))
		copy(tmp, c)
		res = append(res, tmp)
	} else {
		// 剪枝
		// 等同于for循环的终止条件改为：i <= n - (k-len(c)) + 1
		if len(c)+n-k+1 < k {
			return
		}
		for i := start; i <= n; i++ {
			c = append(c, i)
			getCombinations(n, k, i+1, c)
			c = c[:len(c)-1]
		}
	}
}
