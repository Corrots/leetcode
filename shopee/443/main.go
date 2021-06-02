package main

import "strconv"

//https://leetcode-cn.com/problems/string-compression/
func compress(chars []byte) int {
	w, anchor := 0, 0
	for r := 0; r < len(chars); r++ {
		if r+1 == len(chars) || chars[r+1] != chars[r] {
			chars[w] = chars[anchor]
			w++
			if r > anchor {
				num := strconv.Itoa(r + 1 - anchor)
				for _, val := range num {
					chars[w] = byte(val)
					w++
				}
			}
			anchor = r + 1
		}
	}
	return w
}

func main() {

}
