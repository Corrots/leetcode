package main

func main() {

}

// 哈希表记录已保存的链表节点

//https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
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

func getClonedNode(m map[*Node]*Node, node *Node) *Node {
	if v, ok := m[node]; ok {
		return v
	} else {
		var v *Node
		if node != nil {
			// 新建链表节点，并存入哈希表中
			v = &Node{Val: node.Val}
		}
		m[node] = v
		return v
	}
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
