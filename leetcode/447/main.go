package main

import "fmt"

func main() {
	nums := [][]int{{0, 0}, {1, 0}, {2, 0}}
	fmt.Println(numberOfBoomerangs(nums))
}

// https://leetcode-cn.com/problems/number-of-boomerangs/
func numberOfBoomerangs(points [][]int) int {
	var res int
	for i := range points {
		record := make(map[int]int)
		for j := range points {
			if j != i {
				//fmt.Printf("point i: %v, point j: %v\n", points[i], points[j])
				record[distance(points[i], points[j])]++
			}
		}
		for _, v := range record {
			if v >= 2 {
				res += v * (v - 1)
			}
		}
	}
	return res
}

func distance(i []int, j []int) int {
	return (i[0]-j[0])*(i[0]-j[0]) + (i[1]-j[1])*(i[1]-j[1])
}
