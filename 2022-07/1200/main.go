package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{3, 8, -10, 23, 19, -4, -14, 27}
	nums1 := []int{4, 2, 1, 3}

	fmt.Println(minimumAbsDifference(nums))
	fmt.Println(minimumAbsDifference(nums1))
}

// https://leetcode.cn/problems/minimum-absolute-difference/
func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	var res [][]int
	best := math.MaxInt
	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if diff < best {
			best = diff
			res = [][]int{{arr[i], arr[i+1]}}
		} else if diff == best {
			res = append(res, []int{arr[i], arr[i+1]})
		}
	}
	return res
}
