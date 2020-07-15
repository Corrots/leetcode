package main

import (
	"fmt"
	"time"

	"github.com/corrots/leetcode/sort/helper"
)

func main() {
	nums1 := []int{8, 6, 2, 3, 1, 5, 7, 4}
	BubbleSort(nums1)
	fmt.Println(nums1)
	nums2 := helper.GenerateRandomData(100000, 0, 1000000)
	BubbleSort(nums2)
}

func BubbleSort(nums []int) {
	n := len(nums)
	start := time.Now()
	change := true
	for i := 0; i < n-1 && change; i++ {
		change = false
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				change = true
			}
		}
	}
	fmt.Printf("spent: %d\n", time.Since(start).Milliseconds())
}
