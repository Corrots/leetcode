package main

func main() {

}

// 思路
// 左孩子的Next=右孩子
// 右孩子的Next=当前节点的Next的左孩子
// 递归处理当前节点的左子树和右子树

//https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		root.Left.Next = root.Right
	}
	if root.Right != nil {
		if root.Next != nil {
			root.Right.Next = root.Next.Left
		}
	}
	connect(root.Left)
	connect(root.Right)
	return root
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
