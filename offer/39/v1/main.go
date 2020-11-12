package main

func main() {

}

// 哈希表，一次遍历求解

//https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/submissions/
func majorityElement(nums []int) int {
	m := make(map[int]int)
	n := len(nums)
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			m[v]++
		} else {
			m[v]++
		}
		if m[v] > n/2 {
			return v
		}
	}
	return 0
}
