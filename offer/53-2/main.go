package main

import "fmt"

/**
二分查找
*/

func missingNumber(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (r-l)/2 + l
		if nums[mid] > mid {
			r = mid - 1
		} else if nums[mid] == mid {
			l = mid + 1
		}
	}
	return l
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Println(missingNumber(nums))
}
