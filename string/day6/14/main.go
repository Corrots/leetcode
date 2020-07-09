package main

import "fmt"

/**
#14
最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
*/
func main() {
	str1 := []string{"flower", "flow", "flight"}
	str2 := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(str1))
	fmt.Println(longestCommonPrefix(str2))
}

func longestCommonPrefix(strs []string) string {

}
