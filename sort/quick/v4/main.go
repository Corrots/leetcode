package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/corrots/leetcode/sort/helper"
)

// 三路快排
func main() {
	nums1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	QuickSort3Ways(nums1)
	fmt.Println(nums1)
	nums2 := helper.GenerateRandomData(1000000, 0, 10000000)
	nums3 := make([]int, 1000000)
	copy(nums3, nums2)
	QuickSort3Ways(nums2)
	nums4 := helper.GenerateRandomData(1000000, 0, 10)
	QuickSort3Ways(nums4)
}
func QuickSort3Ways(nums []int) {
	start := time.Now()
	rand.Seed(time.Now().UnixNano())
	quickSort3Ways(nums, 0, len(nums)-1)
	fmt.Printf("quick: %d ms\n", time.Since(start).Milliseconds())
}

// 将nums[l:r]分为<v, ==v, >v 三个部分
// 然后对<v和>v部分递归进行三路快排
func quickSort3Ways(nums []int, l, r int) {
	if r-l <= 15 {
		helper.InsertionSort(nums, l, r)
		return
	}
	lt, gt := partition(nums, l, r)
	quickSort3Ways(nums, l, lt-1)
	quickSort3Ways(nums, gt+1, r)
}

// 返回值lt，左侧为<v的元素
// 返回值gt，右侧为>v的元素
func partition(nums []int, l, r int) (int, int) {
	random := rand.Int()%(r-l+1) + l
	nums[l], nums[random] = nums[random], nums[l]
	v := nums[l]
	// 初始化lt, gt
	// 确保 nums[l+1,lt],nums[gt:r], nums[lt+1:i)都为空
	lt, gt := l, r+1
	for i := l + 1; i <= r; i++ {
		if i >= gt {
			break
		} else if nums[i] < v {
			nums[i], nums[lt+1] = nums[lt+1], nums[i]
			lt++
		} else if nums[i] > v {
			nums[i], nums[gt-1] = nums[gt-1], nums[i]
			gt--
		}
	}
	nums[l], nums[lt] = nums[lt], nums[l]
	return lt, gt
}
