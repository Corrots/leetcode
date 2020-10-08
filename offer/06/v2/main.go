package main

import "fmt"

func main() {
	nums := []int{1, 3, 2}
	head := CreateLinkedList(nums, len(nums))
	fmt.Println(reversePrint(head))
}

// 思路：利用栈后进先出的特性，存储数据
//https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
func reversePrint(head *ListNode) []int {
	var res []int
	if head == nil {
		return res
	}
	stack := NewStack()
	for head != nil {
		stack.Push(head.Val)
		head = head.Next
	}
	for !stack.IsEmpty() {
		res = append(res, stack.Pop())
	}
	return res
}

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() int {
	n := len(s.data) - 1
	ret := s.data[n]
	s.data = s.data[:n]
	return ret
}

func (s *Stack) Peek() int {
	return s.data[len(s.data)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
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
