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
	if len(numbs) == 0 {
		return
	}
	var cnt0, cnt1, cnt2 int
	for i := 0; i < len(numbs); i++ {
		if numbs[i] == 0 {
			cnt0++
		} else if numbs[i] == 1 {
			cnt1++
		} else if numbs[i] == 2 {
			cnt2++
		}
	}

	for i := 0; i < cnt0; i++ {
		numbs[i] = 0
	}
	for i := cnt0; i < cnt0+cnt1; i++ {
		numbs[i] = 1
	}
	for i := cnt0 + cnt1; i < len(numbs); i++ {
		numbs[i] = 2
	}
}
