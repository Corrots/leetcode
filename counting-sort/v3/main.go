package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

const R = 3

type Student struct {
	Name  string
	Score int
}

// 稳定的计数排序
func sortColors(numbs []int) {
	// R 处理元素取值范围是 [0, R) 的计数排序
	var counts [R]int
	for _, num := range numbs {
		counts[num]++
	}
	// 计算每个元素在排完序后处于数组中的边界
	// [index[i], index[i+1]) 区间的值为i
	var index [R + 1]int
	for i := 0; i < R; i++ {
		// index[i+1] = index[i] + cnt[i]
		index[i+1] = index[i] + counts[i]
	}

	for i := 0; i+1 < len(index); i++ {
		// [index[i], index[i+1]) 区间的值为i
		for j := index[i]; j < index[i+1]; j++ {
			numbs[j] = i
		}
	}
}
