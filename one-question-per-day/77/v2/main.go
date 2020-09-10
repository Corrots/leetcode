package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
}

//https://leetcode-cn.com/problems/combinations/
func combine(n int, k int) [][]int {
	var res [][]int
	if n <= 0 || k <= 0 || k > n {
		return res
	}
	var dfs func(int, []int)
	dfs = func(start int, path []int) {
		if len(path) == k {
			res = append(res, append([]int{}, path...))
			return
		}
		if len(path)+(n-start+1) < k {
			return
		}
		for i := start; i <= n; i++ {
			tmp := make([]int, len(path))
			copy(tmp, path)
			tmp = append(tmp, i)
			dfs(i+1, tmp)
		}
	}
	dfs(1, []int{})
	return res
}
