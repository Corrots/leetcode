package main

func main() {

}

//https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return nil
	}
	start, end := head, head
	for i := 0; i < k; i++ {
		end = end.Next
	}
	for end != nil {
		start = start.Next
		end = end.Next
	}
	return start
}

type ListNode struct {
	Val  int
	Next *ListNode
}
