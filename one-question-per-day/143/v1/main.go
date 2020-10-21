package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	head := CreateLinkedList(nums1, len(nums1))
	reorderList(head)
	PrintLinkedList(head)
}

// 思路:	1.将链表节点存入数组中，利用数组对应索引来重排链表
// 		2.定义双指针p, q分别位于数组头和尾(对撞指针)
//		3.循环结束,list[p]即为尾节点，list[p].Next=nil
//https://leetcode-cn.com/problems/reorder-list/
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	var list []*ListNode
	cur := head
	for cur != nil {
		list = append(list, cur)
		cur = cur.Next
	}
	p, q := 0, len(list)-1
	for p < q {
		list[p].Next = list[q]
		list[q].Next = list[p+1]
		p++
		q--
	}
	list[p].Next = nil
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
