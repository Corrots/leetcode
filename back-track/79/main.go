package main

import "fmt"

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	//word := "ABCCED"
	word := "ABCB"
	fmt.Println(exist(board, word))
}

// 顺时针的4个方向点
var (
	dirs = [4]point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
)

type point struct {
	i, j int
}

func (p point) next(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

// 检查point是否越界
func (p point) inArea(m, n int) bool {
	return p.i >= 0 && p.i < m && p.j >= 0 && p.j < n
}

// 获取grid的边界(i,j)
func border(board [][]byte) (m int, n int) {
	return len(board), len(board[0])
}

//https://leetcode-cn.com/problems/word-search/
func exist(board [][]byte, word string) bool {
	m, n := border(board)
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		tmp := make([]bool, n)
		visited[i] = append(visited[i], tmp...)
	}
	var search func(int, point) bool
	search = func(index int, p point) bool {
		if index == len(word)-1 {
			return board[p.i][p.j] == word[index]
		}
		if board[p.i][p.j] == word[index] {
			// 标记已经走过的点
			visited[p.i][p.j] = true
			// 从p向四周搜索word的下一个字母
			for i := 0; i < len(dirs); i++ {
				nextPoint := p.next(dirs[i])
				// 检查newPoint是否越界
				if nextPoint.inArea(m, n) && !visited[nextPoint.i][nextPoint.j] {
					if search(index+1, nextPoint) {
						return true
					}
				}
			}
			// 回溯，撤销
			visited[p.i][p.j] = false
		}
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if search(0, point{i, j}) {
				return true
			}
		}
	}
	return false
}
