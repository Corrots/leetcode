package main

import "fmt"

func main() {

}

// https://leetcode-cn.com/problems/insertion-sort-list/
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	// 定义已排好序的链表尾部tail和待排序的链表节点sort
	tail, sort := head, head.Next
	for sort != nil {
		// 待排序元素sort的Val<已排好序链表尾部元素tail
		if sort.Val < tail.Val {
			// 找到sort节点应该在的位置的前驱节点
			prev := dummy
			for prev.Next.Val < sort.Val {
				prev = prev.Next
			}
			tail.Next = sort.Next
			sort.Next = prev.Next
			prev.Next = sort
			sort = tail.Next
		} else {
			tail = tail.Next
			sort = sort.Next
		}
	}
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
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
