package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/reverse-linked-list-ii/
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if m == 1 {
		return reverseN(head, n)
	}
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

var succ *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		succ = head.Next
		return head
	}
	rev := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = succ
	return rev
}
