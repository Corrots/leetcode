package main

import "fmt"

func main() {
	nums1 := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums1))
	nums2 := []int{-3, -1, 0}
	fmt.Println(sortedSquares(nums2))
}

// 错误思路：尝试查找数组中0的索引，若没有0，则查找最接近0的元素索引

// 思路：因为数组是有序的，所以通过遍历，找到的第一个非负数，即为正数和负数的分界线
// 然后使用归并排序思想
// https://leetcode-cn.com/problems/squares-of-a-sorted-array/
func sortedSquares(A []int) []int {
	n := len(A)
	if n == 0 {
		return nil
	}
	mid := -1
	for i := 0; i < n && A[i] < 0; i++ {
		mid = i
	}
	p, q := mid, mid+1
	aux := make([]int, n)
	for i := 0; i < n; i++ {
		if p < 0 {
			aux[i] = pow(A[q])
			q++
		} else if q >= n {
			aux[i] = pow(A[p])
			p--
		} else if pow(A[p]) > pow(A[q]) {
			aux[i] = pow(A[q])
			q++
		} else {
			aux[i] = pow(A[p])
			p--
		}
	}
	return aux
}

func pow(i int) int {
	return i * i
}
