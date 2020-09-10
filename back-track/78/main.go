package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}

//https://leetcode-cn.com/problems/subsets/
func subsets(nums []int) [][]int {
	res := [][]int{[]int{}}
	if len(nums) == 0 {
		return res
	}
	var dfs func(int, []int)
	dfs = func(start int, path []int) {
		for i := start; i < len(nums); i++ {
			tmp := make([]int, len(path))
			copy(tmp, path)
			tmp = append(tmp, nums[i])
			res = append(res, append([]int{}, tmp...))
			dfs(i+1, tmp)
		}
	}
	dfs(0, []int{})
	return res
}
