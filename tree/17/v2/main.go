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

var letterMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
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
	if len(digits) == index {
		res = append(res, s)
	} else {
		char := digits[index]
		letters := letterMap[string(char)]
		for _, l := range letters {
			// findCombinations("23", 1, "a")
			findCombinations(digits, index+1, s+string(l))
		}
	}
}
