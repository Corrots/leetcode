package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

//https://leetcode-cn.com/problems/permutations/
func permute(nums []int) [][]int {
	res = [][]int{}
	used = make(map[int]bool)
	if len(nums) == 0 {
		return res
	}
	getPermutation(nums, 0, []int{})
	return res
}

var (
	res  [][]int
	used map[int]bool
)

func getPermutation(nums []int, index int, s []int) {
	if index == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, s)
		res = append(res, tmp)
	} else {
		for _, v := range nums {
			if !used[v] {
				s = append(s, v)
				used[v] = true
				getPermutation(nums, index+1, s)
				s = (s)[:len(s)-1]
				used[v] = false
			}
		}
	}
}
