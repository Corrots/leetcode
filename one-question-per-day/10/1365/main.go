package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{8, 1, 2, 2, 3}
	fmt.Println(smallerNumbersThanCurrent(nums))
}

//https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number/
func smallerNumbersThanCurrent(nums []int) []int {
	var res []int
	n := len(nums)
	if n == 0 {
		return res
	}
	cp := make([]int, n)
	for i := 0; i < n; i++ {
		cp[i] = nums[i]
	}
	sort.Ints(nums)
	fmt.Println("cp: ", cp)
	fmt.Println("nums: ", nums)
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if j != i && nums[j] < nums[i] {
				count++
			}
		}
		res = append(res, count)
	}
	return res
}
