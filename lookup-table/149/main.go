package main

import (
	"fmt"
	"strings"
)

func main() {
	points := [][]int{{1, 1}, {2, 2}, {3, 3}}
	fmt.Println(maxPoints(points))
}

//https://leetcode-cn.com/problems/max-points-on-a-line/
func maxPoints(points [][]int) int {
	n := len(points)
	if n < 2 {
		return len(points)
	}
	res := 1
	for i := range points {
		record := make(map[string]int)
		samePoint := 0
		for j := range points {
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] {
				samePoint++
			} else {
				record[getPairStr(slope(points[i], points[j]))]++
			}
		}
		res = max(res, samePoint)
		for _, v := range record {
			res = max(res, v+samePoint)
		}
	}
	return res
}

func slope(pa, pb []int) []int {
	dy := pa[1] - pb[1]
	dx := pa[0] - pb[0]
	if dx == 0 {
		return []int{1, 0}
	}
	if dy == 0 {
		return []int{0, 1}
	}
	g := gcd(abs(dy), abs(dx))
	dy /= g
	dx /= g
	if dx < 0 {
		dy = -dy
		dx = -dx
	}
	return []int{dy, dx}
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func getPairStr(p []int) string {
	var str strings.Builder
	str.WriteString(fmt.Sprintf("%d/%d\n", p[0], p[1]))
	return str.String()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
