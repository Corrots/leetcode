package main

import "fmt"

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}

//https://leetcode-cn.com/problems/single-number/
func singleNumber(nums []int) int {
	var single int
	for _, num := range nums {
		// 基于相同元素异或为0求解
		single ^= num
	}
	return single
}
