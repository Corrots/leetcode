package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
}

//https://leetcode-cn.com/problems/verify-preorder-serialization-of-a-binary-tree/
func isValidSerialization(preorder string) bool {
	slots := 1
	n := len(preorder)
	for i := 0; i < n; {
		if slots == 0 {
			return false
		}

		if preorder[i] == ',' {
			i++
		} else if preorder[i] == '#' {
			// 读空节点
			slots--
			i++
		} else {
			// 读数字
			for i < n && preorder[i] != ',' {
				i++
			}
			slots++
		}
	}
	return slots == 0
}
