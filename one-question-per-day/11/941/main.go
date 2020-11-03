package main

import "fmt"

func main() {
	nums1 := []int{2, 0, 2}
	fmt.Println(validMountainArray(nums1))
	nums2 := []int{0, 3, 2, 1}
	fmt.Println(validMountainArray(nums2))
	nums3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(validMountainArray(nums3))
}

// 错误思路：双指针求解

//https://leetcode-cn.com/problems/valid-mountain-array/
func validMountainArray(A []int) bool {
	n := len(A)
	if n < 3 {
		return false
	}
	// 找到山脉最高点的索引
	top := 0
	for top < n-1 && A[top] < A[top+1] {
		top++
	}
	if top == 0 || top == n-1 {
		return false
	}
	// 从最后向后继续遍历
	for top < n-1 && A[top] > A[top+1] {
		top++
	}

	return top == n-1
}
