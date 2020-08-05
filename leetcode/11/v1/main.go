package main

import "fmt"

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(nums))
}

func maxArea(height []int) int {
	n := len(height)
	i, j := 0, n-1
	res := 0
	for i < j {
		area := min(height[i], height[j]) * (j - i)
		res = max(res, area)
		if height[i] <= height[j] {
			i++
		} else {
			j--
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
