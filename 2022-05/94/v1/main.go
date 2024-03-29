package main

func main() {

}

// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
// 中序遍历，递归实现
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
