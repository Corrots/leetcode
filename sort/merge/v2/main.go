package main

import (
	"fmt"
	"time"

	"github.com/corrots/leetcode/sort/helper"
)

// 自底向上
func main() {
	nums1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	MergeSortBU(nums1)
	fmt.Println(nums1)
	nums2 := helper.GenerateRandomData(100000, 0, 1000000)
	MergeSortBU(nums2)
}

func MergeSortBU(nums []int) {
	n := len(nums)
	start := time.Now()
	for sz := 1; sz < n; sz += sz {
		for i := 0; i+sz < n; i += sz + sz {
			edge := i + sz + sz - 1
			if edge > n-1 {
				edge = n - 1
			}
			// 对nums[i:i+sz-1]和nums[i+sz:i+2*sz-1]进行归并
			merge(nums, i, i+sz-1, edge)
		}
	}
	fmt.Printf("spent: %d\n", time.Since(start).Milliseconds())
}

func merge(nums []int, l, mid, r int) {
	aux := make([]int, (r-l)+1)
	for i := l; i <= r; i++ {
		aux[i-l] = nums[i]
	}
	i, j := l, mid+1
	for k := l; k <= r; k++ {
		if i > mid {
			nums[k] = aux[j-l]
			j++
		} else if j > r {
			nums[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l] {
			nums[k] = aux[i-l]
			i++
		} else {
			nums[k] = aux[j-l]
			j++
		}
	}
}
