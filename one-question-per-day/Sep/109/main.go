package main

import "fmt"

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	head := CreateLinkedList(nums, len(nums))
	PrintLinkedList(head)
}

//https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/
func sortedListToBST(head *ListNode) *TreeNode {
	return createBST(head, nil)
}

func createBST(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	mid := getMid(left, right)
	root := &TreeNode{Val: mid.Val}
	root.Left = createBST(left, mid)
	root.Right = createBST(mid.Next, right)
	return root
}

// 找到链表的中点
func getMid(left, right *ListNode) *ListNode {
	slow, fast := left, left
	for fast != right && fast.Next != right {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
