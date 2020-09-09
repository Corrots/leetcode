package main

import "fmt"

func main() {
	digits := "23"
	fmt.Println(digits[0])
	fmt.Println(digits[0] - '0')
	fmt.Println(letterCombinations(digits))
}

/**
digits 代表数字字符串
s(digits)表示digits所能代表的字母字符串
s(digits[0...n-1])
	= letter(digits[0]) + s(digits[1...n-1])
	= letter(digits[0]) + letter(digits[1]) + s(digits[2...n-1])
*/

var letterMap = [10]string{
	" ",
	"", "abc", "def",
	"ghi", "jkl", "mno",
	"pqrs", "tuv", "wxyz",
}

//https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {
	res = []string{}
	if len(digits) == 0 {
		return res
	}
	findCombinations(digits, 0, "")
	return res
}

var res []string

func findCombinations(digits string, index int, s string) {
	if index == len(digits) {
		res = append(res, s)
		return
	}
	// 第一次，digits[0] = 2
	char := digits[index]
	// 获取数字对应的字符串
	letters := letterMap[char-'0']
	for i := 0; i < len(letters); i++ {
		findCombinations(digits, index+1, s+string(letters[i]))
	}
}
