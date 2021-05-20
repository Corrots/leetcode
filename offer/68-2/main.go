package main

func main() {

}

// 最近公共祖先为root的情况：1.p = root 或者 q = root；2.p和q位于根节点的两侧；
// 分别在左右子树上进行递归查找，如果遇到当前节点则直接返回
// 若左右子树都有返回，则root为公共祖先
// 若左子树无返回，则右子树的返回为公共祖先
// 若右子树无返回，则左子树的返回为公共祖先

//https://leetcode-cn.com/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
