package main

func main() {

}

// https://leetcode.cn/problems/shan-chu-lian-biao-de-jie-dian-lcof/
// 剑指 Offer 18. 删除链表的节点
func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	prev, cur := head, head.Next
	for prev != nil && cur.Val != val {
		prev = cur
		cur = cur.Next
	}
	if cur != nil {
		prev.Next = cur.Next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
