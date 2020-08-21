package main

import "fmt"

func main() {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}
	fmt.Println(fourSumCount(A, B, C, D))
}

// https://leetcode-cn.com/problems/4sum-ii
func fourSumCount(A []int, B []int, C []int, D []int) int {
	var res int
	m := make(map[int]int)
	for _, c := range C {
		for _, d := range D {
			m[c+d]++
		}
	}

	for _, a := range A {
		for _, b := range B {
			if val, ok := m[0-a-b]; ok {
				res += val
			}
		}
	}
	return res
}
