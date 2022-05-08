package main

func main() {

}

// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/
// 230. 二叉搜索树中第K小的元素
func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	// 找到第K小的元素即可终止循环，不需要遍历整个树
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
