package main

func main() {

}

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
// 236. 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val == root.Val || q.Val == root.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return right
	}
	return left
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
