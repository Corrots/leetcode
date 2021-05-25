package main

import (
	"fmt"

	"github.com/corrots/leetcode/sort/helper"
)

func main() {
	nums1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	MergeSort(nums1)
	fmt.Println(nums1)
	nums2 := helper.GenerateRandomData(1000000, 0, 10000000)
	//nums3 := make([]int, 100000)
	//copy(nums3, nums2)
	MergeSort(nums2)
	fmt.Println(helper.IsSorted(nums2))
}

func MergeSort(nums []int) {
	mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := (r-l)/2 + l
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)
	if nums[mid] > nums[mid+1] {
		merge(nums, l, mid, r)
	}
}

func merge(nums []int, l, mid, r int) {
	// 用临时数组来存储所有元素
	aux := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		aux[i-l] = nums[i]
	}
	// 定义i和j分别位于要归并的两个数组的起始位置
	i, j := l, mid+1
	for k := l; k <= r; k++ {
		if i > mid {
			nums[k] = aux[j-l]
			j++
		} else if j > r {
			nums[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l] {
			nums[k] = aux[i-l]
			i++
		} else {
			nums[k] = aux[j-l]
			j++
		}
	}
}
