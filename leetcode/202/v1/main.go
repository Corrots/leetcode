package main

import (
	"fmt"
)

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
}

// https://leetcode-cn.com/problems/happy-number/
func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 {
		n = step(n)
		if m[n] {
			return false
		}
		m[n] = true
	}
	return true
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
