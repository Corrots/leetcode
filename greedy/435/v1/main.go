package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	fmt.Println(eraseOverlapIntervals(intervals))
	intervals1 := [][]int{{1, 2}, {1, 2}, {1, 2}}
	fmt.Println(eraseOverlapIntervals(intervals1))
	intervals2 := [][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}
	fmt.Println(eraseOverlapIntervals(intervals2))
}

// 将问题转化为求不重叠子区间的数量
// 动态规划解法
//https://leetcode-cn.com/problems/non-overlapping-intervals/
func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	// 对区间进行排序
	sort.Sort(Interval(intervals))
	// memo[i]表示使用区间intervals[0...i]所能构成的最长的不重叠区间序列
	memo := make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = 1
	}
	//
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			// 只有当区间intervals[i]的起点 >= 之前某个区间intervals[j]的终点时，两个区间才不会重叠
			if intervals[i][0] >= intervals[j][1] {
				memo[i] = max(memo[i], 1+memo[j])
			}
		}
	}
	// 获取不重叠子区间的最大长度
	res := 1
	for i := 0; i < n; i++ {
		res = max(res, memo[i])
	}
	return n - res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

type Interval [][]int

func (s Interval) Len() int {
	return len(s)
}

func (s Interval) Less(i, j int) bool {
	if s[i][0] != s[j][0] {
		return s[i][0] < s[j][0]
	}
	return s[i][1] < s[j][1]
}

func (s Interval) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
