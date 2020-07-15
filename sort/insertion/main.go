package main

import (
	"fmt"
	"time"
)

func main() {
	nums1 := []int{8, 6, 2, 3, 1, 5, 7, 4}
	InsertionSort(nums1)
	fmt.Println(nums1)
	//nums2 := helper.GenerateRandomData(100000, 0, 1000000)
	//InsertionSort(nums2)
}

// []int{8, 6, 2, 3, 1, 5, 7, 4}
// []int{2, 6, 8, 3, 1, 5, 7, 4}
// 插入排序法优化
func InsertionSort(nums []int) {
	n := len(nums)
	start := time.Now()
	for i := 1; i < n; i++ {
		e := nums[i]
		j := 0
		for j = i; j > 0 && nums[j-1] > e; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = e
	}
	fmt.Printf("spent: %d\n", time.Since(start).Milliseconds())
}

/*
// 插入排序
func InsertionSort(nums []int) {
	n := len(nums)
	start := time.Now()
	for i := 1; i < n; i++ {
		for j := i; j > 0 && nums[j-1] > nums[j]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
	fmt.Printf("spent: %d\n", time.Since(start).Milliseconds())
}
*/
