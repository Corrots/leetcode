package main

func main() {

}

//https://leetcode-cn.com/problems/convert-bst-to-greater-tree/
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var inOrder func(*TreeNode) *TreeNode
	inOrder = func(node *TreeNode) *TreeNode {
		if node == nil {
			return node
		}
		if node.Right != nil {
			node.Right = inOrder(node.Right)
		}
		sum += node.Val
		node.Val = sum
		if node.Left != nil {
			node.Left = inOrder(node.Left)
		}
		return node
	}
	return inOrder(root)
}

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
