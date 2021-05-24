package main

// 纵向扫描
//https://leetcode-cn.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		// 从第二行开始，以第一行的每个字符作为参照
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != strs[0][i] || i == len(strs[j]) {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func main() {

}
