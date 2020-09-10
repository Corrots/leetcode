package main

import (
	"fmt"
)

func main() {
	candidates := []int{2, 3, 6, 7}
	fmt.Println(combinationSum(candidates, 7))
	//candidates1 := []int{2, 3, 5}
	//fmt.Println(combinationSum(candidates1, 8))
}

// https://leetcode-cn.com/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	if len(candidates) == 0 {
		return res
	}
	var dfs func(int, int, []int)
	dfs = func(start, target int, path []int) {
		if target < 0 || start == len(candidates) {
			return
		}
		if target == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		// target > 0
		for i := start; i < len(candidates); i++ {
			tmp := make([]int, len(path))
			copy(tmp, path)
			tmp = append(tmp, candidates[i])
			fmt.Printf("start: %d, target: %d, path: %v\n", i, target-candidates[i], tmp)
			dfs(i, target-candidates[i], tmp)
			fmt.Printf("撤销path至: %v\n", path)
		}
	}
	dfs(0, target, []int{})
	return res
}
