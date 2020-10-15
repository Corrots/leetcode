package main

import "fmt"

func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}

// 思路：

// 遍历nums中的元素，求和sum，
// 对于nums[i]，有两种选择：
// 1：nums[i]不加入子集，F(i) = F(i-1, sum)
// 2：nums[i]加入子集，F(i) = F(i-1, sum-nums[i])

//https://leetcode-cn.com/problems/partition-equal-subset-sum/
func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	C := sum / 2
	memo := make([]bool, C+1)
	for i := 0; i < C+1; i++ {
		memo[i] = nums[0] == i
	}

	for i := 1; i < n; i++ {
		for j := C; j >= nums[i]; j-- {
			memo[j] = memo[j] || memo[j-nums[i]]
		}
	}
	return memo[C]
}
