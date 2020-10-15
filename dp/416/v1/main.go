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
	if n == 1 {
		return false
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	// memo[i][c]表示使用索引为[0...i]的元素，能否填充容量为c的背包
	// -1表示未计算，0表示不可填充，1表示可填充
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < sum/2+1; j++ {
			memo[i] = append(memo[i], -1)
		}
	}
	var dfs func(index, sum int) bool
	dfs = func(index, sum int) bool {
		if sum == 0 {
			return true
		}
		if index < 0 || sum < 0 {
			return false
		}
		if memo[index][sum] != -1 {
			return memo[index][sum] == 1
		}
		res := dfs(index-1, sum) || dfs(index-1, sum-nums[index])
		if res {
			memo[index][sum] = 1
		} else {
			memo[index][sum] = 0
		}
		return res
	}
	return dfs(n-1, sum/2)
}
