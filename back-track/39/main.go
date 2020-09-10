package main

import (
	"fmt"
)

func main() {
	candidates := []int{2, 3, 6, 7}
	fmt.Println(combinationSum(candidates, 7))
	candidates1 := []int{2, 3, 5}
	fmt.Println(combinationSum(candidates1, 8))
}

// https://leetcode-cn.com/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	if len(candidates) == 0 {
		return res
	}
	var backTrack func(int, int, []int)
	backTrack = func(i int, target int, c []int) {
		if i == len(candidates) || target < 0 {
			return
		}
		if target == 0 {
			res = append(res, append([]int{}, c...))
			return
		}
		// 1.直接跳过当前元素
		backTrack(i+1, target, c)
		// 2.继续当前元素
		c = append(c, candidates[i])
		backTrack(i, target-candidates[i], c)
		c = c[:len(c)-1]
	}
	backTrack(0, target, []int{})
	return res
}
