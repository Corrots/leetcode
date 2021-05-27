package main

//https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/submissions/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return nil
	}
	// 定义虚拟头节点，其Next为head
	dummy := &ListNode{Next: head}
	start, end := dummy, dummy
	for i := 0; i <= n; i++ {
		end = end.Next
	}

	for end != nil {
		start = start.Next
		end = end.Next
	}
	delNode := start.Next
	start.Next = delNode.Next
	delNode = nil
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
