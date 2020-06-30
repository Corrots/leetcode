package main

import (
	"fmt"
)

/**
#66
加一
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。
*/
func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{4, 3, 2, 1}
	nums3 := []int{9, 9}

	fmt.Println(plusOne(nums1))
	fmt.Println(plusOne(nums2))
	fmt.Println(plusOne(nums3))
}

func plusOne(digits []int) []int {
	n := len(digits) - 1
	for i := n; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	var newDigits = []int{1}
	digits = append(newDigits, digits...)
	return digits
}
