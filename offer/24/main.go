package main

func main() {

}

// https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
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

type ListNode struct {
	Val  int
	Next *ListNode
}
