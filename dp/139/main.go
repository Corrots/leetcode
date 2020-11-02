package main

func main() {

}

// 状态定义：
// 状态转移方程：

//https://leetcode-cn.com/problems/word-break/
func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for _, v := range wordDict {
		m[v] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && m[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
