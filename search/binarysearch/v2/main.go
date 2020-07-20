package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 3, 4, 4, 4, 5, 6}
	fmt.Printf("floor: %d\n", floor(nums1, 2))
	fmt.Printf("ceil: %d\n", ceil(nums1, 2))

}

// 二分查找法, 在有序数组nums中, 查找target
// 如果找到target, 返回第一个target相应的索引index
// 如果没有找到target, 返回比target小的最大值相应的索引, 如果这个最大值有多个, 返回最大索引
// 如果这个target比整个数组的最小元素值还要小, 则不存在这个target的floor值, 返回-1
func floor(nums []int, target int) int {
	n := len(nums)
	// 寻找比target小的最大索引
	l, r := -1, n-1
	for l < r {
		// 使用向上取整避免死循环
		mid := l + (r-l+1)/2
		if target <= nums[mid] {
			r = mid - 1
		} else {
			l = mid
		}
	}
	if l+1 < n && nums[l+1] == target {
		return l + 1
	}
	return l
}

// 二分查找法, 在有序数组arr中, 查找target
// 如果找到target, 返回最后一个target相应的索引index
// 如果没有找到target, 返回比target大的最小值相应的索引, 如果这个最小值有多个, 返回最小的索引
// 如果这个target比整个数组的最大元素值还要大, 则不存在这个target的ceil值, 返回整个数组元素个数n
func ceil(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n
	for l < r {
		mid := l + (r-l)/2
		if target >= nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if r-1 >= 0 && nums[r-1] == target {
		return r - 1
	}
	return r
}
