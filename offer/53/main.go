package main

import "fmt"

/**
1.通过两次二分查找确定第一个target和最后一个target的位置
2.中间元素与target比较，若>target，则在左半段查找；若<target，则在右半段查找；
3.若=target，则比较mid前一个元素(mid-1)，若=target，则第一个target在左半段
4.若!=target，则第一个target就是mid
5.同理，找最后一个target，判断mid+1位置
*/

//https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	start := l
	r = len(nums)
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	end := l
	return end - start
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(search(nums, 8))
}
