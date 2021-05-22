package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums := []int{5, 1, 1, 2, 0, 0}
	rand.Seed(time.Now().UnixNano())
	QuickSort(nums)
	fmt.Println(nums)
}

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(nums, l, r)
	quickSort(nums, l, p-1)
	quickSort(nums, p+1, r)
}

func partition(nums []int, l, r int) int {
	v := nums[l]
	// <v和>v的分界点的初始位置
	j := l
	for i := l + 1; i <= r; i++ {
		if nums[i] < v {
			nums[j+1], nums[i] = nums[i], nums[j+1]
			j++
		}
	}
	nums[j], nums[l] = nums[l], nums[j]
	return j
}

// 随机选取标的点
func pivot(nums []int, l, r int) {
	randIndex := rand.Int()%(r-l+1) + l
	nums[randIndex], nums[l] = nums[l], nums[randIndex]
}
