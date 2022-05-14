package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(minArray(nums))
	fmt.Println(minArray([]int{3, 3, 1, 3}))
}

// https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
// 剑指 Offer 11. 旋转数组的最小数字
// 二分查找
// 1. nums[mid] > nums[r], 则最小值在mid的右侧
// 2. nums[mid] < nums[r], 则最小值在mid的左侧
// 3. 因为存在重复元素，所以nums[mid] = nums[r]，可将r前移一位
// 4. 最后nums[l] 即为最小值
func minArray(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (r-l)/2 + l
		if nums[mid] > nums[r] {
			l = mid + 1
		} else if nums[mid] < nums[r] {
			r = mid
		} else {
			r--
		}
	}
	return nums[l]
}
