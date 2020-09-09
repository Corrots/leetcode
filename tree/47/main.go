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
	used = make([]bool, len(nums))
	if len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	getPermutation(nums, 0, []int{})
	return res
}

var (
	res  [][]int
	used []bool
)

func getPermutation(nums []int, index int, s []int) {
	if index == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, s)
		res = append(res, tmp)
	} else {
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			// 剪枝
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			s = append(s, nums[i])
			used[i] = true
			getPermutation(nums, index+1, s)
			s = s[:len(s)-1]
			used[i] = false
		}
	}
}
