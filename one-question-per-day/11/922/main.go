package main

import "fmt"

func main() {
	//nums := []int{4, 2, 5, 7}
	//fmt.Println(sortArrayByParityII(nums))
	nums1 := []int{3, 4}
	fmt.Println(sortArrayByParityII(nums1))
}

//https://leetcode-cn.com/problems/sort-array-by-parity-ii/
func sortArrayByParityII(A []int) []int {
	n := len(A)
	if n < 2 {
		return nil
	}
	res := make([]int, n)
	p, q := 0, 1
	for i := 0; i < n; i++ {
		if A[i]%2 == 0 && p < n {
			res[p] = A[i]
			p += 2
		} else if q < n {
			res[q] = A[i]
			q += 2
		}
	}
	return res
}
