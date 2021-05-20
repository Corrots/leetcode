package main

func main() {

}

// 对于二叉搜索树而言，左侧的节点都小于当前节点，右侧都大于当前节点；
// 若p和q位于root的两侧，则root即为p和q的最近公共祖先；
// 若p和q都小于root，则通过递归在root的左子树继续寻找其公共祖先；
// 若p和q都大于root，则通过递归在root的右子树继续寻找其公共祖先；

//https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian-lcof/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
