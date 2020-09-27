package main

func main() {

}

/**
1.若p,q < root，则最近公共祖先一定在root的左子树中
2.若p,q > root, 则最近公共祖先一定在root的右子树中
3.root为p,q的最近公共祖先
*/
//https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
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
