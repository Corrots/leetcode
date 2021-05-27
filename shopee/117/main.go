package main

type Node struct {
	val               int
	Left, Right, Next *Node
}

//https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for len(q) > 0 {
		p := q
		q = nil
		for i, node := range p {
			if i < len(p)-1 {
				node.Next = p[i+1]
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
