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
	dfs = func(index int, target int, c []int) {
		if target == 0 {
			res = append(res, append([]int{}, c...))
		}
		n := len(candidates)
		for i := index; i < n; i++ {
			index = i
			if target < 0 {
				return
			}
			for i < n-1 && candidates[i] == candidates[i+1] {
				i++
			}
			tmp := make([]int, len(c))
			copy(tmp, c)
			tmp = append(tmp, candidates[i])
			dfs(index+1, target-candidates[i], tmp)
		}
	}
	dfs(0, target, []int{})
	return res
}
