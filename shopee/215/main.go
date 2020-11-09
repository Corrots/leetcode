package main

import (
	"math/rand"
	"time"
)

func main() {

}

//https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	n := len(nums)
	k = n - k
	l, r := 0, n-1
	for l <= r {
		p := partition(nums, l, r)
		if p == k {
			return nums[p]
		} else if k < p {
			r = p - 1
		} else {
			l = p + 1
		}
	}
	return -1
}

// 随机选取标的点，即随机选择一个元素与nums[l]交换位置
func randPivot(nums []int, l, r int) {
	randIndex := rand.Int()%(r-l+1) + l
	nums[randIndex], nums[l] = nums[l], nums[randIndex]
}

func partition(nums []int, l, r int) int {
	randPivot(nums, l, r)
	v := nums[l]
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
