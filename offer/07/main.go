package main

func main() {

}

// 通过中序遍历获知root节点在中序遍历中的位置，进而知道左右子树的节点数量
// 进而通过递归来构建左右子树
//https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	var rootIndex int
	for k := range inorder {
		if preorder[0] == inorder[k] {
			rootIndex = k
			break
		}
	}
	root := &TreeNode{Val: preorder[0]}
	root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
	return root
}

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
