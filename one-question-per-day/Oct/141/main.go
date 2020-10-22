package main

import "fmt"

func main() {
	nums := []int{3, 2, 0, -4}
	head := CreateLinkedList(nums, len(nums))
	PrintLinkedList(head)
}

// 思路：利用快慢指针，快指针每次走2步，慢指针每次走1步，如果快慢指针相遇，则存在环形
// 我们定义两个指针，一快一满。慢指针每次只移动一步，而快指针每次移动两步。初始时，慢指针在位置 head，而快指针在位置 head.next。
// 这样一来，如果在移动的过程中，快指针反过来追上慢指针，就说明该链表为环形链表。否则快指针将到达链表尾部，该链表不为环形链表。
//https://leetcode-cn.com/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
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
