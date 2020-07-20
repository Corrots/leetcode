package main

import (
	"fmt"

	"github.com/corrots/leetcode/search/tools"
	"github.com/corrots/leetcode/sort/helper"
)

func main() {
	nums2 := helper.GenerateRandomData(1000000, 0, 1000000)
	tools.QuickSort3Ways(nums2)
	index1 := BinarySearch2(nums2, 5)
	if index1 != -1 {
		fmt.Printf("递归二分查找法 %d, index= %d\n", nums2[index1], index1)
	}
	i := BinarySearch(nums2, 5)
	if i != -1 {
		fmt.Printf("find %d, index= %d\n", nums2[i], i)
	}
}

func BinarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func BinarySearch2(nums []int, target int) int {
	return binarySearch2(nums, target, 0, len(nums))
}

// 递归方式实现二分查找法
func binarySearch2(nums []int, target, l, r int) int {
	// 递归退出条件
	if l > r {
		return -1
	}
	mid := l + (r-l)/2
	if nums[mid] == target {
		return mid
	} else if target < nums[mid] {
		return binarySearch2(nums, target, l, mid-1)
	} else {
		return binarySearch2(nums, target, mid+1, r)
	}
}
