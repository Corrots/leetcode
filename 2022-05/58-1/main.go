package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	n := len(nums)
	newNums := make([]int, n)
	for i, val := range nums {
		newNums[(i+k)%n] = val
	}
	copy(nums, newNums)
}
