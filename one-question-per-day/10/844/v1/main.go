package main

import "fmt"

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c"))
}

// 思路：使用栈，遍历字符串：若字符为'#',取出栈顶元素，若不是'#'，推入栈中

//https://leetcode-cn.com/problems/backspace-string-compare/
func backspaceCompare(S string, T string) bool {
	return backSpace(S) == backSpace(T)
}

func backSpace(s string) string {
	var res []byte
	//n := len(s) - 1
	for i := 0; i < len(s); i++ {
		if s[i] != '#' {
			res = append(res, s[i])
		} else if len(res) > 0 {
			res = res[:len(res)-1]
		}
	}
	return string(res)
}
