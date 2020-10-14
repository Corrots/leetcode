package main

import "fmt"

func main() {
	w1 := []int{1, 2, 3}
	v1 := []int{6, 10, 12}
	fmt.Println(Knapsack01(w1, v1, 5))
}

// 问题描述：
// 有一个背包，他的容量为C(Capacity)。现在有n种不同的物品，编号为0…n-1，其中每一件物品的重量为w(i)，价值为v(i)。
// 问可以向这个背包中盛放哪些物品，使得在不超过背包容量的基础上，物品的总价值最大。

// 基础递归实现
func Knapsack01(w, v []int, C int) int {
	n := len(w)
	if n == 0 || C <= 0 {
		return 0
	}
	// 用[0...index]中的物品，填充容量为c的背包的最大价值
	var maxValue func(index, c int) int
	maxValue = func(index, c int) int {
		if index <= 0 || c <= 0 {
			return 0
		}
		res := maxValue(index-1, c)
		if c >= w[index] {
			res = max(res, v[index]+maxValue(index-1, c-w[index]))
		}
		return res
	}
	return maxValue(n-1, C)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
