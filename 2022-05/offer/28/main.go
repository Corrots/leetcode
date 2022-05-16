package main

func main() {

}

// https://leetcode.cn/problems/dui-cheng-de-er-cha-shu-lcof/
// 剑指 Offer 28. 对称的二叉树
// 不能通过现将整个树翻转后与原树比较的方式求解：root == reverse(root) (x)
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

// 参数为树的左右子树
func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	// 左右子树只有一个为nil，则必然不对称
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
