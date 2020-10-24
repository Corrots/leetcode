package main

import (
	"fmt"
	"math"
)

func main() {
	clips := [][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}
	fmt.Println(videoStitching(clips, 10))
}

// 动态规划求解
// 状态转移方程：dp[i] = min{dp[aj]} + 1 (aj < i <= bj)
//https://leetcode-cn.com/problems/video-stitching/
func videoStitching(clips [][]int, T int) int {
	inf := math.MaxInt64 - 1
	dp := make([]int, T+1)
	for i := 0; i < T+1; i++ {
		dp[i] = inf
	}
	dp[0] = 0
	for i := 1; i < T+1; i++ {
		for _, v := range clips {
			l, r := v[0], v[1]
			if l < i && i <= r && dp[l]+1 < dp[i] {
				dp[i] = dp[l] + 1
			}
		}
	}
	if dp[T] == inf {
		return -1
	}
	return dp[T]
}
