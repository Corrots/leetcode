package main

import "fmt"

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid))
}

var (
	dirs = [4]point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
)

type point struct {
	i, j int
}

func (p point) next(d point) point {
	return point{p.i + d.i, p.j + d.j}
}

// 检查point是否越界
func (p point) inArea(m, n int) bool {
	return p.i >= 0 && p.i < m && p.j >= 0 && p.j < n
}

// 获取grid的边界(m, n)
func border(board [][]byte) (m int, n int) {
	m = len(board)
	if m == 0 {
		n = 0
	} else {
		n = len(board[0])
	}
	return
}

//https://leetcode-cn.com/problems/number-of-islands/
func numIslands(grid [][]byte) int {
	var res int
	m, n := border(grid)
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		tmp := make([]bool, n)
		visited[i] = append(visited[i], tmp...)
	}
	// 从point位置开始进行FloodFill
	var dfs func(point)
	dfs = func(p point) {
		visited[p.i][p.j] = true
		for i := 0; i < len(dirs); i++ {
			next := p.next(dirs[i])
			//fmt.Printf("next: %v\n", next)
			if next.inArea(m, n) && !visited[next.i][next.j] && grid[next.i][next.j] == '1' {
				dfs(next)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				res++
				dfs(point{i, j})
			}
		}
	}
	return res
}
