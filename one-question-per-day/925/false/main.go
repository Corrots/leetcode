package main

import "fmt"

func main() {
	fmt.Println(isLongPressedName("alex", "aaleex"))
	fmt.Println(isLongPressedName("kikcxmvzi", "kiikcxxmmvvzz"))
}

// 错误解法：使用哈希表，忽略了字符串中字符的顺序性
//https://leetcode-cn.com/problems/long-pressed-name/
func isLongPressedName(name string, typed string) bool {
	freq := make(map[uint8]int)
	for i := 0; i < len(name); i++ {
		freq[name[i]]++
	}
	for i := 0; i < len(typed); i++ {
		if _, ok := freq[typed[i]]; !ok {
			return false
		}
		if freq[typed[i]] > 0 {
			freq[typed[i]]--
		}
	}
	for _, v := range freq {
		if v > 0 {
			return false
		}
	}
	return true
}
