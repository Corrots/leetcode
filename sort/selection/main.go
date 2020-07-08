package main

import (
	"fmt"
	"time"
)

func main() {
	nums1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	SelectionSort(nums1)
	fmt.Println(nums1)
	//nums2 := helper.GenerateRandomData(100000, 0, 1000000)
	//SelectionSort(nums2)
	//fmt.Println(nums2)
}

// 选择排序
func SelectionSort(nums []int) {
	n := len(nums)
	start := time.Now()
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
	fmt.Printf("spent: %d\n", time.Since(start).Milliseconds())
}
