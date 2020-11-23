package main

import "sort"

func main() {

}

//https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/solution/yong-zui-shao-shu-liang-de-jian-yin-bao-qi-qiu-1-2/
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	// 将数组排序，end最小的元素居于最前
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	maxRight := points[0][1]
	res := 1
	for _, point := range points {
		if point[0] > maxRight {
			maxRight = point[1]
			res++
		}
	}
	return res
}
