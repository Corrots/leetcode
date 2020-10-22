package main

import "fmt"

func main() {
	fmt.Println(isLongPressedName("alex", "aaleex"))
	fmt.Println(isLongPressedName("kikcxmvzi", "kiikcxxmmvvzz"))
}

// 解法：双指针
// 思路	1.定义两个指针分别执行name和typed首位，若name[p] == typed[q]，则p++, q++；
//		2.若name[p]!=typed[q]，但是type[q] == typed[q-1]，则q++
//		3.最后若 p == len(name)，即为true
//https://leetcode-cn.com/problems/long-pressed-name/
func isLongPressedName(name string, typed string) bool {
	var p, q int
	for q < len(typed) {
		if p < len(name) && name[p] == typed[q] {
			p++
			q++
		} else if q-1 >= 0 && typed[q] == typed[q-1] {
			q++
		} else {
			return false
		}
	}
	return p == len(name)
}
