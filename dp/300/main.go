package main

import "fmt"

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
	nums1 := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	fmt.Println(lengthOfLIS(nums1))
}

//暴力解法：找出所有的子序列进行判断，时间复杂度：O(2^n*n)
//思考递归实现：
//* LIS(i)表示以第i个数字为结尾的最长上升子序列的长度；
//* 状态定义：LIS(i) 表示[0…i]范围内，包括数字nums[i]可以获得的最长上升子序列的长度；
//* 状态转移方程：**LIS(i) = max(1 + LIS(j))** (当且仅当：nums[i] > nums[j]时)；

//https://leetcode-cn.com/problems/longest-increasing-subsequence/
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	memo := make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				memo[i] = max(memo[i], 1+memo[j])
			}
		}
	}
	// 遍历memo，返回最大值
	res := 1
	for i := 0; i < n; i++ {
		res = max(res, memo[i])
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
