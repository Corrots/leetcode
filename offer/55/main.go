package main

//https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}
