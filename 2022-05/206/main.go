package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/reverse-linked-list/
// 206. 反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
