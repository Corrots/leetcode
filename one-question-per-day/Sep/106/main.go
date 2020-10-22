package main

func main() {

}

//https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
func buildTree(inorder []int, postorder []int) *TreeNode {
	m := make(map[int]int)
	for k, v := range inorder {
		m[v] = k
	}
	var build func(int, int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		// 后序遍历的最后一个元素即为根节点
		n := len(postorder)
		val := postorder[n-1]
		postorder = postorder[:n-1]
		// 获取根节点在inorder中的索引，其右侧即为右子树的元素，左侧为左子树的元素
		rootIndex := m[val]
		root := &TreeNode{Val: val}
		root.Right = build(rootIndex+1, right)
		root.Left = build(left, rootIndex-1)
		return root
	}
	return build(0, len(inorder)-1)
}

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
