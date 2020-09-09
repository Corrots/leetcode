package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}

//https://leetcode-cn.com/problems/permutations-ii/
func permuteUnique(nums []int) [][]int {
	res = [][]int{}
	used := make([]bool, len(nums))
	if len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	getPermutation(nums, 0, []int{}, used)
	return res
}

var (
	res [][]int
)

func getPermutation(nums []int, index int, s []int, used []bool) {
	if index == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, s)
		res = append(res, tmp)
	} else {
		for k, v := range nums {
			if used[k] {
				continue
			}
			// 剪枝
			if k > 0 && v == nums[k-1] && !used[k-1] {
				continue
			}
			s = append(s, v)
			used[k] = true
			getPermutation(nums, index+1, s, used)
			s = s[:len(s)-1]
			used[k] = false

		}
	}
}
