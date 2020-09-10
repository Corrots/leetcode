package main

import (
	"fmt"
	"sort"
)

func main() {
	//nums := []int{1, 2, 2}
	nums := []int{4, 4, 4, 1, 4}
	s := subsetsWithDup(nums)
	fmt.Printf("%v\nlen: %d\n", s, len(s))
}

//https://leetcode-cn.com/problems/subsets-ii/
func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	res = append(res, []int{})
	if len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	var dfs func(int, []int)
	dfs = func(start int, path []int) {
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			tmp := make([]int, len(path))
			copy(tmp, path)
			tmp = append(tmp, nums[i])
			res = append(res, tmp)
			dfs(i+1, tmp)
		}
	}
	dfs(0, []int{})
	return res
}
