package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/corrots/leetcode/sort/helper"
)

// 优化对于大量重复元素的排序
// 双路快排
func main() {
	nums1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	QuickSort(nums1)
	fmt.Println(nums1)
	nums2 := helper.GenerateRandomData(1000000, 0, 10000000)
	nums3 := make([]int, 1000000)
	copy(nums3, nums2)
	QuickSort(nums2)
	nums4 := helper.GenerateRandomData(1000000, 0, 10)
	QuickSort(nums4)
}
func QuickSort(nums []int) {
	start := time.Now()
	rand.Seed(time.Now().UnixNano())
	quickSort(nums, 0, len(nums)-1)
	fmt.Printf("quick: %d ms\n", time.Since(start).Milliseconds())
}

func quickSort(nums []int, l, r int) {
	//if l >= r {
	//	return
	//}
	if r-l <= 15 {
		helper.InsertionSort(nums, l, r)
		return
	}
	p := partition(nums, l, r)
	quickSort(nums, l, p-1)
	quickSort(nums, p+1, r)
}

// 返回值p使 nums[l:p-1] < nums[p] < nums[p+1: r]
func partition(nums []int, l, r int) (p int) {
	random := rand.Int()%(r-l+1) + l
	nums[l], nums[random] = nums[random], nums[l]
	v := nums[l]
	// nums[l+1,i) <= v; nums(j:r] >= v
	i, j := l+1, r
	for {
		for i <= r && nums[i] < v {
			i++
		}
		for j >= l+1 && nums[j] > v {
			j--
		}
		if i > j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}
