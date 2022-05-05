package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

// #75 颜色分类
// https://leetcode-cn.com/problems/sort-colors/
func sortColors(numbs []int) {
	var counts [3]int
	for _, num := range numbs {
		counts[num]++
	}

	for i := 0; i < counts[0]; i++ {
		numbs[i] = 0
	}
	for i := counts[0]; i < counts[0]+counts[1]; i++ {
		numbs[i] = 1
	}
	for i := counts[0] + counts[1]; i < counts[0]+counts[1]+counts[2]; i++ {
		numbs[i] = 2
	}
}
