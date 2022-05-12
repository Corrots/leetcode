package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Println(missingNumber(nums))
}

// https://leetcode.cn/problems/que-shi-de-shu-zi-lcof/
// 剑指 Offer 53 - II. 0～n-1中缺失的数字
/**
二分查找
1. nums[mid] = mid，则mid及之前的元素都未缺失,l = mid+1
2. nums[mid] != mid，则说明mid前面有元素缺失, r = mid-1
3. l 即为缺失的数字
*/
func missingNumber(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (r-l)/2 + l
		if nums[mid] == mid {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
