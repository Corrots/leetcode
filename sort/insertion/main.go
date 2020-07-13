package main

import (
	"fmt"
	"time"

	"github.com/corrots/leetcode/sort/helper"
)

func main() {
	nums1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	InsertionSort(nums1)
	fmt.Println(nums1)
	nums2 := helper.GenerateRandomData(100000, 0, 1000000)
	InsertionSort(nums2)
}

/*
// 插入排序
func InsertionSort(nums []int) {
	n := len(nums)
	start := time.Now()
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			} else {
				break
			}
		}
	}
	fmt.Printf("spent: %d\n", time.Since(start).Milliseconds())
}
*/
// 插入排序法优化
func InsertionSort(nums []int) {
	n := len(nums)
	start := time.Now()
	for i := 1; i < n; i++ {
		e := nums[i]
		j := 0
		for j := i; j > 0 && nums[j-1] > e; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = e
	}

	fmt.Printf("spent: %d\n", time.Since(start).Milliseconds())
}
