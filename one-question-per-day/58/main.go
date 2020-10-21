package main

import (
	"fmt"
)

func main() {
	//fmt.Println(lengthOfLastWord("Hello World"))
	//fmt.Println(lengthOfLastWord("ABC123"))
	fmt.Println(lengthOfLastWord("        "))
	fmt.Println(lengthOfLastWord("a"))
}

// 题目为返回最后一个单词的长度
// 思路：先从后往前遍历，有空格即终止；再从该位置往前遍历，找到空格再终止；此时两者相减，即为最后单词的长度

//https://leetcode-cn.com/problems/length-of-last-word/
func lengthOfLastWord(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	res := 0
	for i := n - 1; i >= 0; i-- {
		if s[i] != ' ' {
			res++
		}
		if res != 0 && s[i] == ' ' {
			break
		}
	}
	return res
}
