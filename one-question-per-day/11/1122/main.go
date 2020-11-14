package main

import "sort"

func main() {

}

//https://leetcode-cn.com/problems/relative-sort-array/
func relativeSortArray(arr1 []int, arr2 []int) []int {
	rank := make(map[int]int)
	for i, v := range arr2 {
		rank[v] = i
	}
	sort.Slice(arr1, func(i, j int) bool {
		x, y := arr1[i], arr1[j]
		rankX, hasX := rank[x]
		rankY, hasY := rank[y]
		if hasX && hasY {
			return rankX < rankY
		}
		if hasX || hasY {
			return hasX
		}
		return x < y
	})
	return arr1
}
