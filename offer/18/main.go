package main

func main() {

}

//https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil {
		if val == cur.Next.Val {
			cur.Next = cur.Next.Next
			break
		}
		cur = cur.Next
	}
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
