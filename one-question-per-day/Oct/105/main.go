package main

func main() {

}

//
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// 找到根节点在中序遍历数组中的索引
	var rootIndex int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			rootIndex = i
			break
		}
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:rootIndex+1], inorder[:rootIndex]),
		Right: buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:]),
	}
}

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
