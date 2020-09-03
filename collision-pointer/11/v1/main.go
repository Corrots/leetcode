package main

import "fmt"

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(nums))
}

func maxArea(height []int) int {
	n := len(height)
	p, q := 0, n-1
	res := 0
	for p < q {
		area := min(height[p], height[q]) * (q - p)
		res = max(res, area)
		if height[p] <= height[q] {
			p++
		} else {
			q--
		}
	}
	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
