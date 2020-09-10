package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	fmt.Println(combinationSum2(candidates, 8))
}

//https://leetcode-cn.com/problems/combination-sum-ii/
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	if len(candidates) == 0 {
		return res
	}
	sort.Ints(candidates)
	var dfs func(int, int, []int)
	dfs = func(start int, target int, path []int) {
		if target == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		if target < 0 {
			return
		}
		// target > 0
		for i := start; i < len(candidates); i++ {
			// 去重
			if i > start && candidates[i-1] == candidates[i] {
				continue
			}
			tmp := make([]int, len(path))
			copy(tmp, path)
			tmp = append(tmp, candidates[i])
			dfs(i+1, target-candidates[i], tmp)
		}
	}
	dfs(0, target, []int{})
	return res
}
