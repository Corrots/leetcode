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

// 贪心算法求解
//
//https://leetcode-cn.com/problems/non-overlapping-intervals/
func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	// 对区间序列进行排序，将终点小的排在前面
	sort.Sort(Interval(intervals))
	res := 1
	prev := 0
	for i := 1; i < n; i++ {
		// 如果当前区间的起点 >= 前一个区间的终点
		if intervals[i][0] >= intervals[prev][1] {
			res++
			prev = i
		}
	}
	return n - res
}

type Interval [][]int

func (s Interval) Len() int {
	return len(s)
}

func (s Interval) Less(i, j int) bool {
	if s[i][1] != s[j][1] {
		return s[i][1] < s[j][1]
	}
	return s[i][0] < s[j][0]
}

func (s Interval) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
