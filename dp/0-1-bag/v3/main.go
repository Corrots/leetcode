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


// 动态规划
func Knapsack01(w, v []int, C int) int {
	n := len(w)
	if n == 0 || C <= 0 {
		return 0
	}
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < C+1; j++ {
			memo[i] = append(memo[i], -1)
		}
	}
	// 先解决最基本问题
	for j := 0; j < C+1; j++ {
		memo[0][j] = 0
		// j表示背包当前的容量，只有当j>=物品重量时，才可将物品放入背包
		if j >= w[0] {
			memo[0][j] = v[0]
		}
	}
	for i := 1; i < n; i++ {
		for j := 0; j < C+1; j++ {
			// 选择1：物品i不放入背包，则背包价值=memo[i-1][j]
			memo[i][j] = memo[i-1][j]
			// 选择2：将物品i放入背包，则背包价值=物品i的价值+memo[i-1][j-w[i]]
			if j >= w[i] {
				memo[i][j] = v[i] + memo[i-1][j-w[i]]
			}
		}
	}
	return memo[n-1][C]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
