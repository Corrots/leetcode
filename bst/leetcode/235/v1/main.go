package main

func main() {

}

/**
1.若p,q < root，则公共祖先一定在root的左子树中
2.若p,q > root, 则公共祖先一定在root的右子树中

*/
//https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
