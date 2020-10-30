package main

import "fmt"

func main() {
	grid := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	fmt.Println(islandPerimeter(grid))
}

//https://leetcode-cn.com/problems/island-perimeter/
func islandPerimeter(grid [][]int) int {
	var res int
	if len(grid) == 0 {
		return 0
	}
	m, n := border(grid)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				for k := 0; k < 4; k++ {
					next := point{i, j}.next(dirs[k])
					if !next.inArea(m, n) || grid[next.i][next.j] != 1 {
						res++
					}
				}
			}
		}
	}
	return res
}

var dirs = [4]point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

type point struct {
	i, j int
}

// 获取p的下一个点，往上右下左的任意方向
func (p point) next(d point) point {
	return point{
		i: p.i + d.i,
		j: p.j + d.j,
	}
}

// 检查point是否越界
func (p point) inArea(m, n int) bool {
	return p.i >= 0 && p.i < m && p.j >= 0 && p.j < n
}

// 获取grid的边界(m, n)
func border(grid [][]int) (m int, n int) {
	m = len(grid)
	if m == 0 {
		n = 0
	} else {
		n = len(grid[0])
	}
	return
}
