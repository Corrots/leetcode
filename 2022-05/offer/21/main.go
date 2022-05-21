package main

import "fmt"

func main() {
	fmt.Println(exchange([]int{1, 2, 3, 4}))
}

// https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
// 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
func exchange(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	res := make([]int, n)
	p, q := 0, n-1
	for i := 0; i < n && p <= q; i++ {
		if nums[i]%2 == 0 {
			res[q] = nums[i]
			q--
		} else {
			res[p] = nums[i]
			p++
		}
	}
	return res
}
