package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

// 时间复杂度: O(n)
// 空间复杂度: O(n)
func moveZeroes(nums []int) {
	var nonZeroes []int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nonZeroes = append(nonZeroes, nums[i])
		}
	}
	for i := 0; i < len(nonZeroes); i++ {
		nums[i] = nonZeroes[i]
	}
	for i := len(nonZeroes); i < len(nums); i++ {
		nums[i] = 0
	}
}
