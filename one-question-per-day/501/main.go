package main

func main() {

}

//https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/
func findMode(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var inOrder func(*TreeNode)
	orderMap := make(map[int]int)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		orderMap[node.Val]++
		inOrder(node.Right)
	}
	inOrder(root)
	counter := 0
	for k, v := range orderMap {
		if v >= counter {
			counter = v
		}
	}
	return res
}

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
