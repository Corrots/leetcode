package main

import "fmt"

func main() {
	nums := []int{2, 5, 1, 3, 4, 7}
	fmt.Println(shuffle(nums, 3))
}

// https://leetcode.cn/problems/shuffle-the-array/
func shuffle(nums []int, n int) []int {
	res := make([]int, 2*n)
	for i, num := range nums[:n] {
		res[2*i] = num
		res[2*i+1] = nums[n+i]
	}
	return res
}
