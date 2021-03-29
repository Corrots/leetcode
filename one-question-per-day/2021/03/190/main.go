package main

import "fmt"

// https://leetcode-cn.com/problems/reverse-bits/
func reverseBits(num uint32) uint32 {
	var rev uint32
	for i := 0; i < 32 && num > 0; i++ {
		rev |= num & 1 << (31 - i)
		num >>= 1
	}
	return rev
}

func main() {
	fmt.Println(reverseBits(43261596) == 964176192)
	//fmt.Println(reverseBits(11111111111111111111111111111101) == 10111111111111111111111111111111)
	//fmt.Println(reverseBits(00000010100101000001111010011100) == 964176192)
}
