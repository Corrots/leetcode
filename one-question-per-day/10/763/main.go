package main

import "fmt"

func main() {
	s := "ababcbacadefegdehijhklij"
	fmt.Println(partitionLabels(s))
	// output: ababcbaca defegde hijhklij
}

// 贪心算法求解
// 1.先遍历字符串，找到每个字母最后一次出现的索引;
// 2.对于每个字母c,其最后一次出现的位置为lastIndex[S[i]-'a']，则当前片段的end=max(end, lastIndex[S[i]-'a']);
// 3.当遍历到end时，说明当前片段访问完成，将end-start+1加入解集中，同时更新start=end+1，继续下个片段的遍历

//https://leetcode-cn.com/problems/partition-labels/
func partitionLabels(S string) []int {
	var res []int
	if len(S) == 0 {
		return res
	}
	// 获取每个字符最后一次出现的索引
	lastIndex := [26]int{}
	for i := 0; i < len(S); i++ {
		lastIndex[S[i]-'a'] = i
	}
	// 定义双指针来记录每个片段的首尾位置
	var start, end int
	for i, c := range S {
		end = max(end, lastIndex[c-'a'])
		if i == end {
			res = append(res, end-start+1)
			start = end + 1
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
