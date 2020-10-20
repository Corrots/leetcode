package main

import "fmt"

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c"))
}

// 尝试双指针解法
// 思路：定义skip来记录'#'，遍历字符串，若是'#'则skip+1；
// 根据skip进行判断，若skip>0，则说明当前字符需要删除，skip-1；若skip为0，则说明当前字符不需删除
//https://leetcode-cn.com/problems/backspace-string-compare/
func backspaceCompare(S string, T string) bool {
	skipS, skipT := 0, 0
	i, j := len(S)-1, len(T)-1
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if S[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}

		for j >= 0 {
			if T[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 {
			if S[i] != T[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}
	return true
}
