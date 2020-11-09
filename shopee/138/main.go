package main

func main() {

}

// 使用哈希表记录已拷贝的节点

//https://leetcode-cn.com/problems/copy-list-with-random-pointer/
func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	visitedMap := make(map[*Node]*Node)
	oldNode := head
	newNode := getClonedNode(visitedMap, oldNode)
	newHead := newNode
	for oldNode != nil {
		newNode.Next = getClonedNode(visitedMap, oldNode.Next)
		newNode.Random = getClonedNode(visitedMap, oldNode.Random)

		oldNode = oldNode.Next
		newNode = newNode.Next
	}
	return newHead
}

func getClonedNode(visitedMap map[*Node]*Node, node *Node) *Node {
	if v, ok := visitedMap[node]; ok {
		return v
	} else {
		var v *Node
		if node != nil {
			v = &Node{Val: node.Val}
		}
		visitedMap[node] = v
		return v
	}
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
