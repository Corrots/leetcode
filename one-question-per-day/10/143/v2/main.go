package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	head := CreateLinkedList(nums1, len(nums1))
	reorderList(head)
	PrintLinkedList(head)
}

// 尝试O(1)的空间复杂度解决
// 思路: 1.先找到链表的中点，将链表后半部分反转
//		2.归并链表的前后两部分
//https://leetcode-cn.com/problems/reorder-list/
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	mid := getMid(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverse(l2)
	merge(l1, l2)
}

// 获取链表中点
func getMid(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 反转链表
func reverse(head *ListNode) *ListNode {
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

// 将2个链表进行归并
func merge(l1, l2 *ListNode) {
	var tmp1, tmp2 *ListNode
	for l1 != nil && l2 != nil {
		tmp1 = l1.Next
		tmp2 = l2.Next
		l1.Next = l2
		l1 = tmp1

		l2.Next = l1
		l2 = tmp2
	}
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
