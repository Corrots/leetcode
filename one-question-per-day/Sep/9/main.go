package main

import "fmt"

func main() {
	//n := 121
	n := 12321
	fmt.Println(isPalindrome(n))
}

//https://leetcode-cn.com/problems/palindrome-number/
func isPalindrome(x int) bool {
	// 负数或个位数为0的数字不是回文数，排除0
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	rev := 0
	for x > rev {
		rev = rev*10 + x%10
		x /= 10
	}
	return x == rev || x == rev/10
}
