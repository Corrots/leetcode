package main

import (
	"fmt"
	"sort"
)

/**
#56
合并区间
给出一个区间的集合，请合并所有重叠的区间。
*/
func main() {
	nums1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	nums2 := [][]int{{4, 5}, {1, 4}}
	nums3 := [][]int{{1, 4}, {2, 3}}
	fmt.Println(merge(nums1))
	fmt.Println(merge(nums2))
	fmt.Println(merge(nums3))
}

// 排序
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	// 先排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 将第一个区间加入merged
	var merged [][]int
	merged = append(merged, intervals[0])
	for i := 1; i < len(intervals); i++ {
		m := merged[len(merged)-1]
		c := intervals[i]
		// 如果当前区间的左端点大于最后区间的右端点，则当前区间与其他区间无法合并，直接移到最右端
		if c[0] > m[1] {
			merged = append(merged, c)
			continue
		}
		// 如果当前区间的右端点大于最后区间的右端点，合并区间，即将最后区间的右端点更新为当前区间的右端点
		if c[1] > m[1] {
			m[1] = c[1]
		}
	}
	return merged
}
