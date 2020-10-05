package main

import "fmt"

func main() {
	nums1 := []int{3, 2, 1}
	l1 := CreateLinkedList(nums1, len(nums1))
	nums2 := []int{9}
	l2 := CreateLinkedList(nums2, len(nums2))
	PrintLinkedList(addTwoNumbers(l1, l2))
}

// 不需要预先补0
// 在循环遍历链表的过程中直接进行相加和进位操作
//https://leetcode-cn.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	p, q, cur := l1, l2, &ListNode{}
	dummy := cur
	var carry int
	for p != nil || q != nil {
		var sum int
		if p != nil {
			sum += p.Val
			p = p.Next
		}
		if q != nil {
			sum += q.Val
			q = q.Next
		}
		sum += carry
		cur.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		cur = cur.Next
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry % 10}
	}
	return dummy.Next
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
