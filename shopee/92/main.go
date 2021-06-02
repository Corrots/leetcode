package main

import "fmt"

//https://leetcode-cn.com/problems/reverse-linked-list-ii/
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

// 全局后继节点
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

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	head := CreateLinkedList(nums, len(nums))
	PrintLinkedList(head)
	reverseBetween(head, 2, 4)
	PrintLinkedList(head)
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func CreateLinkedList(nums []int, n int) *ListNode {
	if n == 0 {
		return nil
	}
	head := NewListNode(nums[0])
	cur := head
	for i := 1; i < n; i++ {
		cur.Next = NewListNode(nums[i])
		cur = cur.Next
	}
	return head
}

func PrintLinkedList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%d->", cur.Val)
		cur = cur.Next
	}
	fmt.Printf("nil\n")
}
