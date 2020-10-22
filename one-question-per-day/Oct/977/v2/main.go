package main

import "fmt"

func main() {
	nums1 := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums1))
	nums2 := []int{-3, -1, 0}
	fmt.Println(sortedSquares(nums2))
}

// 尝试直接使用双指针，不使用归并方式
// https://leetcode-cn.com/problems/squares-of-a-sorted-array/
func sortedSquares(A []int) []int {
	n := len(A)
	if n == 0 {
		return nil
	}
	aux := make([]int, n)
	p, q := 0, n-1
	for i := n - 1; i >= 0; i-- {
		if v, w := pow(A[p]), pow(A[q]); v > w {
			aux[i] = v
			p++
		} else {
			aux[i] = w
			q--
		}
	}
	return aux
}

func pow(i int) int {
	return i * i
}
