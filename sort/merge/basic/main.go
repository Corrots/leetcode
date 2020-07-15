package main

import (
	"fmt"
	"time"

	"github.com/corrots/leetcode/sort/helper"
)

// 自顶向下
func main() {
	nums1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	MergeSort(nums1)
	fmt.Println(nums1)
	nums2 := helper.GenerateRandomData(1000000, 0, 10000000)
	//nums3 := make([]int, 100000)
	//copy(nums3, nums2)
	MergeSort(nums2)
}

func MergeSort(nums []int) {
	n := len(nums)
	start := time.Now()
	mergeSort(nums, 0, n-1)
	fmt.Printf("Merge sort: %d ms\n", time.Since(start).Milliseconds())
}

// 递归使用归并排序对nums[l:r]范围的元素进行排序
func mergeSort(nums []int, l, r int) {
	//if l >= r {
	//	return
	//}
	if r-l <= 15 {
		helper.InsertionSort(nums, l, r)
		return
	}
	mid := (l + r) / 2
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)
	if nums[mid] <= nums[mid+1] {
		merge(nums, l, mid, r)
	}
}

// 将nums[l:mid]和nums[mid+1:r]进行归并
func merge(nums []int, l, mid, r int) {
	aux := make([]int, (r-l)+1)
	for i := l; i <= r; i++ {
		aux[i-l] = nums[i]
	}
	// i, j分别位于左右数组的起始位置
	i, j := l, mid+1
	for k := l; k <= r; k++ {
		// 避免i/j越界
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
