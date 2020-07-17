package main

import (
	"fmt"
	"time"

	"github.com/corrots/leetcode/heap/maxHeap/heap"
	"github.com/corrots/leetcode/sort/helper"
)

func main() {
	nums := helper.GenerateRandomData(1000000, 0, 1000000)
	HeapSort(nums, len(nums))
	//fmt.Println(nums)
}

func HeapSort(nums []int, n int) {
	start := time.Now()
	maxHeap := heap.New(n)
	for i := 0; i < n; i++ {
		maxHeap.Insert(nums[i])
	}
	// 反向遍历
	for i := n - 1; i >= 0; i-- {
		nums[i] = maxHeap.ExtractMax()
	}
	fmt.Printf("%d ms\n", time.Since(start).Milliseconds())
}
