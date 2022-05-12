package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/corrots/leetcode/sort/helper"
)

func main() {
	nums1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	QuickSort(nums1)
	fmt.Println(nums1)
	nums2 := helper.GenerateRandomData(1000000, 0, 10000000)
	nums3 := make([]int, 1000000)
	copy(nums3, nums2)
	QuickSort(nums2)
	fmt.Println(helper.IsSorted(nums1))
	fmt.Println(helper.IsSorted(nums2))
}

func QuickSort(nums []int) {
	start := time.Now()
	quickSort(nums, 0, len(nums)-1)
	fmt.Printf("spent: %d ms\n", time.Since(start).Milliseconds())
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(nums, l, r)
	quickSort(nums, l, p-1)
	quickSort(nums, p+1, r)
}

// 返回值p使 nums[l:p-1] < nums[p] < nums[p+1: r]
func partition(nums []int, l, r int) (p int) {
	// 不选取第一个元素作为标的点
	//v := nums[l]
	// 随机选取一个[l:r]中的一个元素作为标的点
	random := rand.Int()%(r-l+1) + l
	nums[random], nums[l] = nums[l], nums[random]
	v := nums[l]
	// j作为区分>v和<v的分界点，初始位置为l
	j := l
	for i := l + 1; i <= r; i++ {
		if nums[i] < v {
			nums[i], nums[j+1] = nums[j+1], nums[i]
			j++
		}
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}
