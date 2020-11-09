package main

func main() {

}

//https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/
func printNumbers(n int) []int {
	// 最大数为10^n-1
	max := 1
	for i := 0; i < n; i++ {
		max *= 10
	}
	max--
	var res []int
	for i := 1; i <= max; i++ {
		res = append(res, i)
	}
	return res
}
