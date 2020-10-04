package main

import "fmt"

func main() {
	nums1 := []int{3, 2, 1}
	l1 := CreateLinkedList(nums1, len(nums1))
	nums2 := []int{9}
	l2 := CreateLinkedList(nums2, len(nums2))
	PrintLinkedList(addTwoNumbers(l1, l2))
}

//	1.获取链表l1和l2的长度，将短的链表末尾补0
//	2.依次取出链表中的每一位，来进行相加操作，注意进位
//https://leetcode-cn.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// 将l1和l2中较短链表缺省部分补0
	p, q := l1, l2
	for p.Next != nil || q.Next != nil {
		if p.Next == nil {
			p.Next = &ListNode{Val: 0}
		}
		if q.Next == nil {
			q.Next = &ListNode{Val: 0}
		}
		p = p.Next
		q = q.Next
	}
	//3->2->1->nil
	//9->0->0->nil
	// 每位求和，并进位
	l3 := &ListNode{}
	res := l3
	var carry int
	for l1 != nil {
		sum := l1.Val + l2.Val + carry
		l3.Next = &ListNode{Val: sum % 10} // 2->3
		carry = sum / 10                   // 1->0
		l3 = l3.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	if carry > 0 {
		l3.Next = &ListNode{Val: carry % 10}
	}
	return res.Next
}

//

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
