package main

import "fmt"

func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(minArray(nums))
	nums2 := []int{3, 1}
	fmt.Println(minArray(nums2))
}

// 利用二分查找思想，将数组中mid和right位置的元素进行比较，来缩小搜索范围
// 1.当nums[mid] > nums[r]时，最小元素一定不在左边，l = mid+1
// 2.当nums[mid] < nums[r]时，r = mid，mid也可能为最小元素值
// 3.当nums[mid] == nums[r]时，只能将nums[r]剔除掉,r--,下一轮搜索区间是 [l, r - 1]

//https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
func minArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	l, r := 0, n-1
	for l <= r {
		mid := (r-l)/2 + l
		if nums[mid] > nums[r] {
			l = mid + 1
		} else if nums[mid] == nums[r] {
			r--
		} else {
			r = mid
		}
	}
	return nums[l]
}
