package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 1, 1, 3}
	fmt.Println(uniqueOccurrences(nums))
	nums1 := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	fmt.Println(uniqueOccurrences(nums1))
	nums2 := []int{1, 2}
	fmt.Println(uniqueOccurrences(nums2))
}

//https://leetcode-cn.com/problems/unique-number-of-occurrences/
func uniqueOccurrences(arr []int) bool {
	if len(arr) == 0 {
		return false
	}
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}

	m := make(map[int]int)
	for _, v := range freq {
		if _, ok := m[v]; ok {
			return false
		} else {
			m[v]++
		}
	}
	return true
}
