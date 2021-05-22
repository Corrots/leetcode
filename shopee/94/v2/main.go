package main

func main() {

}

// 中序遍历，迭代实现
// 二叉树的中序遍历：按照访问左子树——根节点——右子树的方式遍历这棵树，而在访问左子树或者右子树的时候我们按照同样的方式遍历，直到遍历完整棵树。
//https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		// 先遍历左子树，将所有节点放入栈中
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 从栈顶取元素依次进行遍历
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		// 遍历右子树
		root = root.Right
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
