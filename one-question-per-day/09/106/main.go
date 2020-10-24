package main

func main() {
	buildTree(nil, nil)
}

// 1.后序遍历的最后一个节点即为根节点的val，通过val找到其在中序遍历的位置，左右两侧分别为左右子树的节点
// 2.重新通过递归构建二叉树

//https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
func buildTree(inorder []int, postorder []int) *TreeNode {
	iLen, pLen := len(inorder), len(postorder)
	if iLen == 0 || pLen == 0 {
		return nil
	}
	var rootIndex int
	rootVal := postorder[pLen-1]
	for i := 0; i < iLen; i++ {
		if inorder[i] == rootVal {
			rootIndex = i
			break
		}
	}

	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(inorder[:rootIndex], postorder[:rootIndex]),
		Right: buildTree(inorder[rootIndex+1:], postorder[rootIndex:iLen-1]),
	}
}

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
