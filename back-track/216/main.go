package main

import "fmt"

func main() {
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum3(3, 9))
}

//https://leetcode-cn.com/problems/combination-sum-iii/
func combinationSum3(k int, n int) [][]int {
	var res [][]int
	if k <= 0 || n <= 0 || n < k {
		return res
	}
	var dfs func(int, int, []int)
	dfs = func(start int, target int, path []int) {
		if len(path) == k {
			if target == 0 {
				res = append(res, append([]int{}, path...))
			}
			return
		} else if len(path) < k {
			if target <= 0 {
				return
			}
			// target > 0
			for i := start; i <= 9; i++ {
				tmp := make([]int, len(path))
				copy(tmp, path)
				tmp = append(tmp, i)
				dfs(i+1, target-i, tmp)
			}
		}
	}
	dfs(1, n, []int{})
	return res
}
