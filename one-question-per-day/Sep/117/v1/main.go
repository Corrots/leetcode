package main

func main() {

}

//思路：使用层序遍历的思想，用队列存储每一层的节点，依次为每个节点设置Next指针
//时间复杂度：O(N)。我们需要遍历这棵树上所有的点，时间复杂度为 O(N)。
//空间复杂度：O(N)。即队列的空间代价。
//https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for len(q) > 0 {
		p := q
		q = nil
		//n := len(p)
		for k, node := range p {
			if k < len(p)-1 {
				node.Next = p[k+1]
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return root
}

//
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
